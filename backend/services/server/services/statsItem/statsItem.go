package statsitem

import (
	"server/entities"
	"server/repository"
)

type StatsItemService struct {
	repo repository.StatsItemlRepository
}

func NewStatsItemService(repo repository.StatsItemlRepository) *StatsItemService {
	statsItemService := &StatsItemService{
		repo: repo,
	}
	return statsItemService
}

func (s *StatsItemService) FirstOrCreate(statsItem *entities.StatisticsItem) (error) {
	return s.repo.FirstOrCreate(statsItem) 
}
