package services

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"server/entities"
	pb "server/proto"
	"strings"
	"sync"
	"time"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v7"
	jsoniter "github.com/json-iterator/go"
	"google.golang.org/grpc"
)

var SCHEDULE_INDEX_NAME = "schedules"

type schedulesService struct {
	conn           *grpc.ClientConn
	es             *elasticsearch.Client
	leaguesService *leaguesService
	matchURLsOnDay entities.MatchURLsOnDay
}

func NewSchedulesService(leagues *leaguesService, conn *grpc.ClientConn, es *elasticsearch.Client) *schedulesService {
	schedulesService := &schedulesService{
		conn:           conn,
		es:             es,
		leaguesService: leagues,
	}
	return schedulesService
}

func (s *schedulesService) GetSchedules(date string) {
	client := pb.NewCrawlerServiceClient(s.conn)

	in := &pb.Date{
		Date: date,
	}
	// send gRPC request to crawler
	pbSchedules, err := client.GetSchedulesOnDay(context.Background(), in)
	if err != nil {
		log.Printf("error occurred while get schedule on day from crawler error %v \n", err)
		return
	}
	// store schedule in elastic search
	var wg sync.WaitGroup
	elasticSchedules := PbSchedulesToScheduleElastic(pbSchedules)
	dateCasted, err := time.Parse("02-01-2006", pbSchedules.GetDateFormated())
	if err != nil {
		log.Println("error when parse date:", err)
		// handler err later
	}
	s.matchURLsOnDay.Date = dateCasted
	for _, schedule := range elasticSchedules {
		wg.Add(1)
		go func(schedule entities.ScheduleElastic) {
			defer wg.Done()

			// push matchUrl on day
			matchUrls := getMatchUrlOnDay(schedule)
			s.matchURLsOnDay.Urls = append(s.matchURLsOnDay.Urls, matchUrls...)
			exist := checkScheduleWithElasticSearch(schedule, s.es)

			if !exist {
				isExisteague := s.checkIfLeagueExist(schedule.LeagueName)
				
				if isExisteague {
					storeScheduleElasticsearch(schedule, s.es)
				}
			}

		}(schedule)

	}
	wg.Wait()
}

func (s *schedulesService) APIGetScheduleOnDay(date time.Time) (entities.ScheduleOnDay, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	var scheduleOnDay entities.ScheduleOnDay
	var buffer bytes.Buffer

	query := querySearchScheduleOnDay(date)

	err := json.NewEncoder(&buffer).Encode(query)
	if err != nil {
		return scheduleOnDay, fmt.Errorf("encode query failed")
	}

	resp, err := s.es.Search(s.es.Search.WithIndex(SCHEDULE_INDEX_NAME), s.es.Search.WithBody(&buffer))
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
	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {

		scheduleOnLeague := hit.(map[string]interface{})["_source"].(map[string]interface{})

		entityScheduleOnDay := newEntitiesScheduleOnLeaguesFromMap(scheduleOnLeague)
		scheduleOnDay.ScheduleOnLeagues = append(scheduleOnDay.ScheduleOnLeagues, entityScheduleOnDay)

	}
	return scheduleOnDay, nil
}

func querySearchScheduleOnDay(dateISO8601 time.Time) map[string]interface{} {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"match": map[string]interface{}{
						"date": dateISO8601,
					},
				},
			},
		},
	}
	return query
}

func checkScheduleWithElasticSearch(schedule entities.ScheduleElastic, es *elasticsearch.Client) bool {
	req := esapi.ExistsRequest{
		Index:      SCHEDULE_INDEX_NAME,
		DocumentID: strings.ToLower(fmt.Sprintf("$DATE=%s,$LEAGUE=%s", schedule.Date.Format("02-01-2006"), schedule.LeagueName)),
	}

	resp, err := req.Do(context.Background(), es)
	if err != nil {
		log.Printf("Error checking if document exists: %s\n", err)
		return false
	}

	status := resp.StatusCode
	if status == 200 {
		log.Println("Document already exist in index", strings.ToLower(fmt.Sprintf("$DATE=%s,$LEAGUE=%s", schedule.Date.Format("02-01-2006"), schedule.LeagueName)))
		return true
	} else if status == 404 {
		log.Printf("Document not found in index %s, creating new one...", SCHEDULE_INDEX_NAME)
		return false
	}

	return false
}

