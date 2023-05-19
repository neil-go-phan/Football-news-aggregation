package services

import (
	"context"
	"crawler/entities"

	crawlerhelpers "crawler/helper"
	pb "crawler/proto"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"
)

var AMOUNT_PAGE_CRAWL = 3

func CrawlArticleAddedCrawler(configCrawler *pb.ConfigCrawler) ([]entities.Article, error) {
	var articles []entities.Article
	var cacheCrawlerArticle = make(map[string]bool)

	if configCrawler.NetxPageType == "scroll" {
		articles, err := crawlWithChromedpScroll(configCrawler, cacheCrawlerArticle)
		if err != nil {
			return articles, err
		}
		return articles, nil
	}

	articles, err := crawlWithGoQuery(configCrawler, cacheCrawlerArticle)
	if err != nil {
		if err.Error() == "maybe this is a javascript render button" {
			log.Println("err")
			articlesDP, err := CrawlWithChromedp(configCrawler, cacheCrawlerArticle)
			if err != nil {
				log.Println(err)
				return articles, err
			}
			articles = articlesDP
			return articles, nil
		}
		return articles, err
	}

	if checkIfEmpty(articles) {
		log.Println("not found article try crawl with chromedp")
		articlesDP, err := CrawlWithChromedp(configCrawler, cacheCrawlerArticle)
		articles = articlesDP
		if err != nil {
			log.Println(err)
			return articles, err
		}

	}

	return articles, nil
}

func checkIfEmpty(articles []entities.Article) bool {
	var count int
	for _, article := range articles {
		if article.Title == "" && article.Description == "" && article.Link == "" {
			count++
		}
	}
	return count == len(articles)
}

func CrawlWithChromedp(configCrawler *pb.ConfigCrawler, cacheCrawlerArticle map[string]bool) ([]entities.Article, error) {
	var articles []entities.Article
	var nextPageCount int
	// configCrawler = &pb.ConfigCrawler{
	// 	Url:          "https://www.football365.com/",
	// 	Div:          "flex-1 h-full sm:mt-1 sm:pb-1.5 flex-shrink flex flex-col justify-between overflow-hidden",
	// 	Title:        "my-1 sm:mb-0 text-sm sm:text-[15px] text-title font-semibold leading-snug sm:leading-5 line-clamp-3 sm:line-clamp-2",
	// 	Description:  "hidden sm:line-clamp-2 mt-1 text-[13px] font-light leading-4 sm:!opacity-75 text-primary",
	// 	Link:         "my-1 sm:mb-0 text-sm sm:text-[15px] text-title font-semibold leading-snug sm:leading-5 line-clamp-3 sm:line-clamp-2",
	// 	NetxPageType: "none",
	// }
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate(configCrawler.Url),
		chromedp.Sleep(3*time.Second),
	)
	if err != nil {
		return articles, err
	}

	for nextPageCount < AMOUNT_PAGE_CRAWL {
		articleList, err := getArticleList(ctx, configCrawler, cacheCrawlerArticle)
		if err != nil {
			return articles, err
		}
		log.Println("articles",  len(articleList) ) 
		articles = append(articles, articleList...)

		var articleNodes []*cdp.Node

		if configCrawler.NetxPageType == "none" {
			return articles, nil
		}

		err = chromedp.Run(ctx, chromedp.Nodes(fmt.Sprintf(`//*[@class='%s']`, configCrawler.NextPage), &articleNodes))
		if err != nil {
			return articles, err
		}
		nodeIDs := []cdp.NodeID{articleNodes[0].NodeID}
		err = chromedp.Run(ctx,
			chromedp.Click(nodeIDs, chromedp.ByNodeID),
			chromedp.Sleep(3*time.Second),
		)
		if err != nil {
			return articles, err
		}
		nextPageCount++
	}

	return articles, nil
}

