package services

import (
	"fmt"
	"io"
	"log"
	"os"
	"server/entities"
	serverhelper "server/helper"

	"github.com/elastic/go-elasticsearch/v7"
	jsoniter "github.com/json-iterator/go"
)

type tagsService struct {
	Tags entities.Tags
	es   *elasticsearch.Client
	path string
}

func NewTagsService(tags entities.Tags, es *elasticsearch.Client, path string) *tagsService {
	tag := &tagsService{
		Tags: tags,
		es:   es,
		path: path,
	}
	return tag
}

func (s *tagsService) AddTag(newTags string) error {
	newTagFormated := serverhelper.FormatVietnamese(newTags)
	_, err := s.checkTagExist(newTagFormated)
	if err == nil {
		return fmt.Errorf("tag %s already exist", newTagFormated)
	}
	s.Tags.Tags = append(s.Tags.Tags, newTagFormated)

	err = s.WriteTagsJSON()
	if err != nil {
		log.Printf("error occurs: %s", err)
		return err
	}
	return nil
}

func (s *tagsService) DeleteTag(tag string) error {
	tagFormated := serverhelper.FormatVietnamese(tag)

	index, err := s.checkTagExist(tagFormated)
	if err != nil {
		return err
	}
	s.Tags.Tags = removeTag(s.Tags.Tags, index)

	err = s.WriteTagsJSON()
	if err != nil {
		log.Printf("error occurs: %s", err)
		return err
	}
	return nil
}

func removeTag(slice []string, index int) []string {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func (s *tagsService) checkTagExist(tagCheck string) (int, error) {
	for index, tag := range s.Tags.Tags {
		if tag == tagCheck {
			return index, nil
		}
	}
	return -1, fmt.Errorf("tag %s not found", tagCheck)
}

func ReadTagsJSON(jsonPath string) (entities.Tags, error) {
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

func (s *tagsService) ListTags() entities.Tags {
	return s.Tags
}

func (s *tagsService) WriteTagsJSON() error {
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
