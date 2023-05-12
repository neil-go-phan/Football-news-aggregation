package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"server/entities"
	"server/handler"
	"server/middlewares"
	"server/migrations"
	"server/repository"
	pb "server/proto"

	"server/routes"
	"server/services"
	adminservices "server/services/admin"
	articlesservices "server/services/articles"
	leaguesservices "server/services/leagues"
	schedulesservices "server/services/schedules"
	tagsservices "server/services/tags"
	clubservices "server/services/club"
	statsitemservices "server/services/statsItem"
	eventservices "server/services/event"
	overviewitemservices "server/services/overviewItem"
	playerservices "server/services/player"
	lineupservices "server/services/lineup"
	matchservices "server/services/match"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/evalphobia/logrus_sentry"
	"github.com/getsentry/sentry-go"
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

var ELASTIC_SEARCH_INDEXES = []string{articlesservices.ARTICLES_INDEX_NAME}

type EnvConfig struct {
	ElasticsearchAddress string `mapstructure:"ELASTICSEARCH_ADDRESS"`
	Port                 string `mapstructure:"PORT"`
	CrawlerAddress       string `mapstructure:"CRAWLER_ADDRESS"`
	JsonPath             string `mapstructure:"JSON_PATH"`
	DBSource             string `mapstructure:"DB_SOURCE"`
}