func getArticleList(ctx context.Context, configCrawler *pb.ConfigCrawler, cacheCrawlerArticle map[string]bool) ([]entities.Article, error) {
	var articleNodes []*cdp.Node
	var articles []entities.Article

	// task := chromedp.Tasks{
	// 	chromedp.Navigate(url.String()),
	// 	chromedp.Sleep(6 * time.Second),
	// 	chromedp.ActionFunc(func(ctx context.Context) error {
	// 		node, err := dom.GetDocument().Do(ctx)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		fmt.Print(node.NodeID)
	// 		htmlContent, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)

	// 		return err
	// 	}),
	// }
	err := chromedp.Run(ctx, chromedp.Nodes(fmt.Sprintf(`//*[@class='%s']`, configCrawler.Div), &articleNodes))
	log.Println("abc", err)
	log.Println("nodes",  len(articleNodes) ) 
	if err != nil {
		return articles, err
	}

	for _, node := range articleNodes {
		var article entities.Article
		// title
		err = chromedp.Run(ctx, chromedp.Text(fmt.Sprintf(`%s//*[@class='%s']`, node.FullXPath(), configCrawler.Title), &article.Title))
		if err != nil {
			log.Println(err)
		}
		log.Println("title", article.Title)
		// Description
		err = chromedp.Run(ctx, chromedp.Text(fmt.Sprintf(`%s//*[@class='%s']`, node.FullXPath(), configCrawler.Description), &article.Description))
		if err != nil {
			log.Println(err)
		}
		log.Println("description", article.Description)
		// there are only one node
		var linkNodes []*cdp.Node
		err = chromedp.Run(ctx, chromedp.Nodes(fmt.Sprintf(`%s//*[@class='%s']`, node.FullXPath(), configCrawler.Link), &linkNodes))
		if err != nil {
			log.Println(err)
		}

		for _, linkNode := range linkNodes {
			link, ok := linkNode.Attribute("href")
			if ok {
				article.Link = formatLink(link, configCrawler.Url)
			} else {
				var linkChilds []*cdp.Node
				err = chromedp.Run(ctx, chromedp.Nodes(fmt.Sprintf(`%s//*`, linkNode.FullXPath()), &linkChilds))
				if err != nil {
					continue
				}
				for _, child := range linkChilds {
					linkChild, ok := child.Attribute("href")
					if ok {
						article.Link = formatLink(linkChild, configCrawler.Url)
						break
					}
				}
			}
		}

		log.Println("Link", article.Link)

		mapKey := fmt.Sprintf("%s-%s", article.Title, article.Link)
		_, ok := cacheCrawlerArticle[mapKey]
		if !ok {
			articles = append(articles, article)
			cacheCrawlerArticle[mapKey] = true
		}

	}

	return articles, nil
}

func crawlWithChromedpScroll(configCrawler *pb.ConfigCrawler, cacheCrawlerArticle map[string]bool) ([]entities.Article, error) {
	var articles []entities.Article
	var nextPageCount int

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate(configCrawler.Url),
		chromedp.Sleep(3*time.Second),
	)
	if err != nil {
		return articles, err
	}

	for nextPageCount < AMOUNT_PAGE_CRAWL {
		articleList, err := getArticleList(ctx, configCrawler, cacheCrawlerArticle)
		if err != nil {
			return articles, err
		}

		articles = append(articles, articleList...)
		log.Println("len", len(articles))
		var articleNodes []*cdp.Node
		err = chromedp.Run(ctx, chromedp.Nodes(fmt.Sprintf(`//*[@class='%s']`, configCrawler.NextPage), &articleNodes))
		if err != nil {
			return articles, err
		}
		nodeIDs := []cdp.NodeID{articleNodes[0].NodeID}
		err = chromedp.Run(ctx,
			chromedp.ScrollIntoView(nodeIDs, chromedp.ByNodeID),
			chromedp.Sleep(3*time.Second),
		)
		if err != nil {
			return articles, err
		}
		nextPageCount++
	}

	return articles, nil
}

