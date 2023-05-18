package articlesservices

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// 	"reflect"
// 	"server/entities"
// 	pb "server/proto"
// 	mock "server/services/articles/mock"
// 	"strings"
// 	"sync"
// 	"testing"

// 	"github.com/elastic/go-elasticsearch/v7"
// 	jsoniter "github.com/json-iterator/go"
// 	"github.com/stretchr/testify/assert"
// )

// func TestValidateArticle_Fail(t *testing.T) {
// 	mockArticle := &pb.Article{
// 		Title: "title 1", Description: "ca cuoc 1", Link: "test.com",
// 	}
// 	assert := assert.New(t)
// 	got := validateArticles(mockArticle)
// 	want := false
// 	assert.Equal(want, got)
// }

// func TestValidateArticle_True(t *testing.T) {
// 	mockArticle := &pb.Article{
// 		Title: "title 1", Description: "descrition 1", Link: "test.com",
// 	}
// 	assert := assert.New(t)
// 	got := validateArticles(mockArticle)
// 	want := true
// 	assert.Equal(want, got)
// }

// func TestQuerySearchAllArticles(t *testing.T) {
// 	input := 0
// 	want := map[string]interface{}{
// 		"query": map[string]interface{}{
// 			"match_all": map[string]interface{}{},
// 		},
// 		"from": 0,
// 		"size": 10,
// 		"sort": map[string]interface{}{
// 			"created_at": map[string]interface{}{
// 				"order": "desc",
// 			},
// 		},
// 	}
// 	got := querySearchAllArticles(input)
// 	if !reflect.DeepEqual(want, got) {
// 		t.Errorf("function querySearchAllArticles with input = '%v' is supose to %#v, but got %#v", input, want, got)
// 	}
// }

// func TestQuerySearchArticlesOnlySearchKeyword(t *testing.T) {
// 	keyword := "test"
// 	from := 0
// 	want := map[string]interface{}{
// 		"query": map[string]interface{}{
// 			"bool": map[string]interface{}{
// 				"must": map[string]interface{}{
// 					"multi_match": map[string]interface{}{
// 						"query":    "test",
// 						"fields":   []string{"title", "description"},
// 						"analyzer": "no_accent",
// 					},
// 				},
// 				"must_not": map[string]interface{}{
// 					"exists": map[string]interface{}{
// 						"field": "deleted",
// 					},
// 				},
// 			},
// 		},
// 		"from": 0,
// 		"size": 10,
// 		"sort": map[string]interface{}{
// 			"created_at": map[string]interface{}{
// 				"order": "desc",
// 			},
// 		},
// 	}
// 	got := querySearchArticlesOnlySearchKeyword(keyword, from)
// 	if !reflect.DeepEqual(want, got) {
// 		t.Errorf("function querySearchArticlesOnlySearchKeyword with input = '%v and %v' is supose to %#v, but got %#v", keyword, from, want, got)
// 	}
// }

// func TestQuerySearchArticlesOnlyTag(t *testing.T) {
// 	filterTagQueryLine1 := map[string]interface{}{"match_phrase": map[string]interface{}{"tags": "tag1"}}
// 	filterTagQueryLine2 := map[string]interface{}{"match_phrase": map[string]interface{}{"tags": "tag2"}}
// 	filterTagQuery := []map[string]interface{}{filterTagQueryLine1, filterTagQueryLine2}
// 	from := 0
// 	want := map[string]interface{}{
// 		"query": map[string]interface{}{
// 			"bool": map[string]interface{}{
// 				"filter": map[string]interface{}{
// 					"bool": map[string]interface{}{
// 						"must": filterTagQuery,
// 					},
// 				},
// 				"must_not": map[string]interface{}{
// 					"exists": map[string]interface{}{
// 						"field": "deleted",
// 					},
// 				},
// 			},
// 		},
// 		"from": from,
// 		"size": 10,
// 		"sort": map[string]interface{}{
// 			"created_at": map[string]interface{}{
// 				"order": "desc",
// 			},
// 		},
// 	}
// 	got := querySearchArticlesOnlyTag(filterTagQuery, from)
// 	if !reflect.DeepEqual(want, got) {
// 		t.Errorf("function querySearchArticlesOnlyTag with input = '%v and %v' is supose to %#v, but got %#v", filterTagQuery, from, want, got)
// 	}
// }

