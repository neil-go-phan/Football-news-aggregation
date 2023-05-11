package entities

type ScheduleOnDay struct {
	Date   string             `json:"date_formated"`
	ScheduleOnLeagues []ScheduleOnLeague `json:"schedule_on_leagues"`
}

type ScheduleOnLeague struct {
	LeagueName string  `json:"league_name"`
	Matches     []Match `json:"matches"`
}

type Match struct {
	Time        string      `json:"time"`
	Round       string      `json:"round"`
	Club1       Club        `json:"club_1"`
	Club2       Club        `json:"club_2"`
	Scores      string      `json:"scores"`
	MatchDetailLink string `json:"match_detail_link"`
}

type HtmlSchedulesClass struct {
	LeagueClass    string         `json:"league_class"`
	Date           string         `json:"date"`
	HtmlMatchClass HtmlMatchClass `json:"html_match_class"`
}

type HtmlMatchClass struct {
	MatchClass    string        `json:"match_class"`
	Time          string        `json:"time"`
	Round         string        `json:"round"`
	Club1         HtmlClubClass `json:"club_1"`
	Club2         HtmlClubClass `json:"club_2"`
	Scores        string        `json:"scores"`
	MatchDetailLink string        `json:"match_detail_link"`
}
