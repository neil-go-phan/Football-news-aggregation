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
	tagsService    *tagsService
}

func NewSchedulesService(leagues *leaguesService, tags *tagsService, conn *grpc.ClientConn, es *elasticsearch.Client) *schedulesService {
	schedulesService := &schedulesService{
		conn:           conn,
		es:             es,
		leaguesService: leagues,
		tagsService:    tags,
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
	for _, schedule := range elasticSchedules {
		wg.Add(1)
		go func(schedule entities.ScheduleElastic) {
			defer wg.Done()
			exist := checkScheduleWithElasticSearch(schedule, s.es)
			if !exist {
				storeScheduleElasticsearch(schedule, s.es)
				// auto store new league
				isNewLeague := s.storeNewLeague(schedule.LeagueName)
				if isNewLeague {
					log.Println("detect a new league: ", schedule.LeagueName)
					err = s.leaguesService.WriteLeaguesJSON()
					if err != nil {
						log.Println("error occurred while overwrite leagueConfig.JSON:", err)
					}
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

func (s *schedulesService) storeNewLeague(newLeague string) bool {
	// detect new league
	if newLeague == "" {
		return false
	}
	for _, league := range s.leaguesService.leagues.Leagues {
		if newLeague == league {
			return false
		}
	}
	s.leaguesService.leagues.Leagues = append(s.leaguesService.leagues.Leagues, newLeague)
	return true
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
	schedules := make([]entities.ScheduleElastic,0)
	date, err := time.Parse("02-01-2006", pbSchedule.GetDateFormated())
	if err != nil {
		log.Println("error when parse date:", err)
		// handler err later
	}

	for _, scheduleOnLeagueResp := range pbSchedule.GetScheduleOnLeagues() {
		schedule := entities.ScheduleElastic{}
		for _, matchResp := range scheduleOnLeagueResp.GetMatchs() {
			match := entities.Match{
				Time:  strings.TrimSpace(matchResp.Time),
				Round: strings.TrimSpace(matchResp.Round),
				Club1: entities.Club{
					Name: strings.TrimSpace(matchResp.Club1.Name),
					Logo: strings.TrimSpace(matchResp.Club1.Logo),
				},
				Club2: entities.Club{
					Name: strings.TrimSpace(matchResp.Club2.Name),
					Logo: strings.TrimSpace(matchResp.Club2.Logo),
				},
				Scores:          strings.TrimSpace(matchResp.Scores),
				MatchDetailLink: strings.TrimSpace(matchResp.MatchDetailLink),
			}
			schedule.Matchs = append(schedule.Matchs, match)
		}
		schedule.Date = date;
		schedule.LeagueName = strings.TrimSpace(scheduleOnLeagueResp.GetLeagueName())
		schedules = append(schedules, schedule)
	}
	
	return schedules
}

func newEntitiesScheduleOnLeaguesFromMap(respScheduleOnLeague map[string]interface{}) entities.ScheduleOnLeague {
	// type assertion []interface{} to []interface {}
	matchsInterface := respScheduleOnLeague["matchs"].([]interface{})

	matchs := make([]entities.Match, len(matchsInterface))
	for i, matchInterface := range matchsInterface {
		matchParse := matchInterface.(map[string]interface{})
		matchs[i].Time = matchParse["time"].(string)
		matchs[i].Round = matchParse["round"].(string)
		matchs[i].Scores = matchParse["scores"].(string)
		matchs[i].MatchDetailLink = matchParse["match_detail_id"].(string)
		matchs[i].Scores = matchParse["scores"].(string)

		club1 := matchParse["club_1"].(map[string]interface{})
		club2 := matchParse["club_2"].(map[string]interface{})
		matchs[i].Club1.Name = club1["name"].(string)
		matchs[i].Club1.Logo = club1["logo"].(string)
		matchs[i].Club2.Name = club2["name"].(string)
		matchs[i].Club2.Logo = club2["logo"].(string)
	}
	scheduleOnLeague := entities.ScheduleOnLeague{
		LeagueName: respScheduleOnLeague["league_name"].(string),
		Matchs:     matchs,
	}
	return scheduleOnLeague
}
