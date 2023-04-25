package services

import (
	"crawler/helper"
	"crawler/entities"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func CrawlArticles(searchUrl string, page int, htmlClasses entities.HtmlArticleClass) ([]entities.Article, error) {
	var articles []entities.Article
	req, err := http.NewRequest("GET", fmt.Sprintf(`%s&start=%d0`, searchUrl, page), nil)
	if err != nil {
		log.Println("can not request to url: ", fmt.Sprintf(`%s&start=%d0`, searchUrl, page), " err: ", err)
		return articles, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("can not request, err:%s\n", err)
		return articles, err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("can not create goquery document, err:%s\n", err)
		return articles, err
	}

	doc.Find("#recaptcha").Each(func(i int, s *goquery.Selection) {
		log.Printf("Google detect you trying to crawl article. Please shut down crawler services or change proxy")
	})

	doc.Find(crawlerhelpers.FormatClassName(htmlClasses.ArticleClass)).Each(func(i int, s *goquery.Selection) {
		var article entities.Article
		article.Title = s.Find(crawlerhelpers.FormatClassName(htmlClasses.ArticleTitleClass)).Text()
		article.Description = s.Find(crawlerhelpers.FormatClassName(htmlClasses.ArticleDescriptionClass)).Text()
		article.Link, _ = s.Find(crawlerhelpers.FormatClassName(htmlClasses.ArticleLinkClass)).Attr("href")
		articles = append(articles, article)
	})
	return articles, nil
}
