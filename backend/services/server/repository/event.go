package repository

import (
	"gorm.io/gorm"
	"server/entities"
)

type eventRepo struct {
	DB *gorm.DB
}

func NewEventRepo(db *gorm.DB) *eventRepo {
	eventRepo := &eventRepo{
		DB: db,
	}
	return eventRepo
}

func (repo *eventRepo) FirstOrCreate(eventRepo *entities.MatchEvent) error {
	err := repo.DB.FirstOrCreate(eventRepo, entities.MatchEvent{MatchID: eventRepo.MatchID, Time: eventRepo.Time, Content: eventRepo.Content}).Error
	if err != nil {
		return err
	}
	return nil
}
