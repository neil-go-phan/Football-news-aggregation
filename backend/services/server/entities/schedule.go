package entities

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	Date       time.Time `json:"date" gorm:"uniqueIndex:idx_schedule_league_name"`
	LeagueName string    `json:"league_name" gorm:"foreignKey:LeagueName;uniqueIndex:idx_schedule_league_name"`
	Matches    []Match   `json:"matches" gorm:"foreignKey:ScheduleID"`
}