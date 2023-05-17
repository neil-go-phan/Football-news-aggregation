// go:build wireinject
// + build wireinject
// go:generate go run github.com/google/wire/cmd/wire@latest
package infras

import (
	"server/handler"
	"server/repository"
	"server/services"

	adminservices "server/services/admin"
	articlesservices "server/services/articles"
	clubservices "server/services/club"
	configcrawler "server/services/configCrawler"
	eventservices "server/services/event"
	leaguesservices "server/services/leagues"
	lineupservices "server/services/lineup"
	matchservices "server/services/match"
	overviewitemservices "server/services/overviewItem"
	playerservices "server/services/player"
	schedulesservices "server/services/schedules"
	statsitemservices "server/services/statsItem"
	tagsservices "server/services/tags"

	pb "server/proto"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/google/wire"
	"gorm.io/gorm"
)


func InitAdmin(db *gorm.DB) *handler.AdminHandler {
	wire.Build(
		repository.NewAdminRepo,
		adminservices.NewAdminService,
		handler.NewAdminHandler,
		wire.Bind(new(repository.AdminRepository), new(*repository.AdminRepo)),
		wire.Bind(new(services.AdminServices), new(*adminservices.AdminService)),
	)
	return &handler.AdminHandler{}
}

func InitTag(db *gorm.DB) *handler.TagsHandler {
	wire.Build(
		repository.NewTagRepo,
		tagsservices.NewTagsService,
		handler.NewTagsHandler,
		wire.Bind(new(repository.TagRepository), new(*repository.TagRepo)),
		wire.Bind(new(services.TagsServices), new(*tagsservices.TagsService)),
	)
	return &handler.TagsHandler{}
}

func InitConfigCrawler(db *gorm.DB) *handler.ConfigCrawlerHandler {
	wire.Build(
		repository.NewConfigCrawlerRepo,
		configcrawler.NewConfigCrawlerService,
		handler.NewConfigCrawlerHandler,
		wire.Bind(new(repository.ConfigCrawlerRepository), new(*repository.ConfigCrawlerRepo)),
		wire.Bind(new(services.ConfigCrawlerServices), new(*configcrawler.ConfigCrawlerService)),
	)
	return &handler.ConfigCrawlerHandler{}
}

func InitLeague(db *gorm.DB) *handler.LeaguesHandler {
	wire.Build(
		repository.NewTagRepo,
		wire.Bind(new(repository.TagRepository), new(*repository.TagRepo)),
		tagsservices.NewTagsService,
		wire.Bind(new(services.TagsServices), new(*tagsservices.TagsService)),

		repository.NewLeaguesRepo,
		wire.Bind(new(repository.LeaguesRepository), new(*repository.LeaguesRepo)),
		leaguesservices.NewleaguesService,
		wire.Bind(new(services.LeaguesServices), new(*leaguesservices.LeaguesService)),

		handler.NewLeaguesHandler,
	)
	return &handler.LeaguesHandler{}
}

func InitArticle(db *gorm.DB, es *elasticsearch.Client, grpcClient pb.CrawlerServiceClient) *handler.ArticleHandler {
	wire.Build(
		repository.NewArticleRepo,
		wire.Bind(new(repository.ArticleRepository), new(*repository.ArticleRepo)),

		repository.NewTagRepo,
		wire.Bind(new(repository.TagRepository), new(*repository.TagRepo)),
		tagsservices.NewTagsService,
		wire.Bind(new(services.TagsServices), new(*tagsservices.TagsService)),

		repository.NewLeaguesRepo,
		wire.Bind(new(repository.LeaguesRepository), new(*repository.LeaguesRepo)),
		leaguesservices.NewleaguesService,
		wire.Bind(new(services.LeaguesServices), new(*leaguesservices.LeaguesService)),

		articlesservices.NewArticleService,
		wire.Bind(new(services.ArticleServices), new(*articlesservices.ArticleService)),

		handler.NewArticleHandler,
	)
	return &handler.ArticleHandler{}
}

