package playerservice

import (
	"server/entities"
	"server/repository"
)

type PlayerService struct {
	repo repository.PlayerRepository
}

func NewPlayerService(repo repository.PlayerRepository) *PlayerService {
	playerService := &PlayerService{
		repo: repo,
	}
	return playerService
}

func (s *PlayerService) FirstOrCreate(player *entities.Player) error  {
	return s.repo.FirstOrCreate(player) 
}
