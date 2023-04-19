package services

import (
	"fmt"
	"io"
		log "github.com/sirupsen/logrus"
	"os"
	"server/entities"

	jsoniter "github.com/json-iterator/go"
)

type leaguesService struct {
	leagues entities.Leagues
	tagsService *tagsService
	path    string
}

var DEFAULT_LEAGUE_NAME = "Tin tức bóng đá"

func NewleaguesService(leagues entities.Leagues,tags *tagsService, path string) *leaguesService {
	keyword := &leaguesService{
		leagues: leagues,
		tagsService: tags,
		path:    path,
	}
	return keyword
}

func (s *leaguesService) AddLeague(newLeaguesName string) {
	newLeague := entities.League{
		LeagueName: newLeaguesName,
		Active: false,
	}
	s.leagues.Leagues = append(s.leagues.Leagues, newLeague)
	err := s.WriteLeaguesJSON()
	if err != nil {
		log.Printf("error occurs: %s", err)
	}
}

func (s *leaguesService) ListLeagues() entities.Leagues {

	return s.leagues
}

func (s *leaguesService) ChangeStatus(leagueName string) (bool, error){
	for index, league := range s.leagues.Leagues {
		if league.LeagueName == leagueName {
			s.leagues.Leagues[index].Active = !league.Active
			err := s.WriteLeaguesJSON()
			if err != nil {
				return false, err
			}
			if s.leagues.Leagues[index].Active {
				err := s.tagsService.AddTag(leagueName)
				if err != nil {
					return s.leagues.Leagues[index].Active, err
				}
			} else {
				err := s.tagsService.DeleteTag(leagueName)
				if err != nil {
					return s.leagues.Leagues[index].Active, err
				}
			}
			return s.leagues.Leagues[index].Active , nil
		}
	}
	return false, fmt.Errorf("league not found")
}

func (s *leaguesService) GetLeaguesName() []string {
	names := make([]string, 0)
	for _,league := range s.leagues.Leagues {
		names = append(names, league.LeagueName)
	}
	return names
}

func (s *leaguesService) GetLeaguesNameActive() []string {
	names := make([]string, 0)
	for _,league := range s.leagues.Leagues {
		if league.Active {
			names = append(names, league.LeagueName)
		}
	}
	return names
}

func ReadleaguesJSON(jsonPath string) (entities.Leagues, error) {
	var leaguesConfig entities.Leagues
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	leaguesConfigJson, err := os.Open(fmt.Sprintf("%sleaguesConfig.json", jsonPath))
	if err != nil {
		log.Println(err)
		return leaguesConfig, err
	}
	defer leaguesConfigJson.Close()

	leaguesConfigByte, err := io.ReadAll(leaguesConfigJson)
	if err != nil {
		log.Println(err)
		return leaguesConfig, err
	}

	err = json.Unmarshal(leaguesConfigByte, &leaguesConfig)
	if err != nil {
		log.Println(err)
		return leaguesConfig, err
	}
	return leaguesConfig, nil
}

func (s *leaguesService) WriteLeaguesJSON() error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	file, err := os.Create(fmt.Sprintf("%sleaguesConfig.json", s.path))
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(s.leagues)
	if err != nil {
		return err
	}

	return nil
}
