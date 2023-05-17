package main

import (
	"fmt"
	"io"
	"os"
	"server/connects"
	"server/db/migrations"
	// "server/db/seed"
	"server/handler"
	serverhelper "server/helper"
	"server/middlewares"
	pb "server/proto"
	"server/repository"

	"server/infras"
	"server/routes"
	"time"
	configcrawler "server/services/configCrawler"
	"github.com/evalphobia/logrus_sentry"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"gorm.io/gorm"
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
	})
	if err == nil {
		log.AddHook(hook)
	}
}

var (
	DB  *gorm.DB
	Gin *gin.Engine
)

func main() {
	env, err := serverhelper.LoadEnv(".")
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

	// sentry
	configSentry()
	defer sentry.Flush(2 * time.Second)
	sentry.CaptureMessage("Connect to server success")

// connect crawler
	conn := connects.ConnectToCrawler(env)
	grpcClient := pb.NewCrawlerServiceClient(conn)

	// connect postgres and migration
	db := repository.ConnectDB(env.DBSource)
	migrations.RunDBMigration(env.MigrationURL, env.DBSource)

	// elastic search
	es, err := connects.ConnectToElasticsearch(env)
	if err != nil {
		log.Errorln("error occurred while connecting to elasticsearch node: ", err)
	}
	connects.CreateElaticsearchIndex(es)

	// declare services
	adminHandler := infras.InitializeAdmin(db)
	adminRoute := routes.NewAdminRoutes(adminHandler)

	tagsHandler := infras.InitializeTag(db)
	tagsRoutes := routes.NewTagsRoutes(tagsHandler)

	leaguesHandler := infras.InitializeLeague(db)
	leaguesRoutes := routes.NewLeaguesRoutes(leaguesHandler)

	articleHandler := infras.InitializeArticle(db, es, grpcClient)
	articleRoute := routes.NewArticleRoutes(articleHandler)

	schedulesHandler := infras.InitializeSchedule(db, es, grpcClient)
	schedulesRoute := routes.NewScheduleRoutes(schedulesHandler)

	matchDetailHandler := infras.InitializeMatch(db, grpcClient)
	matchDetailRoute := routes.NewMatchDetailRoutes(matchDetailHandler)

	a := repository.NewConfigCrawlerRepo(db)
	b := 	configcrawler.NewConfigCrawlerService(a, grpcClient)

	configCrawlerHandler := handler.NewConfigCrawlerHandler(b)
	configCrawlerRoute := routes.NewConfigCrawlerRoutes(configCrawlerHandler)

	// seed.SeedData(articleHandler, schedulesHandler)

	// createArticleCache(articleHandler)

	// cronjob Setup
	go func() {
		cronjob := cron.New()

		articleHandler.SignalToCrawlerAfter10Min(cronjob)
		// articleHandler.RefreshCacheAfter5Min(cronjob)
		go func() {
			schedulesHandler.SignalToCrawlerOnNewDay(cronjob)
		}()

		cronjob.Run()
	}()

	// schedulesHandler.SignalToCrawlerToDay()

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
	configCrawlerRoute.Setup(r)

	err = r.Run(env.Port)
	if err != nil {
		log.Fatalln("error occurred when run server")
	}
}

func createArticleCache(handler *handler.ArticleHandler) {
	handler.RefreshCache()
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
