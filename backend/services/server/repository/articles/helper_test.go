package articlerepo

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"server/entities"
	"strings"
	"sync"
	"testing"

	"github.com/elastic/go-elasticsearch/v7"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	pb "server/proto"
)

func TestFilterDeletedTag(t *testing.T) {
	article := &entities.Article{
		Title:       "Title",
		Description: "Description",
		Link:        "Link",
		Tags:        []string{"tag1", "tag2", "tag3"},
	}

	tagExist := []string{"tag1", "tag2", "tag5"}

	filterDeletedTag(article, tagExist)
	// check article.tag
	want := []string{"tag1", "tag2"}
	got := article.Tags
	if !reflect.DeepEqual(want, got) {
		t.Errorf("FilterDeletedTag  is supose to %#v, but got %#v", want, got)
	}
}

type checkTagExistOutput struct {
	isExist bool
	index   int
}

type checkTagExistTestCase struct {
	name     string
	tagInput string
	output   checkTagExistOutput
}

func assertCheckTestExist(t *testing.T, checkTagExistTestCase checkTagExistTestCase, article *entities.Article) {
	want := checkTagExistTestCase.output
	ok, index := checkTagExist(article, checkTagExistTestCase.tagInput)
	got := checkTagExistOutput{
		isExist: ok,
		index:   index,
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s with input = '%s' is supose to %#v, but got %#v", checkTagExistTestCase.name, checkTagExistTestCase.tagInput, want, got)
	}
}

func TestCheckTestExist(t *testing.T) {
	article := &entities.Article{
		Title:       "Title",
		Description: "Description",
		Link:        "Link",
		Tags:        []string{"tag1", "tag2", "tag3"},
	}

	var checkTagExistTestCases = []checkTagExistTestCase{
		{name: "normal case", tagInput: "tag1", output: checkTagExistOutput{isExist: true, index: 0}},
		{name: "none exist case", tagInput: "tag5", output: checkTagExistOutput{isExist: false, index: -1}},
	}

	for _, c := range checkTagExistTestCases {
		t.Run(c.name, func(t *testing.T) {
			assertCheckTestExist(t, c, article)
		})
	}
}

func TestQuerySearchAllArticles(t *testing.T) {
	input := 0
	want := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
		"from": 0,
		"size": 10,
		"sort": map[string]interface{}{
			"created_at": map[string]interface{}{
				"order": "desc",
			},
		},
	}
	got := querySearchAllArticles(input)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("function querySearchAllArticles with input = '%v' is supose to %#v, but got %#v", input, want, got)
	}
}

func TestQuerySearchArticlesOnlySearchKeyword(t *testing.T) {
	keyword := "test"
	from := 0
	want := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"multi_match": map[string]interface{}{
						"query":    "test",
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
		"from": 0,
		"size": 10,
		"sort": map[string]interface{}{
			"created_at": map[string]interface{}{
				"order": "desc",
			},
		},
	}
	got := querySearchArticlesOnlySearchKeyword(keyword, from)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("function querySearchArticlesOnlySearchKeyword with input = '%v and %v' is supose to %#v, but got %#v", keyword, from, want, got)
	}
}

func TestQuerySearchArticlesOnlyTag(t *testing.T) {
	filterTagQueryLine1 := map[string]interface{}{"match_phrase": map[string]interface{}{"tags": "tag1"}}
	filterTagQueryLine2 := map[string]interface{}{"match_phrase": map[string]interface{}{"tags": "tag2"}}
	filterTagQuery := []map[string]interface{}{filterTagQueryLine1, filterTagQueryLine2}
	from := 0
	want := map[string]interface{}{
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
	got := querySearchArticlesOnlyTag(filterTagQuery, from)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("function querySearchArticlesOnlyTag with input = '%v and %v' is supose to %#v, but got %#v", filterTagQuery, from, want, got)
	}
}

func TestQuerySearchArticlesBothTagAndKeyword(t *testing.T) {
	filterTagQueryLine1 := map[string]interface{}{"match_phrase": map[string]interface{}{"tags": "tag1"}}
	filterTagQueryLine2 := map[string]interface{}{"match_phrase": map[string]interface{}{"tags": "tag2"}}
	filterTagQuery := []map[string]interface{}{filterTagQueryLine1, filterTagQueryLine2}
	from := 0
	keyword := "test"
	want := map[string]interface{}{
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
	got := querySearchArticlesBothTagAndKeyword(keyword, filterTagQuery, from)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("function querySearchArticlesBothTagAndKeyword with input = '%v and %v and %v' is supose to %#v, but got %#v", keyword, filterTagQuery, from, want, got)
	}
}

