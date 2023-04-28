package articlesservices

import (
	"server/entities"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	mock "server/services/articles/mock"
)

type deleteArticleTestCase struct {
	name   string
	input  string
	output error
}

func assertDeleteArticle(t *testing.T, testCase deleteArticleTestCase, mockRepo *mock.MockArticleRepository, service *articleService) {
	assert := assert.New(t)
	want := testCase.output

	got := service.DeleteArticle(testCase.input)

	if !assert.EqualError(got, want.Error()) {
		t.Errorf("%s with input = '%s' is supose to %v, but got %v", testCase.name, testCase.input, want.Error(), got.Error())

	}
}

func TestDeleteArticle(t *testing.T) {
	mockRepo := new(mock.MockArticleRepository)
	service := NewArticleService(mockRepo)
	assert := assert.New(t)

	mockRepo.On("DeleteArticle", "success").Return(nil)
	mockRepo.On("DeleteArticle", "cant send request").Return(errors.New("error getting response for delete request"))
	mockRepo.On("DeleteArticle", "404").Return(errors.New("404"))

	deleteArticleTestCases := []deleteArticleTestCase{
		{name: "delete success", input: "cant send request", output: errors.New("error getting response for delete request")},
		{name: "delete success", input: "404", output: errors.New("404")},
	}

	for _, c := range deleteArticleTestCases {
		t.Run(c.name, func(t *testing.T) {
			assertDeleteArticle(t, c, mockRepo, service)
		})
	}

	// nil case
	got := service.DeleteArticle("success")
	if !assert.Nil(got) {
		t.Errorf("Success case with input = 'success' is supose to nil, but got %v", got.Error())
	}
}

type getArticleCountOutput struct {
	total float64
	today float64
	err   error
}

func TestGetArticleCount(t *testing.T) {
	mockRepo := new(mock.MockArticleRepository)
	service := NewArticleService(mockRepo)
	assert := assert.New(t)

	want := getArticleCountOutput{
		total: 20,
		today: 10,
		err:   nil,
	}
	mockRepo.On("GetArticleCount").Return(want.total, want.today, want.err)

	total, today, err := service.GetArticleCount()
	got := getArticleCountOutput{
		total: total,
		today: today,
		err:   err,
	}
	if !assert.Nil(got.err) {
		t.Errorf("Method GetArticleCount is supose to %#v, but got %#v", want, got)
	}
}

func TestGetArticles(t *testing.T) {
	// This method dont return anything. there is no point to test it
}

func TestAddTagForAllArticle(t *testing.T) {
	mockRepo := new(mock.MockArticleRepository)
	service := NewArticleService(mockRepo)
	assert := assert.New(t)

	mockRepo.On("AddTagForAllArticle", "success").Return(nil)
	mockRepo.On("AddTagForAllArticle", "fail").Return(errors.New("add tag fail"))

	t.Run("success case", func(t *testing.T) {
		got := service.AddTagForAllArticle("success")
		assert.Nil(got)
	})

	t.Run("fail case", func(t *testing.T) {
		got := service.AddTagForAllArticle("fail")
		assert.Errorf(got, "add tag fail")
	})
}

func TestGetFirstPageOfLeagueRelatedArticle(t *testing.T) {
	mockRepo := new(mock.MockArticleRepository)
	service := NewArticleService(mockRepo)
	assert := assert.New(t)

	successCaseArticles := []entities.Article{
		{Title: "title 1", Description: "description 1", Link: "link 1"},
		{Title: "title 2", Description: "description 2", Link: "link 2"},
		{Title: "title 3", Description: "description 3", Link: "link 3"},
	}
	mockRepo.On("GetFirstPageOfLeagueRelatedArticle", "success").Return(successCaseArticles, nil)
	mockRepo.On("GetFirstPageOfLeagueRelatedArticle", "fail").Return([]entities.Article{}, errors.New("cant get articles"))

	t.Run("success case", func(t *testing.T) {
		articles, err := service.GetFirstPageOfLeagueRelatedArticle("success")
		assert.Nil(err)
		assert.Equal(successCaseArticles, articles)
	})

	t.Run("fail case", func(t *testing.T) {
		_, err := service.GetFirstPageOfLeagueRelatedArticle("fail")
		assert.Errorf(err, "cant get articles")
	})
}

func TestSearchArticlesTagsAndKeyword(t *testing.T) {
	mockRepo := new(mock.MockArticleRepository)
	service := NewArticleService(mockRepo)
	assert := assert.New(t)

	successCaseArticles := []entities.Article{
		{Title: "title 1", Description: "description 1", Link: "link 1"},
		{Title: "title 2", Description: "description 2", Link: "link 2"},
		{Title: "title 3", Description: "description 3", Link: "link 3"},
	}

	keyword := "test"
	formatedTags := []string{}
	from:= 0

	mockRepo.On("SearchArticlesTagsAndKeyword", keyword, formatedTags, from).Return(successCaseArticles, 100.5, nil)

	articles, total, err := service.SearchArticlesTagsAndKeyword(keyword, formatedTags, from)

	assert.Nil(err)
	assert.Equal(articles, successCaseArticles)
	assert.Equal(total, 100.5)

}
