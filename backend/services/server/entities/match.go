package entities

import (
	"time"

	"gorm.io/gorm"
)

type Match struct {
	gorm.Model
	Time            string `json:"time"`
	Round           string `json:"round"`
	Club1ID         uint   `gorm:"foreignKey:Club1ID"`
	Club1           Club   `json:"club_1" gorm:"foreignKey:Club1ID"`
	Club2ID         uint   `gorm:"foreignKey:Club2ID"`
	Club2           Club   `json:"club_2" gorm:"foreignKey:Club2ID"`
	Scores          string `json:"scores"`
	MatchDetailLink string `json:"match_detail_link"`
	ScheduleID      uint   `json:"schedule_id"`
	TimeStart       time.Time `gorm:"column:time_start"`
	MatchStatus     string
	Club1Overview   []OverviewItem   `gorm:"foreignKey:MatchID"`
	Club2Overview   []OverviewItem   `gorm:"foreignKey:MatchID"`
	Statistics      []StatisticsItem `gorm:"foreignKey:MatchID"`
	Events          []MatchEvent     `gorm:"foreignKey:MatchID"`
	LineupClub1ID uint     
	LineupClub2ID uint    
}
