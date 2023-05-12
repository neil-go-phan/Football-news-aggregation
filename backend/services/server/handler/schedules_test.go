package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	mock "server/handler/mocks"
	"server/repository"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAPIGetAllScheduleLeagueOnDay_FailToParseDate(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("date", "invalid date")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockScheduleServices := new(mock.MockSchedulesServices)
	shceduleHandler := NewSchedulesHandler(mockScheduleServices)
	assert := assert.New(t)

	shceduleHandler.APIGetAllScheduleLeagueOnDay(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(`{"success":false,"message":"Date invalid"}`, w.Body.String(), "Expected 'Date invalid' message to be returned")
}

func TestAPIGetAllScheduleLeagueOnDay_ServicesFail(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	assert := assert.New(t)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("date", "2006-01-02")
	date, err := time.Parse(DATE_LAYOUT, "2006-01-02")
	assert.Nil(err)

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockScheduleServices := new(mock.MockSchedulesServices)
	shceduleHandler := NewSchedulesHandler(mockScheduleServices)
	
	schedules := repository.ScheduleOnDay{}

	mockScheduleServices.On("GetAllScheduleLeagueOnDay", date).Return(schedules, fmt.Errorf("failed"))
	shceduleHandler.APIGetAllScheduleLeagueOnDay(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(`{"success":false,"message":"Server error"}`, w.Body.String(), "Expected 'Server error' message to be returned")
}

func TestAPIGetAllScheduleLeagueOnDay_Success(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	assert := assert.New(t)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("date", "2006-01-02")
	date, err := time.Parse(DATE_LAYOUT, "2006-01-02")
	assert.Nil(err)

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockScheduleServices := new(mock.MockSchedulesServices)
	shceduleHandler := NewSchedulesHandler(mockScheduleServices)
	
	schedules := repository.ScheduleOnDay{
		Date: date,
		DateWithWeekday: "monday",
		ScheduleOnLeagues: []repository.ScheduleOnLeague{},
	}

	mockScheduleServices.On("GetAllScheduleLeagueOnDay", date).Return(schedules, nil)
	shceduleHandler.APIGetAllScheduleLeagueOnDay(c)

	assert.Equal(http.StatusOK, w.Code, "Expected HTTP status code 200")
	assert.JSONEq(`{"success":true,"schedules":{"date":"2006-01-02T00:00:00Z", "date_with_weekday":"monday", "schedule_on_leagues":[]}}`, w.Body.String(), "Expected success message to be returned")
}

func TestAPIGetScheduleLeagueOnDay_FailToParseDate(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("date", "invalid date")

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockScheduleServices := new(mock.MockSchedulesServices)
	shceduleHandler := NewSchedulesHandler(mockScheduleServices)
	assert := assert.New(t)

	shceduleHandler.APIGetScheduleLeagueOnDay(c)

	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
	assert.JSONEq(`{"success":false,"message":"Date invalid"}`, w.Body.String(), "Expected 'Date invalid' message to be returned")
}

func TestAPIGetScheduleLeagueOnDay_ServicesFail(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	assert := assert.New(t)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("date", "2006-01-02")
	q.Add("league", "league test")
	date, err := time.Parse(DATE_LAYOUT, "2006-01-02")
	assert.Nil(err)

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockScheduleServices := new(mock.MockSchedulesServices)
	shceduleHandler := NewSchedulesHandler(mockScheduleServices)
	
	schedules := repository.ScheduleOnDay{}

	mockScheduleServices.On("GetScheduleLeagueOnDay", date, "league test").Return(schedules, fmt.Errorf("failed"))
	shceduleHandler.APIGetScheduleLeagueOnDay(c)

	assert.Equal(http.StatusInternalServerError, w.Code, "Expected HTTP status code 500")
	assert.JSONEq(`{"success":false,"message":"Server error"}`, w.Body.String(), "Expected 'Server error' message to be returned")
}

func TestAPIGetScheduleLeagueOnDay_Success(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	assert := assert.New(t)

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
	}
	q := req.URL.Query()
	q.Add("date", "2006-01-02")
	q.Add("league", "league test")
	date, err := time.Parse(DATE_LAYOUT, "2006-01-02")
	assert.Nil(err)

	c.Request = req
	c.Request.URL.RawQuery = q.Encode()

	mockScheduleServices := new(mock.MockSchedulesServices)
	shceduleHandler := NewSchedulesHandler(mockScheduleServices)
	
	schedules := repository.ScheduleOnDay{
		Date: date,
		DateWithWeekday: "monday",
		ScheduleOnLeagues: []repository.ScheduleOnLeague{},
	}

	mockScheduleServices.On("GetScheduleLeagueOnDay", date, "league test").Return(schedules, nil)
	shceduleHandler.APIGetScheduleLeagueOnDay(c)

	assert.Equal(http.StatusOK, w.Code, "Expected HTTP status code 200")
	assert.JSONEq(`{"success":true,"schedules":{"date":"2006-01-02T00:00:00Z", "date_with_weekday":"monday", "schedule_on_leagues":[]}}`, w.Body.String(), "Expected success message to be returned")
}