package grpcserver

import (
	"backend/services/crawler/crawl"
	"log"
	"net"
	"sync"

	pb "github.com/neil-go-phan/Football-news-aggregation/backend/grpcfile"
	"google.golang.org/grpc"
)

var PAGES = 10

type articlesServer struct {
	pb.UnimplementedArticleServiceServer
}

func GRPCServer() {
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", "localhost:50005")
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
	keyword := configs.GetKeywords()
	// Search keyword on google and tab "Tin tức" url
	newsUrl, err := crawl.SearchKeyWord(keyword.String())
	if err != nil {
		log.Printf("Faile to search for key word: %s, err: %v \n", keyword.String(), err)
		return err
	}

	htmlClasses := crawl.HtmlArticleClass{
		ArticleClass     : configs.HtmlClasses.ArticleClass,
    TitleClass       : configs.HtmlClasses.TitleClass,
    DescriptionClass : configs.HtmlClasses.DescriptionClass,
    ThumbnailClass   : configs.HtmlClasses.ThumbnailClass,
    LinkClass	: configs.HtmlClasses.LinkClass,
	}

	// Crawl each page, send result to client using gRPC stream
  // use wait group to allow process to be concurrent
	var wg sync.WaitGroup
	for i := 0; i < PAGES; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			newses, err := crawl.CrawlPage(newsUrl, index, htmlClasses)
			if err != nil {
				log.Printf("Error when crawl page: %v, err: %v \n", index, err)
			}
			articles := crawlArticlesToPbActicles(newses)
			err = stream.Send(articles)
			if err != nil {
				log.Println("Error when send response to client: ", err)
			}

		}(i)
	}
	
	if err != nil {
		log.Printf("Faile to crawl for key word: %s, err: %v \n", keyword.String(), err)
		return err
	}
	wg.Wait()
	return nil
}

func crawlArticlesToPbActicles(crawlArticles []crawl.Article) *pb.Articles {
	pbArticles := &pb.Articles{}
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
