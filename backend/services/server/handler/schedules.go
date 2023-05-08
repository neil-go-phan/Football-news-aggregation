package handler

import (
	"fmt"
	"net/http"
	"server/entities"
	"server/services"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

var DATE_LAYOUT = "2006-01-02" // ISO-8601

type ScheduleHandler struct {
	handler services.SchedulesServices
}

func NewSchedulesHandler(handler services.SchedulesServices) *ScheduleHandler {
	schedulesHandler := &ScheduleHandler{
		handler: handler,
	}
	return schedulesHandler
}

func (schedulesHandler *ScheduleHandler) SignalToCrawlerOnNewDay(cronjob *cron.Cron) {
	_, err := cronjob.AddFunc("0 23 * * *", func() {
		var wg sync.WaitGroup
		now := time.Now()
		var matchsToDay entities.MatchURLsWithTimeOnDay
		var DAYOFWEEK = 7
		for i := -1; i <= DAYOFWEEK; i++ {
			wg.Add(1)
			date := now.AddDate(0, 0, i)
			go func(date time.Time, matchsToDay *entities.MatchURLsWithTimeOnDay) {
				defer wg.Done()
				schedulesHandler.handler.GetSchedules(date.Format("02-01-2006"))
				matchUrls := schedulesHandler.handler.GetMatchURLsOnDay()
				schedulesHandler.handler.SignalMatchDetailServiceToCrawl(matchUrls)
				schedulesHandler.handler.ClearMatchURLsOnDay()
				if date.Day() == time.Now().Day() + 1 {
					*matchsToDay = schedulesHandler.handler.GetMatchURLsOnTime()
				}
				schedulesHandler.handler.ClearMatchURLsOnTime()
			}(date, &matchsToDay)
		}
		wg.Wait()
		go func() {
			MakeCronJobCrawlMatch(matchsToDay, schedulesHandler)
		}()
		
	})
	if err != nil {
		log.Println("error occurred while seting up getSchedules cronjob: ", err)
	}
}

func MakeCronJobCrawlMatch(matchsToDay entities.MatchURLsWithTimeOnDay, schedulesHandler *ScheduleHandler) {
	for _, matchsOnTime := range matchsToDay.MatchsOnTimes {
		go func(matchsOnTime entities.MatchURLsOnTime) {
			matchURLs := entities.MatchURLsOnDay(matchsOnTime)
			now := time.Now()
			loc, err := time.LoadLocation("UTC")
			if err != nil {
				log.Error(err)
				return
			}
			utcTime := time.Date(
				now.Year(),
				now.Month(),
				now.Day(),
				now.Hour(),
				now.Minute(),
				now.Second(),
				now.Nanosecond(),
				loc,
			)

			duration := matchsOnTime.Date.Sub(utcTime)

			fmt.Printf("create a cronjob sleep from %v to %v is with duration %v: \n",utcTime,matchsOnTime.Date, duration)

			// time.Sleep(duration)

			log.Printf("Start cronjob crawl match at: %s", matchsOnTime.Date)
			ticker := time.NewTicker(1 * time.Minute)
			done := make(chan bool)
			matchEnded := false
			var wg sync.WaitGroup
			wg.Add(1)

			go func() {
				defer wg.Done()
				for {
					select {
					case <-done:
						return
					case <-ticker.C:
						matchDetails := schedulesHandler.handler.SignalMatchDetailServiceToCrawl(matchURLs)
						log.Printf("Crawl matchs on %v:%v\n", time.Now().Hour(), time.Now().Minute())
						if checkMatchsEnd(matchDetails) {
							matchEnded = true
						}
					}
					// stop go routine
					if matchEnded {
            ticker.Stop()
            done <- true
            close(done)
						// update schedules
						schedulesHandler.handler.GetSchedules(time.Now().Format("02-01-2006"))
            return
        }
				}
			}()
			wg.Wait()
			log.Printf("Stop cronjob crawl match at: %s", time.Now())
		}(matchsOnTime)
	}
}

// check if all match is ended
func checkMatchsEnd(matchDetails []entities.MatchDetail) bool {
	for _, matchDetail := range matchDetails {
		if matchDetail.MatchDetailTitle.MatchStatus != "Kết thúc" {
			return false
		}
	}
	return true
}

func (schedulesHandler *ScheduleHandler) SignalToCrawlerTest() {
	date := time.Now().AddDate(0,0,0)
	schedulesHandler.handler.GetSchedules(date.Format("02-01-2006"))
	matchUrls := schedulesHandler.handler.GetMatchURLsOnDay()
	schedulesHandler.handler.SignalMatchDetailServiceToCrawl(matchUrls)
	schedulesHandler.handler.ClearMatchURLsOnDay()
	MakeCronJobCrawlMatch(schedulesHandler.handler.GetMatchURLsOnTime(), schedulesHandler)
}


func (schedulesHandler *ScheduleHandler) APIGetAllScheduleLeagueOnDay(c *gin.Context) {
	dateString := c.Query("date")
	date, err := time.Parse(DATE_LAYOUT, dateString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "message": "Date invalid"})
		return
	}

	// request to elasticsearch
	schedules, err := schedulesHandler.handler.GetAllScheduleLeagueOnDay(date)
	if err != nil {
		log.Printf("error occurred while services layer request to elastic search to get schedules on date: %s\n", date)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "schedules": schedules})
}

func (schedulesHandler *ScheduleHandler) APIGetScheduleLeagueOnDay(c *gin.Context) {
	dateString := c.Query("date")
	date, err := time.Parse(DATE_LAYOUT, dateString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "message": "Date invalid"})
		return
	}
	league := c.Query("league")

	// request to elasticsearch
	schedules, err := schedulesHandler.handler.GetScheduleLeagueOnDay(date, league)
	if err != nil {
		log.Printf("error occurred while services layer request to elastic search to get schedule of league: %s on date: %s\n", league, date)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "schedules": schedules})
}
