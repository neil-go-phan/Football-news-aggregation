package schedulesservices

import (
	"server/entities"
	pb "server/proto"
	"server/repository"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetMatchUrl(t *testing.T) {
	schedule := entities.Schedule{
		Date:       time.Now(),
		LeagueName: "Premier League",
		Matches: []entities.Match{
			{
				Time:            "21:00",
				Round:           "1",
				Club1:           entities.Club{Name: "Manchester United"},
				Club2:           entities.Club{Name: "Manchester City"},
				Scores:          "1-2",
				MatchDetailLink: "/premier-league/man-utd-vs-man-city-123",
			},
			{
				Time:            "19:00",
				Round:           "1",
				Club1:           entities.Club{Name: "Chelsea"},
				Club2:           entities.Club{Name: "Liverpool"},
				Scores:          "2-1",
				MatchDetailLink: "/premier-league/chelsea-vs-liverpool-456",
			},
		},
	}

	expected := []string{
		"https://bongda24h.vn/premier-league/man-utd-vs-man-city-123",
		"https://bongda24h.vn/premier-league/chelsea-vs-liverpool-456",
	}

	matchUrls := getMatchUrl(schedule)

	assert := assert.New(t)
	assert.Equal(expected, matchUrls, "Expected matchUrls to be %v but got %v", expected, matchUrls)
}

func TestReadTime(t *testing.T) {
	dayTime := time.Date(2022, 5, 4, 0, 0, 0, 0, time.UTC)
	assert := assert.New(t)
	match := entities.Match{
		Time: "FT - 1:0",
	}

	// Kiểm tra trường hợp match đã kết thúc
	exactTime, err := readTime(match, dayTime)
	assert.Error(err)

	if !strings.Contains(err.Error(), "the match is already end") {
		t.Errorf("Expected error message: 'the match is already end', got: %s", err.Error())
	}
	assert.Equal(exactTime, dayTime, "Expected exactTime to be %s, got %s", dayTime.String(), exactTime.String())

	match = entities.Match{
		Time: "11:30 - 04/05",
	}

	// Kiểm tra trường hợp match chưa diễn ra
	exactTime, err = readTime(match, dayTime)
	assert.Nil(err)

	expectedTime := time.Date(2022, 5, 4, 11, 30, 0, 0, time.UTC)
	assert.Equal(expectedTime, exactTime, "Expected expectedTime to be %s, got %s", exactTime.String(), expectedTime.String())

	// Kiểm tra trường hợp time hour không hợp lệ

	match = entities.Match{
		Time: "invalid",
	}
	_, err = readTime(match, dayTime)
	assert.Error(err)

	// Kiểm tra trường hợp time min không hợp lệ

	match = entities.Match{
		Time: "11:invaild - 04/05",
	}
	_, err = readTime(match, dayTime)
	assert.Error(err)
}

func TestAddMatchUrl_AddsNewMatchOnTime(t *testing.T) {
	date := time.Date(2023, 5, 3, 0, 0, 0, 0, time.UTC)
	inputUrl := "/match-1"
	matchUrlsWithTimeOnDay := &repository.MatchURLsWithTimeOnDay{}

	addMatchUrl(date, inputUrl, matchUrlsWithTimeOnDay)
	assert := assert.New(t)

	assert.Equal(1, len(matchUrlsWithTimeOnDay.MatchsOnTimes))
	assert.Equal(date, matchUrlsWithTimeOnDay.MatchsOnTimes[0].Date)
	assert.Equal([]string{"https://bongda24h.vn/match-1"}, matchUrlsWithTimeOnDay.MatchsOnTimes[0].Urls)
}

func TestAddMatchUrl_AddsNewMatchOnSameTime(t *testing.T) {
	date := time.Date(2023, 5, 3, 0, 0, 0, 0, time.UTC)
	inputUrl := "/match-1"
	matchUrlsWithTimeOnDay := &repository.MatchURLsWithTimeOnDay{
		MatchsOnTimes: []repository.MatchURLsOnTime{
			{
				Date: date,
				Urls: []string{"https://bongda24h.vn/match-1"},
			},
			{
				Date: date,
				Urls: []string{"https://bongda24h.vn/match-2"},
			},
		},
	}

	// Act
	addMatchUrl(date, inputUrl, matchUrlsWithTimeOnDay)

	assert := assert.New(t)

	assert.Equal(2, len(matchUrlsWithTimeOnDay.MatchsOnTimes))
	assert.Equal(date, matchUrlsWithTimeOnDay.MatchsOnTimes[0].Date)
	assert.Equal([]string{"https://bongda24h.vn/match-1","https://bongda24h.vn/match-1"}, matchUrlsWithTimeOnDay.MatchsOnTimes[0].Urls)
}

func TestCheckIsScheduleOnActiveLeague(t *testing.T) {
	activeLeagues := []string{"League A", "League B", "League C"}
	assert := assert.New(t)
	// Test case 1: scheduleLeagueName is in activeLeagues
	result := checkIsScheduleOnActiveLeague(activeLeagues, "League B")
	assert.True(result)

	// Test case 2: scheduleLeagueName is not in activeLeagues
	result = checkIsScheduleOnActiveLeague(activeLeagues, "League D")
	assert.False(result)

	// Test case 3: activeLeagues is empty
	activeLeagues = []string{}
	result = checkIsScheduleOnActiveLeague(activeLeagues, "League A")
	assert.False(result)

	// Test case 4: scheduleLeagueName is empty
	activeLeagues = []string{"League A", "League B", "League C"}
	result = checkIsScheduleOnActiveLeague(activeLeagues, "")
	assert.False(result)
}

func TestIsNewLeague(t *testing.T) {
	// Test case 1: newLeaegueName is empty
	leagues := []entities.League{
		{LeagueName: "League A"},
		{LeagueName: "League B"},
		{LeagueName: "League C"},
	}
	assert := assert.New(t)

	result := isNewLeague(leagues, "")
	assert.False(result)

	// Test case 2: newLeaegueName is not in leagues
	result = isNewLeague(leagues, "League D")
	assert.True(result)

	// Test case 3: newLeaegueName is in leagues
	result = isNewLeague(leagues, "League B")
	assert.False(result)

	// Test case 4: leagues is empty
	leagues = []entities.League{}
	result = isNewLeague(leagues, "League A")
	assert.True(result)
}

func TestNewEntitiesSchedule(t *testing.T) {
	now := time.Now()
	assert := assert.New(t)
	schedule := entities.Schedule{
		Date:       now,
		LeagueName: "Premier League",
		Matches: []entities.Match{
			{
				Time:            "21:00",
				Round:           "1",
				Club1:           entities.Club{Name: "Manchester United"},
				Club2:           entities.Club{Name: "Manchester City"},
				Scores:          "1-2",
				MatchDetailLink: "/premier-league/man-utd-vs-man-city-123",
			},
			{
				Time:            "19:00",
				Round:           "1",
				Club1:           entities.Club{Name: "Chelsea"},
				Club2:           entities.Club{Name: "Liverpool"},
				Scores:          "2-1",
				MatchDetailLink: "/premier-league/chelsea-vs-liverpool-456",
			},
		},
	}
	schedulepb := &pb.ScheduleOnLeague{
		LeagueName: "Premier League",
		Matches: []*pb.Match{
			{
				Time:            "21:00",
				Round:           "1",
				Club_1:           &pb.Club{Name: "Manchester United"},
				Club_2:           &pb.Club{Name: "Manchester City"},
				Scores:          "1-2",
				MatchDetailLink: "/premier-league/man-utd-vs-man-city-123",
			},
			{
				Time:            "19:00",
				Round:           "1",
				Club_1:           &pb.Club{Name: "Chelsea"},
				Club_2:           &pb.Club{Name: "Liverpool"},
				Scores:          "2-1",
				MatchDetailLink: "/premier-league/chelsea-vs-liverpool-456",
			},
		},
	}

	got := newEntitiesSchedule(schedulepb, now)

	assert.Equal(schedule, got)
}

func TestNewScheduleOnLeague(t *testing.T) {
	now := time.Now()
	assert := assert.New(t)
	schedule := entities.Schedule{
		Date:       now,
		LeagueName: "Premier League",
		Matches: []entities.Match{
			{
				Time:            "21:00",
				Round:           "1",
				Club1:           entities.Club{Name: "Manchester United"},
				Club2:           entities.Club{Name: "Manchester City"},
				Scores:          "1-2",
				MatchDetailLink: "/premier-league/man-utd-vs-man-city-123",
			},
			{
				Time:            "19:00",
				Round:           "1",
				Club1:           entities.Club{Name: "Chelsea"},
				Club2:           entities.Club{Name: "Liverpool"},
				Scores:          "2-1",
				MatchDetailLink: "/premier-league/chelsea-vs-liverpool-456",
			},
		},
	}
	scheduleLeague := repository.ScheduleOnLeague{
		LeagueName: "Premier League",
		Matches: []repository.Match{
			{
				Time:            "21:00",
				Round:           "1",
				Club1:           repository.Club{Name: "Manchester United"},
				Club2:           repository.Club{Name: "Manchester City"},
				Scores:          "1-2",
				MatchDetailLink: "/premier-league/man-utd-vs-man-city-123",
			},
			{
				Time:            "19:00",
				Round:           "1",
				Club1:           repository.Club{Name: "Chelsea"},
				Club2:           repository.Club{Name: "Liverpool"},
				Scores:          "2-1",
				MatchDetailLink: "/premier-league/chelsea-vs-liverpool-456",
			},
		},
	}

	got := newScheduleOnLeague(schedule)

	assert.Equal(scheduleLeague, got)

}