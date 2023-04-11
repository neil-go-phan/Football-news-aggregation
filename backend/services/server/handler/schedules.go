package handler

import (
	"log"
	"net/http"
	"server/services"
	"time"

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
	_, err := cronjob.AddFunc("0 0 * * *", func() { schedulesHandler.handler.GetSchedules(time.Now().Format("02-01-2006")) })
	if err != nil {
		log.Println("error occurred while seting up getSchedules cronjob: ", err)
	}
}

func (schedulesHandler *ScheduleHandler) APIGetScheduleOnDay(c *gin.Context) {
	dateString := c.Query("date")
	date, err := time.Parse(DATE_LAYOUT, dateString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "message": "Date invalid"})
	}
	// request to elasticsearch
	schedules, err := schedulesHandler.handler.APIGetScheduleOnDay(date)
	if err != nil {
		log.Printf("error occurred while services layer request to elastic search to get schedule on date: %s\n", date)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Server error"})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "schedules": schedules})
}
