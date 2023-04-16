package routes

import (
	"server/handler"
	"server/middlewares"

	"github.com/gin-gonic/gin"
)

type LeaguesRoutes struct {
	leaguesHandler handler.LeaguesHandler
}

func NewLeaguesRoutes(leaguesHandler *handler.LeaguesHandler) *LeaguesRoutes{
	routes := &LeaguesRoutes{
		leaguesHandler: *leaguesHandler,
	}
	return routes
}

func (leagueRoute *LeaguesRoutes)Setup(r *gin.Engine) {
	routes := r.Group("leagues")
	{
		routes.GET("list-name", leagueRoute.leaguesHandler.GetLeaguesName)
		routes.GET("list-all", middlewares.CheckToken(), leagueRoute.leaguesHandler.ListLeagues)
		routes.GET("change-status", middlewares.CheckToken(), leagueRoute.leaguesHandler.ChangeStatus)
	}
}