// use goquery, only apply with web with next page type is button and
func crawlWithGoQuery(configCrawler *pb.ConfigCrawler, cacheCrawlerArticle map[string]bool) ([]entities.Article, error) {
	var articles []entities.Article

	doc, err := getGoqueryDoc(configCrawler.Url)
	if err != nil {
		return articles, nil
	}

	if configCrawler.NetxPageType == "none" {
		crawledArticles, _ := crawlOnePageWithGoQuery(doc, configCrawler, cacheCrawlerArticle)
		articles = append(articles, crawledArticles...)
		return articles, err
	}

	for i := 0; i < AMOUNT_PAGE_CRAWL; i++ {
		crawledArticles, nextPageLink := crawlOnePageWithGoQuery(doc, configCrawler, cacheCrawlerArticle)
		articles = append(articles, crawledArticles...)

		var nextPageDoc *goquery.Document
		// maybe this is a javascript render button
		if nextPageLink == "" {
			return []entities.Article{}, fmt.Errorf("maybe this is a javascript render button")
		}
		nextPageDoc, err = getGoqueryDoc(nextPageLink)
		if err != nil {
			return articles, err
		}

		doc = nextPageDoc
	}

	return articles, nil
}

func getGoqueryDoc(url string) (*goquery.Document, error) {
	client := http.Client{}
	doc := new(goquery.Document)
	log.Println(" url", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorln("can not create when crawl HTTP:", err)
		return doc, err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Errorln("can not do http request:", err)
		return doc, err
	}
	defer resp.Body.Close()

	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("can not create goquery document, err:%s\n", err)
		return doc, err
	}
	return doc, nil
}

func crawlOnePageWithGoQuery(doc *goquery.Document, configCrawler *pb.ConfigCrawler, cacheCrawlerArticle map[string]bool) ([]entities.Article, string) {
	var articles []entities.Article
	doc.Find(crawlerhelpers.FormatClassName(configCrawler.Div)).Each(func(i int, s *goquery.Selection) {
		var article entities.Article
		article.Title = s.Find(crawlerhelpers.FormatClassName(configCrawler.Title)).Text()
		article.Description = s.Find(crawlerhelpers.FormatClassName(configCrawler.Description)).Text()
		link, ok := s.Find(crawlerhelpers.FormatClassName(configCrawler.Link)).Attr("href")
		if ok {
			article.Link = formatLink(link, configCrawler.Url)
		} else {
			s.Find(crawlerhelpers.FormatClassName(configCrawler.Link)).Each(func(i int, s *goquery.Selection) {
				link, ok := s.Find("a").Attr("href")
				if ok {
					article.Link = formatLink(link, configCrawler.Url)
					return
				}
				article.Link = ""
			})
		}

		mapKey := fmt.Sprintf("%s-%s", article.Title, article.Link)
		_, ok = cacheCrawlerArticle[mapKey]
		if !ok {
			articles = append(articles, article)
			cacheCrawlerArticle[mapKey] = true
		}
	})
	if configCrawler.NetxPageType == "none" {
		return articles, ""
	}
	nextPage, ok := doc.Find(crawlerhelpers.FormatClassName(configCrawler.NextPage)).Attr("href")
	if ok {
		return articles, formatLink(nextPage, configCrawler.Url)
	} else {
		doc.Find(crawlerhelpers.FormatClassName(configCrawler.Link)).Each(func(i int, s *goquery.Selection) {
			next, ok := s.Find("a").Attr("href")
			if ok {
				nextPage = next
				return
			}
		})
	}
	return articles, formatLink(nextPage, configCrawler.Url)
}

func formatLink(link string, urlIn string) string {
	// xử lý trường hợp web để link kiểu /hello/halu thay vì example.com/hello/halu
	linkFormated := link
	if link != "" {
		ok := strings.Contains(linkFormated, urlIn)
		if !ok {
			tempUrl, err := url.ParseRequestURI(urlIn)
			if err != nil {
				return fmt.Sprintf("%s%s", urlIn, linkFormated)
			}
			linkFormated = fmt.Sprintf("https://%s%s", tempUrl.Hostname(), linkFormated)
		}
	}

	return linkFormated
}
