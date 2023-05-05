package services

import (
	"crawler/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrawlSchedules_Success(t *testing.T) {
	assert := assert.New(t)

	dateString := "01-05-2023"

	scheduleClasses := entities.HtmlSchedulesClass{
		LeagueClass: "football-header",
		Date: "table-header text-center",
		HtmlMatchClass: entities.HtmlMatchClass{
			MatchClass: "football-match-livescore",
			Time: "columns-time",
			Round: "vongbang",
			Scores: "soccer-scores",
			MatchDetailLink: "columns-detail",
			Club1: entities.HtmlClubClass{
				Name: "name-club club1",
				Logo: "logo-club",
			},
			Club2: entities.HtmlClubClass{
				Name: "name-club club2",
				Logo: "logo-club",
			},
		},
	}

	schedules, err := CrawlSchedules(dateString, scheduleClasses)
	assert.Nil(err)
	assert.Equal(schedules.Date, dateString)
}