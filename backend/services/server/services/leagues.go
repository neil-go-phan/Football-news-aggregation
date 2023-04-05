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
}

func NewleaguesService(leagues entities.Leagues) *leaguesService{
	keyword := &leaguesService{
		leagues: leagues,	
	}
	return keyword
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
