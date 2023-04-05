package entities

type MatchDetail struct {
	ID string
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
	Title string 
	Statistics []StatisticsItem
}

type StatisticsItem struct {
	StatClub1 string 
	StatContent string
	StatClub2 string 
}

type MatchProgress struct {
	Title string
	Events []MatchEvent
}
type MatchEvent struct {
	Time string
	Content string
}

type MatchLineup struct {
	Title string
	LineupClub1 string
	DetailClub1 string
	LineupClub2 string
	DetailClub2 string
}

type HtmlMatchDetail struct {
	MatchDetailTitle HtmlMatchDetailTitle `json:"match_detail_title"`
	MatchOverview HtmlMatchOverview `json:"match_overview"`
	MatchStatistics HtmlMatchStatistics `json:"match_statistics"`
	MatchLineup HtmlMatchLineup `json:"match_lineup"`
	MatchProgress HtmlMatchProgress `json:"match_progress"`
}

type HtmlMatchDetailTitle struct {
	Class string `json:"class"`
	Club1 HtmlClubClass `json:"club_1"`
	Club2 HtmlClubClass `json:"club_2"`
	MatchScore string `json:"match_score_id"`
}

type HtmlMatchOverview struct {
	MatchOverviewID string `json:"id"`
	Club1OverviewClass HtmlOverviewItem `json:"club_1_overview"`
	Club2OverviewClass HtmlOverviewItem `json:"club_2_overview"`
}

type HtmlOverviewItem struct {
	ImageTypeAndTime string `json:"img_time"`
}

type HtmlMatchStatistics struct {
	MatchStatisticsID string `json:"id"`
	Title string `json:"title"`
	StatisticsItem HtmlStatisticsItem `json:"item"`
}

type HtmlStatisticsItem struct {
	Class string `json:"class"`
	StatClub1 string `json:"stat_club1"`
	StatContent string `json:"stat_content"`
	StatClub2 string  `json:"stat_club2"`
}

type HtmlMatchLineup struct {
	Title string	`json:"title_id"`
	Lineup string`json:"lineup"`
	Detail string`json:"detail"`
}

type HtmlMatchProgress struct {
	ID string `json:"id"`
	Title string`json:"title"`
	EventTime string`json:"event_time"`
	EventContent string`json:"event_content"`
}
