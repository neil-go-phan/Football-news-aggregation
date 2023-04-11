package handler

import (
	"net/http"
	"server/services"

	"github.com/gin-gonic/gin"
)

// admin change config via UI

type LeaguesHandler struct {
	handler services.LeaguesServices
}

func NewLeaguesHandler(handler services.LeaguesServices) *LeaguesHandler {
	leaguesHandler := &LeaguesHandler{
		handler: handler,
	}
	return leaguesHandler;
}

func (leaguesHandler *LeaguesHandler) ListLeagues(c *gin.Context) {
	leagues := leaguesHandler.handler.ListLeagues()
	c.JSON(http.StatusOK, gin.H{"success": true, "leagues": leagues})
}

