package repository

import (
	"server/entities"
	"gorm.io/gorm"
)

type OverviewItemRepo struct {
	DB *gorm.DB
}

func NewoOverviewItemRepo(db *gorm.DB) *OverviewItemRepo {
	OverviewItemRepo := &OverviewItemRepo{
		DB: db,
	}
	return OverviewItemRepo
}

func (repo *OverviewItemRepo) FirstOrCreate(overviewItem *entities.OverviewItem) error {
	err := repo.DB.FirstOrCreate(overviewItem, entities.OverviewItem{MatchID: overviewItem.ID, ClubID: overviewItem.ClubID, Time: overviewItem.Time}).Error
	if err != nil {
		return err
	}
	return nil
}