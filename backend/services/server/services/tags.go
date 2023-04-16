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
	path string
}

func NewTagsService(tags entities.Tags, path string) *tagsService{
	tag := &tagsService{
		Tags: tags,	
		path:path,
	}
	return tag
}

func (s *tagsService)AddTag(newTags []string) {
 	s.Tags.Tags = append(s.Tags.Tags, newTags...)
	s.WriteTagsJSON()
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

func (s *tagsService)WriteTagsJSON() error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	file, err := os.Create(fmt.Sprintf("%stagsConfig.json", s.path))
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(s.Tags)
	if err != nil {
		return err
	}

	return nil
}