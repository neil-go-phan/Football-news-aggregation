package articlerepo

import (
	"reflect"
	"server/entities"
	"testing"
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
	name   string
	inputTag  string
	output bool
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
		Title: "Test article title",
		Description: "Test article description",
		Link: "Test",
		Tags: []string{"tag1", "tag2"},
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