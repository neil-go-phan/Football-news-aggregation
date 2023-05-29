package infras

import (
	// "server/db/seed"
	serverproto "server/proto"
	"server/routes"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

func SetupRoute(db *gorm.DB, es *elasticsearch.Client, grpcClient serverproto.CrawlerServiceClient, r *gin.Engine, jobIDMap map[string]cron.EntryID) {
	cronjob := cron.New()

	adminHandler := InitializeAdmin(db)
	adminRoute := routes.NewAdminRoutes(adminHandler)

	tagsHandler := InitializeTag(db)
	tagsRoutes := routes.NewTagsRoutes(tagsHandler)

	leaguesHandler := InitializeLeague(db)
	leaguesRoutes := routes.NewLeaguesRoutes(leaguesHandler)

	articleHandler := InitializeArticle(db, es, grpcClient)
	articleRoute := routes.NewArticleRoutes(articleHandler)

	schedulesHandler := InitializeSchedule(db, es, grpcClient)
	schedulesRoute := routes.NewScheduleRoutes(schedulesHandler)

	matchDetailHandler := InitializeMatch(db, grpcClient)
	matchDetailRoute := routes.NewMatchDetailRoutes(matchDetailHandler)

	crawlerHandler := InitializeCrawler(db, grpcClient, cronjob, es, jobIDMap)
	crawlerRoute := routes.NewConfigCrawlerRoutes(crawlerHandler)

	cronjobHandler := InitializeCronjob(db, grpcClient, cronjob, es, jobIDMap)
	cronjobRoute := routes.NewCronjobRoutes(cronjobHandler)

	// seed.SeedData(articleHandler, schedulesHandler)
	// createArticleCache(articleHandler)

	// schedulesHandler.SignalToCrawlerToDay()

	tagsRoutes.Setup(r)
	leaguesRoutes.Setup(r)
	articleRoute.Setup(r)
	schedulesRoute.Setup(r)
	matchDetailRoute.Setup(r)
	adminRoute.Setup(r)
	crawlerRoute.Setup(r)
	cronjobRoute.Setup(r)

	// cronjob Setup

	go func() {
		cronjobHandler.CreateCronjobGetArticleFromGoogle()
		// cronjobHandler.CreateCronjobRefreshCache()
		crawlerHandler.CreateCustomCrawlerCronjob()
		schedulesHandler.SignalToCrawlerOnNewDay(cronjob, jobIDMap)

		cronjob.Run()
	}()
}
