package services

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"server/entities"
)

func ReadKeywordsJSON() (entities.Keywords, error){
	var keywordsConfig entities.Keywords

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
