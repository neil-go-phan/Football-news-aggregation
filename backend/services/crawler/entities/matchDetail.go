package entities

type MatchDetail struct {
	MatchDetailTitle MatchDetailTitle
	MatchOverview MatchOverview
	MatchStatistics MatchStatistics
	MatchLineup MatchLineup
	MatchProgress MatchProgress
}

type MatchDetailTitle struct {
	Club1 Club
	Club2 Club
	MatchScore string
}

type MatchOverview struct {
	Club1Overview []OverviewItem
	Club2Overview []OverviewItem
}

type OverviewItem struct {
	Info string
	ImageType string
	Time string
}

type MatchStatistics struct {
	Statistics []StatisticsItem
}

type StatisticsItem struct {
	StatClub1 string 
	StatContent string
	StatClub2 string 
}

type MatchProgress struct {
	Events []MatchEvent
}
type MatchEvent struct {
	Time string
	Content string
}

type MatchLineup struct {
	LineupClub1 MatchLineUpDetail
	LineupClub2 MatchLineUpDetail
}

type MatchLineUpDetail struct {
	ClubName string
	Formation string
	PitchRows []PitchRows
}

type PitchRows struct {
	PitchRowsDetail []PitchRowsDetail
}

type PitchRowsDetail struct {
	PlayerName string
	PlayerNumber string
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
