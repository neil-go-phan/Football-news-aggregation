package main

import (
	"fmt"
	"log"
	"server/entities"
	"server/services"

	// pb "github.com/karankumarshreds/GoProto/protofiles"

)

func main() {
	classConfig, keywordsconfig, err := readConfigFromJSON();
	if err != nil {
		log.Fatalln("Fail to read config from JSON: ", err)
	}

	fmt.Printf("%#v \n", classConfig)
	fmt.Printf("%#v \n", keywordsconfig)
	fmt.Printf("%#v \n", err)


}

func readConfigFromJSON() (entities.HtmlArticleClass, entities.Keywords, error) {
	var classConfig entities.HtmlArticleClass
	var keywordsConfig entities.Keywords

	classConfig, err := services.ReadHtmlClassJSON();
	if err != nil {
		log.Println("Fail to read htmlArticleClassConfig.json: ", err)
		return classConfig, keywordsConfig, err
	}

	keywordsconfig, err := services.ReadKeywordsJSON();
	if err != nil {
		log.Println("Fail to read keywordsConfig.json: ", err)
		return classConfig, keywordsconfig, err
	}

	return classConfig, keywordsconfig, nil
}
