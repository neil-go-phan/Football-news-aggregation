package grpcserver

import (
	"backend/services/crawler/crawl"
	"fmt"
	"log"
	"net"
	"sync"

	pb "backend/grpcfile"

	"google.golang.org/grpc"
)

var PAGES = 5

type articlesServer struct {
	pb.UnimplementedArticleServiceServer
}

func GRPCServer() {
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	pb.RegisterArticleServiceServer(s, &articlesServer{})
	log.Println("start server")
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func (s *articlesServer) GetArticles(configs *pb.AllConfigs, stream pb.ArticleService_GetArticlesServer) error {
	keywords := configs.GetKeywords()

	htmlClasses := crawl.HtmlArticleClass{
		ArticleClass     : configs.HtmlClasses.ArticleClass,
    TitleClass       : configs.HtmlClasses.TitleClass,
    DescriptionClass : configs.HtmlClasses.DescriptionClass,
    LinkClass	: configs.HtmlClasses.LinkClass,
	}
	var wg sync.WaitGroup
	log.Println("Start scrapt")
	for _, keyword := range keywords{
		wg.Add(1)
		go func(keyword string) {
			err := crawlAndStreamResult(stream, keyword, htmlClasses)
			if err != nil {
				log.Printf("error occurred while searching for key word: %s, err: %v \n", keyword, err)
			}
			wg.Done()
		}(keyword)
	}
	wg.Wait()
	log.Println("Finish scrapt")
	return nil
}

func crawlAndStreamResult(stream pb.ArticleService_GetArticlesServer, keyword string,htmlClasses crawl.HtmlArticleClass) error{

	newsUrl := fmt.Sprintf("https://www.google.com/search?tbm=nws&q=%s", crawl.FormatKeywords(keyword))
	
	log.Println(newsUrl)

	var wg sync.WaitGroup

	wg.Add(PAGES)

	for i := 0; i < PAGES; i++ {
		go func(index int) {
			defer wg.Done()
			newses, err := crawl.CrawlPage(newsUrl, index, htmlClasses)
			if err != nil {
				log.Printf("error occurred during crawl page process: %v, err: %v \n", index, err)
			}
			articles := crawlArticlesToPbActicles(newses, keyword)
			err = stream.Send(articles)
			if err != nil {
				log.Println("error occurred while sending response to client: ", err)
			}

		}(i)
	}
	log.Println(keyword, ": crawl successfully")
	wg.Wait()
	return nil
}

func crawlArticlesToPbActicles(crawlArticles []crawl.Article, keyword string) *pb.ArticlesReponse {
	pbArticles := &pb.ArticlesReponse{Keyword: keyword}
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
