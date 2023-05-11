package entities

import (
	"time"

	"gorm.io/gorm"
)

type MatchEvent struct {
	gorm.Model
	TimeString string
	Content    string
	Time       time.Time
	MatchID    uint
}
