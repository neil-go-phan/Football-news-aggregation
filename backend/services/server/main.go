package main

import (
	"backend/services/server/entities"
	"backend/services/server/handler"
	"backend/services/server/helper"
	"backend/services/server/middlewares"
	"backend/services/server/routes"
	"backend/services/server/services"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type EnvConfig struct {
	ElasticsearchAddress string `mapstructure:"ELASTICSEARCH_ADDRESS"`
	Port                 string `mapstructure:"PORT"`
	CrawlerAddress       string `mapstructure:"CRAWLER_ADDRESS"`
}

func main() {
	env, err := loadEnv(".")
	if err != nil {
		log.Fatalln("cannot load env")
	}
	// load default config
	classConfig, keywordsconfig, tagsConfig, err := readConfigFromJSON()
	if err != nil {
		log.Fatalln("Fail to read config from JSON: ", err)
	}
	conn := connectToCrawler(env)

	// elastic search
	es, err := connectToElasticsearch(env)
	if err != nil {
		log.Println("error occurred while connecting to elasticsearch node: ", err)
	}
	createElaticsearchIndex(es, keywordsconfig)

	// declare services
	htmlClassesService := services.NewHtmlClassesService(classConfig)
	keywordsService := services.NewKeywordsService(keywordsconfig)
	tagsService := services.NewTagsService(tagsConfig)
	articleService := services.NewArticleService(keywordsService, htmlClassesService, tagsService, conn, es)
	articleHandler := handler.NewArticleHandler(articleService)
	articleRoute := routes.NewArticleRoutes(articleHandler)
	// cronjob Setup

	go func() {
		cronjob := cron.New()

		articleHandler.SignalToCrawler(cronjob)
		cronjob.Run()
	}()

	// app routes
	log.Println("Setup routes")
	r := gin.Default()
	r.Use(middlewares.Cors())

	articleRoute.Setup(r)

	r.Run(":8080")
}

func readConfigFromJSON() (entities.HtmlClasses, entities.Keywords, entities.Tags, error) {
	var classConfig entities.HtmlClasses
	var keywordsConfig entities.Keywords
	var tagsConfig entities.Tags

	classConfig, err := services.ReadHtmlClassJSON()
	if err != nil {
		log.Println("Fail to read htmlClassesConfig.json: ", err)
		return classConfig, keywordsConfig, tagsConfig, err
	}

	keywordsConfig, err = services.ReadKeywordsJSON()
	if err != nil {
		log.Println("Fail to read keywordsConfig.json: ", err)
		return classConfig, keywordsConfig, tagsConfig, err
	}

	tagsConfig, err = services.ReadTagsJSON()
	if err != nil {
		log.Println("Fail to read tagsConfig.json: ", err)
		return classConfig, keywordsConfig, tagsConfig, err
	}

	return classConfig, keywordsConfig, tagsConfig, nil
}

func loadEnv(path string) (env EnvConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&env)
	return
}

func connectToCrawler(env EnvConfig) *grpc.ClientConn {
	// dial server
	conn, err := grpc.Dial(env.CrawlerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}
	return conn
}

func connectToElasticsearch(env EnvConfig) (*elasticsearch.Client, error) {
	var esConfig = elasticsearch.Config{
		Addresses: []string{
			env.ElasticsearchAddress,
		},
	}

	es, err := elasticsearch.NewClient(esConfig)
	if err != nil {
		return nil, err
	}
	return es, nil
}

func createElaticsearchIndex(es *elasticsearch.Client, keywords entities.Keywords) {
	var wg sync.WaitGroup
	for _, elasticIndex := range keywords.Keywords {
		wg.Add(1)
		go func(elasticIndex string) {
			indexName := helper.FormatElasticSearchIndexName(elasticIndex)
			// check if index exist. if not exist, create new one, if exist, skip it
			exists, err := checkIndexExists(es, indexName)
			if err != nil {
				log.Printf("Error checking if the index %s exists: %s\n", indexName, err)
			}

			if !exists {
				log.Printf("Index: %s is not exist, create a new one...\n", indexName)
				err = createIndex(es, indexName)
				if err != nil {
					log.Fatalf("Error createing index: %s", err)
				}
				wg.Done()
				return
			}
			log.Printf("Index: %s is already exist, skip it...\n", indexName)
			wg.Done()
		}(elasticIndex)
	}

	wg.Wait()
}

func checkIndexExists(es *elasticsearch.Client, indexName string) (bool, error) {
	res, err := es.Indices.Exists([]string{indexName})
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, nil
}

func createIndex(es *elasticsearch.Client, indexName string) error {
	res, err := es.Indices.Create(indexName)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error creating index %s: %s", indexName, res.Status())
	}

	return nil
}
