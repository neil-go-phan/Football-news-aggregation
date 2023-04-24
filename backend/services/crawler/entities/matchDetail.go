package entities

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

type XPathMatchDetail struct {
	MatchDetailTitle XPathMatchDetailTitle `json:"match_detail_title"`
	MatchOverview XPathMatchOverview `json:"match_overview"`
	MatchStatistics XPathMatchStatistics `json:"match_statistics"`
	MatchLineup XPathMatchLineup `json:"match_lineup"`
	MatchProgress XPathMatchProgress `json:"match_progress"`
}

type XPathMatchDetailTitle struct {
	Club1 XPathClubClass `json:"club_1"`
	Club2 XPathClubClass `json:"club_2"`
	MatchScore string `json:"match_score_id"`
}

type XPathMatchOverview struct {
	Club1OverviewClass XPathOverviewItem `json:"club_1_overview"`
	Club2OverviewClass XPathOverviewItem `json:"club_2_overview"`
}

type XPathOverviewItem struct {
	List string `json:"list"`
	Info string `json:"info"`
	Time string `json:"time"`
	Img string `json:"img"`
}

type XPathMatchStatistics struct {
	MatchStatisticsListItem string `json:"list_item"`
	StatisticsItem XPathStatisticsItem `json:"item"`
}

type XPathStatisticsItem struct {
	StatClub1 string `json:"stat_club1"`
	StatContent string `json:"stat_content"`
	StatClub2 string  `json:"stat_club2"`
}

type XPathMatchLineup struct {
	Lineup string`json:"lineup"`
	Club1 XPathMatchLineUpDetail`json:"club1"`
	Club2 XPathMatchLineUpDetail`json:"club2"`
}

type XPathMatchLineUpDetail struct {
	ClubName string`json:"club_name"`
	Formation string`json:"formation"`
	ShirtColor string `json:"shirt_color"`
	PitchRows XPathPitchRow`json:"pitch_row"`
}

type XPathPitchRow struct {
	List string `json:"list"`
	ListPlayer XPathPitchRowPlayer `json:"list_player"`
}

type XPathPitchRowPlayer struct {
	List string `json:"list"`
	PlayerNumber string `json:"player_number"`
	PlayerName string `json:"player_name"`
}

type XPathMatchProgress struct {
	Events string `json:"events"`
	EventTime string`json:"event_time"`
	EventContent string`json:"event_content"`
}
