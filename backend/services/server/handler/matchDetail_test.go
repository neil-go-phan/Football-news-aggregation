package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"server/entities"
	mock "server/handler/mock"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAPIGetMatchDetailInvalidDate(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("date", "35-35-2018")
	q.Add("club_1", "team 1")
	q.Add("club_2", "team 2")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockMatchDetailServices := new(mock.MockMatchDetailServices)
	matchDetailHandler := NewMatchDetailHandler(mockMatchDetailServices)
	assert := assert.New(t)

	matchDetailHandler.APIGetMatchDetail(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(`{"success":false,"message":"Date invalid"}`, w.Body.String(), "Expected 'Date invalid' message to be returned")
}

func TestAPIGetMatchDetailFail(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("date", "2023-05-02")
	q.Add("club_1", "team 1")
	q.Add("club_2", "team 2")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockMatchDetailServices := new(mock.MockMatchDetailServices)
	matchDetailHandler := NewMatchDetailHandler(mockMatchDetailServices)
	assert := assert.New(t)

	matchDetail := entities.MatchDetail{}

	date, _ := time.Parse(DATE_LAYOUT, "2023-05-02")
	mockMatchDetailServices.On("GetMatchDetail", date, "team 1", "team 2").Return(matchDetail, fmt.Errorf("cant get match detail"))
	matchDetailHandler.APIGetMatchDetail(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(`{"success":false,"message":"Server error"}`, w.Body.String(), "Expected 'Server error' message to be returned")
}

func TestAPIGetMatchDetailSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("date", "2023-05-02")
	q.Add("club_1", "team 1")
	q.Add("club_2", "team 2")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockMatchDetailServices := new(mock.MockMatchDetailServices)
	matchDetailHandler := NewMatchDetailHandler(mockMatchDetailServices)
	assert := assert.New(t)

	matchDetail := entities.MatchDetail{}

	date, _ := time.Parse(DATE_LAYOUT, "2023-05-02")
	mockMatchDetailServices.On("GetMatchDetail", date, "team 1", "team 2").Return(matchDetail, nil)
	matchDetailHandler.APIGetMatchDetail(c)

	assert.Equal(http.StatusOK, w.Code, "Expected HTTP status code ok")
}