package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"server/entities"

	
)

func main() {
	classConfig, keywordsconfig, err := readConfig();
	fmt.Printf("%#v \n", classConfig)
	fmt.Printf("%#v \n", keywordsconfig)
	fmt.Printf("%#v \n", err)
}

// Read config (admin can config via UI)
func readConfig() (entities.HtmlArticleClass, entities.Keywords, error){
	var classConfig entities.HtmlArticleClass
	var keywordsConfig entities.Keywords

	classConfigJson, err := os.Open("./configs/htmlArticleClassConfig.json")
	if err != nil {
		log.Println(err)
		return classConfig, keywordsConfig, err
	}
	defer classConfigJson.Close()

	keywordsConfigJson, err := os.Open("./configs/keywordsConfig.json")
	if err != nil {
		log.Println(err)
		return classConfig, keywordsConfig, err
	}
	defer keywordsConfigJson.Close()

	classConfigByte, err := io.ReadAll(classConfigJson)
	if err != nil {
		log.Println(err)
		return classConfig, keywordsConfig, err
	}

	keywordsConfigByte, err := io.ReadAll(keywordsConfigJson)
	if err != nil {
		log.Println(err)
		return classConfig, keywordsConfig, err
	}

	err = json.Unmarshal(classConfigByte, &classConfig)
	if err != nil {
		log.Println(err)
		return classConfig, keywordsConfig, err
	}

	err = json.Unmarshal(keywordsConfigByte, &keywordsConfig)
	if err != nil {
		log.Println(err)
		return classConfig, keywordsConfig, err
	}
	return classConfig, keywordsConfig, nil
}
