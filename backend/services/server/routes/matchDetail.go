package routes

import (
	"server/handler"

	"github.com/gin-gonic/gin"
)

type MatchDetailRoutes struct {
	matchDetailRoutes handler.MatchDetailHandler
}

func NewMatchDetailRoutes(matchDetailRoutes *handler.MatchDetailHandler) *MatchDetailRoutes{
	routes := &MatchDetailRoutes{
		matchDetailRoutes: *matchDetailRoutes,
	}
	return routes
}

func (matchDetailRoutes *MatchDetailRoutes)Setup(r *gin.Engine) {
	routes := r.Group("match-detail")
	{
		routes.GET("get",  matchDetailRoutes.matchDetailRoutes.APIGetMatchDetail)
	}
}