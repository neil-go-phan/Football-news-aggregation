package services

import (
	"server/entities"
	"server/helper"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"strings"

	pb "server/proto"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v7"
	jsoniter "github.com/json-iterator/go"

	"google.golang.org/grpc"
)

var PREV_ARTICLES = make(map[string]bool)
var ARTICLES_INDEX_NAME = "articles"
type articleService struct {
	conn               *grpc.ClientConn
	es                 *elasticsearch.Client
	htmlClassesService *htmlClassesService
	leaguesService    *leaguesService
	tagsService        *tagsService
}

func NewArticleService(leagues *leaguesService, htmlClass *htmlClassesService, tags *tagsService, conn *grpc.ClientConn, es *elasticsearch.Client) *articleService {
	articleService := &articleService{
		conn:               conn,
		es:                 es,
		htmlClassesService: htmlClass,
		leaguesService:    leagues,
		tagsService:        tags,
	}
	return articleService
}

func (s *articleService)FrontendSearchArticlesTagsAndKeyword(keyword string, formatedTags []string) ([]entities.Article, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	articles := make([]entities.Article, 0)
	var buffer bytes.Buffer



	var filterQueries []map[string]interface{}
	for _, tag := range formatedTags {
		if tag != "" {
			tagQuery := map[string]interface{}{"match_phrase": map[string]interface{}{"tags": tag}}
			filterQueries = append(filterQueries, tagQuery)
		}
	}

	// default is search all 
	query := querySearchAll()

	if len(filterQueries) == 0 && keyword != "" {
		// search with only keyword
		query = queryWithOnlySearchKeyword(keyword)
	}

	if len(filterQueries) != 0 && keyword == "" {
		// search with only tags
		query = queryWithOnlyTag(filterQueries)
	}

	if len(filterQueries) != 0 && keyword != "" {
		// search with both tags and keyword
		query = queryWithBothTagAndKeyword(keyword, filterQueries)
	}

	json.NewEncoder(&buffer).Encode(query)
	resp, err := s.es.Search(s.es.Search.WithIndex(ARTICLES_INDEX_NAME), s.es.Search.WithBody(&buffer))
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

func (s *articleService)FrontendSearchAll(search_type string, scroll string, size string) ([]entities.Article ,error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	articles := make([]entities.Article, 0)
	var buffer bytes.Buffer

	query := querySearchAll()

	json.NewEncoder(&buffer).Encode(query)
	resp, err := s.es.Search(s.es.Search.WithIndex(ARTICLES_INDEX_NAME), s.es.Search.WithBody(&buffer))
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

func querySearchAll() map[string]interface{} {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all" : map[string]interface{}{},
		},
	}
	return query
}

func queryWithOnlySearchKeyword(keyword string) map[string]interface{}{
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"multi_match": map[string]interface{}{
						"query":  keyword,
						"fields": []string{"title", "description"},
					},
				},
			},
		},
	}
	return query
}

func queryWithOnlyTag(filterQueries []map[string]interface{}) map[string]interface{}{
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": map[string]interface{}{
					"bool": map[string]interface{}{
						"must": filterQueries,
					},
				},
			},
		},
	}
	return query
}

func queryWithBothTagAndKeyword(keyword string, filterQueries []map[string]interface{}) map[string]interface{}{
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"multi_match": map[string]interface{}{
						"query":  strings.TrimSpace(keyword),
						"fields": []string{"title", "description"},
					},
				},
				"filter": map[string]interface{}{
					"bool": map[string]interface{}{
						"must": filterQueries,
					},
				},
			},
		},
	}
	return query
}

