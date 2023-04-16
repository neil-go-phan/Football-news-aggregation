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
	"sync"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ELASTIC_SEARCH_INDEXES = []string{services.ARTICLES_INDEX_NAME, services.SCHEDULE_INDEX_NAME, services.MATCH_DETAIL_INDEX_NAME}

type EnvConfig struct {
	ElasticsearchAddress string `mapstructure:"ELASTICSEARCH_ADDRESS"`
	Port                 string `mapstructure:"PORT"`
	CrawlerAddress       string `mapstructure:"CRAWLER_ADDRESS"`
	JsonPath             string `mapstructure:"JSON_PATH"`
}

func main() {
	env, err := loadEnv(".")
	if err != nil {
		log.Fatalln("cannot load env")
	}
	// load default config
	classConfig, leaguesconfig, tagsConfig, adminConfig, err := readConfigFromJSON(env.JsonPath)
	if err != nil {
		log.Fatalln("Fail to read config from JSON: ", err)
	}
	conn := connectToCrawler(env)

	// elastic search
	es, err := connectToElasticsearch(env)
	if err != nil {
		log.Println("error occurred while connecting to elasticsearch node: ", err)
	}
	amountOfNewIndex := createElaticsearchIndex(es)
	// createElaticsearchIndex(es)
	// declare services
	htmlClassesService := services.NewHtmlClassesService(classConfig)
	leaguesService := services.NewleaguesService(leaguesconfig, env.JsonPath)
	tagsService := services.NewTagsService(tagsConfig, env.JsonPath)
	articleService := services.NewArticleService(leaguesService, htmlClassesService, tagsService, conn, es)
	schedulesService := services.NewSchedulesService(leaguesService, conn, es)
	matchDetailService := services.NewMatchDetailervice(conn, es, articleService)
	adminService := services.NewAdminService(adminConfig, env.JsonPath)

	tagsHandler := handler.NewTagsHandler(tagsService)
	tagsRoutes := routes.NewTagsRoutes(tagsHandler)

	leaguesHandler := handler.NewLeaguesHandler(leaguesService)
	leaguesRoutes := routes.NewLeaguesRoutes(leaguesHandler)

	articleHandler := handler.NewArticleHandler(articleService)
	articleRoute := routes.NewArticleRoutes(articleHandler)

	schedulesHandler := handler.NewSchedulesHandler(schedulesService)
	schedulesRoute := routes.NewScheduleRoutes(schedulesHandler)

	matchDetailHandler := handler.NewMatchDetailHandler(matchDetailService)
	matchDetailRoute := routes.NewMatchDetailRoutes(matchDetailHandler)

	adminHandler := handler.NewAdminHandler(adminService)
	adminRoute := routes.NewAdminRoutes(adminHandler)

  // check is this a first run to add seed data. // condition: amount of new elastic indices create = amount of elastic indices in whole app
	if amountOfNewIndex == len(ELASTIC_SEARCH_INDEXES) {
		log.Println("This is first time you run this project ? Please wait sometime to add seed data. It's gonna be a longtime")
		seedDataFirstRun(articleService, schedulesService, matchDetailService)
	}
	// articleService.GetArticles(make([]string, 0))

	// cronjob Setup
	go func() {
		cronjob := cron.New()

		// articleHandler.SignalToCrawlerAfter10Min(cronjob)
		schedulesHandler.SignalToCrawlerOnNewDay(cronjob)

		cronjob.Run()
	}()

	// app routes
	log.Println("Setup routes")
	r := gin.Default()
	r.Use(middlewares.Cors())

	tagsRoutes.Setup(r)
	leaguesRoutes.Setup(r)
	articleRoute.Setup(r)
	schedulesRoute.Setup(r)
	matchDetailRoute.Setup(r)
	adminRoute.Setup(r)

	err = r.Run(":8080")
	if err != nil {
		log.Fatalln("error occurred when run server")
	}
}

func readConfigFromJSON(JsonPath string) (entities.HtmlClasses, entities.Leagues, entities.Tags, entities.Admin, error) {
	var classConfig entities.HtmlClasses
	var leaguesConfig entities.Leagues
	var tagsConfig entities.Tags
	var adminConfig entities.Admin

	classConfig, err := services.ReadHtmlClassJSON(JsonPath)
	if err != nil {
		log.Println("Fail to read htmlClassesConfig.json: ", err)
		return classConfig, leaguesConfig, tagsConfig, adminConfig, err
	}

	leaguesConfig, err = services.ReadleaguesJSON(JsonPath)
	if err != nil {
		log.Println("Fail to read leaguesConfig.json: ", err)
		return classConfig, leaguesConfig, tagsConfig, adminConfig, err
	}

	tagsConfig, err = services.ReadTagsJSON(JsonPath)
	if err != nil {
		log.Println("Fail to read tagsConfig.json: ", err)
		return classConfig, leaguesConfig, tagsConfig, adminConfig, err
	}

	adminConfig, err = services.ReadAdminJSON(JsonPath)
	if err != nil {
		log.Println("Fail to read adminConfig.json: ", err)
		return classConfig, leaguesConfig, tagsConfig, adminConfig, err
	}

	return classConfig, leaguesConfig, tagsConfig, adminConfig, nil
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

func createElaticsearchIndex(es *elasticsearch.Client) int{
	// check if index exist. if not exist, create new one, if exist, skip it
	var amountOfNewIndex int = 0
	for _, indexName := range ELASTIC_SEARCH_INDEXES {
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
			amountOfNewIndex += 1
			continue
		}

		log.Printf("Index: %s is already exist, skip it...\n", indexName)
	}
	return amountOfNewIndex
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

func seedDataFirstRun(articleService services.ArticleServices, schedulesService services.SchedulesServices, matchDetailService services.MatchDetailServices) {
	// Get schedule from january to current month + 2
	year, currentMonth, _ := time.Now().Date()
	var wg sync.WaitGroup

	for month := 4; month <= int(currentMonth); month++ {
		// loop each day in current month
		wg.Add(1)
		t := time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.UTC)
		go func(t time.Time, month int) {
			for day := 1; day <= t.Day(); day++ {
				date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
				schedulesService.GetSchedules(date.Format("02-01-2006"))

				matchUrls := schedulesService.GetMatchURLsOnDay()
				matchDetailService.GetMatchDetailsOnDayFromCrawler(matchUrls)
				schedulesService.ClearMatchURLsOnDay()
			}
			defer wg.Done()
		}(t, month)

	}
	wg.Wait()

	// Get articles
	articleService.GetArticles(make([]string, 0))
}
