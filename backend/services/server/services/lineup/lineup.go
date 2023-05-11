package lineupservice

import (
	"server/entities"
	"server/repository"
)

type lineupService struct {
	repo repository.LineupRepository
}

func NewLineupService(repo repository.LineupRepository) *lineupService {
	lineupService := &lineupService{
		repo: repo,
	}
	return lineupService
}

func (s *lineupService) GetOrCreate(lineup *entities.MatchLineUp) (*entities.MatchLineUp, error) {
	return s.repo.FirstOrCreate(lineup) 
}

func (s *lineupService) GetLineUps(id1 uint, id2 uint) (*entities.MatchLineUp, *entities.MatchLineUp, error) {
	lineup1 , err := s.repo.Get(id1)
	if err != nil {
		return nil,nil, err
	}
	lineup2 , err := s.repo.Get(id2)
	if err != nil {
		return nil,nil, err
	}
	return lineup1, lineup2, nil
}