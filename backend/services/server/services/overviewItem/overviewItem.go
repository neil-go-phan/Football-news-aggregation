package overviewitem

import (
	"server/entities"
	"server/repository"
)

type OverviewItemService struct {
	repo repository.OverviewItemRepository
}

func NewOverviewItemService(repo repository.OverviewItemRepository) *OverviewItemService {
	overviewItemService := &OverviewItemService{
		repo: repo,
	}
	return overviewItemService
}

func (s *OverviewItemService) FirstOrCreate(overviewItem *entities.OverviewItem) error {
	return s.repo.FirstOrCreate(overviewItem) 
}
