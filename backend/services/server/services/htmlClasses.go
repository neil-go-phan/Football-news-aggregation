package services

import (
	"backend/services/server/entities"
	"io"
	"log"
	"os"

	jsoniter "github.com/json-iterator/go"
)

type htmlClassesService struct {
	HtmlClasses entities.HtmlClasses
}

func NewHtmlClassesService(htmlClassesInput entities.HtmlClasses) *htmlClassesService{
	htmlClasses := &htmlClassesService{
		HtmlClasses: htmlClassesInput,
	}
	return htmlClasses
}

func ReadHtmlClassJSON() (entities.HtmlClasses, error){
	var classConfig entities.HtmlClasses
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	classConfigJson, err := os.Open("configs/htmlClassesConfig.json")
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
