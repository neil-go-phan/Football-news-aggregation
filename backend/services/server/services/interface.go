package services

import (
	"server/entities"
	"time"
	// "github.com/robfig/cron/v3"
)

type HtmlClassesServices interface {

}

type LeaguesServices interface {
	ListLeagues() (entities.Leagues)
}

type TagsServices interface {
	ListTags() (entities.Tags)
}

type ArticleServices interface {
	APISearchArticlesTagsAndKeyword(keyword string, formatedTags []string) ([]entities.Article ,error)
	APISearchAll(search_type string, scroll string, size string) ([]entities.Article ,error)
	GetArticles()
}

type SchedulesServices interface {
	APIGetScheduleOnDay(date time.Time) (entities.ScheduleOnDay, error)
	GetSchedules(date string)
}