// func TestQuerySearchArticlesBothTagAndKeyword(t *testing.T) {
// 	filterTagQueryLine1 := map[string]interface{}{"match_phrase": map[string]interface{}{"tags": "tag1"}}
// 	filterTagQueryLine2 := map[string]interface{}{"match_phrase": map[string]interface{}{"tags": "tag2"}}
// 	filterTagQuery := []map[string]interface{}{filterTagQueryLine1, filterTagQueryLine2}
// 	from := 0
// 	keyword := "test"
// 	want := map[string]interface{}{
// 		"query": map[string]interface{}{
// 			"bool": map[string]interface{}{
// 				"must": map[string]interface{}{
// 					"multi_match": map[string]interface{}{
// 						"query":  keyword,
// 						"fields": []string{"title", "description"},
// 					},
// 				},
// 				"filter": map[string]interface{}{
// 					"bool": map[string]interface{}{
// 						"must": filterTagQuery,
// 					},
// 				},
// 				"must_not": map[string]interface{}{
// 					"exists": map[string]interface{}{
// 						"field": "deleted",
// 					},
// 				},
// 			},
// 		},
// 		"from": from,
// 		"size": 10,
// 		"sort": map[string]interface{}{
// 			"created_at": map[string]interface{}{
// 				"order": "desc",
// 			},
// 		},
// 	}
// 	got := querySearchArticlesBothTagAndKeyword(keyword, filterTagQuery, from)
// 	if !reflect.DeepEqual(want, got) {
// 		t.Errorf("function querySearchArticlesBothTagAndKeyword with input = '%v and %v and %v' is supose to %#v, but got %#v", keyword, filterTagQuery, from, want, got)
// 	}
// }

// type querySearchArticleInput struct {
// 	from         int
// 	keyword      string
// 	formatedTags []string
// }

// type querySearchArticleTestCase struct {
// 	name   string
// 	input  querySearchArticleInput
// 	output map[string]interface{}
// }

// func assertQuerySearchArticle(t *testing.T, testCase querySearchArticleTestCase) {
// 	want := testCase.output
// 	got := querySearchArticle(testCase.input.keyword, testCase.input.formatedTags, testCase.input.from)
// 	if !reflect.DeepEqual(want, got) {
// 		t.Errorf("%s with input = '%#v' is supose to %#v, but got %#v", testCase.name, testCase.input, want, got)
// 	}
// }
// func TestQuerySearchArticle(t *testing.T) {
// 	querySearchArticleTestCases := []querySearchArticleTestCase{
// 		{name: "Only Search Keyword", input: querySearchArticleInput{formatedTags: []string{}, keyword: "test", from: 0}, output: map[string]interface{}{
// 			"query": map[string]interface{}{
// 				"bool": map[string]interface{}{
// 					"must": map[string]interface{}{
// 						"multi_match": map[string]interface{}{
// 							"query":    "test",
// 							"fields":   []string{"title", "description"},
// 							"analyzer": "no_accent",
// 						},
// 					},
// 					"must_not": map[string]interface{}{
// 						"exists": map[string]interface{}{
// 							"field": "deleted",
// 						},
// 					},
// 				},
// 			},
// 			"from": 0,
// 			"size": 10,
// 			"sort": map[string]interface{}{
// 				"created_at": map[string]interface{}{
// 					"order": "desc",
// 				},
// 			},
// 		}},
// 		{name: "Only Tag", input: querySearchArticleInput{formatedTags: []string{"tag1", "tag2"}, keyword: "", from: 0}, output: map[string]interface{}{
// 			"query": map[string]interface{}{
// 				"bool": map[string]interface{}{
// 					"filter": map[string]interface{}{
// 						"bool": map[string]interface{}{
// 							"must": []map[string]interface{}{{"match_phrase": map[string]interface{}{"tags": "tag1"}}, {"match_phrase": map[string]interface{}{"tags": "tag2"}}},
// 						},
// 					},
// 					"must_not": map[string]interface{}{
// 						"exists": map[string]interface{}{
// 							"field": "deleted",
// 						},
// 					},
// 				},
// 			},
// 			"from": 0,
// 			"size": 10,
// 			"sort": map[string]interface{}{
// 				"created_at": map[string]interface{}{
// 					"order": "desc",
// 				},
// 			},
// 		}},
// 		{name: "Both Tag And Keyword", input: querySearchArticleInput{formatedTags: []string{"tag1", "tag2"}, keyword: "test", from: 0}, output: map[string]interface{}{
// 			"query": map[string]interface{}{
// 				"bool": map[string]interface{}{
// 					"must": map[string]interface{}{
// 						"multi_match": map[string]interface{}{
// 							"query":  "test",
// 							"fields": []string{"title", "description"},
// 						},
// 					},
// 					"filter": map[string]interface{}{
// 						"bool": map[string]interface{}{
// 							"must": []map[string]interface{}{{"match_phrase": map[string]interface{}{"tags": "tag1"}}, {"match_phrase": map[string]interface{}{"tags": "tag2"}}},
// 						},
// 					},
// 					"must_not": map[string]interface{}{
// 						"exists": map[string]interface{}{
// 							"field": "deleted",
// 						},
// 					},
// 				},
// 			},
// 			"from": 0,
// 			"size": 10,
// 			"sort": map[string]interface{}{
// 				"created_at": map[string]interface{}{
// 					"order": "desc",
// 				},
// 			}}},
// 	}
// 	for _, c := range querySearchArticleTestCases {
// 		t.Run(c.name, func(t *testing.T) {
// 			assertQuerySearchArticle(t, c)
// 		})
// 	}
// }

