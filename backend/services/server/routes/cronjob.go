package routes

import (
	"server/handler"
	"server/middlewares"

	"github.com/gin-gonic/gin"
)

type CronjobRoutes struct {
	chartHandler *handler.CronjobHandler
}

func NewCronjobRoutes(chartHandler *handler.CronjobHandler) *CronjobRoutes {
	routes := &CronjobRoutes{
		chartHandler: chartHandler,
	}
	return routes
}

func (route *CronjobRoutes) Setup(r *gin.Engine) {
	routes := r.Group("cronjob")
	{
		routes.GET("list-cronjob", middlewares.CheckToken(), route.chartHandler.APIListCrawlerCronjob)
		routes.GET("cronjob-on-hour", middlewares.CheckToken(), route.chartHandler.APIGetCronjobOnHour)
		routes.GET("cronjob-on-day", middlewares.CheckToken(), route.chartHandler.APIGetCronjobOnDay)
	}
}
