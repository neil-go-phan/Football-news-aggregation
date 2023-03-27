package services

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"backend/services/server/entities"
)

func ReadHtmlClassJSON() (entities.HtmlArticleClass, error){
	var classConfig entities.HtmlArticleClass

	classConfigJson, err := os.Open("configs/htmlArticleClassConfig.json")
	if err != nil {
		log.Println(err)
		return classConfig, err
	}
	defer classConfigJson.Close()

	classConfigByte, err := io.ReadAll(classConfigJson)
	if err != nil {
		log.Println(err)
		return classConfig, err
	}

	err = json.Unmarshal(classConfigByte, &classConfig)
	if err != nil {
		log.Println(err)
		return classConfig, err
	}

	return classConfig, nil
}
