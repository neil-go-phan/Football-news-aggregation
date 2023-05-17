package configcrawler

import (
	"fmt"
	"github.com/go-playground/validator"
	"net/url"
	"server/entities"
	pb "server/proto"
)

type ConfigCrawler struct {
	Url                string `json:"url" validate:"required"`
	ArticleList        string `json:"article_list" validate:"required"`
	ArticleDiv         string `json:"article_div" validate:"required"`
	ArticleTitle       string `json:"article_title" validate:"required"`
	ArticleDescription string `json:"article_description" validate:"required"`
	ArticleLink        string `json:"article_link" validate:"required"`
	NextPage           string `json:"next_page"`
	NetxPageType       string `json:"next_page_type" validate:"required"`
}

func validateConfigCrawler(configCrawler *ConfigCrawler) error {
	validate := validator.New()
	err := validate.Struct(configCrawler)
	if err != nil {
		return err
	}
	_, err = url.ParseRequestURI(configCrawler.Url)
	if err != nil {
		return fmt.Errorf("url invalid")
	}
	return nil
}

func newEntityConfigCrawler(configCrawler *ConfigCrawler) *entities.ConfigCrawler {
	return &entities.ConfigCrawler{
		Url:                configCrawler.Url,
		ArticleList:        configCrawler.ArticleList,
		ArticleDiv:         configCrawler.ArticleDiv,
		ArticleTitle:       configCrawler.ArticleTitle,
		ArticleDescription: configCrawler.ArticleDescription,
		ArticleLink:        configCrawler.ArticleLink,
		NextPage:           configCrawler.NextPage,
		NetxPageType:       configCrawler.NetxPageType,
	}
}

func newConfigCrawler(configCrawler *entities.ConfigCrawler) ConfigCrawler {
	return ConfigCrawler{
		Url:                configCrawler.Url,
		ArticleList:        configCrawler.ArticleLink,
		ArticleDiv:         configCrawler.ArticleDiv,
		ArticleTitle:       configCrawler.ArticleTitle,
		ArticleDescription: configCrawler.ArticleDescription,
		ArticleLink:        configCrawler.ArticleLink,
		NextPage:           configCrawler.NextPage,
		NetxPageType:       configCrawler.NetxPageType,
	}
}

func newPbConfigCrawler(configCrawler *ConfigCrawler) *pb.ConfigCrawler {
	return &pb.ConfigCrawler{
		Url:          configCrawler.Url,
		List:         configCrawler.ArticleList,
		Div:          configCrawler.ArticleDiv,
		Title:        configCrawler.ArticleTitle,
		Description:  configCrawler.ArticleDescription,
		Link:         configCrawler.ArticleLink,
		NextPage:     configCrawler.NextPage,
		NetxPageType: configCrawler.NetxPageType,
	}
}

func newEntitiesArticle(respArticle *pb.Article) entities.Article {
	article := entities.Article{
		Title:       respArticle.Title,
		Description: respArticle.Description,
		Link:        respArticle.Link,
	}
	return article
}