func TestQueryGetAmountCrawledArtilceToday(t *testing.T) {
	want := map[string]interface{}{
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
	got := queryGetAmountCrawledArtilceToday()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("function queryGetAmountCrawledArtilceToday is supose to %#v, but got %#v", want, got)
	}
}

func TestQueryGetTotalCrawledArtilce(t *testing.T) {
	want := map[string]interface{}{
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
	got := queryGetTotalCrawledArtilce()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("function queryGetTotalCrawledArtilce is supose to %#v, but got %#v", want, got)
	}
}

type querySearchArticleInput struct {
	from         int
	keyword      string
	formatedTags []string
}

type querySearchArticleTestCase struct {
	name   string
	input  querySearchArticleInput
	output map[string]interface{}
}

func assertQuerySearchArticle(t *testing.T, testCase querySearchArticleTestCase) {
	want := testCase.output
	got := querySearchArticle(testCase.input.keyword, testCase.input.formatedTags, testCase.input.from)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s with input = '%#v' is supose to %#v, but got %#v", testCase.name, testCase.input, want, got)
	}
}
func TestQuerySearchArticle(t *testing.T) {
	querySearchArticleTestCases := []querySearchArticleTestCase{
		{name: "Only Search Keyword", input: querySearchArticleInput{formatedTags: []string{}, keyword: "test", from: 0}, output: map[string]interface{}{
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"must": map[string]interface{}{
						"multi_match": map[string]interface{}{
							"query":    "test",
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
			"from": 0,
			"size": 10,
			"sort": map[string]interface{}{
				"created_at": map[string]interface{}{
					"order": "desc",
				},
			},
		}},
		{name: "Only Tag", input: querySearchArticleInput{formatedTags: []string{"tag1", "tag2"}, keyword: "", from: 0}, output: map[string]interface{}{
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"filter": map[string]interface{}{
						"bool": map[string]interface{}{
							"must": []map[string]interface{}{{"match_phrase": map[string]interface{}{"tags": "tag1"}}, {"match_phrase": map[string]interface{}{"tags": "tag2"}}},
						},
					},
					"must_not": map[string]interface{}{
						"exists": map[string]interface{}{
							"field": "deleted",
						},
					},
				},
			},
			"from": 0,
			"size": 10,
			"sort": map[string]interface{}{
				"created_at": map[string]interface{}{
					"order": "desc",
				},
			},
		}},
		{name: "Both Tag And Keyword", input: querySearchArticleInput{formatedTags: []string{"tag1", "tag2"}, keyword: "test", from: 0}, output: map[string]interface{}{
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"must": map[string]interface{}{
						"multi_match": map[string]interface{}{
							"query":  "test",
							"fields": []string{"title", "description"},
						},
					},
					"filter": map[string]interface{}{
						"bool": map[string]interface{}{
							"must": []map[string]interface{}{{"match_phrase": map[string]interface{}{"tags": "tag1"}}, {"match_phrase": map[string]interface{}{"tags": "tag2"}}},
						},
					},
					"must_not": map[string]interface{}{
						"exists": map[string]interface{}{
							"field": "deleted",
						},
					},
				},
			},
			"from": 0,
			"size": 10,
			"sort": map[string]interface{}{
				"created_at": map[string]interface{}{
					"order": "desc",
				},
			}}},
	}
	for _, c := range querySearchArticleTestCases {
		t.Run(c.name, func(t *testing.T) {
			assertQuerySearchArticle(t, c)
		})
	}
}

type checkTagWhenAddTagTestCase struct {
	name     string
	inputTag string
	output   bool
}

func assertCheckTagWhenAddTag(t *testing.T, testCase checkTagWhenAddTagTestCase, article entities.Article) {
	want := testCase.output
	got := checkTagWhenAddTag(article, testCase.inputTag)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s with input = '%#v' is supose to %#v, but got %#v", testCase.name, testCase.inputTag, want, got)
	}
}

func TestCheckTagWhenAddTagExist(t *testing.T) {
	inputArticle := entities.Article{
		Title:       "Test article title",
		Description: "Test article description",
		Link:        "Test",
		Tags:        []string{"tag1", "tag2"},
	}
	checkTagWhenAddTagTestCases := []checkTagWhenAddTagTestCase{
		{name: "tag already exist", inputTag: "tag1", output: false},
		{name: "article contain tag", inputTag: "title", output: true},
		{name: "article not contain tag", inputTag: "not contain", output: false},
	}

	for _, c := range checkTagWhenAddTagTestCases {
		t.Run(c.name, func(t *testing.T) {
			assertCheckTagWhenAddTag(t, c, inputArticle)
		})
	}
}

