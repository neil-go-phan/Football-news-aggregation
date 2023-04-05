package handler

import (
	"github.com/robfig/cron/v3"
	"log"
	"server/services"
)

// user search article, server query elastic search

type ScheduleOnDayHandler struct {
	handler services.ScheduleOnDayServices
}

func NewScheduleOnDayHandler(handler services.ScheduleOnDayServices) *ScheduleOnDayHandler {
	scheduleOnDayHandler := &ScheduleOnDayHandler{
		handler: handler,
	}
	return scheduleOnDayHandler
}

func (scheduleOnDayHandler *ScheduleOnDayHandler) SignalToCrawler(cronjob *cron.Cron) {
	scheduleOnDayHandler.handler.GetSchedules()
	_, err := cronjob.AddFunc("0 0 * * *", func() { scheduleOnDayHandler.handler.GetSchedules() })
	if err != nil {
		log.Println("error occurred while seting up getSchedules cronjob: ", err)
	}
}
