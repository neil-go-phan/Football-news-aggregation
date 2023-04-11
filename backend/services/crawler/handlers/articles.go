package handlers

import (
	"crawler/services"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"crawler/entities"
	"crawler/helper"
	pb "crawler/proto"

)

var PAGES = 5

func (s *gRPCServer) GetArticles(configs *pb.AllConfigsArticles, stream pb.CrawlerService_GetArticlesServer) error {
	leagues := configs.GetLeagues()

	htmlClasses := entities.HtmlArticleClass{
		ArticleClass:            configs.HtmlClasses.ArticleClass,
		ArticleTitleClass:       configs.HtmlClasses.ArticleTitleClass,
		ArticleDescriptionClass: configs.HtmlClasses.ArticleDescriptionClass,
		ArticleLinkClass:        configs.HtmlClasses.ArticleLinkClass,
	}
	var wg sync.WaitGroup
	log.Println("Start scrapt article")

	proxyList, err := crawlerhelpers.RequestProxyList()
	if err != nil {
		log.Printf("error occurred while get proxy: %s\n", err)
	}

	for _, league := range leagues {
		wg.Add(1)
		time.Sleep(3 *time.Second)
		go func(league string, proxyList []string) {
			err := crawlArticlesAndStreamResult(stream, league, htmlClasses, proxyList)
			if err != nil {
				log.Printf("error occurred while searching for key word: %s, err: %v \n", league, err)
			}
			wg.Done()
		}(league, proxyList)
	}
	wg.Wait()
	log.Println("Finish scrapt article")
	return nil
}

func crawlArticlesAndStreamResult(stream pb.CrawlerService_GetArticlesServer, league string, htmlClasses entities.HtmlArticleClass, proxyList []string) error {

	newsUrl := fmt.Sprintf("https://www.google.com/search?tbm=nws&q=%s", crawlerhelpers.FormatToSearch(league))
	log.Println("Search URL: ", newsUrl)

	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	for index := 0; index < PAGES; index++ {
		randomProxyIndex := random.Intn(len(proxyList) -1)
		newses, err := services.CrawlArticles(newsUrl, index, htmlClasses, proxyList[randomProxyIndex])
		if err != nil {
			log.Printf("error occurred during crawl page process: %v, err: %v \n", index, err)
		}

		articles := crawledArticlesToPbActicles(newses, league)

		err = stream.Send(articles)
		if err != nil {
			log.Println("error occurred while sending response to client: ", err)
		}
	}
	log.Println(league, ": crawl ended")
	return nil
}

func crawledArticlesToPbActicles(crawlArticles []entities.Article, league string) *pb.ArticlesReponse {
	pbArticles := &pb.ArticlesReponse{League: league}
	for _, article := range crawlArticles {
		pbArticle := &pb.Article{
			Title:       article.Title,
			Description: article.Description,
			Link:        article.Link,
		}
		pbArticles.Articles = append(pbArticles.Articles, pbArticle)
	}
	return pbArticles
}
