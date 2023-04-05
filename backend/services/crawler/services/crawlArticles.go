package services

import (
	"crawler/entities"
	"crawler/helpers"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func CrawlArticles(url string, page int, htmlClasses entities.HtmlArticleClass) ([]entities.Article, error) {
	var articles []entities.Article
	
	req, err := http.NewRequest("GET", fmt.Sprintf(`%s&start=%d0`, url, page), nil)
	if err != nil {
		log.Println("can not request to url: ", fmt.Sprintf(`%s&start=%d0`, url, page), " err: ", err)
		return articles, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("can not set header, err:", err)
		return articles, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println("can not create goquery document, err:", err)
		return articles, err
	}

	doc.Find(helpers.FormatClassName(htmlClasses.ArticleClass)).Each(func(i int, s *goquery.Selection) {
		var article entities.Article
		article.Title = s.Find(helpers.FormatClassName(htmlClasses.ArticleTitleClass)).Text()
		article.Description = s.Find(helpers.FormatClassName(htmlClasses.ArticleDescriptionClass)).Text()
		article.Link, _ = s.Find(helpers.FormatClassName(htmlClasses.ArticleLinkClass)).Attr("href")
		articles = append(articles, article)
	})
	return articles, nil
}
