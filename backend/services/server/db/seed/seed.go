package seed

import (
	"server/handler"
	"time"

	log "github.com/sirupsen/logrus"
)

func SeedData(articleHandler *handler.ArticleHandler, schedulesHandler *handler.ScheduleHandler) {
	now := time.Now()
	var DAYOFWEEK = 1

	for i := -DAYOFWEEK; i <= DAYOFWEEK; i++ {
		date := now.AddDate(0, 0, i)
		schedulesHandler.SignalToCrawlerSeed(date)
	}

	// Get articles
	articleHandler.GetArticleSeed()
	log.Printf("Add seed data success\n")
}