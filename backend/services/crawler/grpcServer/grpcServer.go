package grpcserver

import (
	"backend/services/crawler/crawl"
	"log"
	"net"
	"sync"

	pb "backend/grpcfile"
	"google.golang.org/grpc"
)

var PAGES = 1

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
    ThumbnailClass   : configs.HtmlClasses.ThumbnailClass,
    LinkClass	: configs.HtmlClasses.LinkClass,
	}
	var wg sync.WaitGroup

	// Search keyword on google and tab "Tin tức" url
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
	return nil
}

func crawlAndStreamResult(stream pb.ArticleService_GetArticlesServer, keyword string,htmlClasses crawl.HtmlArticleClass) error{

	newsUrl, err := crawl.SearchKeyWord(keyword)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	for i := 0; i < PAGES; i++ {
		wg.Add(1)
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
			Thumbnail:   article.Thumbnail,
			Link:        article.Link,
		}
		pbArticles.Articles = append(pbArticles.Articles, pbArticle)
	}

	return pbArticles
}
