package connects

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	serverhelper "server/helper"
	articlesservices "server/services/articles"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v7"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
)

func ConnectToElasticsearch(env serverhelper.EnvConfig) (*elasticsearch.Client, error) {
	var esConfig = elasticsearch.Config{
		Addresses: []string{
			env.ElasticsearchAddress,
		},
	}
	
	es, err := elasticsearch.NewClient(esConfig)
	if err != nil {
		return nil, err
	}
	return es, nil
}

func CreateElaticsearchIndex(es *elasticsearch.Client) {
	// check if index exist. if not exist, create new one, if exist, skip it
	indexName := articlesservices.ARTICLES_INDEX_NAME
	exists, err := checkIndexExists(es, indexName)
	if err != nil {
		log.Printf("Error checking if the index %s exists: %s", indexName, err)
	}

	if !exists {
		log.Printf("Index: %s is not exist, create a new one...", indexName)
		err = createArticleIndex(es, indexName)
		for err != nil {
			log.Printf("Error createing index: %s, try again in 10 seconds", err)
			time.Sleep(10 * time.Second)
			err = createArticleIndex(es, indexName)
		}
	}

	log.Printf("Index: %s is already exist, skip it...", indexName)
}

func checkIndexExists(es *elasticsearch.Client, indexName string) (bool, error) {
	res, err := es.Indices.Exists([]string{indexName})
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, nil
}

func queryCreateArticleIndex() map[string]interface{} {
	query := map[string]interface{}{
		"settings": map[string]interface{}{
			"analysis": map[string]interface{}{
				"analyzer": map[string]interface{}{
					"no_accent": map[string]interface{}{
						"tokenizer": "standard",
						"filter": []string{
							"lowercase",
							"asciifolding",
						},
					},
				},
			},
		},
		"mappings": map[string]interface{}{
			"properties": map[string]interface{}{
				"title": map[string]interface{}{
					"type":     "text",
					"analyzer": "no_accent",
				},
				"description": map[string]interface{}{
					"type":     "text",
					"analyzer": "no_accent",
				},
			},
		},
	}
	return query
}

func createArticleIndex(es *elasticsearch.Client, indexName string) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	requestBody := queryCreateArticleIndex()
	body, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("error encoding article: %s", err)
	}
	req := esapi.IndicesCreateRequest{
		Index: indexName,
		Body:  strings.NewReader(string(body)),
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		return fmt.Errorf("error creating the index: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error creating index %s: %s", indexName, res.Status())
	}

	return nil
}
