package entities

import "time"

type MatchURLsOnDay struct {
	Date time.Time
	Urls []string
}

// The matchs takes place in a spectific time on a day
type MatchURLsOnTime struct {
	Date time.Time
	Urls []string
}

type MatchURLsWithTimeOnDay struct {
	MatchsOnTimes []MatchURLsOnTime
}


type MatchDetail struct {
	MatchDetailTitle MatchDetailTitle `json:"match_detail_title"`
	MatchOverview MatchOverview `json:"match_overview"`
	MatchStatistics MatchStatistics `json:"match_statistics"`
	MatchLineup MatchLineup `json:"match_lineup"`
	MatchProgress MatchProgress `json:"match_progress"`
}

type MatchDetailTitle struct {
	Club1 Club`json:"club_1"`
	Club2 Club`json:"club_2"`
	MatchScore string`json:"match_score"`
}

type MatchOverview struct {
	Club1Overview []OverviewItem `json:"club_1_overview"`
	Club2Overview []OverviewItem `json:"club_2_overview"`
} 

type OverviewItem struct {
	Info string `json:"info"`
	ImageType string`json:"image_type"`
	Time string`json:"time"`
}

type MatchStatistics struct {
	Statistics []StatisticsItem `json:"statistics"`
}

type StatisticsItem struct {
	StatClub1 string `json:"stat_club_1"`
	StatContent string`json:"stat_content"`
	StatClub2 string `json:"stat_club_2"`
}

type MatchProgress struct {
	Events []MatchEvent`json:"events"`
}
type MatchEvent struct {
	Time string`json:"time"`
	Content string`json:"content"`
}

type MatchLineup struct {
	LineupClub1 MatchLineUpDetail`json:"lineup_club_1"`
	LineupClub2 MatchLineUpDetail`json:"lineup_club_2"`
}

type MatchLineUpDetail struct {
	ClubName string`json:"club_name"`
	Formation string`json:"formation"`
	ShirtColor string `json:"shirt_color"`
	PitchRows []PitchRows`json:"pitch_row"`
}

type PitchRows struct {
	PitchRowsDetail []PitchRowsDetail`json:"pitch_rows_detail"`
}

type PitchRowsDetail struct {
	PlayerName string`json:"player_name"`
	PlayerNumber string`json:"player_number"`
}
