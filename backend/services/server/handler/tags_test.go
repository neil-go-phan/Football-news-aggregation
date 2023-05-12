package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	mock "server/handler/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestListTags_Success(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	mockTagServices := new(mock.MockTagsServices)
	tagHandler := NewTagsHandler(mockTagServices)
	assert := assert.New(t)

	tagOutput := []string{"tag1", "tag2"}

	mockTagServices.On("ListTagsName").Return(tagOutput, nil)
	tagHandler.ListTags(c)

	assert.Equal(http.StatusOK, w.Code, "Expected HTTP status code OK")
	assert.JSONEq(`{"success":true,"tags":{"tags":["tag1","tag2"]}}`, w.Body.String(), "Expected tags list to be returned")
}

func TestListTags_Fail(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	mockTagServices := new(mock.MockTagsServices)
	tagHandler := NewTagsHandler(mockTagServices)
	assert := assert.New(t)

	tagOutput := []string{"tag1", "tag2"}

	mockTagServices.On("ListTagsName").Return(tagOutput, fmt.Errorf("fail"))
	tagHandler.ListTags(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(`{"success":false,"message":"Internal server error"}`, w.Body.String(), "Expected 'Internal server error' message to be returned")
}

func TestDeleteTags_Success(t *testing.T) {
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
	mockTagServices := new(mock.MockTagsServices)
	tagHandler := NewTagsHandler(mockTagServices)
	assert := assert.New(t)

	mockTagServices.On("DeleteTag", "tag1").Return(nil)
	tagHandler.DeleteTag(c)

	assert.Equal(http.StatusOK, w.Code, "Expected HTTP status code OK")
	assert.JSONEq(`{"success":true,"message":"Delete successfull"}`, w.Body.String(), "Expected delete successfull message to be returned")
}

func TestDeleteTagsFail(t *testing.T) {
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
	mockTagServices := new(mock.MockTagsServices)
	tagHandler := NewTagsHandler(mockTagServices)
	assert := assert.New(t)

	mockTagServices.On("DeleteTag", "tag1").Return(fmt.Errorf("fail to delete"))
	tagHandler.DeleteTag(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(`{"success":false,"message":"Delete failed"}`, w.Body.String(), "Expected delete failed message to be returned")
}

func TestAddTags_Success(t *testing.T) {
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
	mockTagServices := new(mock.MockTagsServices)
	tagHandler := NewTagsHandler(mockTagServices)
	assert := assert.New(t)

	mockTagServices.On("AddTag", "tag1").Return(nil)
	tagHandler.AddTag(c)

	assert.Equal(http.StatusOK, w.Code, "Expected HTTP status code OK")
	assert.JSONEq(`{"success":true,"message":"Add tag successfull"}`, w.Body.String(), "Expected add tag successfull message to be returned")
}

func TestAddTags_Fail(t *testing.T) {
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
	mockTagServices := new(mock.MockTagsServices)
	tagHandler := NewTagsHandler(mockTagServices)
	assert := assert.New(t)

	mockTagServices.On("AddTag", "tag1").Return(fmt.Errorf("fail to add tag"))
	tagHandler.AddTag(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(`{"success":false,"message":"Add tag failed"}`, w.Body.String(), "Expected add tag failed message to be returned")
}
