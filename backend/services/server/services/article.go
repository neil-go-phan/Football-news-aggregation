package services

import (
	"backend/services/server/entities"
	"context"
	"io"
	"log"
	"sync"

	pb "backend/grpcfile"

	"google.golang.org/grpc"
)

func GetArticlesWithAllKeyWords(keywords entities.Keywords, htmlClass entities.HtmlArticleClass, conn *grpc.ClientConn) {
	var wg sync.WaitGroup
	for index, keyword := range keywords.Keywords {
		wg.Add(1)
		log.Println("Get article with keyword: ", keyword)
		go func (index int, keyword string)  {
			GetArticles(keyword, htmlClass, conn)
			wg.Done()
		}(index, keyword)
		
	}
	wg.Wait()
}

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
		Keyword: keyword,
	}

	stream, err := client.GetArticles(context.Background(), in)
	if err != nil {
		log.Printf("open stream error %v \n", err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			for index := range resp.GetArticles() {
				log.Printf("Keyword:%s  received: %v\n",keyword, index)
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