// type checkTagWhenAddTagTestCase struct {
// 	name     string
// 	inputTag string
// 	output   bool
// }

// func assertCheckTagWhenAddTag(t *testing.T, testCase checkTagWhenAddTagTestCase, article entities.ElasticArticle) {
// 	want := testCase.output
// 	got := checkTagWhenAddTag(article, testCase.inputTag)
// 	if !reflect.DeepEqual(want, got) {
// 		t.Errorf("%s with input = '%#v' is supose to %#v, but got %#v", testCase.name, testCase.inputTag, want, got)
// 	}
// }

// func TestCheckTagWhenAddTagExist(t *testing.T) {
// 	inputArticle := entities.ElasticArticle{
// 		Title:       "Test article title",
// 		Description: "Test article description",
// 		Tags:        []string{"tag1", "tag2"},
// 	}
// 	checkTagWhenAddTagTestCases := []checkTagWhenAddTagTestCase{
// 		{name: "tag already exist", inputTag: "tag1", output: false},
// 		{name: "article contain tag", inputTag: "title", output: true},
// 		{name: "article not contain tag", inputTag: "not contain", output: false},
// 	}

// 	for _, c := range checkTagWhenAddTagTestCases {
// 		t.Run(c.name, func(t *testing.T) {
// 			assertCheckTagWhenAddTag(t, c, inputArticle)
// 		})
// 	}
// }

// func TestCreateBulkRequestBodyAddTag(t *testing.T) {
// 	assert := assert.New(t)
// 	articles := []entities.ElasticArticle{
// 		{Title: "title 1", Description: "description 1", Tags: []string{"tag1", "tag2"}},
// 		{Title: "title 2", Description: "description 2", Tags: []string{"tag1", "tag3"}},
// 		{Title: "title 3", Description: "description 3", Tags: []string{"tag3", "tag2"}},
// 	}

// 	newTag := "title 1123"

// 	bulkRequestBody, _, err := createBulkRequestBodyAddTag(articles, newTag)
// 	assert.Nil(err)
// 	var json = jsoniter.ConfigCompatibleWithStandardLibrary
// 	expectedResult := []byte{}
// 	for _, article := range articles {
// 		ok := checkTagWhenAddTag(article, newTag)
// 		if ok {
// 			newTagsField := article.Tags
// 			newTagsField = append(newTagsField, newTag)

// 			updateBody := map[string]interface{}{
// 				"doc": map[string]interface{}{
// 					"tags": newTagsField,
// 				},
// 			}

// 			updateHeader := map[string]interface{}{
// 				"update": map[string]interface{}{
// 					"_index": ARTICLES_INDEX_NAME,
// 					"_id":    strings.ToLower(article.Title),
// 				},
// 			}

