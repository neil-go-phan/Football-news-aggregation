package handler

// import (
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"net/url"
// 	"testing"

// 	"server/entities"
// 	mock "server/handler/mocks"
// 	"server/services"
// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )

// func TestAPISearchTagsAndKeywordFailToConvertString(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	req := &http.Request{
// 		URL:    &url.URL{},
// 		Header: make(http.Header),
// 	}
// 	q := req.URL.Query()
// 	q.Add("tags", "tag1")
// 	q.Add("q", "keyword")
// 	q.Add("from", "invalid int")

// 	c.Request = req
// 	c.Request.URL.RawQuery = q.Encode()

// 	mockArticleServices := new(mock.MockArticleServices)
// 	articleHandler := NewArticleHandler(mockArticleServices)
// 	assert := assert.New(t)

// 	articleHandler.APISearchTagsAndKeyword(c)

// 	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
// 	assert.JSONEq(`{"success":false,"message":"Bad request"}`, w.Body.String(), "Expected 'Bad request' message to be returned")
// }

// func TestAPISearchTagsAndKeywordServicesFail(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	req := &http.Request{
// 		URL:    &url.URL{},
// 		Header: make(http.Header),
// 	}
// 	q := req.URL.Query()
// 	keyword := "keyword"
// 	formatedTag := []string{"tag1"}
// 	from := 0
// 	q.Add("tags", "tag1")
// 	q.Add("q", keyword)
// 	q.Add("from", "0")

// 	c.Request = req
// 	c.Request.URL.RawQuery = q.Encode()

// 	mockArticleServices := new(mock.MockArticleServices)
// 	articleHandler := NewArticleHandler(mockArticleServices)
// 	assert := assert.New(t)

// 	articles := []entities.Article{}
// 	mockArticleServices.On("SearchArticles", keyword, formatedTag, from).Return(articles, int64(10), fmt.Errorf("service failed"))
// 	articleHandler.APISearchTagsAndKeyword(c)

// 	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
// 	assert.JSONEq(`{"success":false,"message":"Server error"}`, w.Body.String(), "Expected 'Server error' message to be returned")
// }

// func TestAPISearchTagsAndKeywordServicesSuccess(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	req := &http.Request{
// 		URL:    &url.URL{},
// 		Header: make(http.Header),
// 	}
// 	q := req.URL.Query()
// 	keyword := "keyword"
// 	formatedTag := []string{"tag1"}
// 	from := 0
// 	q.Add("tags", "tag1")
// 	q.Add("q", keyword)
// 	q.Add("from", "0")

// 	c.Request = req
// 	c.Request.URL.RawQuery = q.Encode()

// 	mockArticleServices := new(mock.MockArticleServices)
// 	articleHandler := NewArticleHandler(mockArticleServices)
// 	assert := assert.New(t)

// 	articles := []entities.Article{}
// 	mockArticleServices.On("SearchArticles", keyword, formatedTag, from).Return(articles, int64(10), nil)
// 	articleHandler.APISearchTagsAndKeyword(c)

// 	assert.Equal(http.StatusOK, w.Code, "Expected HTTP status code 200")
// }

// func TestAPICrawlArticleLeague(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	req := &http.Request{
// 		URL:    &url.URL{},
// 		Header: make(http.Header),
// 	}
// 	q := req.URL.Query()
// 	leagueName := []string{"league 1"}
// 	q.Add("league", "league 1")

// 	c.Request = req
// 	c.Request.URL.RawQuery = q.Encode()

// 	mockArticleServices := new(mock.MockArticleServices)
// 	articleHandler := NewArticleHandler(mockArticleServices)
// 	assert := assert.New(t)

// 	mockArticleServices.On("GetArticles", leagueName)
// 	articleHandler.APICrawlArticleLeague(c)

// 	assert.Equal(http.StatusOK, w.Code, "Expected HTTP status code 200")
// 	assert.JSONEq(`{"success":true,"message":"Signal crawl artilce success"}`, w.Body.String(), "Expected 'Signal crawl artilce success' message to be returned")
// }

// func TestAPIAddUpdateNewTagFail(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	req := &http.Request{
// 		URL:    &url.URL{},
// 		Header: make(http.Header),
// 	}
// 	q := req.URL.Query()
// 	q.Add("tag", "tag1")

// 	c.Request = req
// 	c.Request.URL.RawQuery = q.Encode()

// 	mockArticleServices := new(mock.MockArticleServices)
// 	articleHandler := NewArticleHandler(mockArticleServices)
// 	assert := assert.New(t)

// 	mockArticleServices.On("AddTagForAllArticle", "tag1").Return(fmt.Errorf("add tag service failed"))
// 	articleHandler.APIUpdateNewTag(c)

// 	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
// 	assert.JSONEq(`{"success":false,"message":"Update tag failed"}`, w.Body.String(), "Expected 'Update tag failed' message to be returned")
// }

// func TestAPIAddUpdateNewTagSuccess(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	req := &http.Request{
// 		URL:    &url.URL{},
// 		Header: make(http.Header),
// 	}
// 	q := req.URL.Query()
// 	q.Add("tag", "tag1")

// 	c.Request = req
// 	c.Request.URL.RawQuery = q.Encode()

// 	mockArticleServices := new(mock.MockArticleServices)
// 	articleHandler := NewArticleHandler(mockArticleServices)
// 	assert := assert.New(t)

// 	mockArticleServices.On("AddTagForAllArticle", "tag1").Return(nil)
// 	articleHandler.APIUpdateNewTag(c)

