package services

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"server/entities"
	serverhelper "server/helper"
	"strings"
	"sync"
	"time"

	pb "server/proto"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v7"
	jsoniter "github.com/json-iterator/go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

var PREV_ARTICLES = make(map[string]bool)
var ARTICLES_INDEX_NAME = "articles"
var DEFAULT_TAG = "tin tuc bong da"
var PIT_LIVE = "1m"

type work = []entities.Article
type result = []byte

type articleService struct {
	conn               *grpc.ClientConn
	es                 *elasticsearch.Client
	htmlClassesService *htmlClassesService
	leaguesService     *leaguesService
	tagsService        *tagsService
}

func NewArticleService(leagues *leaguesService, htmlClass *htmlClassesService, tags *tagsService, conn *grpc.ClientConn, es *elasticsearch.Client) *articleService {
	articleService := &articleService{
		conn:               conn,
		es:                 es,
		htmlClassesService: htmlClass,
		leaguesService:     leagues,
		tagsService:        tags,
	}
	return articleService
}

func (s *articleService) SearchArticlesTagsAndKeyword(keyword string, formatedTags []string, from int) ([]entities.Article, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	articles := make([]entities.Article, 0)
	var buffer bytes.Buffer

	query := querySearchArticle(keyword, formatedTags, from)

	err := json.NewEncoder(&buffer).Encode(query)
	if err != nil {
		return articles, fmt.Errorf("encode query failed")
	}

	resp, err := s.es.Search(s.es.Search.WithIndex(ARTICLES_INDEX_NAME), s.es.Search.WithBody(&buffer))
	if err != nil {
		return articles, fmt.Errorf("request to elastic search fail")
	}

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return articles, fmt.Errorf("decode respose from elastic search failed")
	}

	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {

		article := hit.(map[string]interface{})["_source"].(map[string]interface{})
		articles = append(articles, newEntitiesArticleFromMap(article))
	}
	return articles, nil
}

func querySearchArticle(keyword string, formatedTags []string, from int) map[string]interface{} {

	var filterTagQuery []map[string]interface{}
	for _, tag := range formatedTags {
		if tag != "" {
			tagQuery := map[string]interface{}{"match_phrase": map[string]interface{}{"tags": tag}}
			filterTagQuery = append(filterTagQuery, tagQuery)
		}
	}
	// default is search all
	query := querySearchAllArticles(from)

	if len(filterTagQuery) == 0 && keyword != "" {
		// search with only keyword
		query = querySearchArticlesOnlySearchKeyword(keyword, from)
	}

	if len(filterTagQuery) != 0 && keyword == "" {
		// search with only tags
		query = querySearchArticlesOnlyTag(filterTagQuery, from)
	}

	if len(filterTagQuery) != 0 && keyword != "" {
		// search with both tags and keyword
		query = querySearchArticlesBothTagAndKeyword(keyword, filterTagQuery, from)
	}
	return query
}

func querySearchAllArticles(from int) map[string]interface{} {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
		"from": from,
		"size": 10,
		"sort": map[string]interface{}{
			"created_at": map[string]interface{}{
				"order": "desc",
			},
		},
	}
	return query
}

func querySearchArticlesOnlySearchKeyword(keyword string, from int) map[string]interface{} {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"multi_match": map[string]interface{}{
						"query":    keyword,
						"fields":   []string{"title", "description"},
						"analyzer": "no_accent",
					},
				},
			},
		},
		"from": from,
		"size": 10,
		"sort": map[string]interface{}{
			"created_at": map[string]interface{}{
				"order": "desc",
			},
		},
	}
	return query
}

func querySearchArticlesOnlyTag(filterTagQuery []map[string]interface{}, from int) map[string]interface{} {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": map[string]interface{}{
					"bool": map[string]interface{}{
						"must": filterTagQuery,
					},
				},
			},
		},
		"from": from,
		"size": 10,
		"sort": map[string]interface{}{
			"created_at": map[string]interface{}{
				"order": "desc",
			},
		},
	}
	return query
}

func querySearchArticlesBothTagAndKeyword(keyword string, filterTagQuery []map[string]interface{}, from int) map[string]interface{} {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"multi_match": map[string]interface{}{
						"query":  keyword,
						"fields": []string{"title", "description"},
					},
				},
				"filter": map[string]interface{}{
					"bool": map[string]interface{}{
						"must": filterTagQuery,
					},
				},
			},
		},
		"from": from,
		"size": 10,
		"sort": map[string]interface{}{
			"created_at": map[string]interface{}{
				"order": "desc",
			},
		},
	}
	return query
}

// refresh cache after expire

// implement search_after query, // implement worker pool

