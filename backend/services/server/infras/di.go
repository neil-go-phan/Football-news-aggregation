package infras

import (
	"server/db/seed"
	"server/handler"
	serverproto "server/proto"
	"server/routes"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

func SetupRoute(db *gorm.DB, es *elasticsearch.Client, grpcClient serverproto.CrawlerServiceClient, r *gin.Engine) {
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

	configCrawlerHandler := InitializeConfigCrawler(db, grpcClient)
	configCrawlerRoute := routes.NewConfigCrawlerRoutes(configCrawlerHandler)

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

	seed.SeedData(articleHandler, schedulesHandler)
	createArticleCache(articleHandler)

	schedulesHandler.SignalToCrawlerToDay()

	tagsRoutes.Setup(r)
	leaguesRoutes.Setup(r)
	articleRoute.Setup(r)
	schedulesRoute.Setup(r)
	matchDetailRoute.Setup(r)
	adminRoute.Setup(r)
	configCrawlerRoute.Setup(r)
}

func createArticleCache(handler *handler.ArticleHandler) {
	handler.RefreshCache()
}