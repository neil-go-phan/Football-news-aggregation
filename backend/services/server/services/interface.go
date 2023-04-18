package services

import (
	"server/entities"
	"time"
)

type HtmlClassesServices interface {
}

type LeaguesServices interface {
	GetLeaguesName() []string
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
	SearchArticlesTagsAndKeyword(keyword string, formatedTags []string, from int) ([]entities.Article, error)
	GetArticles(keywords []string)
}

type SchedulesServices interface {
	GetAllScheduleLeagueOnDay(date time.Time) (entities.ScheduleOnDay, error)
	GetScheduleLeagueOnDay(date time.Time, league string) (entities.ScheduleOnDay, error)
	GetMatchURLsOnDay() entities.MatchURLsOnDay
	ClearMatchURLsOnDay()
	GetSchedules(date string)
}

type MatchDetailServices interface {
	GetMatchDetailsOnDayFromCrawler(matchURLs entities.MatchURLsOnDay)
	GetMatchDetail(date time.Time, club1Name string, club2Name string) (entities.MatchDetail, error)
}

type AdminServices interface {
	GetAdminUsername(username string) (string, error)
	CheckAdminUsernameToken(username string) error
	ChangePassword(admin *AdminWithConfirmPassword, usernameToken string) error
	Login(admin *Admin) (string, error)
}

type Admin struct {
	Username string `json:"username" validate:"required,min=8,max=16"`
	Password string `json:"password" validate:"required"`
}

type AdminWithConfirmPassword struct {
	Username             string `json:"username" validate:"required,min=8,max=16"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required"`
}
