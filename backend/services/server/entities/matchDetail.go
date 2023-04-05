package entities

type MatchDetail struct {
	ID string `json:"match_detail_id"`
	MatchDetailTitle MatchDetailTitle `json:"match_detail_title"`
	MatchOverview MatchOverview `json:"match_overview"`
	MatchStatistics MatchStatistics `json:"match_statistics"`
	MatchLineup MatchLineup `json:"match_lineup"`
	MatchProgress MatchProgress `json:"match_progress"`
}

type MatchDetailTitle struct {
	Club1 Club `json:"club_1"`
	Club2 Club `json:"club_2"`
	MatchScore string `json:"match_score"`
}

type MatchOverview struct {
	Club1Overview []OverviewItem `json:"club_1_overview"`
	Club2Overview []OverviewItem `json:"club_2_overview"`
}

type OverviewItem struct {
	Info string `json:"info"`
	ImageType string `json:"image_type"`
	Time string `json:"time"`
}

type MatchStatistics struct {
	Title string `json:"title"`
	Statistics []StatisticsItem `json:"stats"`
}

type StatisticsItem struct {
	StatClub1 string `json:"stat_club1"`
	StatContent string`json:"content"`
	StatClub2 string `json:"stat_club2"`
}

type MatchProgress struct {
	Title string `json:"title"`
	Events []MatchEvent `json:"events"`
}
type MatchEvent struct {
	Time string `json:"time"`
	Content string `json:"content"`
}

type MatchLineup struct {
	Title string `json:"title"`
	LineupClub1 string `json:"lineup_club1"`
	DetailClub1 string `json:"detail_club1"`
	LineupClub2 string `json:"lineup_club2"`
	DetailClub2 string `json:"detail_club2"`
}
