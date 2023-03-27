package services

import (
	"backend/services/server/entities"
	"context"
	"io"
	"log"

	pb "github.com/neil-go-phan/Football-news-aggregation/backend/grpcfile"
	"google.golang.org/grpc"
)

func GetArticles(keyword string, htmlClass entities.HtmlArticleClass, conn *grpc.ClientConn) {
	client := pb.NewArticleServiceClient(conn)
	in := &pb.AllConfigs{
		HtmlClasses: &pb.HTMLClasses{
			ArticleClass: htmlClass.ArticleClass,
			TitleClass: htmlClass.TitleClass,
			DescriptionClass: htmlClass.DescriptionClass,
			ThumbnailClass: htmlClass.ThumbnailClass,
			LinkClass: htmlClass.LinkClass,
		},
		Keywords: keyword,
	}
	stream, err := client.GetArticles(context.Background(), in)
	if err != nil {
		log.Printf("open stream error %v \n", err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Printf("cannot receive %v\n", err)
			}
			log.Printf("Resp received: %s\n", resp.Articles[0])
		}
	}()

	<-done //we will wait until all response is received
	log.Printf("finished")
}