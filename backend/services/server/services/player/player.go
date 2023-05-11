package playerservice

import (
	"server/entities"
	"server/repository"
)

type playerService struct {
	repo repository.PlayerRepository
}

func NewPlayerService(repo repository.PlayerRepository) *playerService {
	playerService := &playerService{
		repo: repo,
	}
	return playerService
}

func (s *playerService) FirstOrCreate(player *entities.Player) error  {
	return s.repo.FirstOrCreate(player) 
}
