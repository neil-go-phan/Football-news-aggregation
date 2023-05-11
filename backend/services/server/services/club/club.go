package clubservices

import (
	"server/entities"
	"server/repository"
)

type clubService struct {
	repo repository.ClubRepository
}

func NewClubService(repo repository.ClubRepository) *clubService {
	clubService := &clubService{
		repo: repo,
	}
	return clubService
}

func (s *clubService) GetOrCreate(clubName string, logo string) (*entities.Club,error) {
	return s.repo.FirstOrCreate(clubName, logo) 
}

func (s *clubService) GetClubByName(clubName string) (*entities.Club,error){
	return s.repo.GetByName(clubName) 
}