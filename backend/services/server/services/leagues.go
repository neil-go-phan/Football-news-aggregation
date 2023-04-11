package services

import (
	"fmt"
	"io"
	"log"
	"os"
	"server/entities"

	jsoniter "github.com/json-iterator/go"
)

type leaguesService struct {
	leagues entities.Leagues
	path string
}

func NewleaguesService(leagues entities.Leagues, path string) *leaguesService{
	keyword := &leaguesService{
		leagues: leagues,	
		path: path,
	}
	return keyword
}

func (s *leaguesService)ListLeagues() (entities.Leagues) {
	return s.leagues
}

func ReadleaguesJSON(jsonPath string) (entities.Leagues, error){
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

func (s *leaguesService)WriteLeaguesJSON() error {
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