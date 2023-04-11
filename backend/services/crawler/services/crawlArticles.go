package services

import (
	"crawler/helper"
	"crawler/entities"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func CrawlArticles(searchUrl string, page int, htmlClasses entities.HtmlArticleClass, proxy string) ([]entities.Article, error) {
	var articles []entities.Article
	req, err := http.NewRequest("GET", fmt.Sprintf(`%s&start=%d0`, searchUrl, page), nil)
	if err != nil {
		log.Println("can not request to url: ", fmt.Sprintf(`%s&start=%d0`, searchUrl, page), " err: ", err)
		return articles, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36")

	// fmt.Println("proxy ", proxy)
	// create a transport
	// proxyUrl, err := url.Parse(fmt.Sprintf("http://%s", proxy))
	// if err != nil {
	// 	log.Println("error when parse proxy to url:", err)
	// }

	// create a client
	// client := &http.Client{
	// 	Timeout: time.Second * 10,
	// 	Transport: &http.Transport{
	// 		Proxy: http.ProxyURL(proxyUrl),
	// 	},
	// }
	// fmt.Println(client)
	// var resp *http.Response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("can not request, err:%s\n", err)
		return articles, err
	}
	// 	resp, err := client.Do(req)
	// if err != nil {
	// 	log.Printf("can not request, err:%s\n", err)
	// 	return articles, err
	// }
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("can not create goquery document, err:%s\n", err)
		return articles, err
	}

	doc.Find(crawlerhelpers.FormatClassName(htmlClasses.ArticleClass)).Each(func(i int, s *goquery.Selection) {
		var article entities.Article
		article.Title = s.Find(crawlerhelpers.FormatClassName(htmlClasses.ArticleTitleClass)).Text()
		article.Description = s.Find(crawlerhelpers.FormatClassName(htmlClasses.ArticleDescriptionClass)).Text()
		article.Link, _ = s.Find(crawlerhelpers.FormatClassName(htmlClasses.ArticleLinkClass)).Attr("href")
		articles = append(articles, article)
	})
	return articles, nil
}
