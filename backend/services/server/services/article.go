package services

import (
	"backend/services/server/entities"
	"context"
	"io"
	"log"

	pb "backend/grpcfile"
	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
)

type articleService struct {
	conn *grpc.ClientConn
	htmlClassesService *htmlClassesService
	keywordsService *keywordsService
}

func NewArticleService(keywords *keywordsService, htmlClass *htmlClassesService, conn *grpc.ClientConn) *articleService {
	articleService := &articleService{
		conn: conn,
		htmlClassesService: htmlClass,
		keywordsService: keywords,
	}
	return articleService
}

func (s *articleService)GetArticlesEveryMinutes(cronjob *cron.Cron) {
	_, err := cronjob.AddFunc("@every 0h01m", func() { getArticles(s.keywordsService.Keywords, s.htmlClassesService.HtmlClasses, s.conn) })
	if err != nil {
		log.Println("error occurred while seting up getArticle cronjob: ", err)
	}
}

func getArticles(keywords entities.Keywords, htmlClass entities.HtmlClasses, conn *grpc.ClientConn) {
	client := pb.NewArticleServiceClient(conn)

	in := &pb.AllConfigs{
		HtmlClasses: &pb.HTMLClasses{
			ArticleClass:     htmlClass.ArticleClass,
			TitleClass:       htmlClass.TitleClass,
			DescriptionClass: htmlClass.DescriptionClass,
			ThumbnailClass:   htmlClass.ThumbnailClass,
			LinkClass:        htmlClass.LinkClass,
		},
		Keywords: keywords.Keywords,
	}

	stream, err := client.GetArticles(context.Background(), in)
	if err != nil {
		log.Printf("error occurred while openning stream error %v \n", err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			for index := range resp.GetArticles() {
				log.Printf("Keyword:%s  received: %v\n", resp.Keyword, index)
			}
			if err == io.EOF {
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Printf("cannot receive %v\n", err)
			}

		}
	}()

	<-done //we will wait until all response is received
	log.Printf("finished")
}

// First: check with redis, then check 
// func checkSimilarArticle()