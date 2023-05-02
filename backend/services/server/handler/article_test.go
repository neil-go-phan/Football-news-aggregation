package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"server/entities"
	mock "server/handler/mock"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAPISearchTagsAndKeywordFailToConvertString(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("tags", "tag1")
	q.Add("q", "keyword")
	q.Add("from", "invalid int")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mock.MockArticleServices)
	articleHandler := NewArticleHandler(mockArticleServices)
	assert := assert.New(t)

	articleHandler.APISearchTagsAndKeyword(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(`{"success":false,"message":"Bad request"}`, w.Body.String(), "Expected 'Bad request' message to be returned")
}

func TestAPISearchTagsAndKeywordServicesFail(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	keyword := "keyword"
	formatedTag := []string{"tag1"}
	from := 0
	q.Add("tags", "tag1")
	q.Add("q", keyword)
	q.Add("from", "0")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mock.MockArticleServices)
	articleHandler := NewArticleHandler(mockArticleServices)
	assert := assert.New(t)

	articles := []entities.Article{}
	mockArticleServices.On("SearchArticlesTagsAndKeyword", keyword, formatedTag, from).Return(articles, 10.0, fmt.Errorf("service failed"))
	articleHandler.APISearchTagsAndKeyword(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(`{"success":false,"message":"Server error"}`, w.Body.String(), "Expected 'Server error' message to be returned")
}

func TestAPISearchTagsAndKeywordServicesSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	keyword := "keyword"
	formatedTag := []string{"tag1"}
	from := 0
	q.Add("tags", "tag1")
	q.Add("q", keyword)
	q.Add("from", "0")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mock.MockArticleServices)
	articleHandler := NewArticleHandler(mockArticleServices)
	assert := assert.New(t)

	articles := []entities.Article{}
	mockArticleServices.On("SearchArticlesTagsAndKeyword", keyword, formatedTag, from).Return(articles, 10.0, nil)
	articleHandler.APISearchTagsAndKeyword(c)

	assert.Equal(http.StatusOK, w.Code, "Expected HTTP status code 200")
}

func TestAPICrawlArticleLeague(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	leagueName:= []string{"league 1"}
	q.Add("league", "league 1")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mock.MockArticleServices)
	articleHandler := NewArticleHandler(mockArticleServices)
	assert := assert.New(t)


	mockArticleServices.On("GetArticles", leagueName)
	articleHandler.APICrawlArticleLeague(c)

	assert.Equal(http.StatusOK, w.Code, "Expected HTTP status code 200")
	assert.JSONEq(`{"success":true,"message":"Signal crawl artilce success"}`, w.Body.String(), "Expected 'Signal crawl artilce success' message to be returned")
}

func TestAPIAddUpdateNewTagFail(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("tag", "tag1")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mock.MockArticleServices)
	articleHandler := NewArticleHandler(mockArticleServices)
	assert := assert.New(t)

	mockArticleServices.On("AddTagForAllArticle", "tag1").Return(fmt.Errorf("add tag service failed"))
	articleHandler.APIAddUpdateNewTag(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(`{"success":false,"message":"Update tag failed"}`, w.Body.String(), "Expected 'Update tag failed' message to be returned")
}

func TestAPIAddUpdateNewTagSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("tag", "tag1")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockArticleServices := new(mock.MockArticleServices)
	articleHandler := NewArticleHandler(mockArticleServices)
	assert := assert.New(t)

	mockArticleServices.On("AddTagForAllArticle", "tag1").Return(nil)
	articleHandler.APIAddUpdateNewTag(c)

	assert.Equal(http.StatusOK, w.Code, "Expected HTTP status code 200")
	assert.JSONEq(`{"success":true,"message":"Update tag successfull"}`, w.Body.String(), "Expected 'Update tag successfull' message to be returned")
}

func TestAPIGetArticleCountSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	mockArticleServices := new(mock.MockArticleServices)
	articleHandler := NewArticleHandler(mockArticleServices)
	assert := assert.New(t)

	mockArticleServices.On("GetArticleCount").Return(100.0, 100.0, nil)
	articleHandler.APIGetArticleCount(c)

	assert.Equal(http.StatusOK, w.Code, "Expected HTTP status code 200")
	assert.JSONEq(`{"success":true,"message":"Get article count success", "today":100, "total":100}`, w.Body.String(), "Expected 'Get article count success' message to be returned")
}

func TestAPIGetArticleCountFail(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	mockArticleServices := new(mock.MockArticleServices)
	articleHandler := NewArticleHandler(mockArticleServices)
	assert := assert.New(t)

	mockArticleServices.On("GetArticleCount").Return(100.0, 100.0, fmt.Errorf("GetArticleCount service fail"))
	articleHandler.APIGetArticleCount(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(`{"success":false,"message":"Server error"}`, w.Body.String(), "Expected 'Server error' message to be returned")
}
