package articlesservices

import (
	"bytes"
	"context"
	"fmt"
	"server/entities"
	serverhelper "server/helper"
	"strconv"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"

	pb "server/proto"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v7"
	jsoniter "github.com/json-iterator/go"
)

func newEntitiesArticle(respArticle *pb.Article, tags []entities.Tag) entities.Article {
	article := entities.Article{
		Title:       respArticle.Title,
		Description: respArticle.Description,
		Link:        respArticle.Link,
		Tags:        tags,
	}
	return article
}

func storeArticleToElasticsearch(article entities.Article, es *elasticsearch.Client) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	tagNames := []string{}
	for _, tag := range article.Tags {
		tagNames = append(tagNames, tag.TagName)
	}

	doc := entities.ElasticArticle{
		ID:          article.ID,
		Title:       article.Title,
		Description: article.Description,
		Tags:        tagNames,
		CreatedAt:   article.CreatedAt,
	}

	body, err := json.Marshal(doc)
	if err != nil {
		log.Errorf("Error encoding article: %s", err)
	}
	docID := strconv.FormatUint(uint64(article.ID), 10)
	req := esapi.IndexRequest{
		Index:      ARTICLES_INDEX_NAME,
		DocumentID: docID,
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

// return a slice of tag_name which article is tagged
func taggedWhenCrawl(article *pb.Article, tags []string, league string) []string {
	articleTags := make(map[string]bool)
	articleTags[serverhelper.FormatVietnamese(league)] = true

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

// return false if article is a bet web or exist in previous crawl results
func validateArticles(article *pb.Article) bool {
	if checkBetWeb(article) {
		log.Println("Detect a online bet web, skip it")
		return false
	}

	mapKey := fmt.Sprintf("%s-%s", article.Title, article.Link)
	_, ok := PREV_ARTICLES[mapKey]

	return !ok
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

// save articles crawled in this turn
func saveToMapSearchResult(respArticles []*pb.Article, mapSearchResult map[string]bool) {
	for _, article := range respArticles {
		mapKey := fmt.Sprintf("%s-%s", article.Title, article.Link)
		mapSearchResult[mapKey] = true
	}
}

func deleteArticleFromElasticSearch(id uint, es *elasticsearch.Client) error {
	docID := strconv.FormatUint(uint64(id), 10)
	req := esapi.DeleteRequest{
		Index:      ARTICLES_INDEX_NAME,
		DocumentID: docID,
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Errorf("Error getting response for delete request: %s", err)
		return fmt.Errorf("error getting response for delete request")
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Errorf("[%s] Error delete article document with ID='%s'", res.Status(), docID)
		return fmt.Errorf(res.Status())
	} else {
		log.Printf("[%s] Deleted document with id: %s", res.Status(), docID)
	}
	return nil
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

func newElasticArticleFromMap(respArticle map[string]interface{}) entities.ElasticArticle {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	article := entities.ElasticArticle{}

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

func requestFirstPageOfArticleSearchWithTagAsKeyword(tag string, pitID string, es *elasticsearch.Client) (articles []entities.ElasticArticle, newPitID string, searchAfter []interface{}, err error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	var buffer bytes.Buffer

	query := querySearchFirstPageArticlesWithTagAsKeyword(tag, pitID)

	err = json.NewEncoder(&buffer).Encode(query)
	if err != nil {
		return articles, pitID, searchAfter, err
	}

	resp, err := es.Search(es.Search.WithBody(&buffer))
	if err != nil {
		return articles, pitID, searchAfter, err
	}

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return articles, pitID, searchAfter, err
	}

	for index, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		article := hit.(map[string]interface{})["_source"].(map[string]interface{})
		articles = append(articles, newElasticArticleFromMap(article))
		// get last result article sort field
		// elastic search will sent 10 hits by default. if we can not find the 10th article, that mean there is no more result
		if index == 9 {
			searchAfter = hit.(map[string]interface{})["sort"].([]interface{})
		}
	}

	newPitID = result["pit_id"].(string)
	return articles, newPitID, searchAfter, nil
}

func requestNextPageOfArticleWithTagAsKeyword(tag string, pitID string, searchAfterQuery []interface{}, es *elasticsearch.Client) (articles []entities.ElasticArticle, newPitID string, searchAfter []interface{}, err error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	var buffer bytes.Buffer

	query := queryGetNextPageOfArticleWithTagAsKeyword(tag, pitID, searchAfterQuery)

	err = json.NewEncoder(&buffer).Encode(query)
	if err != nil {
		return articles, pitID, searchAfter, fmt.Errorf("encode query failed")
	}

	resp, err := es.Search(es.Search.WithBody(&buffer))
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
		articles = append(articles, newElasticArticleFromMap(article))
		// elastic search will sent 10 hits by default. if we can not find the 10th article, that mean there is no more result
		if index == 9 {
			searchAfter = hit.(map[string]interface{})["sort"].([]interface{})
		}
	}

	newPitID = result["pit_id"].(string)
	return articles, newPitID, searchAfter, nil
}

func requestOpenPointInTime(es *elasticsearch.Client) (string, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	var pointInTime string

	resp, err := es.OpenPointInTime([]string{ARTICLES_INDEX_NAME}, PIT_LIVE)
	if err != nil {
		return pointInTime, err
	}

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return pointInTime, err
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

// producer sending job continuously to job chan until no more job
func producerRequestArticle(jobs chan<- *work, es *elasticsearch.Client, tag string, firstPitID string, firstSearchAfterQuery []interface{}) {
	pitID := firstPitID
	searchAfterQuery := firstSearchAfterQuery

	for {
		articles, newPitID, newSearchAfterQuery, err := requestNextPageOfArticleWithTagAsKeyword(tag, pitID, searchAfterQuery, es)
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
		bulkRequestBody, ids, err := createBulkRequestBodyAddTag(*job, newTag)
		if err != nil {
			log.Errorf("error when worker create bulk request body %s", err)
		}
		// sometime, there is no article match the new tag
		if len(bulkRequestBody) != 0 {
			result := result{
				ids:             ids,
				bulkRequestBody: bulkRequestBody,
			}
			results <- &result
		}

	}
}

// analyzer collect result done by worker
func analyzedResult(results <-chan *result, s *articleService , wg2 *sync.WaitGroup, entityTag *entities.Tag) {
	defer wg2.Done()
	bulkRequestBody := []byte{}
	ids := []uint{}
	for result := range results {
		body := result.bulkRequestBody
		bulkRequestBody = append(bulkRequestBody, body...)
	}
	requestUpdateTagsArticle(bulkRequestBody, s.es)
	err := s.repo.AddTag(ids, entityTag)
	if err != nil {
		log.Error(err)
	}
	log.Printf("analyzer sent request ")
}

func createBulkRequestBodyAddTag(articles []entities.ElasticArticle, newTag string) ([]byte, []uint, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	bulkRequestBody := []byte{}
	ids := []uint{}
	for _, article := range articles {
		ok := checkTagWhenAddTag(article, newTag)
		if ok {
			newTagsField := append(article.Tags, newTag)
			ids = append(ids, article.ID)
			updateBody := map[string]interface{}{
				"doc": map[string]interface{}{
					"tags": newTagsField,
				},
			}
			docID := strconv.FormatUint(uint64(article.ID), 10)
			updateHeader := map[string]interface{}{
				"update": map[string]interface{}{
					"_index": ARTICLES_INDEX_NAME,
					"_id":    docID,
				},
			}

			updateBytes, err := json.Marshal(updateBody)
			if err != nil {
				return bulkRequestBody, ids, err
			}
			updateHeaderBytes, err := json.Marshal(updateHeader)
			if err != nil {
				return bulkRequestBody, ids, err
			}
			bulkRequestBody = append(bulkRequestBody, updateHeaderBytes...)
			bulkRequestBody = append(bulkRequestBody, []byte("\n")...)
			bulkRequestBody = append(bulkRequestBody, updateBytes...)
			bulkRequestBody = append(bulkRequestBody, []byte("\n")...)
		}
	}
	return bulkRequestBody, ids, nil
}

func requestUpdateTagsArticle(bulkRequestBody []byte, es *elasticsearch.Client) {
	req := esapi.BulkRequest{
		Body:    bytes.NewReader(bulkRequestBody),
		Refresh: "true",
	}
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Errorf("can not send a bulk update request to elastic search %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Errorf("Error response: %s", res.String())
	} else {
		log.Printf("Bulk update successful")
	}
}

func checkTagWhenAddTag(article entities.ElasticArticle, newTag string) bool {
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
