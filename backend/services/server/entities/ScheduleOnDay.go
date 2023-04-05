package entities

import "time"

type ScheduleElastic struct {
	Date       time.Time `json:"date"`
	LeagueName string    `json:"league_name"`
	Matchs     []Match   `json:"matchs"`
}

type ScheduleOnDay struct {
	Date              time.Time          `json:"date"`
	DateWithWeekday   string             `json:"date_with_weekday"`
	ScheduleOnLeagues []ScheduleOnLeague `json:"schedule_on_leagues"`
}

type ScheduleOnLeague struct {
	LeagueName string  `json:"league_name"`
	Matchs     []Match `json:"matchs"`
}

type Match struct {
	Time            string `json:"time"`
	Round           string `json:"round"`
	Club1           Club   `json:"club_1"`
	Club2           Club   `json:"club_2"`
	Scores          string `json:"scores"`
	MatchDetailLink string `json:"match_detail_id"`
}

type Club struct {
	Name string `json:"name"`
	Logo string `json:"logo"`
}
