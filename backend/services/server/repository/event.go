package repository

import (
	"gorm.io/gorm"
	"server/entities"
)

type EventRepo struct {
	DB *gorm.DB
}

func NewEventRepo(db *gorm.DB) *EventRepo {
	EventRepo := &EventRepo{
		DB: db,
	}
	return EventRepo
}

func (repo *EventRepo) FirstOrCreate(event *entities.MatchEvent) error {
	err := repo.DB.FirstOrCreate(event, entities.MatchEvent{MatchID: event.MatchID, Time: event.Time, Content: event.Content}).Error
	if err != nil {
		return err
	}
	return nil
}
