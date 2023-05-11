package handler

// import (
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"net/url"
// 	"testing"

// 	"server/entities"
// 	mock "server/handler/mock"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )

// func TestListLeagues(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	mockLeagueServices := new(mock.MockLeaguesServices)
// 	leaguesHandler := NewLeaguesHandler(mockLeagueServices)
// 	assert := assert.New(t)

// 	leaguesOutPut := entities.Leagues{
// 		Leagues: []entities.League{
// 			{LeagueName: "league 1", Active: true},
// 			{LeagueName: "league 2", Active: false},
// 			{LeagueName: "league 3", Active: true},
// 		},
// 	}
// 	mockLeagueServices.On("ListLeagues").Return(leaguesOutPut)
// 	leaguesHandler.ListLeagues(c)

// 	assert.Equal(http.StatusOK, w.Code, "Expected HTTP status code OK")
// 	assert.JSONEq(`{"success":true,"leagues":{"leagues":[{"league_name":"league 1","active":true}, {"league_name":"league 2","active":false}, {"league_name":"league 3","active":true}]}}`, w.Body.String(), "Expected leagues list to be returned")
// }

// func TestGetLeaguesName(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	mockLeagueServices := new(mock.MockLeaguesServices)
// 	leaguesHandler := NewLeaguesHandler(mockLeagueServices)
// 	assert := assert.New(t)

// 	leagues:= []string{"league 1", "league 3"}

// 	mockLeagueServices.On("GetLeaguesNameActive").Return(leagues)
// 	leaguesHandler.GetLeaguesName(c)

// 	assert.Equal(http.StatusOK, w.Code, "Expected HTTP status code OK")
// 	assert.JSONEq(`{"success":true,"leagues":["league 1","league 3"]}`, w.Body.String(), "Expected active leagues name list to be returned")
// }

// func TestChangeStatusSuccess(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	req := &http.Request{
// 		URL:    &url.URL{},
// 		Header: make(http.Header),
// 	}
// 	q := req.URL.Query()
// 	q.Add("league", "league 1")

// 	c.Request = req
// 	c.Request.URL.RawQuery = q.Encode()
// 	mockLeagueServices := new(mock.MockLeaguesServices)
// 	leaguesHandler := NewLeaguesHandler(mockLeagueServices)
// 	assert := assert.New(t)

// 	mockLeagueServices.On("ChangeStatus", "league 1").Return(true, nil)
// 	leaguesHandler.ChangeStatus(c)

// 	assert.Equal(http.StatusOK, w.Code, "Expected HTTP status code OK")
// 	assert.JSONEq(`{"success":true,"message":"Change league active status success","status_active":true}`, w.Body.String(), "Expected change league active status success message to be returned")
// }

// func TestChangeStatusFail(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	req := &http.Request{
// 		URL:    &url.URL{},
// 		Header: make(http.Header),
// 	}
// 	q := req.URL.Query()
// 	q.Add("league", "league 1")

// 	c.Request = req
// 	c.Request.URL.RawQuery = q.Encode()
// 	mockLeagueServices := new(mock.MockLeaguesServices)
// 	leaguesHandler := NewLeaguesHandler(mockLeagueServices)
// 	assert := assert.New(t)

// 	mockLeagueServices.On("ChangeStatus", "league 1").Return(false, fmt.Errorf("fail to change status"))
// 	leaguesHandler.ChangeStatus(c)

// 	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code OK")
// 	assert.JSONEq(`{"success":false,"message":"Change league active status fail","status_active":false}`, w.Body.String(), "Expected change league active status fail message to be returned")
// }