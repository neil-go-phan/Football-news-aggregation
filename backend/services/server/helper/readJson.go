package serverhelper

import (
	"fmt"
	"io"
	"os"
	"server/entities"
	log "github.com/sirupsen/logrus"
	jsoniter "github.com/json-iterator/go"
)

func ReadAdminJSON(jsonPath string) (entities.Admin, error) {
	var adminConfig entities.Admin
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	adminConfigJson, err := os.Open(fmt.Sprintf("%sadminConfig.json", jsonPath))
	if err != nil {
		log.Errorln(err)
		return adminConfig, fmt.Errorf("file json not found")
	}
	defer adminConfigJson.Close()

	adminConfigByte, err := io.ReadAll(adminConfigJson)
	if err != nil {
		log.Errorln(err)
		return adminConfig, err
	}

	err = json.Unmarshal(adminConfigByte, &adminConfig)
	if err != nil {
		log.Errorln(err)
		return adminConfig, err
	}
	return adminConfig, nil
}

func ReadleaguesJSON(jsonPath string) (entities.Leagues, error) {
	var leaguesConfig entities.Leagues
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	leaguesConfigJson, err := os.Open(fmt.Sprintf("%sleaguesConfig.json", jsonPath))
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
	return leaguesConfig, nil
}

func ReadTagsJSON(jsonPath string) (entities.Tags, error) {
	var tagsConfig entities.Tags
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	tagsConfigJson, err := os.Open(fmt.Sprintf("%stagsConfig.json", jsonPath))
	if err != nil {
		log.Errorln(err)
		return tagsConfig, fmt.Errorf("file json not found")
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
	return tagsConfig, nil
}

func ReadHtmlClassJSON(jsonPath string) (entities.HtmlClasses, error) {
	var classConfig entities.HtmlClasses
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	classConfigJson, err := os.Open(fmt.Sprintf("%shtmlClassesConfig.json", jsonPath))
	if err != nil {
		log.Errorln(err)
		return classConfig, fmt.Errorf("file json not found")
	}
	defer classConfigJson.Close()

	classConfigByte, err := io.ReadAll(classConfigJson)
	if err != nil {
		log.Errorln(err)
		return classConfig, err
	}

	err = json.Unmarshal(classConfigByte, &classConfig)
	if err != nil {
		log.Errorln(err)
		return classConfig, err
	}

	return classConfig, nil
}