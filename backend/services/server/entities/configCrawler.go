package entities

import "gorm.io/gorm"

type ConfigCrawler struct {
	gorm.Model
	Url string
	ArticleDiv string
	ArticleTitle string
	ArticleDescription string
	ArticleLink string
	NextPage string
  NetxPageType string
	
}
