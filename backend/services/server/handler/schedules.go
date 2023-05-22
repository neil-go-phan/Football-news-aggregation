package handler

import (
	"net/http"
	serverhelper "server/helper"
	"server/repository"
	"server/services"
	"sync"
	"time"

	pb "server/proto"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
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
		var DAYOFWEEK = 7
		for i := -1; i <= DAYOFWEEK; i++ {
			wg.Add(1)
			date := now.AddDate(0, 0, i)
			go func(date time.Time) {
				defer wg.Done()
				schedulesHandler.handler.GetSchedules(date.Format("02-01-2006"))
				matchUrls := schedulesHandler.handler.GetAllMatchURLs()
				schedulesHandler.handler.SignalMatchDetailServiceToCrawl(matchUrls)
				schedulesHandler.handler.ClearAllMatchURLs()

				if date.Day() == time.Now().Day()+1 {
					MakeCronJobCrawlMatch(schedulesHandler.handler.GetMatchURLsOnTime(), schedulesHandler)
				}

			}(date)
		}
		wg.Wait()
		schedulesHandler.handler.ClearMatchURLsOnTime()
	})
	if err != nil {
		log.Errorln("error occurred while seting up getSchedules cronjob: ", err)
	}
}

func MakeCronJobCrawlMatch(matchsToDay repository.MatchURLsWithTimeOnDay, schedulesHandler *ScheduleHandler) {
	for _, matchsOnTime := range matchsToDay.MatchsOnTimes {
		go func(matchsOnTime repository.MatchURLsOnTime) {
			matchURLs := repository.AllMatchURLsOnDay(matchsOnTime)
			now := time.Now()
			loc, err := time.LoadLocation("UTC")
			if err != nil {
				log.Error(err)
				return
			}
			utcTimeNow := time.Date(
				now.Year(),
				now.Month(),
				now.Day(),
				now.Hour(),
				now.Minute(),
				now.Second(),
				now.Nanosecond(),
				loc,
			)
			utcTimeThen := time.Date(
				matchsOnTime.Date.Year(),
				matchsOnTime.Date.Month(),
				matchsOnTime.Date.Day(),
				matchsOnTime.Date.Hour(),
				matchsOnTime.Date.Minute(),
				matchsOnTime.Date.Second(),
				matchsOnTime.Date.Nanosecond(),
				loc,
			)
			duration := utcTimeThen.Sub(utcTimeNow)
			if duration > time.Hour*time.Duration(24) || duration < 0{
				return
			}
			log.Printf("create a cronjob sleep from %v to %v with duration %v and get %v match\n", utcTimeNow, utcTimeThen, duration, len(matchsOnTime.Urls))

			time.Sleep(duration)

			log.Printf("Start cronjob crawl match at: %s", utcTimeThen)

			ticker := time.NewTicker(1 * time.Minute)

			done := make(chan bool)
			var wg sync.WaitGroup
			wg.Add(1)

			go func(wg *sync.WaitGroup) {
				for {
					select {
					case <-done:
						return
					case <-ticker.C:
						matchDetails := schedulesHandler.handler.SignalMatchDetailServiceToCrawl(matchURLs)
						log.Printf("Crawl matchs on %v:%v\n", time.Now().Hour(), time.Now().Minute())
						if checkMatchsEnd(matchDetails) {
							// stop go routine
							ticker.Stop()
							// update schedules
							schedulesHandler.handler.GetSchedules(time.Now().Format("02-01-2006"))
							wg.Done()
							done <- true
							close(done)
							return
						}
					}
				}
			}(&wg)
			wg.Wait()
			log.Printf("Stop cronjob crawl match at: %s", time.Now())
		}(matchsOnTime)
	}
}

// check if all match is ended
func checkMatchsEnd(matchDetails []*pb.MatchDetail) bool {
	for _, matchDetail := range matchDetails {
		log.Printf("%s vs %s: status %s\n", matchDetail.MatchDetailTitle.Club_1.Name, matchDetail.MatchDetailTitle.Club_2.Name, matchDetail.MatchDetailTitle.MatchStatus)
		if serverhelper.FormatVietnamese(matchDetail.MatchDetailTitle.MatchStatus) != "ket thuc" {
			return false
		}
	}
	return true
}

func (schedulesHandler *ScheduleHandler) SignalToCrawlerToDay() {
	date := time.Now()
	schedulesHandler.handler.GetSchedules(date.Format("02-01-2006"))
	matchUrls := schedulesHandler.handler.GetAllMatchURLs()
	// chunks := serverhelper.MatchUrlsChunk(matchUrls.Urls, 60) 
	// for _, chunk := range chunks {
	// 	matchUrl := repository.AllMatchURLsOnDay{
	// 		Date: matchUrls.Date,
	// 		Urls: chunk,
	// 	}
	// 	schedulesHandler.handler.SignalMatchDetailServiceToCrawl(matchUrl)
	// }
	schedulesHandler.handler.SignalMatchDetailServiceToCrawl(matchUrls)
	schedulesHandler.handler.ClearAllMatchURLs()
	MakeCronJobCrawlMatch(schedulesHandler.handler.GetMatchURLsOnTime(), schedulesHandler)
	schedulesHandler.handler.ClearMatchURLsOnTime()
}

func (schedulesHandler *ScheduleHandler) SignalToCrawlerSeed(date time.Time) {
	schedulesHandler.handler.GetSchedules(date.Format("02-01-2006"))
	matchUrls := schedulesHandler.handler.GetAllMatchURLs()
	log.Printf("seed for date: %v len: %v\n", matchUrls.Date, len(matchUrls.Urls))
	// chunks := serverhelper.MatchUrlsChunk(matchUrls.Urls, 60) 
	// for _, chunk := range chunks {
	// 	matchUrl := repository.AllMatchURLsOnDay{
	// 		Date: matchUrls.Date,
	// 		Urls: chunk,
	// 	}
	// 	schedulesHandler.handler.SignalMatchDetailServiceToCrawl(matchUrl)
	// }
	schedulesHandler.handler.SignalMatchDetailServiceToCrawl(matchUrls)
	schedulesHandler.handler.ClearAllMatchURLs()
	schedulesHandler.handler.ClearMatchURLsOnTime()
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
