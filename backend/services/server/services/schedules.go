package services

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"server/entities"
	pb "server/proto"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v7"
	jsoniter "github.com/json-iterator/go"
	"google.golang.org/grpc"
)

var SCHEDULE_INDEX_NAME = "schedules"

type schedulesService struct {
	conn *grpc.ClientConn
	es   *elasticsearch.Client
}

func NewSchedulesService(conn *grpc.ClientConn, es *elasticsearch.Client) *schedulesService {
	schedulesService := &schedulesService{
		conn: conn,
		es:   es,
	}
	return schedulesService
}

func (s *schedulesService) GetSchedules() {
	client := pb.NewCrawlerServiceClient(s.conn)

	date := time.Now().Format("02-01-2006")

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
	elasticSchedules := PbSchedulesToScheduleElastic(pbSchedules)
	for _,schedule := range elasticSchedules {
		exist := checkScheduleWithElasticSearch(schedule, s.es)
		if !exist {
			storeScheduleElasticsearch(schedule, s.es)
		}
		
	}
}

func (s *schedulesService)APIGetScheduleOnDay(date time.Time) (entities.ScheduleOnDay, error){
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	var scheduleOnDay entities.ScheduleOnDay
	var buffer bytes.Buffer

	query := querySearchScheduleOnDay(date)

	json.NewEncoder(&buffer).Encode(query)
	resp, err := s.es.Search(s.es.Search.WithIndex(SCHEDULE_INDEX_NAME), s.es.Search.WithBody(&buffer))
	if err != nil {
		return scheduleOnDay, fmt.Errorf("request to elastic search fail")
	}

	scheduleOnDay.Date = date
	scheduleOnDay.DateWithWeekday = date.Weekday().String()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {

		scheduleOnLeague := hit.(map[string]interface{})["_source"].(map[string]interface{})

		scheduleOnDay.ScheduleOnLeagues = append(scheduleOnDay.ScheduleOnLeagues, newEntitiesScheduleOnLeaguesFromMap(scheduleOnLeague))
	}
	return scheduleOnDay, nil
}

func querySearchScheduleOnDay(dateISO8601 time.Time) map[string]interface{} {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"match": map[string]interface{}{
						"date":  dateISO8601,
					},
				},
			},
		},
	}
	return query
}

func newEntitiesScheduleOnLeaguesFromMap(respScheduleOnLeague map[string]interface{}) entities.ScheduleOnLeague {
	// type assertion []interface{} to []interface {}
	matchsInterface := respScheduleOnLeague["matchs"].([]interface{})
	
	matchs := make([]entities.Match, len(matchsInterface))
	for i, matchInterface := range matchsInterface {
		matchParse := matchInterface.(map[string]interface {})
		matchs[i].Time = matchParse["time"].(string)
		matchs[i].Round = matchParse["round"].(string)
		matchs[i].Scores = matchParse["scores"].(string)
		matchs[i].MatchDetailLink = matchParse["match_detail_id"].(string)
		matchs[i].Scores = matchParse["scores"].(string)

		club1 := matchParse["club_1"].(map[string]interface {})
		club2 := matchParse["club_2"].(map[string]interface {})
		matchs[i].Club1.Name = club1["name"].(string)
		matchs[i].Club1.Logo = club1["logo"].(string)
		matchs[i].Club2.Name = club2["name"].(string)
		matchs[i].Club1.Logo = club1["logo"].(string)
	}
	scheduleOnLeague := entities.ScheduleOnLeague{
		LeagueName: respScheduleOnLeague["league_name"].(string),
		Matchs: matchs,
	}
	return scheduleOnLeague
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
		log.Println("Document already exist in index", SCHEDULE_INDEX_NAME)
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

func PbSchedulesToScheduleElastic(pbSchedule *pb.SchedulesReponse) []entities.ScheduleElastic {
	schedules := make([]entities.ScheduleElastic, len(pbSchedule.GetScheduleOnLeagues()))
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0,0,0,0, time.UTC)
	for _, scheduleOnLeagueResp := range pbSchedule.GetScheduleOnLeagues() {
		schedule := entities.ScheduleElastic{
			Date: today,
			LeagueName: strings.TrimSpace(scheduleOnLeagueResp.GetLeagueName()),
		}
		for _, matchResp := range scheduleOnLeagueResp.GetMatchs() {
			match := entities.Match{
				Time: strings.TrimSpace(matchResp.Time) ,
				Round: strings.TrimSpace(matchResp.Round),
				Club1: entities.Club{
					Name: strings.TrimSpace(matchResp.Club1.Name) ,
					Logo: strings.TrimSpace(matchResp.Club1.Logo) ,
				},
				Club2: entities.Club{
					Name: strings.TrimSpace(matchResp.Club2.Name),
					Logo: strings.TrimSpace(matchResp.Club2.Logo) ,
				},
				Scores: strings.TrimSpace(matchResp.Scores),
				MatchDetailLink: strings.TrimSpace(matchResp.MatchDetailLink),
			}
			schedule.Matchs = append(schedule.Matchs, match)
		}
		schedules = append(schedules, schedule)
	}
	return schedules
}
