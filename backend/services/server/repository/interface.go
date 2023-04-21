package repository

import (
	"server/entities"
	"time"
	pb "server/proto"
)

type AdminRepository interface {
	WriteAdminJSON(admin *entities.Admin) error 
	ReadAdminJSON() (entities.Admin, error)
	GetAdmin() (entities.Admin)
	SetAdmin(newAdmin entities.Admin) () 
}

type LeaguesRepository interface {
	AddLeague(newLeaguesName string) 
	WriteLeaguesJSON(leagues entities.Leagues) error
	ReadleaguesJSON() (entities.Leagues, error)
	GetLeaguesNameActive() []string 
	GetLeaguesName() []string 
	GetLeagues() entities.Leagues
}

type TagRepository interface {
	ListTags() entities.Tags
	AddTag(newTags string) error
	DeleteTag(tag string) error
}

type MatchDetailRepository interface {
	GetMatchDetailsOnDayFromCrawler(matchURLs entities.MatchURLsOnDay)
	GetMatchDetail(date time.Time, club1Name string, club2Name string) (entities.MatchDetail, error)
}

type SchedulesRepository interface {
	GetAllScheduleLeagueOnDay(date time.Time) (entities.ScheduleOnDay, error)
	GetScheduleLeagueOnDay(date time.Time, league string) (entities.ScheduleOnDay, error)
	GetMatchURLsOnDay() entities.MatchURLsOnDay
	ClearMatchURLsOnDay()
	GetSchedules(date *pb.Date)
}

type HtmlClassesRepository interface {
	GetHtmlClasses() entities.HtmlClasses
}

type ArticleRepository interface {
	SearchArticlesTagsAndKeyword(keyword string, formatedTags []string, from int) ([]entities.Article, error)
	AddTagForAllArticle(tag string) error
	GetFirstPageOfLeagueRelatedArticle(leagueName string) ([]entities.Article, error)
	RefreshCache()
	GetArticles(keywords []string)
	GetArticleCount() (total float64, today float64, err error) 
}

type NotificationRepository interface {
	Send(title string, notiType string, message string)
}