func main() {
	env, err := loadEnv(".")

	if err != nil {
		log.Fatalln("cannot load env")
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

// connect crawler
	conn := connectToCrawler(env)
	grpcClient := pb.NewCrawlerServiceClient(conn)

	// connect postgres
	repository.ConnectDB(env.DBSource)
	db := repository.GetDB()
	err = migrations.Migrate(db)
	if err != nil {
		log.Fatal(err)
	}

	// elastic search
	es, err := connectToElasticsearch(env)
	if err != nil {
		log.Errorln("error occurred while connecting to elasticsearch node: ", err)
	}
	createElaticsearchIndex(es)

	// declare services
	clubRepo := repository.NewClubRepo(db)
	clubService := clubservices.NewClubService(clubRepo)

	statsItemRepo := repository.NewStatsItemRepo(db)
	statsItemService := statsitemservices.NewStatsItemService(statsItemRepo)

	eventRepo := repository.NewEventRepo(db)
	eventService := eventservices.NewEventService(eventRepo)

	ovewviewItemRepo := repository.NewoOverviewItemRepo(db)
	overviewItemServices := overviewitemservices.NewOverviewItemService(ovewviewItemRepo)

	playerRepo := repository.NewPlayerRepo(db)
	playerService := playerservices.NewPlayerService(playerRepo)

	lineupRepo := repository.NewLineupRepo(db)
	lineupService := lineupservices.NewLineupService(lineupRepo)

	tagRepo := repository.NewTagRepo(db)
	tagsService := tagsservices.NewTagsService(tagRepo)

	leaguesRepo := repository.NewLeaguesRepo(db)
	leaguesService := leaguesservices.NewleaguesService(leaguesRepo, tagsService)

	articleRepo := repository.NewArticleRepo(db)
	articleService := articlesservices.NewArticleService(leaguesService, tagsService,grpcClient,es, articleRepo)

	matchRepo := repository.NewMatchRepo(db)
	matchService := matchservices.NewMatchService(grpcClient,matchRepo, clubService, statsItemService, eventService, overviewItemServices, lineupService, playerService)

	schedulesRepo := repository.NewScheduleRepo(db)
	schedulesService := schedulesservices.NewSchedulesService(leaguesService, tagsService, grpcClient, es, schedulesRepo, matchService)

	adminRepo := repository.NewAdminRepo(db)
	adminService := adminservices.NewAdminService(adminRepo)

	tagsHandler := handler.NewTagsHandler(tagsService)
	tagsRoutes := routes.NewTagsRoutes(tagsHandler)

	leaguesHandler := handler.NewLeaguesHandler(leaguesService)
	leaguesRoutes := routes.NewLeaguesRoutes(leaguesHandler)

	articleHandler := handler.NewArticleHandler(articleService)
	articleRoute := routes.NewArticleRoutes(articleHandler)

	schedulesHandler := handler.NewSchedulesHandler(schedulesService)
	schedulesRoute := routes.NewScheduleRoutes(schedulesHandler)

	matchDetailHandler := handler.NewMatchDetailHandler(matchService)
	matchDetailRoute := routes.NewMatchDetailRoutes(matchDetailHandler)

	adminHandler := handler.NewAdminHandler(adminService)
	adminRoute := routes.NewAdminRoutes(adminHandler)

	createArticleCache(articleService)

	seedData(articleService, schedulesService, matchService,leaguesRepo, tagRepo, adminRepo)
	// cronjob Setup
	go func() {
		cronjob := cron.New()

		articleHandler.SignalToCrawlerAfter10Min(cronjob)
		articleHandler.RefreshCacheAfter5Min(cronjob)
		go func() {
			schedulesHandler.SignalToCrawlerOnNewDay(cronjob)
		}()

		cronjob.Run()
	}()

	schedulesHandler.SignalToCrawlerToDay()

	// app routes
	log.Infoln("Setup routes")
	r := gin.Default()
	r.Use(middlewares.Cors())

	tagsRoutes.Setup(r)
	leaguesRoutes.Setup(r)
	articleRoute.Setup(r)
	schedulesRoute.Setup(r)
	matchDetailRoute.Setup(r)
	adminRoute.Setup(r)

	err = r.Run(env.Port)
	if err != nil {
		log.Fatalln("error occurred when run server")
	}
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
			log.Printf("Error checking if the index %s exists: %s", indexName, err)
		}

		if !exists {
			log.Printf("Index: %s is not exist, create a new one...", indexName)
			if indexName != articlesservices.ARTICLES_INDEX_NAME {
				err = createIndex(es, indexName)
			} else {
				err = createArticleIndex(es, indexName)
			}
			for err != nil {
				log.Printf("Error createing index: %s, try again in 10 seconds", err)
				time.Sleep(10 * time.Second)
				err = createIndex(es, indexName)
			}
			amountOfNewIndex += 1
			continue
		}

		log.Printf("Index: %s is already exist, skip it...", indexName)
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

func seedData(articleService services.ArticleServices, schedulesService services.SchedulesServices, matchService services.MatchServices, leagueRepo repository.LeaguesRepository, tagRepo repository.TagRepository, adminRepo repository.AdminRepository) {
	
	err := leagueRepo.Create(&entities.League{
		LeagueName: "Tin tức bóng đá",
		Active: true,
	})
	if err != nil {
		log.Error(err)
	}
	err = tagRepo.Create(&entities.Tag{
		TagName: "tin tuc bong da",
	})
	if err != nil {
		log.Error(err)
	}
	err = adminRepo.Create(&entities.Admin{
		Username: "admin2023",
		Password: "fa585d89c851dd338a70dcf535aa2a92fee7836dd6aff1226583e88e0996293f16bc009c652826e0fc5c706695a03cddce372f139eff4d13959da6f1f5d3eabe",
	})
	if err != nil {
		log.Error(err)
	}
	now := time.Now()
	var DAYOFWEEK = 1

	for i := -DAYOFWEEK; i <= DAYOFWEEK; i++ {
		date := now.AddDate(0, 0, i)
		schedulesService.GetSchedules(date.Format("02-01-2006"))
		matchUrls := schedulesService.GetAllMatchURLs()

		log.Printf("seed for date: %v len: %v\n", matchUrls.Date, len(matchUrls.Urls))
		matchService.GetMatchDetailsOnDayFromCrawler(matchUrls)
		schedulesService.ClearAllMatchURLs()
	}

	// Get articles
	articleService.GetArticles(make([]string, 0))
	log.Printf("Add seed data success\n")
}

func createArticleCache(articleService services.ArticleServices) {
	articleService.RefreshCache()
}

func configSentry() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://4cad04fffc3348dc8d14d1f592f1d014@o4505040225501184.ingest.sentry.io/4505066672947200",
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}
