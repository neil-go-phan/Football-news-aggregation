package leaguesservices

import (
	"fmt"
	"server/entities"
	"server/repository"
)

type leaguesService struct {
	leaguesRepo repository.LeaguesRepository
	tagRepo repository.TagRepository
}

var DEFAULT_LEAGUE_NAME = "Tin tức bóng đá"

func NewleaguesService(leaguesRepo repository.LeaguesRepository, tagRepo repository.TagRepository) *leaguesService {
	leaguesService := &leaguesService{
		leaguesRepo: leaguesRepo,
		tagRepo:        tagRepo,
	}
	return leaguesService
}

func (s *leaguesService) AddLeague(newLeaguesName string) {
	s.leaguesRepo.AddLeague(newLeaguesName)
}

func (s *leaguesService) ListLeagues() entities.Leagues {
	return s.leaguesRepo.GetLeagues()
}

func (s *leaguesService)GetLeaguesNameActive() []string{
	return s.leaguesRepo.GetLeaguesNameActive()
}

func (s *leaguesService) ChangeStatus(leagueName string) (bool, error) {
	repoLeagues := s.leaguesRepo.GetLeagues()
	for index, league := range repoLeagues.Leagues {
		if league.LeagueName == leagueName {
			repoLeagues.Leagues[index].Active = !league.Active
			err := s.leaguesRepo.WriteLeaguesJSON(repoLeagues)
			if err != nil {
				return false, err
			}
			if repoLeagues.Leagues[index].Active {
				err := s.tagRepo.AddTag(leagueName)
				if err != nil {
					return repoLeagues.Leagues[index].Active, err
				}
			} else {
				err := s.tagRepo.DeleteTag(leagueName)
				if err != nil {
					return repoLeagues.Leagues[index].Active, err
				}
			}
			return repoLeagues.Leagues[index].Active, nil
		}
	}
	return false, fmt.Errorf("league not found")
}