func InitMatch(db *gorm.DB, grpcClient pb.CrawlerServiceClient) *handler.MatchDetailHandler {
	wire.Build(
		repository.NewClubRepo,
		wire.Bind(new(repository.ClubRepository), new(*repository.ClubRepo)),
		clubservices.NewClubService,
		wire.Bind(new(services.ClubServices), new(*clubservices.ClubService)),

		repository.NewStatsItemRepo,
		wire.Bind(new(repository.StatsItemlRepository), new(*repository.StatsItemRepo)),
		statsitemservices.NewStatsItemService,
		wire.Bind(new(services.StatsItemServices), new(*statsitemservices.StatsItemService)),

		repository.NewEventRepo,
		wire.Bind(new(repository.EventRepository), new(*repository.EventRepo)),
		eventservices.NewEventService,
		wire.Bind(new(services.EventServices), new(*eventservices.EventService)),

		repository.NewoOverviewItemRepo,
		wire.Bind(new(repository.OverviewItemRepository), new(*repository.OverviewItemRepo)),
		overviewitemservices.NewOverviewItemService,
		wire.Bind(new(services.OverviewItemServices), new(*overviewitemservices.OverviewItemService)),

		repository.NewLineupRepo,
		wire.Bind(new(repository.LineupRepository), new(*repository.LineupRepo)),
		lineupservices.NewLineupService,
		wire.Bind(new(services.LineUpServices), new(*lineupservices.LineupService)),

		repository.NewPlayerRepo,
		wire.Bind(new(repository.PlayerRepository), new(*repository.PlayerRepo)),
		playerservices.NewPlayerService,
		wire.Bind(new(services.PlayerServices), new(*playerservices.PlayerService)),

		repository.NewMatchRepo,
		wire.Bind(new(repository.MatchRepository), new(*repository.MatchRepo)),
		matchservices.NewMatchService,
		wire.Bind(new(services.MatchServices), new(*matchservices.MatchService)),

		handler.NewMatchDetailHandler,
	)
	return &handler.MatchDetailHandler{}
}

func InitSchedule(db *gorm.DB, es *elasticsearch.Client, grpcClient pb.CrawlerServiceClient) *handler.ScheduleHandler {
	wire.Build(
		repository.NewClubRepo,
		wire.Bind(new(repository.ClubRepository), new(*repository.ClubRepo)),
		clubservices.NewClubService,
		wire.Bind(new(services.ClubServices), new(*clubservices.ClubService)),

		repository.NewStatsItemRepo,
		wire.Bind(new(repository.StatsItemlRepository), new(*repository.StatsItemRepo)),
		statsitemservices.NewStatsItemService,
		wire.Bind(new(services.StatsItemServices), new(*statsitemservices.StatsItemService)),

		repository.NewEventRepo,
		wire.Bind(new(repository.EventRepository), new(*repository.EventRepo)),
		eventservices.NewEventService,
		wire.Bind(new(services.EventServices), new(*eventservices.EventService)),

		repository.NewoOverviewItemRepo,
		wire.Bind(new(repository.OverviewItemRepository), new(*repository.OverviewItemRepo)),
		overviewitemservices.NewOverviewItemService,
		wire.Bind(new(services.OverviewItemServices), new(*overviewitemservices.OverviewItemService)),

		repository.NewLineupRepo,
		wire.Bind(new(repository.LineupRepository), new(*repository.LineupRepo)),
		lineupservices.NewLineupService,
		wire.Bind(new(services.LineUpServices), new(*lineupservices.LineupService)),

		repository.NewPlayerRepo,
		wire.Bind(new(repository.PlayerRepository), new(*repository.PlayerRepo)),
		playerservices.NewPlayerService,
		wire.Bind(new(services.PlayerServices), new(*playerservices.PlayerService)),

		repository.NewMatchRepo,
		wire.Bind(new(repository.MatchRepository), new(*repository.MatchRepo)),
		matchservices.NewMatchService,
		wire.Bind(new(services.MatchServices), new(*matchservices.MatchService)),
		
		repository.NewTagRepo,
		wire.Bind(new(repository.TagRepository), new(*repository.TagRepo)),
		tagsservices.NewTagsService,
		wire.Bind(new(services.TagsServices), new(*tagsservices.TagsService)),

		repository.NewLeaguesRepo,
		wire.Bind(new(repository.LeaguesRepository), new(*repository.LeaguesRepo)),
		leaguesservices.NewleaguesService,
		wire.Bind(new(services.LeaguesServices), new(*leaguesservices.LeaguesService)),

		repository.NewScheduleRepo,
		wire.Bind(new(repository.SchedulesRepository), new(*repository.ScheduleRepo)),
		schedulesservices.NewSchedulesService,
		wire.Bind(new(services.SchedulesServices), new(*schedulesservices.SchedulesService)),

		handler.NewSchedulesHandler,
	)
	return &handler.ScheduleHandler{}
}
