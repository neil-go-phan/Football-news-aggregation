package services

import (
	jsoniter "github.com/json-iterator/go"
	"io"
	"log"
	"os"
	"backend/services/server/entities"
)

type keywordsService struct {
	Keywords entities.Keywords
}

func NewKeywordsService(keywords entities.Keywords) *keywordsService{
	keyword := &keywordsService{
		Keywords: keywords,	
	}
	return keyword
}


func ReadKeywordsJSON() (entities.Keywords, error){
	var keywordsConfig entities.Keywords
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	keywordsConfigJson, err := os.Open("configs/keywordsConfig.json")
	if err != nil {
		log.Println(err)
		return keywordsConfig, err
	}
	defer keywordsConfigJson.Close()

	keywordsConfigByte, err := io.ReadAll(keywordsConfigJson)
	if err != nil {
		log.Println(err)
		return keywordsConfig, err
	}

	err = json.Unmarshal(keywordsConfigByte, &keywordsConfig)
	if err != nil {
		log.Println(err)
		return keywordsConfig, err
	}
	return keywordsConfig, nil
}
