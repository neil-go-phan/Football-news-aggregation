package handler

import (
	"net/http"
	"server/entities"
	"server/services"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)


type MatchDetailHandler struct {
	handler services.MatchDetailServices
}

func NewMatchDetailHandler(handler services.MatchDetailServices) *MatchDetailHandler {
	matchDetailHandler := &MatchDetailHandler{
		handler: handler,
	}
	return matchDetailHandler
}

func (matchDetailHandler *MatchDetailHandler) SignalToCrawler(matchUrl entities.MatchURLsOnDay) {
	matchDetailHandler.handler.GetMatchDetailsOnDayFromCrawler(matchUrl)
}

func (matchDetailHandler *MatchDetailHandler) APIGetMatchDetail(c *gin.Context) {
	dateString := c.Query("date")
	club1Name := c.Query("club_1")
	club2Name := c.Query("club_2")

	date, err := time.Parse(DATE_LAYOUT, dateString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "message": "Date invalid"})
	}
	// request to elasticsearch
	matchDetail, err := matchDetailHandler.handler.GetMatchDetail(date, club1Name, club2Name)
	if err != nil {
		log.Printf("error occurred while services layer request to elastic search to get match detail: %s\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Server error"})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "match_detail": matchDetail})
}
