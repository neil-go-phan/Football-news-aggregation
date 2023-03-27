package main

import (
	"backend/services/server/entities"
	"backend/services/server/services"
	"fmt"
	"log"

	"google.golang.org/grpc"
	// pb "github.com/karankumarshreds/GoProto/protofiles"
)

func main() {
	classConfig, keywordsconfig, err := readConfig();
	if err != nil {
		log.Fatalln("Fail to read config from JSON: ", err)
	}
	conn := connectToCrawler()
	services.GetArticlesWithAllKeyWords(keywordsconfig, classConfig, conn)
	fmt.Printf("%#v \n", classConfig)
	fmt.Printf("%#v \n", keywordsconfig)
	fmt.Printf("%#v \n", err)


}

func readConfig() (entities.HtmlArticleClass, entities.Keywords, error) {
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

func connectToCrawler() (*grpc.ClientConn){
		// dial server
		conn, err := grpc.Dial(":8000", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("can not connect with server %v", err)
		}
	return conn
}