// 			updateBytes, _ := json.Marshal(updateBody)
// 			updateHeaderBytes, _ := json.Marshal(updateHeader)
// 			expectedResult = append(expectedResult, updateHeaderBytes...)
// 			expectedResult = append(expectedResult, []byte("\n")...)
// 			expectedResult = append(expectedResult, updateBytes...)
// 			expectedResult = append(expectedResult, []byte("\n")...)
// 		}
// 	}

// 	assert.True(reflect.DeepEqual(bulkRequestBody, expectedResult))
// }

// func TestWorkerAddTagBulkRequest(t *testing.T) {
// 	jobs := make(chan *work, 2)
// 	results := make(chan *result, 2)
// 	newTag := "title"
// 	var wg sync.WaitGroup
// 	id := 1

// 	jobs <- &work{{Title: "title 1", Description: "description 1", Tags: []string{"tag1", "tag2"}}}
// 	jobs <- &work{{Title: "title 2", Description: "description 2", Tags: []string{"tag1", "tag3"}}}
// 	close(jobs)

// 	wg.Add(1)
// 	go workerAddTagBulkRequest(jobs, results, newTag, &wg, id)
// 	wg.Wait()
// 	close(results)

// 	var resultsArr []*result
// 	for result := range results {
// 		resultsArr = append(resultsArr, result)
// 	}

// 	if len(resultsArr) != 2 {
// 		t.Errorf("expected 2 results, but got %d", len(resultsArr))
// 	}
// }

// func TestRequestAddTagArticleSuccess(t *testing.T) {
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		fmt.Fprintln(w, `{"took":1,"errors":false,"items":[{"update":{"_index":"articles","_type":"_doc","_id":"1","_version":1,"result":"updated","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":0,"_primary_term":1,"status":200}}]}`)
// 	}))
// 	defer server.Close()
// 	assert := assert.New(t)

// 	cfg := elasticsearch.Config{
// 		Addresses: []string{server.URL},
// 	}
// 	es, err := elasticsearch.NewClient(cfg)
// 	assert.Nil(err)

// 	bulkRequestBody := []byte(`{"index":{"_index":"articles","_type":"_doc","_id":"1"}}`)

// 	requestUpdateTagsArticle(bulkRequestBody, es)

// 	res, err := server.Client().Get(server.URL)
// 	assert.Nil(err)
// 	defer res.Body.Close()

// 	_, err = io.ReadAll(res.Body)
// 	assert.Nil(err)
// }

// func TestDeleteArticleFromElasticSearch(t *testing.T) {
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		fmt.Fprintln(w, `{"took":1,"errors":false,"items":[{"update":{"_index":"articles","_type":"_doc","_id":"1","_version":1,"result":"updated","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":0,"_primary_term":1,"status":200}}]}`)
// 	}))
// 	defer server.Close()
// 	assert := assert.New(t)

// 	cfg := elasticsearch.Config{
// 		Addresses: []string{server.URL},
// 	}
// 	es, err := elasticsearch.NewClient(cfg)
// 	assert.Nil(err)

// 	id := uint(1)

// 	err = deleteArticleFromElasticSearch(id, es)
// 	assert.Nil(err)
// 	res, err := server.Client().Get(server.URL)
// 	assert.Nil(err)
// 	defer res.Body.Close()

// 	_, err = io.ReadAll(res.Body)
// 	assert.Nil(err)
// }

// func TestAnalyzedResult(t *testing.T) {
// 	// mock elastic client
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		fmt.Fprintln(w, `{"took":1,"errors":false,"items":[{"update":{"_index":"articles","_type":"_doc","_id":"1","_version":1,"result":"updated","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":0,"_primary_term":1,"status":200}}]}`)
// 	}))
// 	defer server.Close()
// 	assert := assert.New(t)

// 	cfg := elasticsearch.Config{
// 		Addresses: []string{server.URL},
// 	}
// 	es, err := elasticsearch.NewClient(cfg)
// 	assert.Nil(err)

// 	mockRepo := new(mock.MockArticleRepository)
// 	mockLeague := new(mock.MockLeaguesServices)
// 	mockTag := new(mock.MockTagsServices)
// 	mockCrawler := new(mock.MockCrawlerServiceClient)

