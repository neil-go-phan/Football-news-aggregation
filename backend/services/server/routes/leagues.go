package routes

import (
	"server/handler"

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
	}
}