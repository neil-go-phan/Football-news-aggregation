package handler

import (
	"net/http"
	"server/handler/presenter"
	"server/services"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

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
	leaguesNames, err := leaguesHandler.handler.GetLeaguesNameActive()
	if err != nil {
		log.Errorln("error occurs when response League name active to frontend: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "leagues": leaguesNames})
}

func (leaguesHandler *LeaguesHandler) ListLeagues(c *gin.Context) {
	leagues, err := leaguesHandler.handler.ListLeagues()
	if err != nil {
		log.Errorln("Error occurs when response list leagues to frontend: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Internal server error"})
		return
	}
	// parse to presenter
	resposeLeagues := new(presenter.Leagues)
	for _,league := range *leagues {
		if league.LeagueName != "Tin tức bóng đá" {
			presenterLeague := presenter.League{
				LeagueName: league.LeagueName,
				Active: league.Active,
			}
			resposeLeagues.Leagues = append(resposeLeagues.Leagues, presenterLeague)
		}
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "leagues": resposeLeagues})
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
