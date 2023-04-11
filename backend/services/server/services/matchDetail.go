package services

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"server/entities"
	"strings"
	"time"

	pb "server/proto"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v7"
	jsoniter "github.com/json-iterator/go"

	"google.golang.org/grpc"
)

// var PREV_ARTICLES = make(map[string]bool)
var MATCH_DETAIL_INDEX_NAME = "match_detail"

type matchDetailService struct {
	conn *grpc.ClientConn
	es   *elasticsearch.Client
}

func NewMatchDetailervice(conn *grpc.ClientConn, es *elasticsearch.Client) *matchDetailService {
	matchDetailService := &matchDetailService{
		conn: conn,
		es:   es,
	}
	return matchDetailService
}

func (s *matchDetailService) GetMatchDetailFromCrawler(matchURLs entities.MatchURLsOnDay) {
	client := pb.NewCrawlerServiceClient(s.conn)

	in := &pb.MatchURLs{
		Url: matchURLs.Urls,
	}
	// send gRPC request to crawler
	stream, err := client.GetMatchDetail(context.Background(), in)
	if err != nil {
		log.Printf("error occurred while openning stream error %v \n", err)
		return
	}

	done := make(chan bool)
	log.Printf("Start get stream of match detail...\n")
	// recieve stream of article from crawler
	go func(date time.Time) {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Printf("cannot receive %v\n", err)
			}
			
			matchDetail := pbMatchDetailToEntityMatchDetail(resp)
			upsertMatchDetailElastic(matchDetail, s.es, date)
			// fmt.Printf("entity: %v", a.MatchOverview)

		}
	}(matchURLs.Date)

	<-done
	log.Printf("finished.")
}

func (s *matchDetailService)APIGetMatchDetail(date time.Time, club1Name string, club2Name string) (entities.MatchDetail, error) {
	matchDetail := entities.MatchDetail{}
	return matchDetail, nil
}

func pbMatchDetailToEntityMatchDetail(pbMatchDetail *pb.MatchDetail) entities.MatchDetail {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	matchDetail := entities.MatchDetail{}
	matchDetailByte, err := json.Marshal(pbMatchDetail)
	if err != nil {
		log.Printf("error occrus when marshal pb.MatchDetail: %s", err)
	}
	err = json.Unmarshal(matchDetailByte, &matchDetail)
	if err != nil {
		log.Printf("error occrus when unmarshal pb.MatchDetail to entities.MatchDetail: %s", err)
	}

	return matchDetail
}

func upsertMatchDetailElastic(matchDetail entities.MatchDetail, es *elasticsearch.Client, date time.Time) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	docID := strings.ToLower(fmt.Sprintf("$date=%s$match=%s vs %s", date, matchDetail.MatchDetailTitle.Club1.Name, matchDetail.MatchDetailTitle.Club2.Name))

	var buffer bytes.Buffer

	query := map[string]interface{}{
		"doc": matchDetail,
		"doc_as_upsert": true,
	}
	json.NewEncoder(&buffer).Encode(query)

	req := esapi.UpdateRequest{
		Index:      MATCH_DETAIL_INDEX_NAME,
		DocumentID: docID,
		Body:       &buffer,
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Printf("Error getting response: %s\n", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error indexing document\n", res.Status())
	} else {
		log.Printf("[%s] Upsert document with index: %s \n", res.Status(), MATCH_DETAIL_INDEX_NAME)
	}
}
