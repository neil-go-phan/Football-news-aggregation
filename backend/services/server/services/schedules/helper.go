package schedulesservices

import (
	"fmt"
	"server/entities"
	pb "server/proto"
	"server/repository"
	"strconv"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
)

func newEntitiesSchedule(pbSchedule *pb.ScheduleOnLeague, date time.Time) entities.Schedule {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	schedule := entities.Schedule{}
	scheduleByte, err := json.Marshal(pbSchedule)
	if err != nil {
		log.Errorf("error occrus when marshal crawled schedules: %s", err)
	}
	err = json.Unmarshal(scheduleByte, &schedule)
	if err != nil {
		log.Errorf("error occrus when unmarshal crawled schedules to proto.Schedules: %s", err)
	}
	schedule.Date = date
	return schedule
}

func newScheduleOnLeague(entityScheduel entities.Schedule) repository.ScheduleOnLeague {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	schedule := repository.ScheduleOnLeague{}
	scheduleByte, err := json.Marshal(entityScheduel)
	if err != nil {
		log.Errorf("error occrus when marshal crawled schedules: %s", err)
	}
	err = json.Unmarshal(scheduleByte, &schedule)
	if err != nil {
		log.Errorf("error occrus when unmarshal crawled schedules to proto.Schedules: %s", err)
	}

	return schedule
}

func readTime(match entities.Match, dayTime time.Time) (time.Time, error) {
	exactTime := dayTime
	timeStr := strings.Split(match.Time, "-")
	if strings.Trim(timeStr[0], " ") == "FT" {
		return exactTime, fmt.Errorf("the match is already end")
	}
	timeDetail := strings.Split(timeStr[0], ":")
	hours, err := strconv.Atoi(strings.Trim(timeDetail[0], " "))
	if err != nil {
		return exactTime, fmt.Errorf("can not parse hour to set cronjob err: %s", err)
	}
	mins, err := strconv.Atoi(strings.Trim(timeDetail[1], " "))
	if err != nil {
		return exactTime, fmt.Errorf("can not parse minute to set cronjob err: %s", err)
	}
	dateDetail := strings.Split(timeStr[1], "/")
	day, err := strconv.Atoi(strings.Trim(dateDetail[0], " "))
	if err != nil {
		return exactTime, fmt.Errorf("can not parse day to set cronjob err: %s", err)
	}
	month, err := strconv.Atoi(strings.Trim(dateDetail[1], " "))
	if err != nil {
		return exactTime, fmt.Errorf("can not parse month to set cronjob err: %s", err)
	}
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		log.Error(err)
	}
	time1 := time.Date(dayTime.Year(), time.Month(month), day, hours, mins, 0, 0, loc)

	return time1, nil
}

func getMatchUrl(schedule entities.Schedule) []string {
	matchUrls := make([]string, 0)
	for _, match := range schedule.Matches {
		url := fmt.Sprintf("https://bongda24h.vn%s", match.MatchDetailLink)
		matchUrls = append(matchUrls, url)
	}
	return matchUrls
}

func addMatchUrl(time time.Time, inputurl string, matchsOnTime *repository.MatchURLsWithTimeOnDay) {
	url := fmt.Sprintf("https://bongda24h.vn%s", inputurl)
	for index := range matchsOnTime.MatchsOnTimes {
		if matchsOnTime.MatchsOnTimes[index].Date == time {
			matchsOnTime.MatchsOnTimes[index].Urls = append(matchsOnTime.MatchsOnTimes[index].Urls, url)
			return
		}
	}
	newMatchOnTimes := repository.MatchURLsOnTime{
		Date: time,
		Urls: []string{url},
	}
	matchsOnTime.MatchsOnTimes = append(matchsOnTime.MatchsOnTimes, newMatchOnTimes)
}

func checkIsScheduleOnActiveLeague(activeLeaguesName []string, scheduleLeagueName string) bool {
	for _, name := range activeLeaguesName {
		if name == scheduleLeagueName {
			return true
		}
	}
	return false
}

func isNewLeague(leagues []entities.League, newLeagueName string) bool {
	// detect new league
	if newLeagueName == "" {
		return false
	}
	for _, league := range leagues {
		if newLeagueName == league.LeagueName {
			return false
		}
	}
	return true
}
