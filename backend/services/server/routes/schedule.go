package routes

import (
	"server/handler"

	"github.com/gin-gonic/gin"
)

type ScheduleRoutes struct {
	schedulesHandler handler.ScheduleHandler
}

func NewScheduleRoutes(schedulesHandler *handler.ScheduleHandler) *ScheduleRoutes{
	routes := &ScheduleRoutes{
		schedulesHandler: *schedulesHandler,
	}
	return routes
}

func (scheduleRoutes *ScheduleRoutes)Setup(r *gin.Engine) {
	routes := r.Group("schedules")
	{
		routes.GET("league-on-day",  scheduleRoutes.schedulesHandler.APIGetScheduleLeagueOnDay)
		routes.GET("all-league-on-day",  scheduleRoutes.schedulesHandler.APIGetAllScheduleLeagueOnDay)
	}
}