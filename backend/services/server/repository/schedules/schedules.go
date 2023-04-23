package schedulesrepo

import (
	"bytes"
	"context"
	"fmt"
	"server/entities"
	serverhelper "server/helper"
	pb "server/proto"
	"server/repository"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/elastic/go-elasticsearch/v7"
	jsoniter "github.com/json-iterator/go"
	"google.golang.org/grpc"
)

var SCHEDULE_INDEX_NAME = "schedules"

var NOTI_COMPLETE_CRAWL_TITLE = "Crawl schedule success"
var NOTI_COMPLETE_CRAWL_TYPE = "INFO"

type schedulesRepo struct {
	conn           *grpc.ClientConn
	es             *elasticsearch.Client
	leagueRepo     repository.LeaguesRepository
	tagsRepo       repository.TagRepository
	matchURLsOnDay entities.MatchURLsOnDay
	notification    repository.NotificationRepository
}

func NewSchedulesRepo(leagueRepo repository.LeaguesRepository, tagsRepo repository.TagRepository,	notification    repository.NotificationRepository, conn *grpc.ClientConn, es *elasticsearch.Client) *schedulesRepo {
	schedulesRepo := &schedulesRepo{
		conn:       conn,
		es:         es,
		leagueRepo: leagueRepo,
		tagsRepo:   tagsRepo,
		notification: notification,
	}
	return schedulesRepo
}

func (repo *schedulesRepo) GetSchedules(date *pb.Date) {
	client := pb.NewCrawlerServiceClient(repo.conn)

	// send gRPC request to crawler
	pbSchedules, err := client.GetSchedulesOnDay(context.Background(), date)
	if err != nil {
		log.Errorf("error occurred while get schedule on day from crawler error %v \n", err)
		return
	}
	// store schedule in elastic search
	var wg sync.WaitGroup
	elasticSchedules := PbSchedulesToScheduleElastic(pbSchedules)
	dateCasted, err := time.Parse("02-01-2006", pbSchedules.GetDateFormated())
	if err != nil {
		log.Errorf("error when parse date:", err)
	}
	repo.matchURLsOnDay.Date = dateCasted
	for _, schedule := range elasticSchedules {
		wg.Add(1)
		go func(schedule entities.ScheduleElastic) {
			defer wg.Done()

			matchUrls := getMatchUrlOnDay(schedule)
			repo.matchURLsOnDay.Urls = append(repo.matchURLsOnDay.Urls, matchUrls...)

			// auto store new league
			isNewLeague := isNewLeague(repo.leagueRepo.GetLeagues().Leagues, schedule.LeagueName)
			if isNewLeague {
				log.Println("detect a new league: ", schedule.LeagueName)
				repo.leagueRepo.AddLeague(schedule.LeagueName)

			} else {
				exist := checkScheduleWithElasticSearch(schedule, repo.es)

				if !exist {
					storeScheduleElasticsearch(schedule, repo.es)
					// check is schedule active // push matchUrl on day
					isActive := isLeagueActive(repo.leagueRepo.GetLeagues().Leagues, schedule.LeagueName)
					if isActive {
						isTagExist := isLeagueTagExist(repo.tagsRepo.ListTags().Tags, schedule.LeagueName)
						if !isTagExist {
							tagFormated := serverhelper.FormatVietnamese(schedule.LeagueName)
							log.Printf("Detect %s is new tag, add...", tagFormated)
							err := repo.tagsRepo.AddTag(tagFormated)
							if err != nil {
								log.Errorf("error occrus %s \n", err)
							}
						}

					}
				}

			}

		}(schedule)

	}
	wg.Wait()

	// repo.notification.Send(NOTI_COMPLETE_CRAWL_TITLE, NOTI_COMPLETE_CRAWL_TYPE, fmt.Sprintf("Crawler scrape schedule on %s success", date.GetDate()))
}

func (repo *schedulesRepo) GetAllScheduleLeagueOnDay(date time.Time) (entities.ScheduleOnDay, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	var scheduleOnDay entities.ScheduleOnDay
	var buffer bytes.Buffer

	query := querySearchAllScheduleOnDay(date)

	err := json.NewEncoder(&buffer).Encode(query)
	if err != nil {
		return scheduleOnDay, fmt.Errorf("encode query failed")
	}

	resp, err := repo.es.Search(repo.es.Search.WithIndex(SCHEDULE_INDEX_NAME), repo.es.Search.WithBody(&buffer))
	if err != nil {
		return scheduleOnDay, fmt.Errorf("request to elastic search fail")
	}

	scheduleOnDay.Date = date
	scheduleOnDay.DateWithWeekday = date.Weekday().String()

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return scheduleOnDay, fmt.Errorf("decode respose from elastic search failed")
	}

	activeLeague := repo.leagueRepo.GetLeaguesNameActive()

	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {

		scheduleOnLeague := hit.(map[string]interface{})["_source"].(map[string]interface{})

		entityScheduleOnDay := newEntitiesScheduleOnLeaguesFromMap(scheduleOnLeague)
		if checkIsScheduleOnActiveLeague(activeLeague, entityScheduleOnDay.LeagueName) {
			scheduleOnDay.ScheduleOnLeagues = append(scheduleOnDay.ScheduleOnLeagues, entityScheduleOnDay)
		}
	}
	return scheduleOnDay, nil
}

func (repo *schedulesRepo) GetScheduleLeagueOnDay(date time.Time, league string) (entities.ScheduleOnDay, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	var scheduleOnDay entities.ScheduleOnDay
	var buffer bytes.Buffer

	query := querySearchScheduleLeagueOnDay(date, league)

	err := json.NewEncoder(&buffer).Encode(query)
	if err != nil {
		return scheduleOnDay, fmt.Errorf("encode query failed")
	}

	resp, err := repo.es.Search(repo.es.Search.WithIndex(SCHEDULE_INDEX_NAME), repo.es.Search.WithBody(&buffer))
	if err != nil {
		return scheduleOnDay, fmt.Errorf("request to elastic search fail")
	}

	scheduleOnDay.Date = date
	scheduleOnDay.DateWithWeekday = date.Weekday().String()

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return scheduleOnDay, fmt.Errorf("decode respose from elastic search failed")
	}

	activeLeague := repo.leagueRepo.GetLeaguesNameActive()

	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {

		scheduleOnLeague := hit.(map[string]interface{})["_source"].(map[string]interface{})

		entityScheduleOnDay := newEntitiesScheduleOnLeaguesFromMap(scheduleOnLeague)
		if checkIsScheduleOnActiveLeague(activeLeague, league) {
			scheduleOnDay.ScheduleOnLeagues = append(scheduleOnDay.ScheduleOnLeagues, entityScheduleOnDay)
		}
	}
	return scheduleOnDay, nil
}

func (repo *schedulesRepo) GetMatchURLsOnDay() entities.MatchURLsOnDay {
	return repo.matchURLsOnDay
}

func (repo *schedulesRepo) ClearMatchURLsOnDay() {
	repo.matchURLsOnDay = entities.MatchURLsOnDay{}
}
