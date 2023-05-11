package migrations

import (
	"server/entities"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) (err error) {
	return db.AutoMigrate(
		&entities.League{},
		&entities.Admin{},
		&entities.Tag{},
		&entities.Article{},
		&entities.Club{},
		&entities.Schedule{},
		&entities.MatchLineUp{},
		&entities.Match{},
		&entities.OverviewItem{},
		&entities.Player{},
		&entities.MatchEvent{},
		&entities.StatisticsItem{},
	)
}