package repository

import (
	"server/entities"
	"gorm.io/gorm"
)

type StatsItemRepo struct {
	DB *gorm.DB
}

func NewStatsItemRepo(db *gorm.DB) *StatsItemRepo {
	StatsItemRepo := &StatsItemRepo{
		DB: db,
	}
	return StatsItemRepo
}

func (repo *StatsItemRepo) FirstOrCreate(statsItem *entities.StatisticsItem) error {
	err := repo.DB.FirstOrCreate(statsItem, entities.StatisticsItem{MatchID: statsItem.MatchID, StatContent: statsItem.StatContent}).Error
	if err != nil {
		return err
	}
	return nil
}