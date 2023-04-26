package leaguesrepo

import (
	"fmt"
	"io"
	"os"
	"server/entities"

	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
)

type leaguesRepo struct {
	leagues entities.Leagues
	path    string
}

var DEFAULT_LEAGUE_NAME = "Tin tức bóng đá"

func NewLeaguesRepo(leagues entities.Leagues, path string) *leaguesRepo {
	leaguesRepo := &leaguesRepo{
		leagues: leagues,
		path:    path,
	}
	return leaguesRepo
}

func (repo *leaguesRepo) GetLeagues() entities.Leagues {
	return repo.leagues
}

func (repo *leaguesRepo) GetLeaguesName() []string {
	names := make([]string, 0)
	for _,league := range repo.leagues.Leagues {
		names = append(names, league.LeagueName)
	}
	return names
}

func (repo *leaguesRepo) GetLeaguesNameActive() []string {
	names := make([]string, 0)
	for _,league := range repo.leagues.Leagues {
		if league.Active {
			names = append(names, league.LeagueName)
		}
	}
	return names
}

func (repo *leaguesRepo)AddLeague(newLeaguesName string) {
	newLeague := entities.League{
		LeagueName: newLeaguesName,
		Active:     false,
	}
	repo.leagues.Leagues = append(repo.leagues.Leagues, newLeague)
	err := repo.WriteLeaguesJSON(repo.leagues)
	if err != nil {
		log.Printf("error occurs: %s", err)
	}
}


func (repo *leaguesRepo)ReadleaguesJSON() (entities.Leagues, error) {
	var leaguesConfig entities.Leagues
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	leaguesConfigJson, err := os.Open(fmt.Sprintf("%sleaguesConfig.json", repo.path))
	if err != nil {
		log.Errorln(err)
		return leaguesConfig, fmt.Errorf("file json not found")
	}
	defer leaguesConfigJson.Close()

	leaguesConfigByte, err := io.ReadAll(leaguesConfigJson)
	if err != nil {
		log.Errorln(err)
		return leaguesConfig, err
	}

	err = json.Unmarshal(leaguesConfigByte, &leaguesConfig)
	if err != nil {
		log.Errorln(err)
		return leaguesConfig, err
	}
	repo.leagues = leaguesConfig
	return leaguesConfig, nil
}

func (repo *leaguesRepo) WriteLeaguesJSON(leagues entities.Leagues) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	file, err := os.Create(fmt.Sprintf("%sleaguesConfig.json", repo.path))
	if err != nil {
		return fmt.Errorf("file json not found")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(leagues)
	if err != nil {
		return err
	}
	repo.leagues = leagues
	return nil
}
