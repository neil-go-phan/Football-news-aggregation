package services

import (
	jsoniter "github.com/json-iterator/go"
	"io"
	"log"
	"os"
	"backend/services/server/entities"
)

type tagsService struct {
	Tags entities.Tags
}

func NewTagsService(tags entities.Tags) *tagsService{
	tag := &tagsService{
		Tags: tags,	
	}
	return tag
}


func ReadTagsJSON() (entities.Tags, error){
	var tagsConfig entities.Tags
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	tagsConfigJson, err := os.Open("configs/tagsConfig.json")
	if err != nil {
		log.Println(err)
		return tagsConfig, err
	}
	defer tagsConfigJson.Close()

	tagsConfigByte, err := io.ReadAll(tagsConfigJson)
	if err != nil {
		log.Println(err)
		return tagsConfig, err
	}

	err = json.Unmarshal(tagsConfigByte, &tagsConfig)
	if err != nil {
		log.Println(err)
		return tagsConfig, err
	}
	return tagsConfig, nil
}
