package schedulesrepo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"server/entities"
	"strings"
	"testing"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/stretchr/testify/assert"
)

func TestGetMatchUrl(t *testing.T) {
	schedule := entities.ScheduleElastic{
		Date:       time.Now(),
		LeagueName: "Premier League",
		Matchs: []entities.Match{
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
	matchUrlsWithTimeOnDay := &entities.MatchURLsWithTimeOnDay{}

	addMatchUrl(date, inputUrl, matchUrlsWithTimeOnDay)
	assert := assert.New(t)

	assert.Equal(1, len(matchUrlsWithTimeOnDay.MatchsOnTimes))
	assert.Equal(date, matchUrlsWithTimeOnDay.MatchsOnTimes[0].Date)
	assert.Equal([]string{"https://bongda24h.vn/match-1"}, matchUrlsWithTimeOnDay.MatchsOnTimes[0].Urls)
}

func TestAddMatchUrl_AddsNewMatchOnSameTime(t *testing.T) {
	date := time.Date(2023, 5, 3, 0, 0, 0, 0, time.UTC)
	inputUrl := "/match-1"
	matchUrlsWithTimeOnDay := &entities.MatchURLsWithTimeOnDay{
			MatchsOnTimes: []entities.MatchURLsOnTime{
					{
							Date: date,
							Urls: []string{"https://bongda24h.vn/match-1"},
					},
			},
	}

	// Act
	addMatchUrl(date, inputUrl, matchUrlsWithTimeOnDay)

	assert := assert.New(t)

	assert.Equal(1, len(matchUrlsWithTimeOnDay.MatchsOnTimes))
	assert.Equal(date, matchUrlsWithTimeOnDay.MatchsOnTimes[0].Date)
	assert.Equal([]string{"https://bongda24h.vn/match-1"}, matchUrlsWithTimeOnDay.MatchsOnTimes[0].Urls)
}

func TestQuerySearchAllScheduleOnDay(t *testing.T) {
	expectedQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": map[string]interface{}{
					"term": map[string]interface{}{
						"date": time.Date(2023, time.May, 1, 0, 0, 0, 0, time.UTC),
					},
				},
			},
		},
		"size": 1000,
	}
	date := time.Date(2023, time.May, 1, 0, 0, 0, 0, time.UTC)
	resultQuery := querySearchAllScheduleOnDay(date)
	assert := assert.New(t)
	assert.Equal(expectedQuery, resultQuery, "Result query does not match expected query.")
}

func TestQuerySearchScheduleLeagueOnDay(t *testing.T) {
	expectedQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": []map[string]interface{}{
					{
						"term": map[string]interface{}{
							"league_name.keyword": "Premier League",
						},
					},
					{
						"term": map[string]interface{}{
							"date": time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
						},
					},
				},
			},
		},
	}

	dateISO8601 := time.Date(2022, 01, 01, 0, 0, 0, 0, time.UTC)
	league := "Premier League"

	result := querySearchScheduleLeagueOnDay(dateISO8601, league)
	assert := assert.New(t)
	assert.Equal(expectedQuery, result, "Result query does not match expected query.")
}

func TestCheckScheduleWithElasticSearchFail(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, `{"id":"pitID"}`)
	}))
	defer server.Close()
	assert := assert.New(t)

	cfg := elasticsearch.Config{
		Addresses: []string{server.URL},
	}
	es, err := elasticsearch.NewClient(cfg)
	assert.Nil(err)

	date := time.Date(2023, time.May, 1, 0, 0, 0, 0, time.UTC)
	schedule := entities.ScheduleElastic{
		Date: date,
		LeagueName: "League 1",
		Matchs: []entities.Match{},
	}

	ok := checkScheduleWithElasticSearch(schedule, es)
	assert.Equal(false, ok)
}

func TestCheckScheduleWithElasticSearchOK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"id":"pitID"}`)
	}))
	defer server.Close()
	assert := assert.New(t)

	cfg := elasticsearch.Config{
		Addresses: []string{server.URL},
	}
	es, err := elasticsearch.NewClient(cfg)
	assert.Nil(err)

	date := time.Date(2023, time.May, 1, 0, 0, 0, 0, time.UTC)
	schedule := entities.ScheduleElastic{
		Date: date,
		LeagueName: "League 1",
		Matchs: []entities.Match{},
	}

	ok := checkScheduleWithElasticSearch(schedule, es)
	assert.Equal(true, ok)
}

func TestNewEntitiesScheduleOnLeaguesFromMapSuccess(t *testing.T) {
	respScheduleOnLeague := map[string]interface{}{
		"league_name": "Premier League",
		"matchs": []interface{}{
			map[string]interface{}{
				"time":             "23:30 - FT",
				"round":            "38/38",
				"club_1":           map[string]interface{}{"name": "Liverpool", "logo": "https://bongda24h.vn/image/upload/score/02-1629365283-15.jpg"},
				"club_2":           map[string]interface{}{"name": "Wolves", "logo": "https://bongda24h.vn/image/upload/score/02-1629365283-15.jpg"},
				"scores":           "2-0",
				"match_detail_link":"https://bongda24h.vn/bong-da-anh/liverpool-vs-wolves-23h00-ngay-23-5-tbd110623.html",
			},
			map[string]interface{}{
				"time":             "02:00",
				"round":            "1/38",
				"club_1":           map[string]interface{}{"name": "Man United", "logo": "https://bongda24h.vn/image/upload/score/02-1629365283-15.jpg"},
				"club_2":           map[string]interface{}{"name": "Chelsea", "logo": "https://bongda24h.vn/image/upload/score/02-1629365283-15.jpg"},
				"scores":           "",
				"match_detail_link":"https://bongda24h.vn/bong-da-anh/manchester-united-vs-chelsea-02h00-ngay-23-5-tbd110551.html",
			},
		},
	}

	want := entities.ScheduleOnLeague{
		LeagueName: "Premier League",
		Matchs: []entities.Match{
			{
				Time:            "23:30 - FT",
				Round:           "38/38",
				Club1:           entities.Club{Name: "Liverpool", Logo: "https://bongda24h.vn/image/upload/score/02-1629365283-15.jpg"},
				Club2:           entities.Club{Name: "Wolves", Logo: "https://bongda24h.vn/image/upload/score/02-1629365283-15.jpg"},
				Scores:          "2-0",
				MatchDetailLink: "https://bongda24h.vn/bong-da-anh/liverpool-vs-wolves-23h00-ngay-23-5-tbd110623.html",
			},
			{
				Time:            "02:00",
				Round:           "1/38",
				Club1:           entities.Club{Name: "Man United", Logo: "https://bongda24h.vn/image/upload/score/02-1629365283-15.jpg"},
				Club2:           entities.Club{Name: "Chelsea", Logo: "https://bongda24h.vn/image/upload/score/02-1629365283-15.jpg"},
				Scores:          "",
				MatchDetailLink: "https://bongda24h.vn/bong-da-anh/manchester-united-vs-chelsea-02h00-ngay-23-5-tbd110551.html",
			},
		},
	}

	got := newEntitiesScheduleOnLeaguesFromMap(respScheduleOnLeague)

	assert := assert.New(t)

	assert.Equal(want, got)
}

func TestNewEntitiesScheduleOnLeaguesFromMapFail(t *testing.T) {
	invalidMap := map[string]interface{}{
		"key1": make(chan int),
		"key2": "value2",
	}

	want := entities.ScheduleOnLeague{}

	got := newEntitiesScheduleOnLeaguesFromMap(invalidMap)

	assert := assert.New(t)

	assert.Equal(want, got)
}