package clubservices

import (
	"server/entities"
	"server/repository"
)

type ClubService struct {
	repo repository.ClubRepository
}

func NewClubService(repo repository.ClubRepository) *ClubService {
	clubService := &ClubService{
		repo: repo,
	}
	return clubService
}

func (s *ClubService) GetOrCreate(clubName string, logo string) (*entities.Club,error) {
	return s.repo.FirstOrCreate(clubName, logo) 
}

func (s *ClubService) GetClubByName(clubName string) (*entities.Club,error){
	return s.repo.GetByName(clubName) 
}