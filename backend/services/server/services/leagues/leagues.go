package leaguesservices

import (
	"server/entities"
	serverhelper "server/helper"
	"server/repository"
	"server/services"
)

type LeaguesService struct {
	leaguesRepo repository.LeaguesRepository
	tagService  services.TagsServices
}

var DEFAULT_LEAGUE_NAME = "Tin tức bóng đá"

func NewleaguesService(leaguesRepo repository.LeaguesRepository, tagService services.TagsServices) *LeaguesService {
	leaguesService := &LeaguesService{
		leaguesRepo: leaguesRepo,
		tagService:     tagService,
	}
	return leaguesService
}

func (s *LeaguesService) CreateLeague(newLeaguesName string) error {
	newLeague := &entities.League{
		LeagueName: newLeaguesName,
		Active:     false,
	}
	return s.leaguesRepo.Create(newLeague)
}

func (s *LeaguesService) ListLeagues() (*[]entities.League, error) {
	return s.leaguesRepo.List()
}

func (s *LeaguesService) GetLeaguesNameActive() ([]string, error) {
	leagueName := []string{}
	leagues, err := s.leaguesRepo.GetLeaguesNameActive()
	if err != nil {
		return leagueName, err
	}

	for _, league := range *leagues {
		leagueName = append(leagueName, league.LeagueName)
	}
	return leagueName, nil
}

func (s *LeaguesService) GetLeaguesName() ([]string, error) {
	leagueName := []string{}
	leagues, err := s.leaguesRepo.GetLeaguesName()
	if err != nil {
		return leagueName, err
	}

	for _, league := range *leagues {
		leagueName = append(leagueName, league.LeagueName)
	}
	return leagueName, nil
}

// ChangeStatus: change league.active
func (s *LeaguesService) ChangeStatus(leagueName string) (bool, error) {
	tagFromLeague := serverhelper.FormatVietnamese(leagueName)
	league, err := s.leaguesRepo.GetByName(leagueName)
	if err != nil {
		return false, err
	}
	league.Active = !league.Active

	err = s.leaguesRepo.Update(league)
	if err != nil {
		return false, err
	}
	// add new tag if league.active = true
	if league.Active {
		err := s.tagService.AddTag(tagFromLeague)
		if err != nil {
			return league.Active, err
		}
	} else {
		// delete tag if league.active = false
		err := s.tagService.DeleteTag(tagFromLeague)
		if err != nil {
			return league.Active, err
		}
	}
	return league.Active, nil
}
