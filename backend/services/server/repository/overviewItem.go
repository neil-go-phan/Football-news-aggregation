package repository

import (
	"server/entities"
	"gorm.io/gorm"
)

type overviewItemRepo struct {
	DB *gorm.DB
}

func NewoOverviewItemRepo(db *gorm.DB) *overviewItemRepo {
	overviewItemRepo := &overviewItemRepo{
		DB: db,
	}
	return overviewItemRepo
}

func (repo *overviewItemRepo) FirstOrCreate(overviewItem *entities.OverviewItem) error {
	err := repo.DB.FirstOrCreate(overviewItem, entities.OverviewItem{MatchID: overviewItem.ID, ClubID: overviewItem.ClubID, Time: overviewItem.Time}).Error
	if err != nil {
		return err
	}
	return nil
}