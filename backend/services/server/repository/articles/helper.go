package articlerepo

import (
	"bytes"
	"context"
	"fmt"
	"server/entities"
	serverhelper "server/helper"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	pb "server/proto"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v7"
	jsoniter "github.com/json-iterator/go"
)

func filterDeletedTag(article *entities.Article, tags []string) {
	var wg sync.WaitGroup
	for _, tag := range tags {
		wg.Add(1)
		go func(tag string) {
			articleTag := make([]string, 0)
			for _, tag := range tags {
				isExist, index := checkTagExist(article, tag)
				if isExist {
					articleTag = append(articleTag, article.Tags[index])
				}
			}
			article.Tags = articleTag
			wg.Done()
		}(tag)
	}
	wg.Wait()
}

func checkTagExist(article *entities.Article, tag string) (bool, int) {
	for index, articleTag := range article.Tags {
		if articleTag == tag {
			return true, index
		}
	}
	return false, -1
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
				"must_not": map[string]interface{}{
					"exists": map[string]interface{}{
						"field": "deleted",
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
				"must_not": map[string]interface{}{
					"exists": map[string]interface{}{
						"field": "deleted",
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
				"must_not": map[string]interface{}{
					"exists": map[string]interface{}{
						"field": "deleted",
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

func queryGetAmountCrawledArtilceToday() map[string]interface{} {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"range": map[string]interface{}{
				"created_at": map[string]interface{}{
					"gte": "now-1d/d",
					"lte": "now/d",
				},
			},
		},
		"size": 0,
	}
	return query
}

func queryGetTotalCrawledArtilce() map[string]interface{} {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"range": map[string]interface{}{
				"created_at": map[string]interface{}{
					"gte": "now-30d/d",
					"lte": "now/d",
				},
			},
		},
		"size": 0,
	}
	return query
}

// producer sending job continuously to job chan until no more job
func producerRequestArticle(jobs chan<- *work, s *articleRepo, tag string, firstPitID string, firstSearchAfterQuery []interface{}) {
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
func workerAddTagBulkRequest(jobs <-chan *work, results chan<- *result, newTag string, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for job := range jobs {
		log.Printf("worker #%d received: %v article", id, len(*job))
		bulkRequestBody, err := createBulkRequestBodyAddTag(*job, newTag)
		if err != nil {
			log.Errorf("error when worker create bulk request body %s", err)
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
	requestAddTagArticle(bulkRequestBody, es)
	log.Printf("analyzer send request ")
}

func createBulkRequestBodyAddTag(articles []entities.Article, newTag string) ([]byte, error) {
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

func requestAddTagArticle(bulkRequestBody []byte, es *elasticsearch.Client) {
	req := esapi.BulkRequest{
		Body:    bytes.NewReader(bulkRequestBody),
		Refresh: "true",
	}
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Errorf("can not send a bulk update request to elastic search %s", err)
		return
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Errorf("Error response: %s", res.String())
	} else {
		log.Printf("Bulk update successful")
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

// Nếu gửi từng bài bào lên elastic check thì mỗi lần tìm sẽ gửi vài ngàn request
// Solution: Server lưu kết quả cào ở lần trước đó, sau đó lấy kết quả mới so sánh với cũ, nếu có bài báo nào mới thì sẽ check lại với elasticsearch. Elasticsearch chưa có thì thêm vào

func checkSimilarArticles(respArticles []*pb.Article, es *elasticsearch.Client, league string, tags []string) {
	// Condition: similar title
	for _, article := range respArticles {
		
		// check if it a bet web
		if checkBetWeb(article) {
			log.Println("Detect a online bet web, skip it")
			continue
		}
		mapKey := fmt.Sprintf("%s-%s", article.Title ,article.Link)
		// check if it exist in previous results
		_, ok := PREV_ARTICLES[mapKey]

		if !ok {
			exist := checkArtilceWithElasticSearch(article, es)
			if !exist {
				entityArticle := newEntitiesArticleFromPb(article, tags, league)
				
				if entityArticle.Title != "" && entityArticle.Link != "" && entityArticle.Description != "" {
					storeArticleInElasticsearch(entityArticle, es)
				}
			}
		}
	}
}

func checkBetWeb(article *pb.Article) bool {
	betKeywords := []string{"ca cuoc", "nha cai", "keo bong da", "danh bac", "song bac", "casino", "xo so"}
	for _, keyword := range betKeywords {
		articleInfo := serverhelper.FormatVietnamese(article.Description) + serverhelper.FormatVietnamese(article.Title)
		if strings.Contains(articleInfo, keyword) {
			return true
		}
	}
	return false
}

func checkArtilceWithElasticSearch(article *pb.Article, es *elasticsearch.Client) bool {
	req := esapi.ExistsRequest{
		Index:      ARTICLES_INDEX_NAME,
		DocumentID: strings.ToLower(article.Title),
	}

	resp, err := req.Do(context.Background(), es)
	if err != nil {
		log.Errorf("Error checking if document exists: %s", err)
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
		mapKey := fmt.Sprintf("%s-%s", article.Title ,article.Link)
		mapSearchResult[mapKey] = true
	}
}

func newEntitiesArticleFromMap(respArticle map[string]interface{}) entities.Article {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	article := entities.Article{}

	articleByte, err := json.Marshal(respArticle)
	if err != nil {
		log.Errorf("error occrus when marshal elastic response article: %s", err)
	}

	err = json.Unmarshal(articleByte, &article)
	if err != nil {
		log.Errorf("error occrus when unmarshal elastic response to entity article: %s", err)
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
		log.Errorf("Error encoding article: %s", err)
	}

	req := esapi.IndexRequest{
		Index:      ARTICLES_INDEX_NAME,
		DocumentID: strings.ToLower(article.Title),
		Body:       strings.NewReader(string(body)),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Errorf("Error getting response: %s", err)
		return
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Errorf("[%s] Error indexing article document with id='%s'", res.Status(), strings.ToLower(article.Title))
	} else {
		log.Printf("[%s] Indexed document with index: %s", res.Status(), ARTICLES_INDEX_NAME)
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
				articleTags[formatedTag] = true
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
