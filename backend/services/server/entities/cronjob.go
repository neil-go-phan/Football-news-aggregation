package entities

import (
	"time"

	"gorm.io/gorm"
)

type Cronjob struct {
	gorm.Model
	Name        string
	StartAt     time.Time
	EndAt       time.Time
	CrawlerID   uint    `gorm:"foreignKey:CrawlerID"`
	Crawler     Crawler `gorm:"foreignKey:CrawlerID"`
	RunEveryMin int
}