func TestCreateBulkRequestBodyAddTag(t *testing.T) {
	assert := assert.New(t)
	articles := []entities.Article{
		{Title: "title 1", Description: "description 1", Link: "test.com", Tags: []string{"tag1", "tag2"}},
		{Title: "title 2", Description: "description 2", Link: "test.com", Tags: []string{"tag1", "tag3"}},
		{Title: "title 3", Description: "description 3", Link: "test.com", Tags: []string{"tag3", "tag2"}},
	}

	newTag := "title 1"

	bulkRequestBody, err := createBulkRequestBodyAddTag(articles, newTag)
	assert.Nil(err)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	expectedResult := []byte{}
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

			updateBytes, _ := json.Marshal(updateBody)
			updateHeaderBytes, _ := json.Marshal(updateHeader)
			expectedResult = append(expectedResult, updateHeaderBytes...)
			expectedResult = append(expectedResult, []byte("\n")...)
			expectedResult = append(expectedResult, updateBytes...)
			expectedResult = append(expectedResult, []byte("\n")...)
		}
	}

	assert.True(reflect.DeepEqual(bulkRequestBody, expectedResult))
}

func TestWorkerAddTagBulkRequest(t *testing.T) {
	jobs := make(chan *work, 2)
	results := make(chan *result, 2)
	newTag := "newTag"
	var wg sync.WaitGroup
	id := 1

	jobs <- &work{{Title: "title 1", Description: "description 1", Link: "test.com", Tags: []string{"tag1", "tag2"}}}
	jobs <- &work{{Title: "title 2", Description: "description 2", Link: "test.com", Tags: []string{"tag1", "tag3"}}}
	close(jobs)

	wg.Add(1)
	go workerAddTagBulkRequest(jobs, results, newTag, &wg, id)
	wg.Wait()
	close(results)

	var resultsArr []*result
	for result := range results {
		resultsArr = append(resultsArr, result)
	}

	if len(resultsArr) != 2 {
		t.Errorf("expected 2 results, but got %d", len(resultsArr))
	}
}

func TestRequestAddTagArticleSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"took":1,"errors":false,"items":[{"update":{"_index":"articles","_type":"_doc","_id":"1","_version":1,"result":"updated","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":0,"_primary_term":1,"status":200}}]}`)
	}))
	defer server.Close()
	assert := assert.New(t)

	cfg := elasticsearch.Config{
		Addresses: []string{server.URL},
	}
	es, err := elasticsearch.NewClient(cfg)
	assert.Nil(err)

	bulkRequestBody := []byte(`{"index":{"_index":"articles","_type":"_doc","_id":"1"}}`)

	requestAddTagArticle(bulkRequestBody, es)

	res, err := server.Client().Get(server.URL)
	assert.Nil(err)
	defer res.Body.Close()

	_, err = io.ReadAll(res.Body)
	assert.Nil(err)
}

func TestRequestAddTagArticleFail(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, `{"_index":"articles","_type":"_doc","_id":"123456","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":0,"_primary_term":1}`)
	}))
	defer server.Close()

	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{server.URL},
	})
	if err != nil {
		t.Fatalf("Error creating Elasticsearch client: %s", err)
	}

	bulkRequestBody := []byte(`{"index":{"_index":"articles","_type":"_doc","_id":"1"}}`)

	requestAddTagArticle(bulkRequestBody, es)
	assert := assert.New(t)

	res, err := server.Client().Get(server.URL)
	assert.Nil(err)
	defer res.Body.Close()

	_, err = io.ReadAll(res.Body)
	assert.Nil(err)
}

func TestAnalyzedResult(t *testing.T) {
	// mock elastic client
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"took":1,"errors":false,"items":[{"update":{"_index":"articles","_type":"_doc","_id":"1","_version":1,"result":"updated","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":0,"_primary_term":1,"status":200}}]}`)
	}))
	defer server.Close()
	assert := assert.New(t)

	cfg := elasticsearch.Config{
		Addresses: []string{server.URL},
	}
	es, err := elasticsearch.NewClient(cfg)
	assert.Nil(err)

	// Setup channels for results and sync.WaitGroup
	results := make(chan *result)
	var wg2 sync.WaitGroup
	wg2.Add(1)

	// Start goroutine to call analyzedResult with mocked params
	go func() {
		analyzedResult(results, es, &wg2)
	}()

	// Send mock results to channel
	var testResult result
	results <- &testResult
	results <- &testResult

	// Close results channel and wait for goroutine to finish
	close(results)
	wg2.Wait()

	if len(results) != 0 {
		t.Errorf("expected results chan is empty, but got %d", len(results))
	}
}

