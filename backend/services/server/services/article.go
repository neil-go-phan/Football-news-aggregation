package services

import (
	"backend/services/server/entities"
	"backend/services/server/helper"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"strings"

	pb "backend/grpcfile"

	"github.com/520MianXiangDuiXiang520/MapSize"
	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v7"
	jsoniter "github.com/json-iterator/go"
	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
)

var PREV_ARTICLES = make(map[string]bool)

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

func (s *articleService) FrontendSearchWithIndex(keyword string, indexName string) ([]entities.Article, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	articles := make([]entities.Article, 0)
	var buffer bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  keyword,
				"fields": []string{"title", "discription"},
			},
		},
	}
	json.NewEncoder(&buffer).Encode(query)
	resp, err := s.es.Search(s.es.Search.WithIndex(indexName), s.es.Search.WithBody(&buffer))
	if err != nil {
		return articles, fmt.Errorf("request to elastic search fail")
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		article := hit.(map[string]interface{})["_source"].(map[string]interface{})
		articles = append(articles, newEntitiesArticleFromMap(article))
	}
	return articles, nil
}

// func printAr(a entities.Article) {
// 	fmt.Println("Title: ", a.Title)
// 	fmt.Println("Description: ", a.Description)
// 	fmt.Println("Link: ", a.Link)
// 	fmt.Println("Thumbnail: ", a.Thumbnail)
// }

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
		return
	}

	done := make(chan bool)
	var mapSearchResult = make(map[string]bool)
	log.Printf("PREV_MAP before: len: %v, memory: %v\n", len(PREV_ARTICLES), mapsize.Size(PREV_ARTICLES))
	go func() {
		for {
			resp, err := stream.Recv()

			keyword := resp.GetKeyword()
			respArticles := resp.GetArticles()

			checkSimilarArticle(respArticles, s.es, keyword)

			saveToMapSearchResult(respArticles, mapSearchResult)

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
	PREV_ARTICLES = mapSearchResult
	log.Printf("PREV_MAP before: len: %v, memory: %v\n", len(PREV_ARTICLES), mapsize.Size(PREV_ARTICLES))
	log.Printf("finished")
}

// Nếu gửi từng bài bào lên elastic check thì mỗi lần tìm sẽ gửi vài ngàn request
// C1: Server lưu kết quả cào ở lần trước đó, sau đó lấy kết quả mới so sánh với cũ, nếu có bài báo nào mới thì sẽ check lại với elasticsearch. Elasticsearch chưa có thì thêm vào

func checkSimilarArticle(respArticles []*pb.Article, es *elasticsearch.Client, keyword string) {
	indexName := helper.FormatElasticSearchIndexName(keyword)
	// Condition: similar title
	for _, article := range respArticles {
		// check if it a bet web
		if strings.Contains(strings.ToLower(article.Description), "cá cược") {
			continue
		}
		_, ok := PREV_ARTICLES[article.Title]
		if !ok {
			checkWithElasticSearch(article, indexName, es)
			// checkWithRedis(article)
		}
	}
}

func checkWithElasticSearch(article *pb.Article, indexName string, es *elasticsearch.Client) {
	req := esapi.ExistsRequest{
		Index:      indexName,
		DocumentID: strings.ToLower(article.Title),
	}
	resp, err := req.Do(context.Background(), es)
	if err != nil {
		log.Printf("Error checking if document exists: %s\n", err)
		return
	}
	status := resp.StatusCode
	if status == 200 {
		log.Println("Document already exist in elastic search")
	} else if status == 404 {
		log.Println("Document not found, creating new one...")
		storeElasticsearch(article, indexName, es)
	}
}

// func checkWithRedis(){}

func saveToMapSearchResult(respArticles []*pb.Article, mapSearchResult map[string]bool) {
	for _, article := range respArticles {
		mapSearchResult[article.Title] = true
	}
}

func newEntitiesArticleFromMap(respArticle map[string]interface{}) entities.Article {
	article := entities.Article{
		Title:       respArticle["title"].(string),
		Description: respArticle["description"].(string),
		Thumbnail:   respArticle["thumbnail"].(string),
		Link:        respArticle["link"].(string),
	}
	return article
}

func newEntitiesArticleFromPb(respArticle *pb.Article) entities.Article {
	article := entities.Article{
		Title:       respArticle.Title,
		Description: respArticle.Description,
		Thumbnail:   respArticle.Thumbnail,
		Link:        respArticle.Link,
	}
	return article
}

func storeElasticsearch(article *pb.Article, indexName string, es *elasticsearch.Client) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	body, err := json.Marshal(newEntitiesArticleFromPb(article))
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
		log.Printf("[%s] Indexed document with index: %s \n", res.Status(), indexName)
	}
}
