package handlers

import (
	"context"
	"crawler/entities"
	pb "crawler/proto"
	"crawler/services"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func (s *gRPCServer) GetArticlesFromAddedCrawler(ctx context.Context, configCrawler *pb.ConfigCrawler) (*pb.ArticleAddedCrawler, error) {

	log.Println("Start scrapt article")

	schedules, err := crawlArticleAddedCrawlerAndParse(configCrawler)

	if err != nil {
		log.Println(err)
	}

	log.Println("Finish scrapt article from custom crawler")
	return schedules, nil
}

func crawlArticleAddedCrawlerAndParse(configCrawler *pb.ConfigCrawler) (*pb.ArticleAddedCrawler, error) {
	articles := new(pb.ArticleAddedCrawler) 
	articlesCrawl, err := services.CrawlArticleAddedCrawler(configCrawler)
	if err != nil {
		log.Errorf("error occurred during crawl article with custom crawler, err: %v", err)
	}
	articles = crawledArticlesToPbActiclesAddedCrawler(articlesCrawl, configCrawler.Url)
	if err != nil {
		return articles, fmt.Errorf("error occurred while sending response to client: %v", err)
	}

	log.Println("crawl article with custom crawler successfully")
	return articles, nil
}

func crawledArticlesToPbActiclesAddedCrawler(crawlArticles []entities.Article, url string) *pb.ArticleAddedCrawler {
	pbArticles := &pb.ArticleAddedCrawler{}
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