package handlers

import (
	"context"
	"crawler/entities"
	pb "crawler/proto"
	"crawler/services"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

func (s *gRPCServer) GetArticlesFromAddedCrawler(ctx context.Context, configCrawler *pb.ConfigCrawler) (*pb.ArticleAddedCrawler, error) {

	log.Println("Start scrapt article")

	schedules, err := crawlArticleAddedCrawlerAndParse(configCrawler)

	if err != nil {
		log.Println(err)
	}

	log.Println("Finish scrapt schedule")
	return schedules, nil
}

func crawlArticleAddedCrawlerAndParse(configCrawler *pb.ConfigCrawler) (*pb.ArticleAddedCrawler, error) {

	articlesCrawl, err := services.CrawlArticleAddedCrawler(configCrawler)
	if err != nil {
		return nil, fmt.Errorf("error occurred during crawl article with custom crawler, err: %v", err)
	}

	articles := crawledArticlesToPbActiclesAddedCrawler(articlesCrawl, configCrawler.Url)
	if err != nil {
		return nil, fmt.Errorf("error occurred while sending response to client: %v", err)
	}

	log.Println("crawl article with custom crawler successfully")
	return articles, nil
}

func crawledArticlesToPbActiclesAddedCrawler(crawlArticles []entities.Article, url string) *pb.ArticleAddedCrawler {
	pbArticles := &pb.ArticleAddedCrawler{}
	for _, article := range crawlArticles {
		// xử lý trường hợp web để link kiểu /hello/halu thay vì example.com/hello/halu
		link := article.Link
		if article.Link != "" {
			ok := strings.Contains(link, url)
			if!ok {
				url = strings.TrimRight(url, "/")
				if !strings.HasPrefix(article.Link, "/") {
					article.Link = "/" + strings.TrimLeft(article.Link, "/")
				}
				link = fmt.Sprintf("%s%s",url, article.Link) 
			}
		}
		pbArticle := &pb.Article{
			Title:       article.Title,
			Description: article.Description,
			Link:        link,
		}
		pbArticles.Articles = append(pbArticles.Articles, pbArticle)
	}
	return pbArticles
}