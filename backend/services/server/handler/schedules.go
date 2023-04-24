package handler

import (
	"net/http"
	"server/services"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

var DATE_LAYOUT = "2006-01-02" // ISO-8601

type ScheduleHandler struct {
	handler services.SchedulesServices
}

func NewSchedulesHandler(handler services.SchedulesServices) *ScheduleHandler {
	schedulesHandler := &ScheduleHandler{
		handler: handler,
	}
	return schedulesHandler
}

func (schedulesHandler *ScheduleHandler) SignalToCrawlerOnNewDay(cronjob *cron.Cron) {
	_, err := cronjob.AddFunc("0 23 * * *", func() {
		var wg sync.WaitGroup
		now := time.Now()
		var DAYOFWEEK = 7
		for i := -DAYOFWEEK; i <= DAYOFWEEK; i++ {
			wg.Add(1)
			date := now.AddDate(0, 0, i)
			go func(date time.Time) {
				defer wg.Done()
				schedulesHandler.handler.GetSchedules(date.Format("02-01-2006"))
				matchUrls := schedulesHandler.handler.GetMatchURLsOnDay()
				schedulesHandler.handler.SignalMatchDetailServiceToCrawl(matchUrls)
				schedulesHandler.handler.ClearMatchURLsOnDay()
			}(date)
		}
		wg.Wait()
	})
	if err != nil {
		log.Println("error occurred while seting up getSchedules cronjob: ", err)
	}
}

func (schedulesHandler *ScheduleHandler) APIGetAllScheduleLeagueOnDay(c *gin.Context) {
	dateString := c.Query("date")
	date, err := time.Parse(DATE_LAYOUT, dateString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "message": "Date invalid"})
	}

	// request to elasticsearch
	schedules, err := schedulesHandler.handler.GetAllScheduleLeagueOnDay(date)
	if err != nil {
		log.Printf("error occurred while services layer request to elastic search to get schedules on date: %s\n", date)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Server error"})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "schedules": schedules})
}

func (schedulesHandler *ScheduleHandler) APIGetScheduleLeagueOnDay(c *gin.Context) {
	dateString := c.Query("date")
	date, err := time.Parse(DATE_LAYOUT, dateString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "message": "Date invalid"})
	}
	league := c.Query("league")

	// request to elasticsearch
	schedules, err := schedulesHandler.handler.GetScheduleLeagueOnDay(date, league)
	if err != nil {
		log.Printf("error occurred while services layer request to elastic search to get schedule of league: %s on date: %s\n", league, date)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Server error"})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "schedules": schedules})
}
