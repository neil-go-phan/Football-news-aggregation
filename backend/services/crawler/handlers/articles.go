package handlers

import (
	"crawler/services"
	"fmt"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"

	"crawler/entities"
	"crawler/helper"
	pb "crawler/proto"
)

var PAGES = 2

func (s *gRPCServer) GetArticles(configs *pb.KeywordToSearch, stream pb.CrawlerService_GetArticlesServer) error {
	keywords := configs.GetKeyword()

	htmlClasses, err := crawlerhelpers.ReadHtmlArticlesClassJSON()
	if err != nil {
		log.Println("can not read file htmlArticleClass.json, err: ", err)
	}

	var wg sync.WaitGroup
	log.Println("Start scrapt article")

	for _, keyword := range keywords {
		wg.Add(1)
		time.Sleep(3 * time.Second)

		go func(keyword string) {
			err := crawlArticlesAndStreamResult(stream, keyword, htmlClasses)
			if err != nil {
				log.Printf("error occurred while searching for key word: %s, err: %v \n", keyword, err)
			}
			wg.Done()
		}(keyword)
	}
	wg.Wait()
	log.Println("Finish scrapt article")
	return nil
}

func crawlArticlesAndStreamResult(stream pb.CrawlerService_GetArticlesServer, keyword string, htmlClasses entities.HtmlArticleClass) error {

	newsUrl := fmt.Sprintf("https://www.google.com/search?tbm=nws&q=%s", crawlerhelpers.FormatToSearch(keyword))
	log.Println("Search URL: ", newsUrl)


	for index := 0; index < PAGES; index++ {
		newses, err := services.CrawlArticles(newsUrl, index, htmlClasses)
		if err != nil {
			log.Printf("error occurred during crawl page process: %v, err: %v \n", index, err)
		}

		articles := crawledArticlesToPbActicles(newses, keyword)

		err = stream.Send(articles)
		if err != nil {
			log.Println("error occurred while sending response to client: ", err)
		}
	}
	log.Println(keyword, ": crawl ended")
	return nil
}

func crawledArticlesToPbActicles(crawlArticles []entities.Article, keyword string) *pb.ArticlesReponse {
	pbArticles := &pb.ArticlesReponse{League: keyword}
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
