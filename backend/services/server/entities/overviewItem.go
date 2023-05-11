package entities

import (
	"time"

	"gorm.io/gorm"
)

type OverviewItem struct {
	gorm.Model
	Info       string
	ImageType  string
	TimeString string
	Time       time.Time
	MatchID    uint
	ClubID  uint
}
