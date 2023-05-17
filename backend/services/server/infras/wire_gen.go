// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package infras

import (
	"github.com/elastic/go-elasticsearch/v7"
	"gorm.io/gorm"
	"server/handler"
	"server/proto"
	"server/repository"
	"server/services/admin"
	"server/services/articles"
	"server/services/club"
	// "server/services/configCrawler"
	"server/services/event"
	"server/services/leagues"
	"server/services/lineup"
	"server/services/match"
	"server/services/overviewItem"
	"server/services/player"
	"server/services/schedules"
	"server/services/statsItem"
	"server/services/tags"
)

// Injectors from wire.go:

func InitializeAdmin(db *gorm.DB) *handler.AdminHandler {
	adminRepo := repository.NewAdminRepo(db)
	adminService := adminservices.NewAdminService(adminRepo)
	adminHandler := handler.NewAdminHandler(adminService)
	return adminHandler
}

func InitializeTag(db *gorm.DB) *handler.TagsHandler {
	tagRepo := repository.NewTagRepo(db)
	tagsService := tagsservices.NewTagsService(tagRepo)
	tagsHandler := handler.NewTagsHandler(tagsService)
	return tagsHandler
}

// func InitializeConfigCrawler(db *gorm.DB) *handler.ConfigCrawlerHandler {
// 	configCrawlerRepo := repository.NewConfigCrawlerRepo(db)
// 	configCrawlerService := configcrawler.NewConfigCrawlerService(configCrawlerRepo)
// 	configCrawlerHandler := handler.NewConfigCrawlerHandler(configCrawlerService)
// 	return configCrawlerHandler
// }

func InitializeLeague(db *gorm.DB) *handler.LeaguesHandler {
	leaguesRepo := repository.NewLeaguesRepo(db)
	tagRepo := repository.NewTagRepo(db)
	tagsService := tagsservices.NewTagsService(tagRepo)
	leaguesService := leaguesservices.NewleaguesService(leaguesRepo, tagsService)
	leaguesHandler := handler.NewLeaguesHandler(leaguesService)
	return leaguesHandler
}

func InitializeArticle(db *gorm.DB, es *elasticsearch.Client, grpcClient serverproto.CrawlerServiceClient) *handler.ArticleHandler {
	leaguesRepo := repository.NewLeaguesRepo(db)
	tagRepo := repository.NewTagRepo(db)
	tagsService := tagsservices.NewTagsService(tagRepo)
	leaguesService := leaguesservices.NewleaguesService(leaguesRepo, tagsService)
	articleRepo := repository.NewArticleRepo(db)
	articleService := articlesservices.NewArticleService(leaguesService, tagsService, grpcClient, es, articleRepo)
	articleHandler := handler.NewArticleHandler(articleService)
	return articleHandler
}

func InitializeMatch(db *gorm.DB, grpcClient serverproto.CrawlerServiceClient) *handler.MatchDetailHandler {
	matchRepo := repository.NewMatchRepo(db)
	clubRepo := repository.NewClubRepo(db)
	clubService := clubservices.NewClubService(clubRepo)
	statsItemRepo := repository.NewStatsItemRepo(db)
	statsItemService := statsitem.NewStatsItemService(statsItemRepo)
	eventRepo := repository.NewEventRepo(db)
	eventService := eventservice.NewEventService(eventRepo)
	overviewItemRepo := repository.NewoOverviewItemRepo(db)
	overviewItemService := overviewitem.NewOverviewItemService(overviewItemRepo)
	lineupRepo := repository.NewLineupRepo(db)
	lineupService := lineupservice.NewLineupService(lineupRepo)
	playerRepo := repository.NewPlayerRepo(db)
	playerService := playerservice.NewPlayerService(playerRepo)
	matchService := matchservices.NewMatchService(grpcClient, matchRepo, clubService, statsItemService, eventService, overviewItemService, lineupService, playerService)
	matchDetailHandler := handler.NewMatchDetailHandler(matchService)
	return matchDetailHandler
}

func InitializeSchedule(db *gorm.DB, es *elasticsearch.Client, grpcClient serverproto.CrawlerServiceClient) *handler.ScheduleHandler {
	leaguesRepo := repository.NewLeaguesRepo(db)
	tagRepo := repository.NewTagRepo(db)
	tagsService := tagsservices.NewTagsService(tagRepo)
	leaguesService := leaguesservices.NewleaguesService(leaguesRepo, tagsService)
	scheduleRepo := repository.NewScheduleRepo(db)
	matchRepo := repository.NewMatchRepo(db)
	clubRepo := repository.NewClubRepo(db)
	clubService := clubservices.NewClubService(clubRepo)
	statsItemRepo := repository.NewStatsItemRepo(db)
	statsItemService := statsitem.NewStatsItemService(statsItemRepo)
	eventRepo := repository.NewEventRepo(db)
	eventService := eventservice.NewEventService(eventRepo)
	overviewItemRepo := repository.NewoOverviewItemRepo(db)
	overviewItemService := overviewitem.NewOverviewItemService(overviewItemRepo)
	lineupRepo := repository.NewLineupRepo(db)
	lineupService := lineupservice.NewLineupService(lineupRepo)
	playerRepo := repository.NewPlayerRepo(db)
	playerService := playerservice.NewPlayerService(playerRepo)
	matchService := matchservices.NewMatchService(grpcClient, matchRepo, clubService, statsItemService, eventService, overviewItemService, lineupService, playerService)
	schedulesService := schedulesservices.NewSchedulesService(leaguesService, tagsService, grpcClient, es, scheduleRepo, matchService)
	scheduleHandler := handler.NewSchedulesHandler(schedulesService)
	return scheduleHandler
}
