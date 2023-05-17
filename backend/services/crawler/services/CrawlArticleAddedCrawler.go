package services

import (
	"crawler/entities"
	crawlerhelpers "crawler/helper"
	pb "crawler/proto"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
)

func CrawlArticleAddedCrawler(configCrawler *pb.ConfigCrawler) ([]entities.Article, error) {
	var articles []entities.Article
	isScroll := checkTypeNextPageIsScrolll(configCrawler.NetxPageType)
	// handle later
	log.Println(isScroll)

	articles, err := crawlWithGoQuery(configCrawler)
	if err != nil {
		return articles, err
	}

	return articles, nil
}

// check type next page, if it is 'scroll' then we use chromedp to crawl
func checkTypeNextPageIsScrolll(netxPageType string) bool {
	return netxPageType == "scroll"
}

func crawlWithGoQuery(configCrawler *pb.ConfigCrawler) ([]entities.Article, error) {
	var articles []entities.Article
	client := http.Client{}
	req, err := http.NewRequest("GET", configCrawler.Url, nil)
	if err != nil {
		log.Errorln("can not create when crawl HTTP:", err)
		return articles, err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Errorln("can not do http request:", err)
		return articles, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("can not create goquery document, err:%s\n", err)
		return articles, err
	}

	// check is SPA web
	// crawl
	doc.Find(crawlerhelpers.FormatClassName(configCrawler.List)).Each(func(i int, s *goquery.Selection) {
		s.Find(crawlerhelpers.FormatClassName(configCrawler.Div)).Each(func(i int, s *goquery.Selection) {
			var article entities.Article
			article.Title = s.Find(crawlerhelpers.FormatClassName(configCrawler.Title)).Text()
			article.Description = s.Find(crawlerhelpers.FormatClassName(configCrawler.Description)).Text()
			link, ok := s.Find(crawlerhelpers.FormatClassName(configCrawler.Link)).Attr("href")
			if ok {
				article.Link = link
			} else {
				s.Find(crawlerhelpers.FormatClassName(configCrawler.Link)).Each(func(i int, s *goquery.Selection) {
					link, ok := s.Find("a").Attr("href")
					if ok {
						article.Link = link
						return
					}
					article.Link = ""
				})
			}
			articles = append(articles, article)
		})
	})
	return articles, nil
}