func storeScheduleElasticsearch(schedule entities.ScheduleElastic, es *elasticsearch.Client) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	body, err := json.Marshal(schedule)
	if err != nil {
		log.Printf("Error encoding schedule: %s\n", err)
	}
	req := esapi.IndexRequest{
		Index:      SCHEDULE_INDEX_NAME,
		DocumentID: strings.ToLower(fmt.Sprintf("$DATE=%s,$LEAGUE=%s", schedule.Date.Format("02-01-2006"), schedule.LeagueName)),
		Body:       strings.NewReader(string(body)),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Printf("Error getting response: %s\n", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error indexing document\n", res.Status())
	} else {
		log.Printf("[%s] Indexed document with index: %s \n", res.Status(), SCHEDULE_INDEX_NAME)
	}
}

func getMatchUrlOnDay(schedule entities.ScheduleElastic) []string {
	matchUrls := make([]string, 0)
	for _, match := range schedule.Matchs {
		url := fmt.Sprintf("https://bongda24h.vn%s", match.MatchDetailLink)
		matchUrls = append(matchUrls, url)
	}
	return matchUrls
}

func (s *schedulesService) GetMatchURLsOnDay() entities.MatchURLsOnDay {
	return s.matchURLsOnDay
}

func (s *schedulesService) ClearMatchURLsOnDay() {
	s.matchURLsOnDay = entities.MatchURLsOnDay{}
}

func (s *schedulesService) checkIfLeagueExist(checkLeague string) bool {
	// detect new league
	if checkLeague == "" {
		return false
	}
	for _, league := range s.leaguesService.leagues.Leagues {
		if strings.EqualFold(checkLeague, league) {
			return true
		}
	}
	return false
}

// func (s *schedulesService) storeNewTag(newTags string) bool {
// 	if newTags == "" {
// 		return false
// 	}
// 	// detect new league
// 	tagFormat := helper.FormatVietnamese(newTags)
// 	for _, tag := range s.tagsService.Tags.Tags {
// 		if tagFormat == tag {
// 			return false
// 		}
// 	}
// 	s.tagsService.Tags.Tags = append(s.tagsService.Tags.Tags, tagFormat)
// 	return true
// }

func PbSchedulesToScheduleElastic(pbSchedule *pb.SchedulesReponse) []entities.ScheduleElastic {
	schedules := make([]entities.ScheduleElastic, 0)
	date, err := time.Parse("02-01-2006", pbSchedule.GetDateFormated())
	if err != nil {
		log.Println("error when parse date:", err)
		// handler err later
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	for _, scheduleOnLeagueResp := range pbSchedule.GetScheduleOnLeagues() {
		schedule := entities.ScheduleElastic{}
		scheduleByte, err := json.Marshal(scheduleOnLeagueResp)
		if err != nil {
			log.Printf("error occrus when marshal crawled schedules: %s", err)
		}
		err = json.Unmarshal(scheduleByte, &schedule)
		if err != nil {
			log.Printf("error occrus when unmarshal crawled schedules to proto.Schedules: %s", err)
		}
		schedule.Date = date
		schedules = append(schedules, schedule)
	}

	return schedules
}

func newEntitiesScheduleOnLeaguesFromMap(respScheduleOnLeague map[string]interface{}) entities.ScheduleOnLeague {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	scheduleOnLeague := entities.ScheduleOnLeague{}
	scheduleByte, err := json.Marshal(respScheduleOnLeague)
	if err != nil {
		log.Printf("error occrus when marshal elastic response schedules: %s", err)
	}
	err = json.Unmarshal(scheduleByte, &scheduleOnLeague)
	if err != nil {
		log.Printf("error occrus when unmarshal elastic response to entity schedules: %s", err)
	}

	return scheduleOnLeague
}
