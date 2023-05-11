package overviewitem

import (
	"server/entities"
	"server/repository"
)

type overviewItemService struct {
	repo repository.OverviewItemRepository
}

func NewOverviewItemService(repo repository.OverviewItemRepository) *overviewItemService {
	overviewItemService := &overviewItemService{
		repo: repo,
	}
	return overviewItemService
}

func (s *overviewItemService) FirstOrCreate(overviewItem *entities.OverviewItem) error {
	return s.repo.FirstOrCreate(overviewItem) 
}
