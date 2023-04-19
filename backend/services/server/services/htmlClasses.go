package services

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"server/entities"

	jsoniter "github.com/json-iterator/go"
)

type htmlClassesService struct {
	HtmlClasses entities.HtmlClasses
}

func NewHtmlClassesService(htmlClassesInput entities.HtmlClasses) *htmlClassesService {
	htmlClasses := &htmlClassesService{
		HtmlClasses: htmlClassesInput,
	}
	return htmlClasses
}

func ReadHtmlClassJSON(jsonPath string) (entities.HtmlClasses, error) {
	var classConfig entities.HtmlClasses
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	classConfigJson, err := os.Open(fmt.Sprintf("%shtmlClassesConfig.json", jsonPath))
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
