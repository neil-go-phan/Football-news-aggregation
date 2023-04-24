package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"server/entities"
	"server/handler"
	serverhelper "server/helper"
	"server/middlewares"

	adminrepo "server/repository/admin"
	articlerepo "server/repository/articles"
	htmlclassesrepo "server/repository/htmlClasses"
	leaguesrepo "server/repository/leagues"
	matchdetailrepo "server/repository/matchDetail"
	// notificationrepo "server/repository/notification"
	schedulesrepo "server/repository/schedules"
	tagsrepo "server/repository/tags"
	"server/routes"
	"server/services"
	adminservices "server/services/admin"
	articlesservices "server/services/articles"
	leaguesservices "server/services/leagues"
	matchdetailservices "server/services/matchDetail"
	// notificationservices "server/services/notification"
	schedulesservices "server/services/schedules"
	tagsservices "server/services/tags"

	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/evalphobia/logrus_sentry"
	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	// setup log
	log.SetLevel(log.InfoLevel)
	format := &easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%]: %time% - %msg%\n",
	}
	log.SetFormatter(format)

  hook, err := logrus_sentry.NewSentryHook("https://4cad04fffc3348dc8d14d1f592f1d014@o4505040225501184.ingest.sentry.io/4505066672947200", []log.Level{
    log.PanicLevel,
    log.FatalLevel,
    log.ErrorLevel,
  })
	if err == nil {
    log.AddHook(hook)
  }
}

var ELASTIC_SEARCH_INDEXES = []string{articlerepo.ARTICLES_INDEX_NAME, schedulesrepo.SCHEDULE_INDEX_NAME, matchdetailrepo.MATCH_DETAIL_INDEX_NAME}

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

	// write log file
	logFile, err := os.OpenFile("serverlog.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile" + "serverlog.log")
		panic(err)
	}
	defer logFile.Close()
	log.SetOutput(io.MultiWriter(logFile, os.Stdout))

	configSentry()
	defer sentry.Flush(2 * time.Second)
	sentry.CaptureMessage("Connect to server success")
	conn := connectToCrawler(env)

	// elastic search
	es, err := connectToElasticsearch(env)
	if err != nil {
		log.Errorln("error occurred while connecting to elasticsearch node: ", err)
	}
	amountOfNewIndex := createElaticsearchIndex(es)
	// declare services
	htmlclassesRepo := htmlclassesrepo.NewHtmlClassesRepo(classConfig)

	// notificationChan := make(chan entities.Notification)

	// notificationRepo := notificationrepo.NewNotificationRepo(notificationChan)
	// notificationService := notificationservices.NewNotificationService(notificationRepo)

	tagRepo := tagsrepo.NewTagsRepo(tagsConfig, env.JsonPath)
	tagsService := tagsservices.NewTagsService(tagRepo)

	leaguesRepo := leaguesrepo.NewLeaguesRepo(leaguesconfig, env.JsonPath)
	leaguesService := leaguesservices.NewleaguesService(leaguesRepo, tagRepo)

	articleRepo := articlerepo.NewArticleRepo(leaguesRepo, htmlclassesRepo, tagRepo, conn, es)
	articleService := articlesservices.NewArticleService(articleRepo)

	matchDetailRepo := matchdetailrepo.NewMatchDetailRepo(conn, es)
	matchDetailService := matchdetailservices.NewMatchDetailervice(matchDetailRepo)

	schedulesRepo := schedulesrepo.NewSchedulesRepo(leaguesRepo, tagRepo, conn, es)
	schedulesService := schedulesservices.NewSchedulesService(schedulesRepo, matchDetailRepo)

	adminRepo := adminrepo.NewAdminRepo(adminConfig, env.JsonPath)
	adminService := adminservices.NewAdminService(adminRepo)

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

	// notificationHandler := handler.NewNotificationHandler(notificationService)
	// notificationRoute := routes.NewNotificationRoutes(notificationHandler)

	// check is this a first run to add seed data. // condition: amount of new elastic indices create = amount of elastic indices in whole app
	if amountOfNewIndex == len(ELASTIC_SEARCH_INDEXES) {
		log.Infoln("This is first time you run this project ? Please wait sometime to add seed data. It's gonna be a longtime")
		seedDataFirstRun(articleService, schedulesService, matchDetailService)
	}
	createArticleCache(articleService)

	// cronjob Setup
	go func() {
		cronjob := cron.New()

		articleHandler.SignalToCrawlerAfter10Min(cronjob)
		articleHandler.RefreshCacheAfter5Min(cronjob)
		schedulesHandler.SignalToCrawlerOnNewDay(cronjob)

		cronjob.Run()
	}()

	// app routes
	log.Infoln("Setup routes")
	r := gin.Default()
	r.Use(middlewares.Cors())

	// TODO: Fix realtime push notification
	// go notificationRoute.Setup(r)

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

	classConfig, err := serverhelper.ReadHtmlClassJSON(JsonPath)
	if err != nil {
		log.Println("Fail to read htmlClassesConfig.json: ", err)
		return classConfig, leaguesConfig, tagsConfig, adminConfig, err
	}

	leaguesConfig, err = serverhelper.ReadleaguesJSON(JsonPath)
	if err != nil {
		log.Println("Fail to read leaguesConfig.json: ", err)
		return classConfig, leaguesConfig, tagsConfig, adminConfig, err
	}

	tagsConfig, err = serverhelper.ReadTagsJSON(JsonPath)
	if err != nil {
		log.Println("Fail to read tagsConfig.json: ", err)
		return classConfig, leaguesConfig, tagsConfig, adminConfig, err
	}

	adminConfig, err = serverhelper.ReadAdminJSON(JsonPath)
	if err != nil {
		log.Println("Fail to read adminConfig.json: ", err)
		return classConfig, leaguesConfig, tagsConfig, adminConfig, err
	}

	return classConfig, leaguesConfig, tagsConfig, adminConfig, nil
}

