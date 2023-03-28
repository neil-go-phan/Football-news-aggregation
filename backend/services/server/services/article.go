package services

import (
	"backend/services/server/helper"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	pb "backend/grpcfile"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
)

type articleService struct {
	conn               *grpc.ClientConn
	es                 *elasticsearch.Client
	htmlClassesService *htmlClassesService
	keywordsService    *keywordsService
}

func NewArticleService(keywords *keywordsService, htmlClass *htmlClassesService, conn *grpc.ClientConn, es *elasticsearch.Client) *articleService {
	articleService := &articleService{
		conn:               conn,
		es:                 es,
		htmlClassesService: htmlClass,
		keywordsService:    keywords,
	}
	return articleService
}

func (s *articleService) GetArticlesEveryMinutes(cronjob *cron.Cron) {
	s.getArticles()
	_, err := cronjob.AddFunc("@every 0h01m", func() { s.getArticles() })
	if err != nil {
		log.Println("error occurred while seting up getArticle cronjob: ", err)
	}
}

// Get scrap data result and store it in elastic search,
func (s *articleService) getArticles() {
	client := pb.NewArticleServiceClient(s.conn)

	in := &pb.AllConfigs{
		HtmlClasses: &pb.HTMLClasses{
			ArticleClass:     s.htmlClassesService.HtmlClasses.ArticleClass,
			TitleClass:       s.htmlClassesService.HtmlClasses.TitleClass,
			DescriptionClass: s.htmlClassesService.HtmlClasses.DescriptionClass,
			ThumbnailClass:   s.htmlClassesService.HtmlClasses.ThumbnailClass,
			LinkClass:        s.htmlClassesService.HtmlClasses.LinkClass,
		},
		Keywords: s.keywordsService.Keywords.Keywords,
	}

	stream, err := client.GetArticles(context.Background(), in)
	if err != nil {
		log.Printf("error occurred while openning stream error %v \n", err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			log.Printf("Keywords: %s\n", resp.GetKeyword())

			checkSimilarArticle(resp, s.es)
			// for _, article := range resp.GetArticles() {
			// 	storeInElasticsearch(article, helper.FormatElasticSearchIndexName(resp.GetKeyword()), s.es)
			// }
			if err == io.EOF {
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Printf("cannot receive %v\n", err)
			}
		}
	}()

	<-done //we will wait until all response is received
	log.Printf("finished")
}

// Condition: similar discription
func checkSimilarArticle(resp *pb.ArticlesReponse, es *elasticsearch.Client) bool {
	indexName := helper.FormatElasticSearchIndexName(resp.GetKeyword())
	fmt.Println(indexName)
	var body []byte
// 		body = []byte(`{"index":"cupc1"}
// {"query" : {"match" : {"description": "sao bayern munich muốn vô địch cúp c1"} }}
// `)

	for _, article := range resp.GetArticles() {
		body = append(body, []byte(fmt.Sprintln("{\"index\":\"", resp.GetKeyword(),"\"}\n{\"query\" : {\"match\" : {\"description\": \"",article.Description,"\"} }}"))...)
	}
	log.Println(string(body[:]))
	res, _ := es.API.Msearch(bytes.NewReader(body))
	defer res.Body.Close()
	// log.Printf("%#v\n", res)

	// if res.IsError() {
	// 	panic(fmt.Sprintf("error checking documents: %s", res.Status()))
	// }
	var response map[string]interface{}
	json.NewDecoder(res.Body).Decode(&response)
	log.Printf("%#v\n", response)

	return true
}

func storeElasticsearch(article *pb.Article, indexName string, es *elasticsearch.Client) {
	body, err := json.Marshal(article)
	if err != nil {
		log.Printf("Error encoding article: %s\n", err)
	}

	req := esapi.IndexRequest{
		Index:      indexName,
		DocumentID: strings.ToLower(article.Title),
		Body:       strings.NewReader(string(body)),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Printf("Error getting response: %s\n", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error indexing document\n", res.Status())
	} else {
		log.Printf("[%s] Indexed document\n", res.Status())
	}
}
