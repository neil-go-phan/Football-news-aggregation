package matchservices

import (
	"fmt"
	"server/entities"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

func readTimeStart(match entities.Match, dayTime time.Time) (time.Time, error) {
	exactTime := dayTime
	timeStr := strings.Split(match.Time, "-")
	if strings.Trim(timeStr[0], " ") == "FT" {
		return exactTime, nil
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

	exactTime = exactTime.Add(time.Hour*time.Duration(hours) + time.Minute*time.Duration(mins))

	if hours == 0 && mins == 0 {
		exactTime = exactTime.AddDate(0, 0, 1)
	}

	return exactTime, nil
}

// change time string crawled to time.Time. example" readTimeEvent("45+1'", 23h) -> 23h46p
func readTimeEvent(timeString string, dayTime time.Time) (time.Time, error) {
	exactTime := dayTime
	timeStr := strings.ReplaceAll(timeString, "'", "")
	timeDetail := strings.Split(timeStr, "+")
	mins, err := strconv.Atoi(strings.Trim(timeDetail[0], " "))
	if err != nil {
		return exactTime, fmt.Errorf("can not parse min to set event time err: %s", err)
	}
	extraTime := 0
	if len(timeDetail) != 1 {
		extraTime, err = strconv.Atoi(strings.Trim(timeDetail[1], " "))
		if err != nil {
			log.Errorf("can not parse extra time to set cronjob err: %s", err)
		}
	}

	exactTime = exactTime.Add(time.Minute*time.Duration(mins + extraTime))

	return exactTime, nil
}