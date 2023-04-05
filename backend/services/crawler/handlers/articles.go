package handlers

import (
	"crawler/services"
	"fmt"
	"log"
	"sync"

	"crawler/entities"
	"crawler/helpers"
	pb "crawler/proto"
)

var PAGES = 10

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

	for _, league := range leagues {
		wg.Add(1)
		go func(league string) {
			err := crawlArticlesAndStreamResult(stream, league, htmlClasses)
			if err != nil {
				log.Printf("error occurred while searching for key word: %s, err: %v \n", league, err)
			}
			wg.Done()
		}(league)
	}
	wg.Wait()
	log.Println("Finish scrapt article")
	return nil
}

func crawlArticlesAndStreamResult(stream pb.CrawlerService_GetArticlesServer, league string, htmlClasses entities.HtmlArticleClass) error {

	newsUrl := fmt.Sprintf("https://www.google.com/search?tbm=nws&q=%s", helpers.FormatToSearch(league))

	log.Println("Search URL: ", newsUrl)

	var wg sync.WaitGroup

	for i := 0; i < PAGES; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			newses, err := services.CrawlArticles(newsUrl, index, htmlClasses)
			if err != nil {
				log.Printf("error occurred during crawl page process: %v, err: %v \n", index, err)
			}
			articles := crawledArticlesToPbActicles(newses, league)
			err = stream.Send(articles)
			if err != nil {
				log.Println("error occurred while sending response to client: ", err)
			}

		}(i)
	}
	log.Println(league, ": crawl successfully")
	wg.Wait()
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