func TestRequestOpenPointInTime(t *testing.T) {
	// mock elastic client
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"id":"pitID"}`)
	}))
	defer server.Close()
	assert := assert.New(t)

	cfg := elasticsearch.Config{
		Addresses: []string{server.URL},
	}
	es, err := elasticsearch.NewClient(cfg)
	assert.Nil(err)

	got, err := requestOpenPointInTime(es)
	assert.Nil(err)
	assert.Equal("pitID", got)
}

func TestQuerySearchFirstPageArticlesWithTagAsKeyword(t *testing.T) {
	tag := "test"
	pitID := "test_pit"
	expectedQuery := map[string]interface{}{
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

	result := querySearchFirstPageArticlesWithTagAsKeyword(tag, pitID)

	if !reflect.DeepEqual(result, expectedQuery) {
		t.Errorf("querySearchFirstPageArticlesWithTagAsKeyword() returned wrong query, expected: %v, got: %v", expectedQuery, result)
	}
}

func TestQueryGetNextPageOfArticleWithTagAsKeyword(t *testing.T) {
	tag := "test tag"
	pitID := "test pit ID"
	searchAfter := []interface{}{"test search after value"}

	expectedQuery := map[string]interface{}{
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

	actualQuery := queryGetNextPageOfArticleWithTagAsKeyword(tag, pitID, searchAfter)

	if !reflect.DeepEqual(expectedQuery, actualQuery) {
		t.Errorf("Expected query: %v, but got: %v", expectedQuery, actualQuery)
	}
}

func TestCheckBetWebReturnTrue(t *testing.T) {
	mockArticle := &pb.Article{
		Title: "title 1", Description: "ca cuoc 1", Link: "test.com",
	}
	assert := assert.New(t)

	got := checkBetWeb(mockArticle)
	want := true
	assert.Equal(want, got)
}

func TestCheckBetWebReturnFail(t *testing.T) {
	mockArticle := &pb.Article{
		Title: "title 1", Description: "description 1", Link: "test.com",
	}
	assert := assert.New(t)

	got := checkBetWeb(mockArticle)
	want := false
	assert.Equal(want, got)
}

func TestCheckArtilceWithElasticSearchSuccess(t *testing.T) {
	// mock elastic client
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"id":"pitID"}`)
	}))
	defer server.Close()
	assert := assert.New(t)

	mockArticle := &pb.Article{
		Title: "title 1", Description: "description 1", Link: "test.com",
	}

	cfg := elasticsearch.Config{
		Addresses: []string{server.URL},
	}
	es, err := elasticsearch.NewClient(cfg)
	assert.Nil(err)

	got := checkArtilceWithElasticSearch(mockArticle, es)
	want := true
	assert.Equal(want, got)
}

func TestCheckArtilceWithElasticSearchFail(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, `{"id":"pitID"}`)
	}))
	defer server.Close()
	assert := assert.New(t)

	mockArticle := &pb.Article{
		Title: "title 1", Description: "description 1", Link: "test.com",
	}

	cfg := elasticsearch.Config{
		Addresses: []string{server.URL},
	}
	es, err := elasticsearch.NewClient(cfg)
	assert.Nil(err)

	got := checkArtilceWithElasticSearch(mockArticle, es)
	want := false
	assert.Equal(want, got)
}

func TestSaveToMapSearchResult(t *testing.T) {
	respArticles := []*pb.Article{
		{Title: "Article 1", Link: "https://example.com/article1"},
		{Title: "Article 2", Link: "https://example.com/article2"},
		{Title: "Article 3", Link: "https://example.com/article3"},
	}
	mapSearchResult := make(map[string]bool)

	saveToMapSearchResult(respArticles, mapSearchResult)

	assert.Equal(t, true, mapSearchResult["Article 1-https://example.com/article1"])
	assert.Equal(t, true, mapSearchResult["Article 2-https://example.com/article2"])
	assert.Equal(t, true, mapSearchResult["Article 3-https://example.com/article3"])
	assert.Equal(t, 3, len(mapSearchResult))

	saveToMapSearchResult(respArticles, mapSearchResult)
	assert.Equal(t, true, mapSearchResult["Article 1-https://example.com/article1"])
	assert.Equal(t, true, mapSearchResult["Article 2-https://example.com/article2"])
	assert.Equal(t, true, mapSearchResult["Article 3-https://example.com/article3"])
	assert.Equal(t, 3, len(mapSearchResult))
}

