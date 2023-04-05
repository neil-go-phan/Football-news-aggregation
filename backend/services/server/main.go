package main

import (
	"fmt"
	"log"
	"net/http"
	"server/entities"
	"server/handler"
	"server/middlewares"
	"server/routes"
	"server/services"
	"time"

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
	JsonPath string `mapstructure:"JSON_PATH"`
}

func main() {
	env, err := loadEnv(".")
	if err != nil {
		log.Fatalln("cannot load env")
	}
	// load default config
	classConfig, leaguesconfig, tagsConfig, err := readConfigFromJSON(env.JsonPath)
	if err != nil {
		log.Fatalln("Fail to read config from JSON: ", err)
	}
	conn := connectToCrawler(env)

	// elastic search
	es, err := connectToElasticsearch(env)
	if err != nil {
		log.Println("error occurred while connecting to elasticsearch node: ", err)
	}
	createElaticsearchIndex(es)

	// declare services
	htmlClassesService := services.NewHtmlClassesService(classConfig)
	leaguesService := services.NewleaguesService(leaguesconfig)
	tagsService := services.NewTagsService(tagsConfig)
	articleService := services.NewArticleService(leaguesService, htmlClassesService, tagsService, conn, es)
	schedulesService := services.NewScheduleOnDayService(conn, es)

	tagsHandler := handler.NewTagsHandler(tagsService)
	tagsRoutes := routes.NewTagsRoutes(tagsHandler)

	articleHandler := handler.NewArticleHandler(articleService)
	articleRoute := routes.NewArticleRoutes(articleHandler)

	schedulesHandler := handler.NewScheduleOnDayHandler(schedulesService)
	// cronjob Setup

	go func() {
		cronjob := cron.New()

		articleHandler.SignalToCrawler(cronjob)
		schedulesHandler.SignalToCrawler(cronjob)

		cronjob.Run()
	}()

	// app routes
	log.Println("Setup routes")
	r := gin.Default()
	r.Use(middlewares.Cors())

	tagsRoutes.Setup(r)
	articleRoute.Setup(r)

	r.Run(":8080")
}

func readConfigFromJSON(JsonPath string) (entities.HtmlClasses, entities.Leagues, entities.Tags, error) {
	var classConfig entities.HtmlClasses
	var leaguesConfig entities.Leagues
	var tagsConfig entities.Tags

	classConfig, err := services.ReadHtmlClassJSON(JsonPath)
	if err != nil {
		log.Println("Fail to read htmlClassesConfig.json: ", err)
		return classConfig, leaguesConfig, tagsConfig, err
	}

	leaguesConfig, err = services.ReadleaguesJSON(JsonPath)
	if err != nil {
		log.Println("Fail to read leaguesConfig.json: ", err)
		return classConfig, leaguesConfig, tagsConfig, err
	}

	tagsConfig, err = services.ReadTagsJSON(JsonPath)
	if err != nil {
		log.Println("Fail to read tagsConfig.json: ", err)
		return classConfig, leaguesConfig, tagsConfig, err
	}

	return classConfig, leaguesConfig, tagsConfig, nil
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

func createElaticsearchIndex(es *elasticsearch.Client) {
	// check if index exist. if not exist, create new one, if exist, skip it
	indexNames := []string{"articles", "scheduleonday"}
	for _, indexName := range indexNames {
		exists, err := checkIndexExists(es, indexName)
		if err != nil {
			log.Printf("Error checking if the index %s exists: %s\n", indexName, err)
		}

		if !exists {
			log.Printf("Index: %s is not exist, create a new one...\n", indexName)
			err = createIndex(es, indexName)
			for err != nil {
				log.Printf("Error createing index: %s, try again in 10 seconds\n", err)
				time.Sleep(10 * time.Second)
				err = createIndex(es, indexName)
			}
		}
		log.Printf("Index: %s is already exist, skip it...\n", indexName)
	}

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