func (s *articleService) AddTagForAllArticle(tag string) error {
	tagFormated := serverhelper.FormatVietnamese(tag)
	pitID, err := requestOpenPointInTime(s.es)
	if err != nil {
		return err
	}
	firstPageArticles, newPitID, searchAfterQuery, err := s.requestFirstPageOfArticleSearchWithTagAsKeyword(tagFormated, pitID)
	if err != nil {
		return err
	}
	bulkRequestBody, err := createBulkRequestBodyUpdateTag(firstPageArticles, tagFormated)
	if err != nil {
		log.Printf("error when worker create bulk request body %s\n", err)
	}
	requestUpdateTagArticle(bulkRequestBody, s.es)

	if len(searchAfterQuery) == 0 {
		return nil
	}

	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup

	// worker pool
	jobs := make(chan *work, 5)
	results := make(chan *result, 5)
	// 5 worker
	for i := 0; i < 5; i++ {
		wg1.Add(1)
		go workerUpdateBulkRequest(jobs, results, tagFormated, &wg1, i)
	}

	go producerRequestArticle(jobs, s, tagFormated, newPitID, searchAfterQuery)

	wg2.Add(1)
	go analyzedResult(results,s.es, &wg2)

	wg1.Wait() 

	close(results)

	wg2.Wait()
	log.Printf("add tag success")
	return nil
}

// producer sending job continuously to job chan until no more job
func producerRequestArticle(jobs chan<- *work, s *articleService, tag string, firstPitID string, firstSearchAfterQuery []interface{}) {
	pitID := firstPitID
	searchAfterQuery := firstSearchAfterQuery

	for {
		articles, newPitID, newSearchAfterQuery, err := s.requestNextPageOfArticleWithTagAsKeyword(tag, pitID, searchAfterQuery)
		if err != nil {
			break
		}

		pitID = newPitID
		searchAfterQuery = newSearchAfterQuery

		jobs <- &articles

		if len(searchAfterQuery) == 0 {
			break
		}
	}
	close(jobs)
}

// worker take job, do something and send result to result chan
func workerUpdateBulkRequest(jobs <-chan *work, results chan<- *result, newTag string, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for job := range jobs {
		log.Printf("worker #%d received: %v article\n", id, len(*job))
		bulkRequestBody, err := createBulkRequestBodyUpdateTag(*job, newTag)
		if err != nil {
			log.Printf("error when worker create bulk request body %s\n", err)
		}
		result := bulkRequestBody
		results <- &result
	}
}

// analyzer collect result done by worker
func analyzedResult(results <-chan *result, es *elasticsearch.Client, wg2 *sync.WaitGroup) {
	defer wg2.Done()
	bulkRequestBody := []byte{}
	for body := range results {
		bulkRequestBody = append(bulkRequestBody, *body...)
	}
	requestUpdateTagArticle(bulkRequestBody, es)
	log.Printf("analyzer send request \n")
}

func createBulkRequestBodyUpdateTag(articles []entities.Article, newTag string) ([]byte, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	bulkRequestBody := []byte{}
	for _, article := range articles {
		ok := checkTagWhenAddTag(article, newTag)
		if ok {
			newTagsField := article.Tags
			newTagsField = append(newTagsField, newTag)

			updateBody := map[string]interface{}{
				"doc": map[string]interface{}{
					"tags": newTagsField,
				},
			}

			updateHeader := map[string]interface{}{
				"update": map[string]interface{}{
					"_index": ARTICLES_INDEX_NAME,
					"_id":    strings.ToLower(article.Title),
				},
			}

			updateBytes, err := json.Marshal(updateBody)
			if err != nil {
				return bulkRequestBody, err
			}
			updateHeaderBytes, err := json.Marshal(updateHeader)
			if err != nil {
				return bulkRequestBody, err
			}
			bulkRequestBody = append(bulkRequestBody, updateHeaderBytes...)
			bulkRequestBody = append(bulkRequestBody, []byte("\n")...)
			bulkRequestBody = append(bulkRequestBody, updateBytes...)
			bulkRequestBody = append(bulkRequestBody, []byte("\n")...)
		}
	}
	return bulkRequestBody, nil
}

func requestUpdateTagArticle(bulkRequestBody []byte, es *elasticsearch.Client) {
	req := esapi.BulkRequest{
		Body:    bytes.NewReader(bulkRequestBody),
		Refresh: "true",
	}
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Printf("can not send a bulk update request to elastic search %s\n", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("Error response: %s\n", res.String())
	} else {
		log.Printf("Bulk update successful\n")
	}
}

func checkTagWhenAddTag(article entities.Article, newTag string) bool {
	// check if tag already exist on article
	for _, tag := range article.Tags {
		if tag == newTag {
			return false
		}
	}
	if strings.Contains(serverhelper.FormatVietnamese(article.Description), newTag) || strings.Contains(serverhelper.FormatVietnamese(article.Title), newTag) {
		return true
	}
	return false
}