// 	assert.Equal(http.StatusOK, w.Code, "Expected HTTP status code 200")
// 	assert.JSONEq(`{"success":true,"message":"Update tag successfull"}`, w.Body.String(), "Expected 'Update tag successfull' message to be returned")
// }

// func TestAPIGetArticleCountSuccess(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	mockArticleServices := new(mock.MockArticleServices)
// 	articleHandler := NewArticleHandler(mockArticleServices)
// 	assert := assert.New(t)

// 	mockArticleServices.On("GetArticleCount").Return(int64(100), int64(100), nil)
// 	articleHandler.APIGetArticleCount(c)

// 	assert.Equal(http.StatusOK, w.Code, "Expected HTTP status code 200")
// 	assert.JSONEq(`{"success":true,"message":"Get article count success", "today":100, "total":100}`, w.Body.String(), "Expected 'Get article count success' message to be returned")
// }

// func TestAPIGetArticleCountFail(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	mockArticleServices := new(mock.MockArticleServices)
// 	articleHandler := NewArticleHandler(mockArticleServices)
// 	assert := assert.New(t)

// 	mockArticleServices.On("GetArticleCount").Return(int64(100),int64(100), fmt.Errorf("GetArticleCount service fail"))
// 	articleHandler.APIGetArticleCount(c)

// 	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
// 	assert.JSONEq(`{"success":false,"message":"Server error"}`, w.Body.String(), "Expected 'Server error' message to be returned")
// }

// func TestAPIDeleteArticleBindJSONFail(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)
// 	c.Request = &http.Request{
// 		Header: make(http.Header),
// 	}
// 	invalidArticleJson := "invalid"

// 	mockJsonPost(c, invalidArticleJson)

// 	mockArticleServices := new(mock.MockArticleServices)
// 	articleHandler := NewArticleHandler(mockArticleServices)
// 	assert := assert.New(t)

// 	mockArticleServices.On("DeleteArticle", "Article title").Return(fmt.Errorf("Delete failed"))

// 	articleHandler.APIDeleteArticle(c)

// 	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
// 	assert.JSONEq(`{"success":false,"message":"Delete article failed"}`, w.Body.String(), "Expected message 'Delete article failed' to be returned")
// }

// func TestAPIDeleteArticleFail(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)
// 	c.Request = &http.Request{
// 		Header: make(http.Header),
// 	}
// 	inputArticleJson := map[string]interface{}{"id": 1}

// 	mockJsonPost(c, inputArticleJson)

// 	mockArticleServices := new(mock.MockArticleServices)
// 	articleHandler := NewArticleHandler(mockArticleServices)
// 	assert := assert.New(t)

// 	mockArticleServices.On("DeleteArticle", uint(1)).Return(fmt.Errorf("Delete failed"))

// 	articleHandler.APIDeleteArticle(c)

// 	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 400")
// 	assert.JSONEq(`{"success":false,"message":"Delete article failed"}`, w.Body.String(), "Expected message 'Delete article failed' to be returned")
// }

// func TestAPIDeleteArticleSuccess(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)
// 	c.Request = &http.Request{
// 		Header: make(http.Header),
// 	}
// 	inputArticleJson := map[string]interface{}{"id": 1}

// 	mockJsonPost(c, inputArticleJson)

// 	mockArticleServices := new(mock.MockArticleServices)
// 	articleHandler := NewArticleHandler(mockArticleServices)
// 	assert := assert.New(t)

// 	mockArticleServices.On("DeleteArticle", uint(1)).Return(nil)

// 	articleHandler.APIDeleteArticle(c)

// 	assert.Equal(http.StatusOK, w.Code, "Expected HTTP status code 400")
// 	assert.JSONEq(`{"success":true,"message":"Delete article successfull"}`, w.Body.String(), "Expected message 'Delete article successfull' to be returned")
// }

// func TestAPIGetFirstPageOfLeagueRelatedArticleSuccess(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	req := &http.Request{
// 		URL:    &url.URL{},
// 		Header: make(http.Header),
// 	}
// 	q := req.URL.Query()
// 	q.Add("league", "league test")

// 	c.Request = req
// 	c.Request.URL.RawQuery = q.Encode()

// 	articles := []services.ArticleCache{}
// 	mockArticleServices := new(mock.MockArticleServices)
// 	articleHandler := NewArticleHandler(mockArticleServices)
// 	assert := assert.New(t)

// 	mockArticleServices.On("GetFirstPageOfLeagueRelatedArticle", "league test").Return(articles, nil)
// 	articleHandler.APIGetFirstPageOfLeagueRelatedArticle(c)

// 	assert.Equal(http.StatusOK, w.Code, "Expected HTTP status code 200")
// 	assert.JSONEq(`{"success":true,"articles":[]}`, w.Body.String(), "Expected 'articles' to be returned")
// }

// func TestAPIGetFirstPageOfLeagueRelatedArticleFail(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	req := &http.Request{
// 		URL:    &url.URL{},
// 		Header: make(http.Header),
// 	}
// 	q := req.URL.Query()
// 	q.Add("league", "league test")

// 	c.Request = req
// 	c.Request.URL.RawQuery = q.Encode()

// 	articles := []services.ArticleCache{}
// 	mockArticleServices := new(mock.MockArticleServices)
// 	articleHandler := NewArticleHandler(mockArticleServices)
// 	assert := assert.New(t)

// 	mockArticleServices.On("GetFirstPageOfLeagueRelatedArticle", "league test").Return(articles, fmt.Errorf("Cant get articles"))
// 	articleHandler.APIGetFirstPageOfLeagueRelatedArticle(c)

// 	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
// 	assert.JSONEq(`{"success":false,"message":"Server error"}`, w.Body.String(), "Expected 'Server error' to be returned")
// }