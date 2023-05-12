package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"server/entities"
	mock "server/handler/mocks"

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

	mockMatchDetailServices := new(mock.MockMatchServices)
	matchDetailHandler := NewMatchDetailHandler(mockMatchDetailServices)
	assert := assert.New(t)

	matchDetailHandler.APIGetMatchDetail(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(`{"success":false,"message":"Date invalid"}`, w.Body.String(), "Expected 'Date invalid' message to be returned")
}

func TestAPIGetMatchDetail_MatchFail(t *testing.T) {
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

	mockMatchDetailServices := new(mock.MockMatchServices)
	matchDetailHandler := NewMatchDetailHandler(mockMatchDetailServices)
	assert := assert.New(t)

	matchDetail := &entities.Match{}

	date, _ := time.Parse(DATE_LAYOUT, "2023-05-02")
	mockMatchDetailServices.On("GetMatch", date, "team 1", "team 2").Return(matchDetail, fmt.Errorf("cant get match detail"))
	
	
	matchDetailHandler.APIGetMatchDetail(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(`{"success":false,"message":"Server error"}`, w.Body.String(), "Expected 'Server error' message to be returned")
}

func TestAPIGetMatchDetail_LineUpFail(t *testing.T) {
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

	mockMatchDetailServices := new(mock.MockMatchServices)
	matchDetailHandler := NewMatchDetailHandler(mockMatchDetailServices)
	assert := assert.New(t)

	matchDetail := &entities.Match{
		LineupClub1ID: uint(1),
		LineupClub2ID: uint(2),
	}

	lineup1 := &entities.MatchLineUp{}
	lineup2 := &entities.MatchLineUp{}

	date, _ := time.Parse(DATE_LAYOUT, "2023-05-02")
	mockMatchDetailServices.On("GetMatch", date, "team 1", "team 2").Return(matchDetail, nil)
	mockMatchDetailServices.On("GetLineUps", uint(1), uint(2)).Return(lineup1, lineup2, fmt.Errorf("error"))

	matchDetailHandler.APIGetMatchDetail(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(`{"success":false,"message":"Server error"}`, w.Body.String(), "Expected 'Server error' message to be returned")
}

func TestAPIGetMatchDetail_Success(t *testing.T) {
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

	mockMatchDetailServices := new(mock.MockMatchServices)
	matchDetailHandler := NewMatchDetailHandler(mockMatchDetailServices)
	assert := assert.New(t)

	matchDetail := &entities.Match{
		LineupClub1ID: uint(1),
		LineupClub2ID: uint(2),
	}

	lineup1 := &entities.MatchLineUp{}
	lineup2 := &entities.MatchLineUp{}

	date, _ := time.Parse(DATE_LAYOUT, "2023-05-02")
	mockMatchDetailServices.On("GetMatch", date, "team 1", "team 2").Return(matchDetail, nil)
	mockMatchDetailServices.On("GetLineUps", uint(1), uint(2)).Return(lineup1, lineup2, nil)

	matchDetailHandler.APIGetMatchDetail(c)

	assert.Equal(http.StatusOK, w.Code, "Expected HTTP status code 200")
}