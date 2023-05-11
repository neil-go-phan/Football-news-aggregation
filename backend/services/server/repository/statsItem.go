package repository

import (
	"server/entities"
	"gorm.io/gorm"
)

type statsItemRepo struct {
	DB *gorm.DB
}

func NewStatsItemRepo(db *gorm.DB) *statsItemRepo {
	statsItemRepo := &statsItemRepo{
		DB: db,
	}
	return statsItemRepo
}

func (repo *statsItemRepo) FirstOrCreate(statsItem *entities.StatisticsItem) error {
	err := repo.DB.FirstOrCreate(statsItem, entities.StatisticsItem{MatchID: statsItem.MatchID, StatContent: statsItem.StatContent}).Error
	if err != nil {
		return err
	}
	return nil
}