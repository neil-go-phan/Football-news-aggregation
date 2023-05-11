package eventservice

import (
	"server/entities"
	"server/repository"
)

type eventService struct {
	repo repository.EventRepository
}

func NewEventService(repo repository.EventRepository) *eventService {
	eventService := &eventService{
		repo: repo,
	}
	return eventService
}

func (s *eventService) FirstOrCreate(eventRepo *entities.MatchEvent) error  {
	return s.repo.FirstOrCreate(eventRepo) 
}