// GetArticles request crawler to scrapt data and sync new data with redis and elastic search
func (s *articleService) GetArticles() {
	client := pb.NewCrawlerServiceClient(s.conn)

	in := &pb.AllConfigsArticles{
		HtmlClasses: &pb.HTMLClasses{
			ArticleClass:     s.htmlClassesService.HtmlClasses.ArticleClass,
			ArticleTitleClass:       s.htmlClassesService.HtmlClasses.ArticleTitleClass,
			ArticleDescriptionClass: s.htmlClassesService.HtmlClasses.ArticleDescriptionClass,
			ArticleLinkClass:        s.htmlClassesService.HtmlClasses.ArticleLinkClass,
		},
		Leagues: s.leaguesService.leagues.Leagues,
	}
	// send gRPC request to crawler
	stream, err := client.GetArticles(context.Background(), in)
	if err != nil {
		log.Printf("error occurred while openning stream error %v \n", err)
		return
	}

	done := make(chan bool)
	var mapSearchResult = make(map[string]bool)
	log.Printf("Start get stream of article...\n")
	// recieve stream of article from crawler
	go func() {
		for {
			resp, err := stream.Recv()

			league := resp.GetLeague()

			respArticles := resp.GetArticles()
			tags := s.tagsService.Tags.Tags

			checkSimilarArticles(respArticles, s.es, league, tags)

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

	<-done
	PREV_ARTICLES = mapSearchResult
	log.Printf("finished.")
}

// Nếu gửi từng bài bào lên elastic check thì mỗi lần tìm sẽ gửi vài ngàn request
// C1: Server lưu kết quả cào ở lần trước đó, sau đó lấy kết quả mới so sánh với cũ, nếu có bài báo nào mới thì sẽ check lại với elasticsearch. Elasticsearch chưa có thì thêm vào

func checkSimilarArticles(respArticles []*pb.Article, es *elasticsearch.Client, league string, tags []string) {

	// Condition: similar title
	for _, article := range respArticles {

		// check if it a bet web
		if strings.Contains(strings.ToLower(article.Description), "cá cược") {
			log.Println("Detect a online bet web, skip it")
			continue
		}

		// check if it exist in previous results
		_, ok := PREV_ARTICLES[article.Title]
		if !ok {
			exist := checkWithElasticSearch(article, es)
			if !exist {
				entityArticle := newEntitiesArticleFromPb(article, tags, league)
				storeElasticsearch(entityArticle, es)
			}
			// checkWithRedis(article)
		}
	}
}

func checkWithElasticSearch(article *pb.Article, es *elasticsearch.Client) bool {
	req := esapi.ExistsRequest{
		Index:      "articles",
		DocumentID: strings.ToLower(article.Title),
	}

	resp, err := req.Do(context.Background(), es)
	if err != nil {
		log.Printf("Error checking if document exists: %s\n", err)
		return false
	}

	status := resp.StatusCode
	if status == 200 {
		log.Println("Document already exist in elastic search")
		return true
	} else if status == 404 {
		log.Println("Document not found, creating new one...")
		return false
	}

	return false
}

// func checkWithRedis(){}

func saveToMapSearchResult(respArticles []*pb.Article, mapSearchResult map[string]bool) {
	for _, article := range respArticles {
		mapSearchResult[article.Title] = true
	}
}

func newEntitiesArticleFromMap(respArticle map[string]interface{}) entities.Article {
	// type assertion []interface{} to []string
	tagInterface := respArticle["tags"].([]interface{})
	tags := make([]string, len(tagInterface))
	for i, tag := range tagInterface {
		tags[i] = tag.(string)
	}

	article := entities.Article{
		Title:       respArticle["title"].(string),
		Description: respArticle["description"].(string),
		Link:        respArticle["link"].(string),
		Tags:        tags,
	}
	return article
}

func newEntitiesArticleFromPb(respArticle *pb.Article, tags []string, league string) entities.Article {
	articleTags := checkTags(respArticle, tags, league)

	article := entities.Article{
		Title:       respArticle.Title,
		Description: respArticle.Description,
		Link:        respArticle.Link,
		Tags:        articleTags,
	}
	return article
}

func storeElasticsearch(article entities.Article, es *elasticsearch.Client) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	body, err := json.Marshal(article)
	if err != nil {
		log.Printf("Error encoding article: %s\n", err)
	}

	req := esapi.IndexRequest{
		Index:      "articles",
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
		log.Printf("[%s] Indexed document with index: %s \n", res.Status(), "articles")
	}
}

func checkTags(article *pb.Article, tags []string, keyword string) []string {
	articleTags := make(map[string]bool)
	articleTags[helper.FormatVietnamese(keyword)] = true

	for _, tag := range tags {
		formatedTag := helper.FormatVietnamese(tag)
		_, ok := articleTags[formatedTag]

		if !ok {
			if strings.Contains(helper.FormatVietnamese(article.Description), formatedTag) || strings.Contains(helper.FormatVietnamese(article.Title), formatedTag) {
				articleTags[tag] = true
			}
		}
	}

	articleTagsSlice := make([]string, 0)

	for tag := range articleTags {
		articleTagsSlice = append(articleTagsSlice, tag)
	}

	return articleTagsSlice
}
