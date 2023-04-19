package handler

import (
	log "github.com/sirupsen/logrus"
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
	return leaguesHandler
}

func (leaguesHandler *LeaguesHandler) GetLeaguesName(c *gin.Context) {
	leaguesNames := leaguesHandler.handler.GetLeaguesNameActive()
	c.JSON(http.StatusOK, gin.H{"success": true, "leagues": leaguesNames})
}

func (leaguesHandler *LeaguesHandler) ListLeagues(c *gin.Context) {
	
	leagues := leaguesHandler.handler.ListLeagues()
	// remove default league 
	c.JSON(http.StatusOK, gin.H{"success": true, "leagues": leagues})
}

func (leaguesHandler *LeaguesHandler) ChangeStatus(c *gin.Context) {
	leagueName := c.Query("league")

	statusActive, err := leaguesHandler.handler.ChangeStatus(leagueName)
	if err != nil {
		log.Printf("Error occus whene trying to change league active status: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Change league active status fail", "status_active": statusActive})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Change league active status success", "status_active": statusActive})
}
