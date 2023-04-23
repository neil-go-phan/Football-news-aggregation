package repository

import (
	"server/entities"
	"time"
	pb "server/proto"
)

//go:generate mockery --name AdminRepository
type AdminRepository interface {
	WriteAdminJSON(admin *entities.Admin) error 
	ReadAdminJSON() (entities.Admin, error)
	GetAdmin() (entities.Admin)
	SetAdmin(newAdmin entities.Admin) () 
}

//go:generate mockery --name LeaguesRepository
type LeaguesRepository interface {
	AddLeague(newLeaguesName string) 
	WriteLeaguesJSON(leagues entities.Leagues) error
	ReadleaguesJSON() (entities.Leagues, error)
	GetLeaguesNameActive() []string 
	GetLeaguesName() []string 
	GetLeagues() entities.Leagues
}

//go:generate mockery --name TagRepository
type TagRepository interface {
	ListTags() entities.Tags
	AddTag(newTags string) error
	DeleteTag(tag string) error
}

//go:generate mockery --name MatchDetailRepository
type MatchDetailRepository interface {
	GetMatchDetailsOnDayFromCrawler(matchURLs entities.MatchURLsOnDay)
	GetMatchDetail(date time.Time, club1Name string, club2Name string) (entities.MatchDetail, error)
}

//go:generate mockery --name SchedulesRepository
type SchedulesRepository interface {
	GetAllScheduleLeagueOnDay(date time.Time) (entities.ScheduleOnDay, error)
	GetScheduleLeagueOnDay(date time.Time, league string) (entities.ScheduleOnDay, error)
	GetMatchURLsOnDay() entities.MatchURLsOnDay
	ClearMatchURLsOnDay()
	GetSchedules(date *pb.Date)
}

//go:generate mockery --name HtmlClassesRepository
type HtmlClassesRepository interface {
	GetHtmlClasses() entities.HtmlClasses
}

//go:generate mockery --name ArticleRepository
type ArticleRepository interface {
	SearchArticlesTagsAndKeyword(keyword string, formatedTags []string, from int) ([]entities.Article, error)
	AddTagForAllArticle(tag string) error
	GetFirstPageOfLeagueRelatedArticle(leagueName string) ([]entities.Article, error)
	RefreshCache()
	GetArticles(keywords []string)
	GetArticleCount() (total float64, today float64, err error) 
}

//go:generate mockery --name NotificationRepository
type NotificationRepository interface {
	Send(title string, notiType string, message string)
}