// 	mockArticle :=  NewArticleService(mockLeague, mockTag, mockCrawler, es, mockRepo)
// 	tag := &entities.Tag{
// 		TagName: "title 1",
// 	}
// 	mockRepo.On("AddTag", []uint{}, tag).Return(nil)
// 	// Setup channels for results and sync.WaitGroup
// 	results := make(chan *result)
// 	var wg2 sync.WaitGroup
// 	wg2.Add(1)

// 	// Start goroutine to call analyzedResult with mocked params
// 	go func() {
// 		analyzedResult(results,mockArticle,&wg2, tag)
// 	}()

// 	// Send mock results to channel
// 	var testResult result
// 	results <- &testResult
// 	results <- &testResult

// 	// Close results channel and wait for goroutine to finish
// 	close(results)
// 	wg2.Wait()

// 	if len(results) != 0 {
// 		t.Errorf("expected results chan is empty, but got %d", len(results))
// 	}
// }

// func TestRequestOpenPointInTime(t *testing.T) {
// 	// mock elastic client
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		fmt.Fprintln(w, `{"id":"pitID"}`)
// 	}))
// 	defer server.Close()
// 	assert := assert.New(t)

// 	cfg := elasticsearch.Config{
// 		Addresses: []string{server.URL},
// 	}
// 	es, err := elasticsearch.NewClient(cfg)
// 	assert.Nil(err)

// 	got, err := requestOpenPointInTime(es)
// 	assert.Nil(err)
// 	assert.Equal("pitID", got)
// }

// func TestQuerySearchFirstPageArticlesWithTagAsKeyword(t *testing.T) {
// 	tag := "test"
// 	pitID := "test_pit"
// 	expectedQuery := map[string]interface{}{
// 		"query": map[string]interface{}{
// 			"bool": map[string]interface{}{
// 				"must": map[string]interface{}{
// 					"multi_match": map[string]interface{}{
// 						"query":    tag,
// 						"fields":   []string{"title", "description"},
// 						"analyzer": "no_accent",
// 					},
// 				},
// 			},
// 		},
// 		"pit": map[string]interface{}{
// 			"id":         pitID,
// 			"keep_alive": PIT_LIVE,
// 		},
// 		"sort": map[string]interface{}{
// 			"created_at": map[string]interface{}{
// 				"order": "desc",
// 			},
// 		},
// 	}

// 	result := querySearchFirstPageArticlesWithTagAsKeyword(tag, pitID)

// 	if !reflect.DeepEqual(result, expectedQuery) {
// 		t.Errorf("querySearchFirstPageArticlesWithTagAsKeyword() returned wrong query, expected: %v, got: %v", expectedQuery, result)
// 	}
// }

// func TestQueryGetNextPageOfArticleWithTagAsKeyword(t *testing.T) {
// 	tag := "test tag"
// 	pitID := "test pit ID"
// 	searchAfter := []interface{}{"test search after value"}

// 	expectedQuery := map[string]interface{}{
// 		"query": map[string]interface{}{
// 			"bool": map[string]interface{}{
// 				"must": map[string]interface{}{
// 					"multi_match": map[string]interface{}{
// 						"query":    tag,
// 						"fields":   []string{"title", "description"},
// 						"analyzer": "no_accent",
// 					},
// 				},
// 			},
// 		},
// 		"pit": map[string]interface{}{
// 			"id":         pitID,
// 			"keep_alive": PIT_LIVE,
// 		},
// 		"search_after": searchAfter,
// 		"sort": map[string]interface{}{
// 			"created_at": map[string]interface{}{
// 				"order": "desc",
// 			},
// 		},
// 	}

// 	actualQuery := queryGetNextPageOfArticleWithTagAsKeyword(tag, pitID, searchAfter)

// 	if !reflect.DeepEqual(expectedQuery, actualQuery) {
// 		t.Errorf("Expected query: %v, but got: %v", expectedQuery, actualQuery)
// 	}
// }

// func TestCheckBetWebReturnTrue(t *testing.T) {
// 	mockArticle := &pb.Article{
// 		Title: "title 1", Description: "ca cuoc 1", Link: "test.com",
// 	}
// 	assert := assert.New(t)

// 	got := checkBetWeb(mockArticle)
// 	want := true
// 	assert.Equal(want, got)
// }

