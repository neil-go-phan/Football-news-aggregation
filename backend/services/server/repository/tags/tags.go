package tagsrepo

import (
	"fmt"
	"io"
	"os"
	"server/entities"
	serverhelper "server/helper"

	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
)

type tagsRepo struct {
	Tags entities.Tags
	path string
}

func NewTagsRepo(tags entities.Tags, path string) *tagsRepo {
	tag := &tagsRepo{
		Tags: tags,
		path: path,
	}
	return tag
}

func (repo *tagsRepo) AddTag(newTags string) error {
	newTagFormated := serverhelper.FormatVietnamese(newTags)
	_, err := repo.checkTagExist(newTagFormated)
	if err == nil {
		return fmt.Errorf("tag %s already exist", newTagFormated)
	}
	repo.Tags.Tags = append(repo.Tags.Tags, newTagFormated)
	// frontend will send a request to update all tag
	err = repo.WriteTagsJSON(repo.Tags)
	if err != nil {
		log.Errorf("error occurs: %s", err)
		return err
	}
	return nil
}

func (repo *tagsRepo) DeleteTag(tag string) error {
	tagFormated := serverhelper.FormatVietnamese(tag)

	index, err := repo.checkTagExist(tagFormated)
	if err != nil {
		return err
	}
	repo.Tags.Tags = removeTag(repo.Tags.Tags, index)
	// no need to delete tag from article in elastic search, we just filter that deleted tag when query article
	err = repo.WriteTagsJSON(repo.Tags)
	if err != nil {
		log.Errorf("error occurs: %s", err)
		return err
	}
	return nil
}

func removeTag(slice []string, index int) []string {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func (repo *tagsRepo) checkTagExist(tagCheck string) (int, error) {
	for index, tag := range repo.Tags.Tags {
		if tag == tagCheck {
			return index, nil
		}
	}
	return -1, fmt.Errorf("tag %s not found", tagCheck)
}

func (repo *tagsRepo)ReadTagsJSON() (entities.Tags, error) {
	var tagsConfig entities.Tags
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	tagsConfigJson, err := os.Open(fmt.Sprintf("%stagsConfig.json", repo.path))
	if err != nil {
		log.Errorln(err)
		return tagsConfig, err
	}
	defer tagsConfigJson.Close()

	tagsConfigByte, err := io.ReadAll(tagsConfigJson)
	if err != nil {
		log.Errorln(err)
		return tagsConfig, err
	}

	err = json.Unmarshal(tagsConfigByte, &tagsConfig)
	if err != nil {
		log.Errorln(err)
		return tagsConfig, err
	}
	repo.Tags = tagsConfig
	return tagsConfig, nil
}

func (repo *tagsRepo) ListTags() entities.Tags {
	return repo.Tags
}

func (repo *tagsRepo)WriteTagsJSON(newTag entities.Tags) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	file, err := os.Create(fmt.Sprintf("%stagsConfig.json", repo.path))
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(newTag)
	if err != nil {
		return err
	}

	return nil
}