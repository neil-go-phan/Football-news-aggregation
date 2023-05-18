package eventservice

import (
	"server/entities"
	"server/repository"
)

type EventService struct {
	repo repository.EventRepository
}

func NewEventService(repo repository.EventRepository) *EventService {
	eventService := &EventService{
		repo: repo,
	}
	return eventService
}

func (s *EventService) FirstOrCreate(eventRepo *entities.MatchEvent) error  {
	return s.repo.FirstOrCreate(eventRepo) 
}
