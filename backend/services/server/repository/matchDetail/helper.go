package matchdetailrepo
import (
	"bytes"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"server/entities"
	serverhelper "server/helper"
	"strings"
	"time"

	pb "server/proto"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v7"
	jsoniter "github.com/json-iterator/go"

)
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
		log.Errorf("error occrus when marshal elastic response match detail: %s\n", err)
	}

	err = json.Unmarshal(matchDetailByte, &matchDetail)
	if err != nil {
		log.Errorf("error occrus when unmarshal elastic response to entity match detail: %s\n", err)
	}
	return matchDetail
}

func pbMatchDetailToEntityMatchDetail(pbMatchDetail *pb.MatchDetail) entities.MatchDetail {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	matchDetail := entities.MatchDetail{}
	matchDetailByte, err := json.Marshal(pbMatchDetail)
	if err != nil {
		log.Errorf("error occrus when marshal pb.MatchDetail: %s", err)
	}
	err = json.Unmarshal(matchDetailByte, &matchDetail)
	if err != nil {
		log.Errorf("error occrus when unmarshal pb.MatchDetail to entities.MatchDetail: %s", err)
	}

	return matchDetail
}

func upsertMatchDetailElastic(matchDetail entities.MatchDetail, es *elasticsearch.Client, date time.Time) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	docID := strings.ToLower(fmt.Sprintf("$date=%s$match=%svs%s", date, serverhelper.FormatElasticSearchIndexName(matchDetail.MatchDetailTitle.Club1.Name), serverhelper.FormatElasticSearchIndexName(matchDetail.MatchDetailTitle.Club2.Name)))

	var buffer bytes.Buffer

	query := map[string]interface{}{
		"doc":           matchDetail,
		"doc_as_upsert": true,
	}
	err := json.NewEncoder(&buffer).Encode(query)
	if err != nil {
		log.Errorf("error occrus when encoding query: %s\n", err)
	}

	req := esapi.UpdateRequest{
		Index:      MATCH_DETAIL_INDEX_NAME,
		DocumentID: docID,
		Body:       &buffer,
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Errorf("Error getting response: %s\n", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Errorf("[%s] Error indexing document\n", res.Status())
	} else {
		log.Errorf("[%s] Upsert document with index: %s \n", res.Status(), MATCH_DETAIL_INDEX_NAME)
	}
}
