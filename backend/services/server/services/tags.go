package services

import (
	"fmt"
	"io"
	"log"
	"os"
	"server/entities"

	jsoniter "github.com/json-iterator/go"
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


func ReadTagsJSON(jsonPath string) (entities.Tags, error){
	var tagsConfig entities.Tags
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	tagsConfigJson, err := os.Open(fmt.Sprintf("%stagsConfig.json", jsonPath))
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

func (s *tagsService)ListTags() (entities.Tags) {
	return s.Tags
}