// func TestCheckBetWebReturnFail(t *testing.T) {
// 	mockArticle := &pb.Article{
// 		Title: "title 1", Description: "description 1", Link: "test.com",
// 	}
// 	assert := assert.New(t)

// 	got := checkBetWeb(mockArticle)
// 	want := false
// 	assert.Equal(want, got)
// }

// func TestSaveToMapSearchResult(t *testing.T) {
// 	respArticles := []*pb.Article{
// 		{Title: "Article 1", Link: "https://example.com/article1"},
// 		{Title: "Article 2", Link: "https://example.com/article2"},
// 		{Title: "Article 3", Link: "https://example.com/article3"},
// 	}
// 	mapSearchResult := make(map[string]bool)

// 	saveToMapSearchResult(respArticles, mapSearchResult)

// 	assert.Equal(t, true, mapSearchResult["Article 1-https://example.com/article1"])
// 	assert.Equal(t, true, mapSearchResult["Article 2-https://example.com/article2"])
// 	assert.Equal(t, true, mapSearchResult["Article 3-https://example.com/article3"])
// 	assert.Equal(t, 3, len(mapSearchResult))

// 	saveToMapSearchResult(respArticles, mapSearchResult)
// 	assert.Equal(t, true, mapSearchResult["Article 1-https://example.com/article1"])
// 	assert.Equal(t, true, mapSearchResult["Article 2-https://example.com/article2"])
// 	assert.Equal(t, true, mapSearchResult["Article 3-https://example.com/article3"])
// 	assert.Equal(t, 3, len(mapSearchResult))
// }

// // func TestNewEntitiesArticleFromMap(t *testing.T) {
// // 	articleMap := map[string]interface{}{
// // 		"title":       "Test article",
// // 		"link":        "https://test.com",
// // 		"description": "This is a test article",
// // 		"created_at":  "2022-05-03T10:00:00.000Z",
// // 	}

// // 	article := newEntitiesArticleFromMap(articleMap)

// // 	assert.Equal(t, "Test article", article.Title)
// // 	assert.Equal(t, "https://test.com", article.Link)
// // 	assert.Equal(t, "This is a test article", article.Description)
// // }

// func TestTaggedWhenCrawl(t *testing.T) {
// 	article := &pb.Article{
// 		Title:       "Apple unveils new products at annual event",
// 		Description: "Apple has announced new products including a new MacBook Pro and AirPods",
// 		Link:        "https://example.com/apple-event",
// 	}

// 	tags := []string{"Apple", "MacBook Pro", "AirPods", "iPhone"}

// 	keyword := "apple"

// 	expectedTags := []string{"tin tuc bong da", "apple", "macbook pro", "airpods"}

// 	resultTags := taggedWhenCrawl(article, tags, keyword)

// 	for _, tag := range resultTags {
// 		if !contains(expectedTags, tag) {
// 			t.Errorf("Expected tags: %v, but got: %v", expectedTags, resultTags)
// 		}
// 	}

// }

// func contains(s []string, str string) bool {
// 	for _, v := range s {
// 		if v == str {
// 			return true
// 		}
// 	}

// 	return false
// }

// func TestNewElasticArticleFromMap(t *testing.T) {
// 	articleMap := map[string]interface{}{
// 		"title":       "Test article",
// 		"link":        "https://test.com",
// 		"description": "This is a test article",
// 		"created_at":  "2022-05-03T10:00:00.000Z",
// 	}

// 	article := newElasticArticleFromMap(articleMap)

// 	assert.Equal(t, "Test article", article.Title)
// 	assert.Equal(t, "This is a test article", article.Description)
// }

// func TestStoreArticleInElasticsearch(t *testing.T) {
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		fmt.Fprintln(w, `{"_index":"articles","_type":"_doc","_id":"123456","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":0,"_primary_term":1}`)
// 	}))
// 	defer server.Close()

// 	es, err := elasticsearch.NewClient(elasticsearch.Config{
// 		Addresses: []string{server.URL},
// 	})
// 	if err != nil {
// 		t.Fatalf("Error creating Elasticsearch client: %s", err)
// 	}

// 	article := entities.Article{
// 		Title:       "Sample Article",
// 		Description: "This is a sample article",
// 		Link:        "https://example.com/sample-article",
// 	}

// 	storeArticleToElasticsearch(article, es)
// 	// not knowing what to do
// }