func requestOpenPointInTime(es *elasticsearch.Client) (string, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	var pointInTime string

	resp, err := es.OpenPointInTime([]string{ARTICLES_INDEX_NAME}, PIT_LIVE)
	if err != nil {
		return pointInTime, fmt.Errorf("request to elastic search fail")
	}

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return pointInTime, fmt.Errorf("decode respose from elastic search failed")
	}

	pointInTime = result["id"].(string)

	return pointInTime, nil
}

func (s *articleService) requestFirstPageOfArticleSearchWithTagAsKeyword(tag string, pitID string) (articles []entities.Article, newPitID string, searchAfter []interface{}, err error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	var buffer bytes.Buffer

	query := querySearchFirstPageArticlesWithTagAsKeyword(tag, pitID)

	err = json.NewEncoder(&buffer).Encode(query)
	if err != nil {
		return articles, pitID, searchAfter, fmt.Errorf("encode query failed")
	}

	resp, err := s.es.Search(s.es.Search.WithBody(&buffer))
	if err != nil {
		return articles, pitID, searchAfter, fmt.Errorf("request to elastic search fail")
	}

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return articles, pitID, searchAfter, fmt.Errorf("decode respose from elastic search failed")
	}

	for index, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		article := hit.(map[string]interface{})["_source"].(map[string]interface{})
		articles = append(articles, newEntitiesArticleFromMap(article))
		// get last result article sort field
		// elastic search will sent 10 hits by default. if we can not find the 10th article, that mean there is no more result
		if index == 9 {
			searchAfter = hit.(map[string]interface{})["sort"].([]interface{})
		}
	}

	newPitID = result["pit_id"].(string)
	return articles, newPitID, searchAfter, nil
}

func (s *articleService) requestNextPageOfArticleWithTagAsKeyword(tag string, pitID string, searchAfterQuery []interface{}) (articles []entities.Article, newPitID string, searchAfter []interface{}, err error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	var buffer bytes.Buffer

	query := queryGetNextPageOfArticleWithTagAsKeyword(tag, pitID, searchAfterQuery)

	err = json.NewEncoder(&buffer).Encode(query)
	if err != nil {
		return articles, pitID, searchAfter, fmt.Errorf("encode query failed")
	}

	resp, err := s.es.Search(s.es.Search.WithBody(&buffer))
	if err != nil {
		return articles, pitID, searchAfter, fmt.Errorf("request to elastic search fail")
	}

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return articles, pitID, searchAfter, fmt.Errorf("decode respose from elastic search failed")
	}

	for index, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		article := hit.(map[string]interface{})["_source"].(map[string]interface{})
		articles = append(articles, newEntitiesArticleFromMap(article))
		// get last result article sort field
		// elastic search will sent 10 hits by default. if we can not find the 10th article, that mean there is no more result
		if index == 9 {
			searchAfter = hit.(map[string]interface{})["sort"].([]interface{})
		}
	}

	newPitID = result["pit_id"].(string)
	return articles, newPitID, searchAfter, nil
}

func querySearchFirstPageArticlesWithTagAsKeyword(tag string, pitID string) map[string]interface{} {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"multi_match": map[string]interface{}{
						"query":    tag,
						"fields":   []string{"title", "description"},
						"analyzer": "no_accent",
					},
				},
			},
		},
		"pit": map[string]interface{}{
			"id":         pitID,
			"keep_alive": PIT_LIVE,
		},
		"sort": map[string]interface{}{
			"created_at": map[string]interface{}{
				"order": "desc",
			},
		},
	}
	return query
}

func queryGetNextPageOfArticleWithTagAsKeyword(tag string, pitID string, searchAfter []interface{}) map[string]interface{} {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"multi_match": map[string]interface{}{
						"query":    tag,
						"fields":   []string{"title", "description"},
						"analyzer": "no_accent",
					},
				},
			},
		},
		"pit": map[string]interface{}{
			"id":         pitID,
			"keep_alive": PIT_LIVE,
		},
		"search_after": searchAfter,
		"sort": map[string]interface{}{
			"created_at": map[string]interface{}{
				"order": "desc",
			},
		},
	}
	return query
}

