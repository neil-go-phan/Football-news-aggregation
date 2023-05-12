package articlesservices

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"server/entities"
	pb "server/proto"
	mock "server/services/articles/mock"
	"testing"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/stretchr/testify/assert"
)

func TestDeleteArticle(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"took":1,"errors":false,"items":[{"update":{"_index":"articles","_type":"_doc","_id":"1","_version":1,"result":"updated","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":0,"_primary_term":1,"status":200}}]}`)
	}))
	defer server.Close()
	assert := assert.New(t)
	cfg := elasticsearch.Config{
		Addresses: []string{server.URL},
	}
	es, err := elasticsearch.NewClient(cfg)
	assert.Nil(err)

	mockRepo := new(mock.MockArticleRepository)
	mockLeague := new(mock.MockLeaguesServices)
	mockTag := new(mock.MockTagsServices)
	mockCrawler := new(mock.MockCrawlerServiceClient)

	mockArticle :=  NewArticleService(mockLeague, mockTag, mockCrawler, es, mockRepo)

	mockRepo.On("Delete", uint(1)).Return(nil)
	err = mockArticle.DeleteArticle(uint(1))
	assert.Nil(err)
}


func TestGetArticleCount(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"took":1,"errors":false,"items":[{"update":{"_index":"articles","_type":"_doc","_id":"1","_version":1,"result":"updated","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":0,"_primary_term":1,"status":200}}]}`)
	}))
	defer server.Close()
	assert := assert.New(t)
	cfg := elasticsearch.Config{
		Addresses: []string{server.URL},
	}
	es, err := elasticsearch.NewClient(cfg)
	assert.Nil(err)

	mockRepo := new(mock.MockArticleRepository)
	mockLeague := new(mock.MockLeaguesServices)
	mockTag := new(mock.MockTagsServices)
	mockCrawler := new(mock.MockCrawlerServiceClient)

	mockArticle :=  NewArticleService(mockLeague, mockTag, mockCrawler, es, mockRepo)

	mockRepo.On("GetCrawledArticleToday").Return(int64(100),nil)
	mockRepo.On("GetTotalCrawledArticle").Return(int64(100),nil)
	total, today, err := mockArticle.GetArticleCount()
	assert.Nil(err)
	assert.Equal(total, int64(100))
	assert.Equal(today, int64(100))
}

func TestStoreArticle(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"took":1,"errors":false,"items":[{"update":{"_index":"articles","_type":"_doc","_id":"1","_version":1,"result":"updated","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":0,"_primary_term":1,"status":200}}]}`)
	}))
	defer server.Close()
	assert := assert.New(t)
	cfg := elasticsearch.Config{
		Addresses: []string{server.URL},
	}
	es, err := elasticsearch.NewClient(cfg)
	assert.Nil(err)

	mockRepo := new(mock.MockArticleRepository)
	mockLeague := new(mock.MockLeaguesServices)
	mockTag := new(mock.MockTagsServices)
	mockCrawler := new(mock.MockCrawlerServiceClient)

	mockArticle :=  NewArticleService(mockLeague, mockTag, mockCrawler, es, mockRepo)
	respArticles :=  []*pb.Article{
		{Title: "title 1", Description: "description 1", Link: "test.com"},
		{Title: "title 2", Description: "description 2",Link: "test.com"},
		{Title: "title 3", Description: "description 3",Link: "test.com"},
	}
	tagName := []string{"tag1", "tag2"}

	tags := &[]entities.Tag{}

	article1 := &entities.Article{
		Title: "title 1", Description: "description 1", Link: "test.com", Tags: []entities.Tag{},
	}
	article2 := &entities.Article{
		Title: "title 2", Description: "description 2",Link: "test.com",Tags: []entities.Tag{},
	}
	article3 := &entities.Article{
		Title: "title 3", Description: "description 3",Link: "test.com",Tags: []entities.Tag{},
	}

	mockTag.On("ListTagsName").Return(tagName, nil)
	mockTag.On("GetTagsByTagNames", []string{"tin tuc bong da", "league"}).Return(tags, nil)
	mockRepo.On("FirstOrCreate", article1).Return(nil)
	mockRepo.On("FirstOrCreate", article2).Return(nil)
	mockRepo.On("FirstOrCreate", article3).Return(nil)


	mockArticle.storeArticles(respArticles, "league")
}