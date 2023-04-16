package handler

import (
	"log"
	"net/http"
	"server/entities"
	"server/services"
	"time"

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

// func (matchDetailHandler *MatchDetailHandler) SignalToCrawler(cronjob *cron.Cron) {
// 	_, err := cronjob.AddFunc("0 0 * * *", func() { matchDetailHandler.handler })
// 	if err != nil {
// 		log.Println("error occurred while seting up getSchedules cronjob: ", err)
// 	}
// }

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
	matchDetail, err := matchDetailHandler.handler.APIGetMatchDetail(date, club1Name, club2Name)
	if err != nil {
		log.Printf("error occurred while services layer request to elastic search to get match detail: %s\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Server error"})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "match_detail": matchDetail})
}
