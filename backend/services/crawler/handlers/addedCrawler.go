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

	articles, err := crawlArticleAddedCrawlerAndParse(configCrawler)

	if err != nil {
		log.Println("Finish scrapt article from custom crawler")
		return articles, err
	}

	log.Println("Finish scrapt article from custom crawler")
	return articles, nil
}

func crawlArticleAddedCrawlerAndParse(configCrawler *pb.ConfigCrawler) (*pb.ArticleAddedCrawler, error) {
	var articles *pb.ArticleAddedCrawler
	isNextButtonWork := true
	articlesCrawl, err := services.CrawlArticleAddedCrawler(configCrawler)
	if err != nil {
		if err.Error() == "next page button not work" {
			isNextButtonWork = false
			err = nil
		} else {
			return articles, err
		}
	}

	articles = crawledArticlesToPbActiclesAddedCrawler(articlesCrawl, configCrawler.Url, isNextButtonWork)
	if err != nil {
		return articles, fmt.Errorf("error occurred while parse response to client: %v", err)
	}

	log.Println("crawl article with custom crawler successfully")
	return articles, nil
}

func crawledArticlesToPbActiclesAddedCrawler(crawlArticles []entities.Article, url string, isNextButtonWork bool) *pb.ArticleAddedCrawler {
	pbArticles := &pb.ArticleAddedCrawler{
		IsNextButtonWork: isNextButtonWork,
	}
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