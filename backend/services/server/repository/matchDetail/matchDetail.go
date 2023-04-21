package matchdetailrepo

import (
	"bytes"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"server/entities"
	serverhelper "server/helper"
	"strings"
	"time"

	pb "server/proto"

	"github.com/elastic/go-elasticsearch/v7"
	jsoniter "github.com/json-iterator/go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// var PREV_ARTICLES = make(map[string]bool)
var MATCH_DETAIL_INDEX_NAME = "match_detail"

type matchDetailRepo struct {
	conn *grpc.ClientConn
	es   *elasticsearch.Client
}

func NewMatchDetailRepo(conn *grpc.ClientConn, es *elasticsearch.Client) *matchDetailRepo {
	matchDetailRepo := &matchDetailRepo{
		conn: conn,
		es:   es,
	}
	return matchDetailRepo
}

func (repo *matchDetailRepo) GetMatchDetailsOnDayFromCrawler(matchURLs entities.MatchURLsOnDay) {
	client := pb.NewCrawlerServiceClient(repo.conn)

	in := &pb.MatchURLs{
		Url: matchURLs.Urls,
	}
	// send gRPC request to crawler
	stream, err := client.GetMatchDetail(context.Background(), in)
	if err != nil {
		log.Errorf("error occurred while openning stream error %v \n", err)
		return
	}

	done := make(chan bool)
	log.Printf("Start get stream of match detail...\n")

	go func(date time.Time) {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Errorf("cannot receive %v\n", err)
				status, _ := status.FromError(err)
				if status.Code().String() == "Unavailable" {
					log.Errorf("gRPC server is down ! %v\n", err)
					done <- true //means stream is finished
					return
				}
			}

			matchDetail := pbMatchDetailToEntityMatchDetail(resp)
			upsertMatchDetailElastic(matchDetail, repo.es, date)

		}
	}(matchURLs.Date)

	<-done

	log.Printf("finished crawl match detail")
}

func (repo *matchDetailRepo) GetMatchDetail(date time.Time, club1Name string, club2Name string) (entities.MatchDetail, error) {
	var matchDetail entities.MatchDetail

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	var buffer bytes.Buffer

	docID := strings.ToLower(fmt.Sprintf("$date=%s$match=%svs%s", date, serverhelper.FormatElasticSearchIndexName(club1Name), serverhelper.FormatElasticSearchIndexName(club2Name)))
	fmt.Println(docID)

	query := querySearchMatchDetailByID(docID)

	err := json.NewEncoder(&buffer).Encode(query)
	if err != nil {
		return matchDetail, fmt.Errorf("encode query failed")
	}

	resp, err := repo.es.Search(repo.es.Search.WithIndex(MATCH_DETAIL_INDEX_NAME), repo.es.Search.WithBody(&buffer))
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
