package services

import (
	"backend/services/server/entities"

	// "github.com/robfig/cron/v3"
)

type HtmlClassesServices interface {

}

type KeywordsServices interface {

}

type TagsServices interface {
	ListTags() (entities.Tags)
}

type ArticleServices interface {
	// GetArticlesEveryMinutes(cronjob *cron.Cron)
	FrontendSearchArticlesTagsAndKeyword(keyword string, formatedTags []string) ([]entities.Article ,error)
	GetArticles()
}