func loadEnv(path string) (env EnvConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

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

func createElaticsearchIndex(es *elasticsearch.Client) int {
	// check if index exist. if not exist, create new one, if exist, skip it
	var amountOfNewIndex int = 0
	for _, indexName := range ELASTIC_SEARCH_INDEXES {
		exists, err := checkIndexExists(es, indexName)
		if err != nil {
			log.Printf("Error checking if the index %s exists: %s\n", indexName, err)
		}

		if !exists {
			log.Printf("Index: %s is not exist, create a new one...\n", indexName)
			if indexName != articlerepo.ARTICLES_INDEX_NAME {
				err = createIndex(es, indexName)
			} else {
				err = createArticleIndex(es, indexName)
			}
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

func createArticleIndex(es *elasticsearch.Client, indexName string) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	requestBody := map[string]interface{}{
		"settings": map[string]interface{}{
			"analysis": map[string]interface{}{
				"analyzer": map[string]interface{}{
					"no_accent": map[string]interface{}{
						"tokenizer": "standard",
						"filter": []string{
							"lowercase",
							"asciifolding",
						},
					},
				},
			},
		},
		"mappings": map[string]interface{}{
			"properties": map[string]interface{}{
				"title": map[string]interface{}{
					"type":     "text",
					"analyzer": "no_accent",
				},
				"description": map[string]interface{}{
					"type":     "text",
					"analyzer": "no_accent",
				},
			},
		},
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("error encoding article: %s", err)
	}
	req := esapi.IndicesCreateRequest{
		Index: indexName,
		Body:  strings.NewReader(string(body)),
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		return fmt.Errorf("error creating the index: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error creating index %s: %s", indexName, res.Status())
	}

	return nil
}

func seedDataFirstRun(articleService services.ArticleServices, schedulesService services.SchedulesServices, matchDetailService services.MatchDetailServices) {
	// crawl data on previous 7 days and the following 7 days
	// var wg sync.WaitGroup
	now := time.Now()
	var DAYOFWEEK = 7
	for i := -DAYOFWEEK; i <= DAYOFWEEK; i++ {
		// wg.Add(1)
		date := now.AddDate(0, 0, i)
		// go func(date time.Time, wg *sync.WaitGroup) {
			// defer wg.Done()
			schedulesService.GetSchedules(date.Format("02-01-2006"))

			matchUrls := schedulesService.GetMatchURLsOnDay()
			matchDetailService.GetMatchDetailsOnDayFromCrawler(matchUrls)
			schedulesService.ClearMatchURLsOnDay()
		// }(date, &wg)
	}
	// wg.Wait()

	// Get articles
	articleService.GetArticles(make([]string, 0))
	log.Printf("Add seed data success\n")
}

func createArticleCache(articleService services.ArticleServices) {
	articleService.RefreshCache()
}

func configSentry() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://4cad04fffc3348dc8d14d1f592f1d014@o4505040225501184.ingest.sentry.io/4505066672947200",
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	sentry.CaptureMessage("It works!")
}