// GetArticles request crawler to scrapt data and sync elastic search
func (s *articleService) GetArticles(keywords []string) {
	client := pb.NewCrawlerServiceClient(s.conn)

	in := &pb.AllConfigsArticles{
		HtmlClasses: &pb.HTMLClasses{
			ArticleClass:            s.htmlClassesService.HtmlClasses.ArticleClass,
			ArticleTitleClass:       s.htmlClassesService.HtmlClasses.ArticleTitleClass,
			ArticleDescriptionClass: s.htmlClassesService.HtmlClasses.ArticleDescriptionClass,
			ArticleLinkClass:        s.htmlClassesService.HtmlClasses.ArticleLinkClass,
		},
		Leagues: s.leaguesService.GetLeaguesNameActive(),
	}

	if len(keywords) != 0 {
		in.Leagues = keywords
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
			if err == io.EOF {
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Printf("cannot receive %v\n", err)
				status, _ := status.FromError(err)
				if status.Code().String() == "Unavailable" {
					done <- true //means stream is finished
					return
				}
			}

			league := resp.GetLeague()

			respArticles := resp.GetArticles()
			tags := s.tagsService.Tags.Tags

			checkSimilarArticles(respArticles, s.es, league, tags)

			saveToMapSearchResult(respArticles, mapSearchResult)

		}
	}()

	<-done
	PREV_ARTICLES = mapSearchResult
	log.Printf("finished.")
}

// Nếu gửi từng bài bào lên elastic check thì mỗi lần tìm sẽ gửi vài ngàn request
// Solution: Server lưu kết quả cào ở lần trước đó, sau đó lấy kết quả mới so sánh với cũ, nếu có bài báo nào mới thì sẽ check lại với elasticsearch. Elasticsearch chưa có thì thêm vào

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
			exist := checkArtilceWithElasticSearch(article, es)
			if !exist {
				entityArticle := newEntitiesArticleFromPb(article, tags, league)
				storeArticleInElasticsearch(entityArticle, es)
			}
		}
	}
}

func checkArtilceWithElasticSearch(article *pb.Article, es *elasticsearch.Client) bool {
	req := esapi.ExistsRequest{
		Index:      ARTICLES_INDEX_NAME,
		DocumentID: strings.ToLower(article.Title),
	}

	resp, err := req.Do(context.Background(), es)
	if err != nil {
		log.Printf("Error checking if document exists: %s\n", err)
		return false
	}

	status := resp.StatusCode
	if status == 200 {
		log.Println("Document already exist in index", ARTICLES_INDEX_NAME)
		return true
	} else if status == 404 {
		log.Printf("Document not found in index %s, creating new one...", ARTICLES_INDEX_NAME)
		return false
	}

	return false
}

func saveToMapSearchResult(respArticles []*pb.Article, mapSearchResult map[string]bool) {
	for _, article := range respArticles {
		mapSearchResult[article.Title] = true
	}
}

func newEntitiesArticleFromMap(respArticle map[string]interface{}) entities.Article {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	article := entities.Article{}

	articleByte, err := json.Marshal(respArticle)
	if err != nil {
		log.Printf("error occrus when marshal elastic response article: %s", err)
	}

	err = json.Unmarshal(articleByte, &article)
	if err != nil {
		log.Printf("error occrus when unmarshal elastic response to entity article: %s", err)
	}
	return article
}

func newEntitiesArticleFromPb(respArticle *pb.Article, tags []string, league string) entities.Article {
	articleTags := taggedWhenCrawl(respArticle, tags, league)

	article := entities.Article{
		Title:       respArticle.Title,
		Description: respArticle.Description,
		Link:        respArticle.Link,
		Tags:        articleTags,
	}
	return article
}

func storeArticleInElasticsearch(article entities.Article, es *elasticsearch.Client) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	doc := entities.ElasticArticle{
		Article:   article,
		CreatedAt: time.Now(),
	}

	body, err := json.Marshal(doc)
	if err != nil {
		log.Printf("Error encoding article: %s\n", err)
	}

	req := esapi.IndexRequest{
		Index:      ARTICLES_INDEX_NAME,
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
		log.Printf("[%s] Indexed document with index: %s \n", res.Status(), ARTICLES_INDEX_NAME)
	}
}

func taggedWhenCrawl(article *pb.Article, tags []string, keyword string) []string {
	articleTags := make(map[string]bool)
	articleTags[serverhelper.FormatVietnamese(keyword)] = true

	for _, tag := range tags {
		formatedTag := serverhelper.FormatVietnamese(tag)
		_, ok := articleTags[formatedTag]

		if !ok {
			if strings.Contains(serverhelper.FormatVietnamese(article.Description), formatedTag) || strings.Contains(serverhelper.FormatVietnamese(article.Title), formatedTag) {
				articleTags[tag] = true
			}
		}
	}

	articleTagsSlice := make([]string, 0)

	_, ok := articleTags[DEFAULT_TAG]
	if !ok {
		articleTagsSlice = append(articleTagsSlice, DEFAULT_TAG)
	}

	for tag := range articleTags {
		articleTagsSlice = append(articleTagsSlice, tag)
	}

	return articleTagsSlice
}
