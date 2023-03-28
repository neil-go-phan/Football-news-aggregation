package services

import (


	"github.com/robfig/cron/v3"
)

type HtmlClassesServices interface {

}

type KeywordsServices interface {

}

type ArticleServices interface {
	GetArticlesEveryMinutes(cronjob *cron.Cron)
}