package matchdetailservices

import (
	"server/entities"
	"testing"
	"time"

	mock "server/services/matchDetail/mock"

	"github.com/stretchr/testify/assert"
)

func TestGetMatchDetail(t *testing.T) {
	mockRepoMatchDetail := new(mock.MockMatchDetailRepository)
	service := NewMatchDetailervice(mockRepoMatchDetail)
	assert := assert.New(t)

	want := entities.MatchDetail{
		MatchDetailTitle: entities.MatchDetailTitle{
			MatchScore: "1-1",
			Club1: entities.Club{
				Name: "Team 1",
				Logo: "Logo1",
			},
			Club2: entities.Club{
				Name: "Team 2",
				Logo: "Logo2",
			},
		},
		MatchOverview: entities.MatchOverview{
			Club1Overview: []entities.OverviewItem{
				{Info: "player 1 goal", ImageType: "ghi ban", Time: "20"},
				{Info: "player 2 goal", ImageType: "ghi ban", Time: "50"},
			},
			Club2Overview: []entities.OverviewItem{
				{Info: "player 3 goal", ImageType: "ghi ban", Time: "70"},
				{Info: "player 4 goal", ImageType: "ghi ban", Time: "60"},
			},
		},
		MatchStatistics: entities.MatchStatistics{
			Statistics: []entities.StatisticsItem{
				{
					StatClub1:   "60",
					StatContent: "Giu bong",
					StatClub2:   "40",
				},
			},
		},
		MatchLineup:   entities.MatchLineup{},
		MatchProgress: entities.MatchProgress{Events: []entities.MatchEvent{{Time: "20", Content: "yeah"}}},
	}
	dateString := "2021-11-22"
	date, _ := time.Parse("2006-01-02", dateString)
	mockRepoMatchDetail.On("GetMatchDetail", date, "Team 1", "Team 2").Return(want, nil)

	got, err := service.GetMatchDetail(date, "Team 1", "Team 2")

	assert.Equal(want, got)
	assert.Nil(err, "no error")
}
