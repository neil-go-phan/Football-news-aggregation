package matchservices

// import (
// 	"fmt"
// 	"server/entities"
// 	mock "server/services/match/mock"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// )

// func TestGetMatch(t *testing.T) {
// 	mockRepo := new(mock.MockMatchRepository)
// 	mockCrawler := new(mock.MockCrawlerServiceClient)
// 	mockClub := new(mock.MockClubServices)
// 	mockEvent := new(mock.MockEventServices)
// 	mockLineUp := new(mock.MockLineUpServices)
// 	mockOverview := new(mock.MockOverviewItemServices)
// 	mockPlayer := new(mock.MockPlayerServices)
// 	mockStats := new(mock.MockStatsItemServices)
// 	assert := assert.New(t)

// 	matchService := NewMatchService(mockCrawler, mockRepo, mockClub, mockStats, mockEvent, mockOverview, mockLineUp, mockPlayer)
// 	now := time.Now()
// 	match := &entities.Match{
// 		TimeStart: now,
// 		Club1ID:   uint(1),
// 		Club1: entities.Club{
// 			Name: "club1",
// 			Logo: "logo1",
// 		},
// 		Club2ID: uint(2),
// 		Club2: entities.Club{
// 			Name: "club2",
// 			Logo: "logo2",
// 		},
// 	}
// 	mockRepo.On("GetIDWithDateAndClubName", now, "club1", "club2").Return(match, nil)

// 	mockRepo.On("GetMatch", match).Return(match, nil)

// 	got, err := matchService.GetMatch(now, "club1", "club2")

// 	assert.Nil(err)
// 	assert.Equal(match, got)
// }

// func TestCreateMatch(t *testing.T) {
// 	mockRepo := new(mock.MockMatchRepository)
// 	mockCrawler := new(mock.MockCrawlerServiceClient)
// 	mockClub := new(mock.MockClubServices)
// 	mockEvent := new(mock.MockEventServices)
// 	mockLineUp := new(mock.MockLineUpServices)
// 	mockOverview := new(mock.MockOverviewItemServices)
// 	mockPlayer := new(mock.MockPlayerServices)
// 	mockStats := new(mock.MockStatsItemServices)
// 	assert := assert.New(t)

// 	matchService := NewMatchService(mockCrawler, mockRepo, mockClub, mockStats, mockEvent, mockOverview, mockLineUp, mockPlayer)
// 	now := time.Now()

// 	club1 := entities.Club{
// 		Name: "club1",
// 		Logo: "logo1",
// 	}
// 	club2 := entities.Club{
// 		Name: "club2",
// 		Logo: "logo2",
// 	}

// 	match := &entities.Match{
// 		TimeStart: now,
// 		Club1ID:   uint(1),
// 		Club1:     club1,
// 		Club2ID:   uint(2),
// 		Club2:     club2,
// 	}
// 	mockClub.On("GetOrCreate", club1.Name, club1.Logo).Return(&club1, nil)
// 	mockClub.On("GetOrCreate", club2.Name, club2.Logo).Return(&club1, nil)
// 	mockRepo.On("Create", match).Return(nil)

// 	err := matchService.createMatch(match)

// 	assert.Nil(err)
// }

// func TestStoreMatch_ScheduleCrawl_NewMatch(t *testing.T) {
// 	mockRepo := new(mock.MockMatchRepository)
// 	mockCrawler := new(mock.MockCrawlerServiceClient)
// 	mockClub := new(mock.MockClubServices)
// 	mockEvent := new(mock.MockEventServices)
// 	mockLineUp := new(mock.MockLineUpServices)
// 	mockOverview := new(mock.MockOverviewItemServices)
// 	mockPlayer := new(mock.MockPlayerServices)
// 	mockStats := new(mock.MockStatsItemServices)
// 	assert := assert.New(t)

// 	matchService := NewMatchService(mockCrawler, mockRepo, mockClub, mockStats, mockEvent, mockOverview, mockLineUp, mockPlayer)
// 	dayTime := time.Date(2022, 5, 4, 0, 0, 0, 0, time.UTC)

// 	club1 := entities.Club{
// 		Name: "club1",
// 		Logo: "logo1",
// 	}
// 	club2 := entities.Club{
// 		Name: "club2",
// 		Logo: "logo2",
// 	}

// 	match := &entities.Match{
// 		ScheduleID: uint(20),
// 		Time: "FT - 11/05",
// 		TimeStart: dayTime,
// 		Club1ID:   uint(1),
// 		Club1:     club1,
// 		Club2ID:   uint(2),
// 		Club2:     club2,
// 	}
// 	mockClub.On("GetOrCreate", club1.Name, club1.Logo).Return(&club1, nil)
// 	mockClub.On("GetOrCreate", club2.Name, club2.Logo).Return(&club2, nil)
// 	mockRepo.On("GetIDWithDateAndClubName", dayTime, "club1", "club2").Return(match, fmt.Errorf("not found"))
// 	mockRepo.On("Create", match).Return(nil)

// 	err := matchService.StoreMatch_ScheduleCrawl(*match, uint(20), dayTime)

// 	assert.Nil(err)
// }

// func TestStoreMatch_ScheduleCrawl_UpdateMatch(t *testing.T) {
// 	mockRepo := new(mock.MockMatchRepository)
// 	mockCrawler := new(mock.MockCrawlerServiceClient)
// 	mockClub := new(mock.MockClubServices)
// 	mockEvent := new(mock.MockEventServices)
// 	mockLineUp := new(mock.MockLineUpServices)
// 	mockOverview := new(mock.MockOverviewItemServices)
// 	mockPlayer := new(mock.MockPlayerServices)
// 	mockStats := new(mock.MockStatsItemServices)
// 	assert := assert.New(t)

// 	matchService := NewMatchService(mockCrawler, mockRepo, mockClub, mockStats, mockEvent, mockOverview, mockLineUp, mockPlayer)
// 	dayTime := time.Date(2022, 5, 4, 0, 0, 0, 0, time.UTC)

// 	club1 := entities.Club{
// 		Name: "club1",
// 		Logo: "logo1",
// 	}
// 	club2 := entities.Club{
// 		Name: "club2",
// 		Logo: "logo2",
// 	}

// 	match := &entities.Match{
// 		ScheduleID: uint(20),
// 		Time: "FT - 11/05",
// 		TimeStart: dayTime,
// 		Club1ID:   uint(1),
// 		Club1:     club1,
// 		Club2ID:   uint(2),
// 		Club2:     club2,
// 	}
// 	mockClub.On("GetOrCreate", club1.Name, club1.Logo).Return(&club1, nil)
// 	mockClub.On("GetOrCreate", club2.Name, club2.Logo).Return(&club2, nil)
// 	mockRepo.On("GetIDWithDateAndClubName", dayTime, "club1", "club2").Return(match, nil)
// 	mockRepo.On("UpdateWhenScheduleCrawl", match).Return(nil)

// 	err := matchService.StoreMatch_ScheduleCrawl(*match, uint(20), dayTime)

// 	assert.Nil(err)
// }