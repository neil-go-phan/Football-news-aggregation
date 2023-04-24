package services

import (
	"server/entities"
	"time"
	"server/services/admin"
)

type HtmlClassesServices interface {
}

type LeaguesServices interface {
	GetLeaguesNameActive() []string
	ListLeagues() entities.Leagues
	ChangeStatus(leagueName string) (bool, error)
	AddLeague(newLeaguesName string)
}

type TagsServices interface {
	ListTags() entities.Tags
	AddTag(newTags string) error
	DeleteTag(tag string) error
}

type ArticleServices interface {
	SearchArticlesTagsAndKeyword(keyword string, formatedTags []string, from int) ([]entities.Article, float64, error)
	AddTagForAllArticle(tag string) error
	GetFirstPageOfLeagueRelatedArticle(leagueName string) ([]entities.Article, error)
	RefreshCache()
	GetArticles(keywords []string)
	GetArticleCount() (total float64, today float64, err error) 
	DeleteArticle(title string) error
}

type SchedulesServices interface {
	GetAllScheduleLeagueOnDay(date time.Time) (entities.ScheduleOnDay, error)
	GetScheduleLeagueOnDay(date time.Time, league string) (entities.ScheduleOnDay, error)
	GetMatchURLsOnDay() entities.MatchURLsOnDay
	ClearMatchURLsOnDay()
	GetSchedules(date string)
	SignalMatchDetailServiceToCrawl(matchURLs entities.MatchURLsOnDay)
}

type MatchDetailServices interface {
	GetMatchDetailsOnDayFromCrawler(matchURLs entities.MatchURLsOnDay)
	GetMatchDetail(date time.Time, club1Name string, club2Name string) (entities.MatchDetail, error)
}

type AdminServices interface {
	GetAdminUsername(username string) (string, error)
	CheckAdminUsernameToken(username string) error
	ChangePassword(admin *adminservices.AdminWithConfirmPassword, usernameToken string) error
	Login(admin *adminservices.Admin) (string, error)
}

type NotificationServices interface {
	Send(title string, notiType string, message string)
}