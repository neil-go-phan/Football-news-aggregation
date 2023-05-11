package statsitem

import (
	"server/entities"
	"server/repository"
)

type statsItemService struct {
	repo repository.StatsItemlRepository
}

func NewStatsItemService(repo repository.StatsItemlRepository) *statsItemService {
	statsItemService := &statsItemService{
		repo: repo,
	}
	return statsItemService
}

func (s *statsItemService) FirstOrCreate(statsItem *entities.StatisticsItem) (error) {
	return s.repo.FirstOrCreate(statsItem) 
}