func TestNewEntitiesArticleFromMap(t *testing.T) {
	articleMap := map[string]interface{}{
		"title":       "Test article",
		"link":        "https://test.com",
		"description": "This is a test article",
		"created_at":  "2022-05-03T10:00:00.000Z",
	}

	article := newEntitiesArticleFromMap(articleMap)

	assert.Equal(t, "Test article", article.Title)
	assert.Equal(t, "https://test.com", article.Link)
	assert.Equal(t, "This is a test article", article.Description)
}

func TestTaggedWhenCrawl(t *testing.T) {
	article := &pb.Article{
		Title:       "Apple unveils new products at annual event",
		Description: "Apple has announced new products including a new MacBook Pro and AirPods",
		Link:        "https://example.com/apple-event",
	}

	tags := []string{"Apple", "MacBook Pro", "AirPods", "iPhone"}

	keyword := "apple"

	expectedTags := []string{"tin tuc bong da", "apple", "macbook pro", "airpods"}

	resultTags := taggedWhenCrawl(article, tags, keyword)

	for _, tag := range resultTags {
		if !contains(expectedTags, tag) {
			t.Errorf("Expected tags: %v, but got: %v", expectedTags, resultTags)
		}
	}

}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func TestNewEntitiesArticleFromPb(t *testing.T) {
	respArticle := &pb.Article{
		Title:       "Some title",
		Description: "Some description",
		Link:        "https://some-article.com",
	}

	tags := []string{"tag1", "tag2"}

	league := "Some league"

	article := newEntitiesArticleFromPb(respArticle, tags, league)

	if article.Title != respArticle.Title {
		t.Errorf("Title does not match, expected: %s, got: %s", respArticle.Title, article.Title)
	}

	if article.Description != respArticle.Description {
		t.Errorf("Description does not match, expected: %s, got: %s", respArticle.Description, article.Description)
	}

	if article.Link != respArticle.Link {
		t.Errorf("Link does not match, expected: %s, got: %s", respArticle.Link, article.Link)
	}

	if len(article.Tags) != 2 {
		t.Errorf("Expected 2 tags, got: %d", len(article.Tags))
	}

	if article.Tags[0] != DEFAULT_TAG {
		t.Errorf("First tag is not default tag")
	}

	if article.Tags[1] != "some league" {
		t.Errorf("Tag does not match, expected: %s, got: %s", "some league", article.Tags[1])
	}
}

func TestStoreArticleInElasticsearch(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"_index":"articles","_type":"_doc","_id":"123456","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":0,"_primary_term":1}`)
	}))
	defer server.Close()

	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{server.URL},
	})
	if err != nil {
		t.Fatalf("Error creating Elasticsearch client: %s", err)
	}

	article := entities.Article{
		Title:       "Sample Article",
		Description: "This is a sample article",
		Link:        "https://example.com/sample-article",
		Tags:        []string{"sample", "article"},
	}

	storeArticleInElasticsearch(article, es)
	// not knowing what to do
}

func TestCheckSimilarArticles(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, `{"id":"pitID"}`)
	}))
	defer server.Close()
	assert := assert.New(t)

	cfg := elasticsearch.Config{
		Addresses: []string{server.URL},
	}
	es, err := elasticsearch.NewClient(cfg)
	assert.Nil(err)

	// Test case 1: empty article list
	checkSimilarArticles([]*pb.Article{}, es, "league", []string{"tag1", "tag2"})
	// Expect nothing to happen

	// Test case 2: article list with one article, already in PREV_ARTICLES
	PREV_ARTICLES = map[string]bool{"Test article-https://test.com": true}
	article := &pb.Article{
		Title:       "Test article",
		Description: "This is a test article",
		Link:        "https://test.com",
	}
	checkSimilarArticles([]*pb.Article{article}, es, "league", []string{"tag1", "tag2"})
	// Expect nothing to happen

	// Test case 3: bet article
	PREV_ARTICLES = map[string]bool{}
	article = &pb.Article{
		Title:       "ca cuoc article",
		Description: "This is a test article",
		Link:        "https://test.com",
	}

	checkSimilarArticles([]*pb.Article{article}, es, "league", []string{"tag1", "tag2"})

	// Test case 4: success

	article = &pb.Article{
		Title:       "Test article",
		Description: "This is a test article",
		Link:        "https://test.com",
	}

	PREV_ARTICLES = map[string]bool{}
	checkSimilarArticles([]*pb.Article{article}, es, "league", []string{"tag1", "tag2"})

}
