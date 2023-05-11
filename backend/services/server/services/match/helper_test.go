package matchservices

// import (
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"server/entities"
// 	pb "server/proto"
// 	"testing"
// 	"time"

// 	"github.com/elastic/go-elasticsearch/v7"
// 	"github.com/stretchr/testify/assert"
// )

// func TestQuerySearchMatchDetailByID(t *testing.T) {
// 	expectedQuery := map[string]interface{}{
// 		"query": map[string]interface{}{
// 			"term": map[string]interface{}{
// 				"_id": "test_id",
// 			},
// 		},
// 	}

// 	result := querySearchMatchDetailByID("test_id")

// 	assert.Equal(t, expectedQuery, result)
// }

// func TestNewEntitiesMatchDetailFromMap(t *testing.T) {
// 	matchDetailMap := map[string]interface{}{
// 		"match_detail_title": map[string]interface{}{
// 			"club_1": map[string]interface{}{
// 				"name": "Real Madrid",
// 				"logo": "https://example.com/logo.png",
// 			},
// 			"club_2": map[string]interface{}{
// 				"name": "Barcelona",
// 				"logo": "https://example.com/logo.png",
// 			},
// 			"match_score": "2-1",
// 		},
// 		"match_overview": map[string]interface{}{
// 			"club_1_overview": []interface{}{
// 				map[string]interface{}{
// 					"info":       "ghi ban",
// 					"image_type": "icon",
// 					"time":       "90'",
// 				},
// 				map[string]interface{}{
// 					"info":       "ghi ban",
// 					"image_type": "icon",
// 					"time":       "91'",
// 				},
// 			},
// 			"club_2_overview": []interface{}{
// 				map[string]interface{}{
// 					"info":       "ghi ban",
// 					"image_type": "icon",
// 					"time":       "80'",
// 				},
// 				map[string]interface{}{
// 					"info":       "ghi ban",
// 					"image_type": "icon",
// 					"time":       "81'",
// 				},
// 			},
// 		},
// 		"match_statistics": map[string]interface{}{
// 			"statistics": []interface{}{
// 				map[string]interface{}{
// 					"stat_club_1":  "50",
// 					"stat_content": "Giu bong",
// 					"stat_club_2":  "50",
// 				},
// 				map[string]interface{}{
// 					"stat_club_1":  "40",
// 					"stat_content": "Mat bong",
// 					"stat_club_2":  "60",
// 				},
// 			},
// 		},
// 		"match_lineup": map[string]interface{}{
// 			"lineup_club_1": map[string]interface{}{
// 				"club_name":   "Real Madrid",
// 				"formation":   "4-4-2",
// 				"shirt_color": "white",
// 				"pitch_row": []interface{}{
// 					map[string]interface{}{
// 						"pitch_rows_detail": []interface{}{
// 							map[string]interface{}{
// 								"player_name":   "Player A",
// 								"player_number": "10",
// 							},
// 							map[string]interface{}{
// 								"player_name":   "Player B",
// 								"player_number": "9",
// 							},
// 						},
// 					},
// 				},
// 			},
// 			"lineup_club_2": map[string]interface{}{
// 				"club_name":   "Barcelona",
// 				"formation":   "4-3-3",
// 				"shirt_color": "red",
// 				"pitch_row": []interface{}{
// 					map[string]interface{}{
// 						"pitch_rows_detail": []interface{}{
// 							map[string]interface{}{
// 								"player_name":   "Player C",
// 								"player_number": "10",
// 							},
// 							map[string]interface{}{
// 								"player_name":   "Player D",
// 								"player_number": "9",
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 		"match_progress": map[string]interface{}{
// 			"events": []interface{}{
// 				map[string]interface{}{
// 					"time":    "15'",
// 					"content": "Goal for Real Madrid",
// 				},
// 				map[string]interface{}{
// 					"time":    "30'",
// 					"content": "Yellow card for Player C",
// 				},
// 			},
// 		},
// 	}
// 	assert := assert.New(t)

// 	want := entities.MatchDetail{
// 		MatchDetailTitle: entities.MatchDetailTitle{
// 			Club1: entities.Club{
// 				Name: "Real Madrid",
// 				Logo: "https://example.com/logo.png",
// 			},
// 			Club2: entities.Club{
// 				Name: "Barcelona",
// 				Logo: "https://example.com/logo.png",
// 			},
// 			MatchScore: "2-1",
// 		},
// 		MatchOverview: entities.MatchOverview{
// 			Club1Overview: []entities.OverviewItem{
// 				{Info: "ghi ban", ImageType: "icon", Time: "90'"},
// 				{Info: "ghi ban", ImageType: "icon", Time: "91'"},
// 			},
// 			Club2Overview: []entities.OverviewItem{
// 				{Info: "ghi ban", ImageType: "icon", Time: "80'"},
// 				{Info: "ghi ban", ImageType: "icon", Time: "81'"},
// 			},
// 		},
// 		MatchStatistics: entities.MatchStatistics{
// 			Statistics: []entities.StatisticsItem{
// 				{StatClub1: "50", StatContent: "Giu bong", StatClub2: "50"},
// 				{StatClub1: "40", StatContent: "Mat bong", StatClub2: "60"},
// 			},
// 		},
// 		MatchLineup: entities.MatchLineup{
// 			LineupClub1: entities.MatchLineUpDetail{
// 				ClubName:   "Real Madrid",
// 				Formation:  "4-4-2",
// 				ShirtColor: "white",
// 				PitchRows: []entities.PitchRows{
// 					{
// 						PitchRowsDetail: []entities.PitchRowsDetail{
// 							{
// 								PlayerName:   "Player A",
// 								PlayerNumber: "10",
// 							},
// 							{
// 								PlayerName:   "Player B",
// 								PlayerNumber: "9",
// 							},
// 						},
// 					},
// 				},
// 			},
// 			LineupClub2: entities.MatchLineUpDetail{
// 				ClubName:   "Barcelona",
// 				Formation:  "4-3-3",
// 				ShirtColor: "red",
// 				PitchRows: []entities.PitchRows{
// 					{
// 						PitchRowsDetail: []entities.PitchRowsDetail{
// 							{
// 								PlayerName:   "Player C",
// 								PlayerNumber: "10",
// 							},
// 							{
// 								PlayerName:   "Player D",
// 								PlayerNumber: "9",
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 		MatchProgress: entities.MatchProgress{
// 			Events: []entities.MatchEvent{
// 				{Time: "15'",
// 					Content: "Goal for Real Madrid"},
// 				{Time: "30'",
// 					Content: "Yellow card for Player C"},
// 			},
// 		},
// 	}

// 	got := newEntitiesMatchDetailFromMap(matchDetailMap)

// 	assert.Equal(want, got)
// }

// func TestNewEntitiesMatchDetailFromMap_InvalidJson(t *testing.T) {
// 	invalidMap := map[string]interface{}{
// 		"key1": make(chan int), // giá trị của key1 không thể marshalled thành json
// 		"key2": "value2",
// 	}

// 	result := newEntitiesMatchDetailFromMap(invalidMap)
// 	assert.Equal(t, entities.MatchDetail{}, result, "result should be an empty MatchDetail")
// }

// func TestPbMatchDetailToEntityMatchDetail(t *testing.T) {
// 	assert := assert.New(t)

// 	matchDetailPb := pb.MatchDetail{
// 		MatchDetailTitle: &pb.MatchDetailTitle{
// 			Club_1: &pb.Club{
// 				Name: "Real Madrid",
// 				Logo: "https://example.com/logo.png",
// 			},
// 			Club_2: &pb.Club{
// 				Name: "Barcelona",
// 				Logo: "https://example.com/logo.png",
// 			},
// 			MatchScore: "2-1",
// 		},
// 		MatchOverview: &pb.MatchOverview{
// 			Club_1Overview: []*pb.OverviewItem{
// 				{Info: "ghi ban", ImageType: "icon", Time: "90'"},
// 				{Info: "ghi ban", ImageType: "icon", Time: "91'"},
// 			},
// 			Club_2Overview: []*pb.OverviewItem{
// 				{Info: "ghi ban", ImageType: "icon", Time: "80'"},
// 				{Info: "ghi ban", ImageType: "icon", Time: "81'"},
// 			},
// 		},
// 		MatchStatistics: &pb.MatchStatistics{
// 			Statistics: []*pb.StatisticsItem{
// 				{StatClub_1: "50", StatContent: "Giu bong", StatClub_2: "50"},
// 				{StatClub_1: "40", StatContent: "Mat bong", StatClub_2: "60"},
// 			},
// 		},
// 		MatchLineup: &pb.MatchLineup{
// 			LineupClub_1: &pb.MatchLineUpDetail{
// 				ClubName:   "Real Madrid",
// 				Formation:  "4-4-2",
// 				ShirtColor: "white",
// 				PitchRow: []*pb.PitchRows{
// 					{
// 						PitchRowsDetail: []*pb.PitchRowsDetail{
// 							{
// 								PlayerName:   "Player A",
// 								PlayerNumber: "10",
// 							},
// 							{
// 								PlayerName:   "Player B",
// 								PlayerNumber: "9",
// 							},
// 						},
// 					},
// 				},
// 			},
// 			LineupClub_2: &pb.MatchLineUpDetail{
// 				ClubName:   "Barcelona",
// 				Formation:  "4-3-3",
// 				ShirtColor: "red",
// 				PitchRow: []*pb.PitchRows{{
// 					PitchRowsDetail: []*pb.PitchRowsDetail{
// 						{
// 							PlayerName:   "Player C",
// 							PlayerNumber: "10",
// 						},
// 						{
// 							PlayerName:   "Player D",
// 							PlayerNumber: "9",
// 						},
// 					},
// 				},
// 				},
// 			},
// 		},
// 		MatchProgress: &pb.MatchProgress{
// 			Events: []*pb.MatchEvent{
// 				{Time: "15'",
// 					Content: "Goal for Real Madrid"},
// 				{Time: "30'",
// 					Content: "Yellow card for Player C"},
// 			},
// 		},
// 	}

// 	want := entities.MatchDetail{
// 		MatchDetailTitle: entities.MatchDetailTitle{
// 			Club1: entities.Club{
// 				Name: "Real Madrid",
// 				Logo: "https://example.com/logo.png",
// 			},
// 			Club2: entities.Club{
// 				Name: "Barcelona",
// 				Logo: "https://example.com/logo.png",
// 			},
// 			MatchScore: "2-1",
// 		},
// 		MatchOverview: entities.MatchOverview{
// 			Club1Overview: []entities.OverviewItem{
// 				{Info: "ghi ban", ImageType: "icon", Time: "90'"},
// 				{Info: "ghi ban", ImageType: "icon", Time: "91'"},
// 			},
// 			Club2Overview: []entities.OverviewItem{
// 				{Info: "ghi ban", ImageType: "icon", Time: "80'"},
// 				{Info: "ghi ban", ImageType: "icon", Time: "81'"},
// 			},
// 		},
// 		MatchStatistics: entities.MatchStatistics{
// 			Statistics: []entities.StatisticsItem{
// 				{StatClub1: "50", StatContent: "Giu bong", StatClub2: "50"},
// 				{StatClub1: "40", StatContent: "Mat bong", StatClub2: "60"},
// 			},
// 		},
// 		MatchLineup: entities.MatchLineup{
// 			LineupClub1: entities.MatchLineUpDetail{
// 				ClubName:   "Real Madrid",
// 				Formation:  "4-4-2",
// 				ShirtColor: "white",
// 				PitchRows: []entities.PitchRows{
// 					{
// 						PitchRowsDetail: []entities.PitchRowsDetail{
// 							{
// 								PlayerName:   "Player A",
// 								PlayerNumber: "10",
// 							},
// 							{
// 								PlayerName:   "Player B",
// 								PlayerNumber: "9",
// 							},
// 						},
// 					},
// 				},
// 			},
// 			LineupClub2: entities.MatchLineUpDetail{
// 				ClubName:   "Barcelona",
// 				Formation:  "4-3-3",
// 				ShirtColor: "red",
// 				PitchRows: []entities.PitchRows{
// 					{
// 						PitchRowsDetail: []entities.PitchRowsDetail{
// 							{
// 								PlayerName:   "Player C",
// 								PlayerNumber: "10",
// 							},
// 							{
// 								PlayerName:   "Player D",
// 								PlayerNumber: "9",
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 		MatchProgress: entities.MatchProgress{
// 			Events: []entities.MatchEvent{
// 				{Time: "15'",
// 					Content: "Goal for Real Madrid"},
// 				{Time: "30'",
// 					Content: "Yellow card for Player C"},
// 			},
// 		},
// 	}

// 	got := pbMatchDetailToEntityMatchDetail(&matchDetailPb)

// 	assert.Equal(want, got)
// }

// func TestUpsertMatchDetailElasticRequestFail(t *testing.T) {
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprintln(w, `{"id":"pitID"}`)
// 	}))
// 	defer server.Close()
// 	assert := assert.New(t)

// 	cfg := elasticsearch.Config{
// 		Addresses: []string{server.URL},
// 	}
// 	es, err := elasticsearch.NewClient(cfg)
// 	assert.Nil(err)

// 	matchDetail := entities.MatchDetail{
// 		MatchDetailTitle: entities.MatchDetailTitle{
// 			Club1: entities.Club{
// 				Name: "Real Madrid",
// 				Logo: "https://example.com/logo.png",
// 			},
// 			Club2: entities.Club{
// 				Name: "Barcelona",
// 				Logo: "https://example.com/logo.png",
// 			},
// 			MatchScore: "2-1",
// 		},
// 		MatchOverview: entities.MatchOverview{
// 			Club1Overview: []entities.OverviewItem{
// 				{Info: "ghi ban", ImageType: "icon", Time: "90'"},
// 				{Info: "ghi ban", ImageType: "icon", Time: "91'"},
// 			},
// 			Club2Overview: []entities.OverviewItem{
// 				{Info: "ghi ban", ImageType: "icon", Time: "80'"},
// 				{Info: "ghi ban", ImageType: "icon", Time: "81'"},
// 			},
// 		},
// 		MatchStatistics: entities.MatchStatistics{
// 			Statistics: []entities.StatisticsItem{
// 				{StatClub1: "50", StatContent: "Giu bong", StatClub2: "50"},
// 				{StatClub1: "40", StatContent: "Mat bong", StatClub2: "60"},
// 			},
// 		},
// 		MatchLineup: entities.MatchLineup{
// 			LineupClub1: entities.MatchLineUpDetail{
// 				ClubName:   "Real Madrid",
// 				Formation:  "4-4-2",
// 				ShirtColor: "white",
// 				PitchRows: []entities.PitchRows{
// 					{
// 						PitchRowsDetail: []entities.PitchRowsDetail{
// 							{
// 								PlayerName:   "Player A",
// 								PlayerNumber: "10",
// 							},
// 							{
// 								PlayerName:   "Player B",
// 								PlayerNumber: "9",
// 							},
// 						},
// 					},
// 				},
// 			},
// 			LineupClub2: entities.MatchLineUpDetail{
// 				ClubName:   "Barcelona",
// 				Formation:  "4-3-3",
// 				ShirtColor: "red",
// 				PitchRows: []entities.PitchRows{
// 					{
// 						PitchRowsDetail: []entities.PitchRowsDetail{
// 							{
// 								PlayerName:   "Player C",
// 								PlayerNumber: "10",
// 							},
// 							{
// 								PlayerName:   "Player D",
// 								PlayerNumber: "9",
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 		MatchProgress: entities.MatchProgress{
// 			Events: []entities.MatchEvent{
// 				{Time: "15'",
// 					Content: "Goal for Real Madrid"},
// 				{Time: "30'",
// 					Content: "Yellow card for Player C"},
// 			},
// 		},
// 	}

// 	date := time.Now()

// 	upsertMatchDetailElastic(matchDetail, es, date)

// }

// func TestUpsertMatchDetailElasticSuccess(t *testing.T) {
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		fmt.Fprintln(w, `{"id":"pitID"}`)
// 	}))
// 	defer server.Close()
// 	assert := assert.New(t)

// 	cfg := elasticsearch.Config{
// 		Addresses: []string{server.URL},
// 	}
// 	es, err := elasticsearch.NewClient(cfg)
// 	assert.Nil(err)

// 	matchDetail := entities.MatchDetail{
// 		MatchDetailTitle: entities.MatchDetailTitle{
// 			Club1: entities.Club{
// 				Name: "Real Madrid",
// 				Logo: "https://example.com/logo.png",
// 			},
// 			Club2: entities.Club{
// 				Name: "Barcelona",
// 				Logo: "https://example.com/logo.png",
// 			},
// 			MatchScore: "2-1",
// 		},
// 		MatchOverview: entities.MatchOverview{
// 			Club1Overview: []entities.OverviewItem{
// 				{Info: "ghi ban", ImageType: "icon", Time: "90'"},
// 				{Info: "ghi ban", ImageType: "icon", Time: "91'"},
// 			},
// 			Club2Overview: []entities.OverviewItem{
// 				{Info: "ghi ban", ImageType: "icon", Time: "80'"},
// 				{Info: "ghi ban", ImageType: "icon", Time: "81'"},
// 			},
// 		},
// 		MatchStatistics: entities.MatchStatistics{
// 			Statistics: []entities.StatisticsItem{
// 				{StatClub1: "50", StatContent: "Giu bong", StatClub2: "50"},
// 				{StatClub1: "40", StatContent: "Mat bong", StatClub2: "60"},
// 			},
// 		},
// 		MatchLineup: entities.MatchLineup{
// 			LineupClub1: entities.MatchLineUpDetail{
// 				ClubName:   "Real Madrid",
// 				Formation:  "4-4-2",
// 				ShirtColor: "white",
// 				PitchRows: []entities.PitchRows{
// 					{
// 						PitchRowsDetail: []entities.PitchRowsDetail{
// 							{
// 								PlayerName:   "Player A",
// 								PlayerNumber: "10",
// 							},
// 							{
// 								PlayerName:   "Player B",
// 								PlayerNumber: "9",
// 							},
// 						},
// 					},
// 				},
// 			},
// 			LineupClub2: entities.MatchLineUpDetail{
// 				ClubName:   "Barcelona",
// 				Formation:  "4-3-3",
// 				ShirtColor: "red",
// 				PitchRows: []entities.PitchRows{
// 					{
// 						PitchRowsDetail: []entities.PitchRowsDetail{
// 							{
// 								PlayerName:   "Player C",
// 								PlayerNumber: "10",
// 							},
// 							{
// 								PlayerName:   "Player D",
// 								PlayerNumber: "9",
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 		MatchProgress: entities.MatchProgress{
// 			Events: []entities.MatchEvent{
// 				{Time: "15'",
// 					Content: "Goal for Real Madrid"},
// 				{Time: "30'",
// 					Content: "Yellow card for Player C"},
// 			},
// 		},
// 	}

// 	date := time.Now()

// 	upsertMatchDetailElastic(matchDetail, es, date)

// }
