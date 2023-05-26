package crawler

import (
	"fmt"
	"net/url"
	"server/entities"
	pb "server/proto"
	"server/services"
	"strings"

	"github.com/go-playground/validator"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

func validateConfigCrawler(configCrawler *services.Crawler) error {
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

func trimConfigCrawler(configCrawler *services.Crawler) *services.Crawler {
	configCrawler.ArticleDescription = strings.TrimSpace(configCrawler.ArticleDescription)
	configCrawler.ArticleDiv = strings.TrimSpace(configCrawler.ArticleDiv)
	configCrawler.ArticleLink = strings.TrimSpace(configCrawler.ArticleLink)
	configCrawler.ArticleTitle = strings.TrimSpace(configCrawler.ArticleTitle)
	configCrawler.NextPage = strings.TrimSpace(configCrawler.NextPage)
	configCrawler.Url = strings.TrimSpace(configCrawler.Url)

	return configCrawler
}

func newEntityConfigCrawler(configCrawler *services.Crawler) *entities.Crawler {
	return &entities.Crawler{
		Url:                configCrawler.Url,
		ArticleDiv:         configCrawler.ArticleDiv,
		ArticleTitle:       configCrawler.ArticleTitle,
		ArticleDescription: configCrawler.ArticleDescription,
		ArticleLink:        configCrawler.ArticleLink,
		NextPage:           configCrawler.NextPage,
		NetxPageType:       configCrawler.NetxPageType,
	}
}

func newConfigCrawler(crawler *entities.Crawler) services.Crawler {
	return services.Crawler{
		Url:                crawler.Url,
		ArticleDiv:         crawler.ArticleDiv,
		ArticleTitle:       crawler.ArticleTitle,
		ArticleDescription: crawler.ArticleDescription,
		ArticleLink:        crawler.ArticleLink,
		NextPage:           crawler.NextPage,
		NetxPageType:       crawler.NetxPageType,
	}
}

func newPbConfigCrawler(configCrawler *services.Crawler) *pb.ConfigCrawler {
	return &pb.ConfigCrawler{
		Url:          configCrawler.Url,
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

func removeScriptTags(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "script" {
		removeNode(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		removeScriptTags(c)
	}
}

func removeNode(n *html.Node) {
	if n.PrevSibling != nil {
		n.PrevSibling.NextSibling = n.NextSibling
	}
	if n.NextSibling != nil {
		n.NextSibling.PrevSibling = n.PrevSibling
	}
	if n.Parent != nil {
		if n.Parent.FirstChild == n {
			n.Parent.FirstChild = n.NextSibling
		}
		if n.Parent.LastChild == n {
			n.Parent.LastChild = n.PrevSibling
		}
	}
}

func renderNode(n *html.Node) string {
	var sb strings.Builder
	err := html.Render(&sb, n)
	if err != nil {
		log.Error(err)
	}
	return sb.String()
}

func newMapKey(url string, runEveryMin int) string {
	cronjobName := fmt.Sprintf("Crawl url :%s, every %v min", url, runEveryMin)
	return fmt.Sprintf("%s$%s$%v",cronjobName, url, runEveryMin)
}