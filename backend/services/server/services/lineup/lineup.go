package lineupservice

import (
	"server/entities"
	"server/repository"
)

type LineupService struct {
	repo repository.LineupRepository
}

func NewLineupService(repo repository.LineupRepository) *LineupService {
	lineupService := &LineupService{
		repo: repo,
	}
	return lineupService
}

func (s *LineupService) GetOrCreate(lineup *entities.MatchLineUp) (*entities.MatchLineUp, error) {
	return s.repo.FirstOrCreate(lineup) 
}

func (s *LineupService) GetLineUps(id1 uint, id2 uint) (*entities.MatchLineUp, *entities.MatchLineUp, error) {
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