package schedulesrepo

import (
	"context"
	"fmt"
	"server/entities"
	serverhelper "server/helper"
	pb "server/proto"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v7"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
)

func getMatchUrlOnDay(schedule entities.ScheduleElastic) []string {
	matchUrls := make([]string, 0)
	for _, match := range schedule.Matchs {
		url := fmt.Sprintf("https://bongda24h.vn%s", match.MatchDetailLink)
		matchUrls = append(matchUrls, url)
	}
	return matchUrls
}

func querySearchAllScheduleOnDay(dateISO8601 time.Time) map[string]interface{} {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": map[string]interface{}{
					"term": map[string]interface{}{
						"date": dateISO8601,
					},
				},
			},
		},
		"size": 1000,
	}
	return query
}

func querySearchScheduleLeagueOnDay(dateISO8601 time.Time, league string) map[string]interface{} {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": []map[string]interface{}{
					{
						"term": map[string]interface{}{
							"league_name.keyword": league,
						},
					},
					{
						"term": map[string]interface{}{
							"date": dateISO8601,
						},
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
		log.Errorf("Error checking if document exists: %s\n", err)
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
		log.Errorf("Error encoding schedule: %s\n", err)
	}
	req := esapi.IndexRequest{
		Index:      SCHEDULE_INDEX_NAME,
		DocumentID: strings.ToLower(fmt.Sprintf("$DATE=%s,$LEAGUE=%s", schedule.Date.Format("02-01-2006"), schedule.LeagueName)),
		Body:       strings.NewReader(string(body)),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Errorf("Error getting response: %s\n", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error indexing document\n", res.Status())
	} else {
		log.Printf("[%s] Indexed document with index: %s \n", res.Status(), SCHEDULE_INDEX_NAME)
	}
}

func PbSchedulesToScheduleElastic(pbSchedule *pb.SchedulesReponse) []entities.ScheduleElastic {
	schedules := make([]entities.ScheduleElastic, 0)
	date, err := time.Parse("02-01-2006", pbSchedule.GetDateFormated())
	if err != nil {
		log.Errorln("error when parse date:", err)
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	for _, scheduleOnLeagueResp := range pbSchedule.GetScheduleOnLeagues() {
		schedule := entities.ScheduleElastic{}
		scheduleByte, err := json.Marshal(scheduleOnLeagueResp)
		if err != nil {
			log.Errorf("error occrus when marshal crawled schedules: %s", err)
		}
		err = json.Unmarshal(scheduleByte, &schedule)
		if err != nil {
			log.Errorf("error occrus when unmarshal crawled schedules to proto.Schedules: %s", err)
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
		log.Errorf("error occrus when marshal elastic response schedules: %s", err)
	}
	err = json.Unmarshal(scheduleByte, &scheduleOnLeague)
	if err != nil {
		log.Errorf("error occrus when unmarshal elastic response to entity schedules: %s", err)
	}

	return scheduleOnLeague
}

func checkIsScheduleOnActiveLeague(activeLeaguesName []string, scheduleLeagueName string) bool {
	for _, name := range activeLeaguesName {
		if name == scheduleLeagueName {
			return true
		}
	}
	return false
}

func isNewLeague(leagues []entities.League, newLeaegueName string) bool {
	// detect new league
	if newLeaegueName == "" {
		return false
	}
	for _, league := range leagues {
		if newLeaegueName == league.LeagueName {
			return false
		}
	}
	return true
}

func  isLeagueActive(leagues []entities.League, leaegueName string) bool {
	// detect new league
	for _, league := range leagues  {
		if leaegueName == league.LeagueName && league.Active {
			return true
		}
	}
	return false
}

func isLeagueTagExist(tags []string, newTag string) bool {
	// detect new tag
	tagFormated := serverhelper.FormatVietnamese(newTag)
	if tagFormated == "" {
		return true
	}
	for _, tag := range tags {
		if tag == tagFormated {
			return true
		}
	}
	return false
}
