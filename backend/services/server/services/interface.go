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

}

type ArticleServices interface {
	// GetArticlesEveryMinutes(cronjob *cron.Cron)
	FrontendSearchWithIndex(keyword string, indexName string) ([]entities.Article ,error)
	GetArticles()
}