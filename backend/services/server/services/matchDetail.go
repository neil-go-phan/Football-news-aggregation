package services

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"server/entities"
	serverhelper "server/helper"
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

func (s *matchDetailService) APIGetMatchDetail(date time.Time, club1Name string, club2Name string) (entities.MatchDetail, error) {
	var matchDetail entities.MatchDetail

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	var buffer bytes.Buffer

	docID := strings.ToLower(fmt.Sprintf("$date=%s$match=%svs%s", date, serverhelper.FormatElasticSearchIndexName(club1Name) , serverhelper.FormatElasticSearchIndexName(club2Name)))
	fmt.Println(docID)

	query := querySearchMatchDetailByID(docID)

	err := json.NewEncoder(&buffer).Encode(query)
	if err != nil {
		return matchDetail, fmt.Errorf("encode query failed")
	}

	resp, err := s.es.Search(s.es.Search.WithIndex(MATCH_DETAIL_INDEX_NAME), s.es.Search.WithBody(&buffer))
	if err != nil {
		return matchDetail, fmt.Errorf("request to elastic search fail")
	}

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return matchDetail, fmt.Errorf("decode respose from elastic search failed")
	}
	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {

		matchDetailElastic := hit.(map[string]interface{})["_source"].(map[string]interface{})
		matchDetail = newEntitiesMatchDetailFromMap(matchDetailElastic)
	}

	return matchDetail, nil
}

func querySearchMatchDetailByID(docID string) map[string]interface{} {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"_id": docID,
			},
		},
	}
	return query
}

func newEntitiesMatchDetailFromMap(respMatchDetail map[string]interface{}) entities.MatchDetail {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	matchDetail := entities.MatchDetail{}

	matchDetailByte, err := json.Marshal(respMatchDetail)
	if err != nil {
		log.Printf("error occrus when marshal elastic response match detail: %s\n", err)
	}
	
	err = json.Unmarshal(matchDetailByte, &matchDetail)
	if err != nil {
		log.Printf("error occrus when unmarshal elastic response to entity match detail: %s\n", err)
	}
	return matchDetail
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

	docID := strings.ToLower(fmt.Sprintf("$date=%s$match=%svs%s", date, serverhelper.FormatElasticSearchIndexName(matchDetail.MatchDetailTitle.Club1.Name) , serverhelper.FormatElasticSearchIndexName(matchDetail.MatchDetailTitle.Club2.Name)))

	var buffer bytes.Buffer

	query := map[string]interface{}{
		"doc":           matchDetail,
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
