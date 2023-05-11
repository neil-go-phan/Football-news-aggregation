package schedulesservices

// import (
// 	"testing"
// 	"time"

// 	"server/entities"
// 	mock "server/services/schedules/mock"

// 	"github.com/stretchr/testify/assert"
// )

// func TestGetAllScheduleLeagueOnDay(t *testing.T) {
// 	mockRepoMatchDetail := new(mock.MockMatchDetailRepository)
// 	mockRepoSchedule := new(mock.MockSchedulesRepository)
// 	service := NewSchedulesService(mockRepoSchedule, mockRepoMatchDetail)
// 	assert := assert.New(t)

// 	dateString := "2021-11-22"
// 	date, _ := time.Parse("2006-01-02", dateString)

// 	want := entities.ScheduleOnDay{
// 		Date:              date,
// 		DateWithWeekday:   "Monday",
// 		ScheduleOnLeagues: []entities.ScheduleOnLeague{},
// 	}

// 	mockRepoSchedule.On("GetAllScheduleLeagueOnDay", date).Return(want, nil)

// 	got, err := service.GetAllScheduleLeagueOnDay(date)
// 	assert.Nil(err, "no error")
// 	assert.Equal(want, got)
// }

// func TestGetScheduleLeagueOnDay(t *testing.T) {
// 	mockRepoMatchDetail := new(mock.MockMatchDetailRepository)
// 	mockRepoSchedule := new(mock.MockSchedulesRepository)
// 	service := NewSchedulesService(mockRepoSchedule, mockRepoMatchDetail)
// 	assert := assert.New(t)

// 	dateString := "2021-11-22"
// 	date, _ := time.Parse("2006-01-02", dateString)

// 	want := entities.ScheduleOnDay{
// 		Date:              date,
// 		DateWithWeekday:   "Monday",
// 		ScheduleOnLeagues: []entities.ScheduleOnLeague{},
// 	}

// 	mockRepoSchedule.On("GetScheduleLeagueOnDay", date, "league 1").Return(want, nil)

// 	got, err := service.GetScheduleLeagueOnDay(date, "league 1")
// 	assert.Nil(err, "no error")
// 	assert.Equal(want, got)
// }

// func TestGetMatchURLsOnDay(t *testing.T) {
// 	mockRepoMatchDetail := new(mock.MockMatchDetailRepository)
// 	mockRepoSchedule := new(mock.MockSchedulesRepository)
// 	service := NewSchedulesService(mockRepoSchedule, mockRepoMatchDetail)
// 	assert := assert.New(t)

// 	dateString := "2021-11-22"
// 	date, _ := time.Parse("2006-01-02", dateString)

// 	want := entities.MatchURLsOnDay{
// 		Date: date,
// 		Urls: []string{},
// 	}

// 	mockRepoSchedule.On("GetMatchURLsOnDay").Return(want, nil)

// 	got := service.GetMatchURLsOnDay()
// 	assert.Equal(want, got)
// }

// func TestGetMatchURLsOnTime(t *testing.T) {
// 	mockRepoMatchDetail := new(mock.MockMatchDetailRepository)
// 	mockRepoSchedule := new(mock.MockSchedulesRepository)
// 	service := NewSchedulesService(mockRepoSchedule, mockRepoMatchDetail)
// 	assert := assert.New(t)

// 	want := entities.MatchURLsWithTimeOnDay{
// 		MatchsOnTimes: []entities.MatchURLsOnTime{},
// 	}

// 	mockRepoSchedule.On("GetMatchURLsOnTime").Return(want, nil)

// 	got := service.GetMatchURLsOnTime()
// 	assert.Equal(want, got)
// }