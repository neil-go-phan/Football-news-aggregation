package articlerepo

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"server/entities"
	serverhelper "server/helper"
	"server/repository"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	pb "server/proto"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v7"
	jsoniter "github.com/json-iterator/go"
	"github.com/patrickmn/go-cache"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

var PREV_ARTICLES = make(map[string]bool)

var ARTICLES_INDEX_NAME = "articles"
var DEFAULT_TAG = "tin tuc bong da"
var PIT_LIVE = "1m"
var ARTICLE_CACHE_KEY_PREFIX = "article_cache"
var CACHE_EXPIRATION_TIME = 5 * time.Minute
var CACHE_CLEAR_UP_TIME = 5 * time.Minute
var ARTICLE_CACHE = cache.New(CACHE_EXPIRATION_TIME, CACHE_CLEAR_UP_TIME)

var NOTI_COMPLETE_CRAWL_TITLE = "Crawl article success"
var NOTI_COMPLETE_CRAWL_TYPE = "INFO"
var NOTI_COMPLETE_CRAWL_MESSAGE = "Crawler scrape articles from google success"

var NOTI_CRAWLER_SHUTDOWN_TITLE = "Crawler shutdown"
var NOTI_CRAWLER_SHUTDOWN_TYPE = "ERROR"
var NOTI_CRAWLER_SHUTDOWN_MESSAGE = "Crawler is unavaiable while scraping data"

var NOTI_SERVER_CAN_NOT_RECEIVE_ARTICLE_TITLE = "Can not receive articles"
var NOTI_SERVER_CAN_NOT_RECEIVE_ARTICLE_TYPE = "ERROR"
var NOTI_SERVER_CAN_NOT_RECEIVE_ARTICLE_MESSAGE = "Server can not receive articles from crawler. Please checking for connection"

var NOTI_GOOGLE_CAPTCHA_TITLE = "Google captcha"
var NOTI_GOOGLE_CAPTCHA_TYPE = "WARNING"
var NOTI_GOOGLE_CAPTCHA_MESSAGE = "Google detect your crawler. Please shutdown crawler for 12 hour"

type work = []entities.Article
type result = []byte

type articleRepo struct {
	conn            *grpc.ClientConn
	es              *elasticsearch.Client
	htmlClassesRepo repository.HtmlClassesRepository
	leaguesRepo     repository.LeaguesRepository
	tagsRepo        repository.TagRepository
}

func NewArticleRepo(leaguesRepo repository.LeaguesRepository, htmlClassesRepo repository.HtmlClassesRepository, tagsRepo repository.TagRepository, conn *grpc.ClientConn, es *elasticsearch.Client) *articleRepo {
	articleRepo := &articleRepo{
		conn:            conn,
		es:              es,
		leaguesRepo:     leaguesRepo,
		htmlClassesRepo: htmlClassesRepo,
		tagsRepo:        tagsRepo,
	}
	return articleRepo
}

func (repo *articleRepo) SearchArticlesTagsAndKeyword(keyword string, formatedTags []string, from int) ([]entities.Article, float64, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	articles := make([]entities.Article, 0)
	var buffer bytes.Buffer
	var total float64
	query := querySearchArticle(keyword, formatedTags, from)

	err := json.NewEncoder(&buffer).Encode(query)
	if err != nil {
		return articles, total, fmt.Errorf("encode query failed")
	}

	resp, err := repo.es.Search(repo.es.Search.WithIndex(ARTICLES_INDEX_NAME), repo.es.Search.WithBody(&buffer))
	if err != nil {
		return articles, total, fmt.Errorf("request to elastic search fail")
	}

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return articles, total, fmt.Errorf("decode respose from elastic search failed")
	}

	value := result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"]
	total = value.(float64)

	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {

		newArticle := hit.(map[string]interface{})["_source"].(map[string]interface{})
		article := newEntitiesArticleFromMap(newArticle)
		// usecase: when admin add a new tag. then some article will tagged with that new tag.
		// then admin delete that tag. this removeNotExistTag function will detect and filter all deleted tag from return articles
		filterDeletedTag(&article, repo.tagsRepo.ListTags().Tags)

		articles = append(articles, article)

	}

	return articles, total, nil
}

func (repo *articleRepo) GetFirstPageOfLeagueRelatedArticle(leagueName string) ([]entities.Article, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	articles := []entities.Article{}

	key := fmt.Sprintf("%s_%s", ARTICLE_CACHE_KEY_PREFIX, serverhelper.FormatCacheKey(leagueName))
	articleInterface, found := ARTICLE_CACHE.Get(key)
	if found {
		log.Println("FOUND CACHE: ", key)
		articleByte, err := json.Marshal(articleInterface)
		if err != nil {
			log.Printf("error occrus when marshal cache articles: %s", err)
		}

		err = json.Unmarshal(articleByte, &articles)
		if err != nil {
			log.Printf("error occrus when unmarshal cache articles to entity articles: %s", err)
		}
		return articles, nil
	}
	log.Println("MISS CACHE: ", key)
	// Cache miss, fetch article from elastic
	formatedTags := serverhelper.FortmatTagsFromRequest(leagueName)
	articles, _, err := repo.SearchArticlesTagsAndKeyword("", formatedTags, 0)
	if err != nil {
		return nil, err
	}

	repo.RefreshCache()

	return articles, nil
}

func (repo *articleRepo) RefreshCache() {
	log.Printf("refresh cache...")
	var wg sync.WaitGroup
	for _, league := range repo.leaguesRepo.GetLeaguesNameActive() {
		wg.Add(1)
		go func(league string) {
			defer wg.Done()
			tagFromLeague := serverhelper.FormatVietnamese(league)
			firstPageArticleLeague, _, err := repo.SearchArticlesTagsAndKeyword("", []string{tagFromLeague}, 0)
			if err != nil {
				log.Printf("can not request to elastic to reset tag")
			}
			key := fmt.Sprintf("%s_%s", ARTICLE_CACHE_KEY_PREFIX, serverhelper.FormatCacheKey(league))
			ARTICLE_CACHE.Set(key, firstPageArticleLeague, cache.DefaultExpiration)
			log.Printf("refresh cache key %s\n", key)
		}(league)
	}
	wg.Wait()
	log.Printf("refresh cache end.")
}

func (repo *articleRepo) GetArticleCount() (total float64, today float64, err error) {
	today, err = repo.GetCrawledArticleToday()
	if err != nil {
		return total, today, err
	}
	total, err = repo.GetTotalCrawledArticle()
	if err != nil {
		return total, today, err
	}
	return total, today, nil
}

func (repo *articleRepo) GetCrawledArticleToday() (float64, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	var buffer bytes.Buffer

	query := queryGetAmountCrawledArtilceToday()

	err := json.NewEncoder(&buffer).Encode(query)
	if err != nil {
		return 0, err
	}

	resp, err := repo.es.Search(repo.es.Search.WithIndex(ARTICLES_INDEX_NAME), repo.es.Search.WithBody(&buffer))
	if err != nil {
		return 0, err
	}

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return 0, err
	}

	value := result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"]
	today := value.(float64)
	return today, nil
}

func (repo *articleRepo) GetTotalCrawledArticle() (float64, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	var buffer bytes.Buffer

	query := queryGetTotalCrawledArtilce()

	err := json.NewEncoder(&buffer).Encode(query)
	if err != nil {
		return 0, err
	}

	resp, err := repo.es.Search(repo.es.Search.WithIndex(ARTICLES_INDEX_NAME), repo.es.Search.WithBody(&buffer))
	if err != nil {
		return 0, err
	}

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return 0, err
	}

	value := result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"]
	total := value.(float64)
	return total, nil
}

// add tag // implement search_after query, // implement worker pool
func (repo *articleRepo) AddTagForAllArticle(tag string) error {
	tagFormated := serverhelper.FormatVietnamese(tag)
	pitID, err := requestOpenPointInTime(repo.es)
	if err != nil {
		return err
	}
	firstPageArticles, newPitID, searchAfterQuery, err := repo.requestFirstPageOfArticleSearchWithTagAsKeyword(tagFormated, pitID)
	if err != nil {
		return err
	}
	bulkRequestBody, err := createBulkRequestBodyAddTag(firstPageArticles, tagFormated)
	if err != nil {
		log.Printf("error when worker create bulk request body %s\n", err)
	}
	requestAddTagArticle(bulkRequestBody, repo.es)

	if len(searchAfterQuery) == 0 {
		return nil
	}

	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup

	// worker pool
	jobs := make(chan *work, 5)
	results := make(chan *result, 5)
	// 5 worker
	for i := 0; i < 5; i++ {
		wg1.Add(1)
		go workerAddTagBulkRequest(jobs, results, tagFormated, &wg1, i)
	}

	go producerRequestArticle(jobs, repo, tagFormated, newPitID, searchAfterQuery)

	wg2.Add(1)
	go analyzedResult(results, repo.es, &wg2)

	wg1.Wait()

	close(results)

	wg2.Wait()
	log.Printf("add tag success")

	repo.RefreshCache()
	return nil
}

func (repo *articleRepo) requestFirstPageOfArticleSearchWithTagAsKeyword(tag string, pitID string) (articles []entities.Article, newPitID string, searchAfter []interface{}, err error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	var buffer bytes.Buffer

	query := querySearchFirstPageArticlesWithTagAsKeyword(tag, pitID)

	err = json.NewEncoder(&buffer).Encode(query)
	if err != nil {
		return articles, pitID, searchAfter, fmt.Errorf("encode query failed")
	}

	resp, err := repo.es.Search(repo.es.Search.WithBody(&buffer))
	if err != nil {
		return articles, pitID, searchAfter, fmt.Errorf("request to elastic search fail")
	}

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return articles, pitID, searchAfter, fmt.Errorf("decode respose from elastic search failed")
	}

	for index, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		article := hit.(map[string]interface{})["_source"].(map[string]interface{})
		articles = append(articles, newEntitiesArticleFromMap(article))
		// get last result article sort field
		// elastic search will sent 10 hits by default. if we can not find the 10th article, that mean there is no more result
		if index == 9 {
			searchAfter = hit.(map[string]interface{})["sort"].([]interface{})
		}
	}

	newPitID = result["pit_id"].(string)
	return articles, newPitID, searchAfter, nil
}

func (repo *articleRepo) requestNextPageOfArticleWithTagAsKeyword(tag string, pitID string, searchAfterQuery []interface{}) (articles []entities.Article, newPitID string, searchAfter []interface{}, err error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	var buffer bytes.Buffer

	query := queryGetNextPageOfArticleWithTagAsKeyword(tag, pitID, searchAfterQuery)

	err = json.NewEncoder(&buffer).Encode(query)
	if err != nil {
		return articles, pitID, searchAfter, fmt.Errorf("encode query failed")
	}

	resp, err := repo.es.Search(repo.es.Search.WithBody(&buffer))
	if err != nil {
		return articles, pitID, searchAfter, fmt.Errorf("request to elastic search fail")
	}

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return articles, pitID, searchAfter, fmt.Errorf("decode respose from elastic search failed")
	}

	for index, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		article := hit.(map[string]interface{})["_source"].(map[string]interface{})
		articles = append(articles, newEntitiesArticleFromMap(article))
		// get last result article sort field
		// elastic search will sent 10 hits by default. if we can not find the 10th article, that mean there is no more result
		if index == 9 {
			searchAfter = hit.(map[string]interface{})["sort"].([]interface{})
		}
	}

	newPitID = result["pit_id"].(string)
	return articles, newPitID, searchAfter, nil
}

// GetArticles request crawler to scrapt data and sync elastic search
func (repo *articleRepo) GetArticles(keywords []string) {
	client := pb.NewCrawlerServiceClient(repo.conn)

	htmlClass := repo.htmlClassesRepo.GetHtmlClasses()

	in := &pb.AllConfigsArticles{
		HtmlClasses: &pb.HTMLClasses{
			ArticleClass:            htmlClass.ArticleClass,
			ArticleTitleClass:       htmlClass.ArticleTitleClass,
			ArticleDescriptionClass: htmlClass.ArticleDescriptionClass,
			ArticleLinkClass:        htmlClass.ArticleLinkClass,
		},
		Leagues: repo.leaguesRepo.GetLeaguesNameActive(),
	}

	if len(keywords) != 0 {
		in.Leagues = keywords
	}

	// send gRPC request to crawler
	stream, err := client.GetArticles(context.Background(), in)
	if err != nil {
		log.Printf("error occurred while openning stream error %v \n", err)
		return
	}

	done := make(chan bool)
	var mapSearchResult = make(map[string]bool)
	log.Printf("Start get stream of article...\n")
	// recieve stream of article from crawler
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Printf("cannot receive %v\n", err)
				status, _ := status.FromError(err)
				if status.Code().String() == "Unavailable" {
					// repo.notification.Send(NOTI_CRAWLER_SHUTDOWN_TITLE, NOTI_CRAWLER_SHUTDOWN_TYPE, NOTI_CRAWLER_SHUTDOWN_MESSAGE)

					done <- true //means stream is finished
					return
				}
			}

			league := resp.GetLeague()

			respArticles := resp.GetArticles()
			tags := repo.tagsRepo.ListTags().Tags

			checkSimilarArticles(respArticles, repo.es, league, tags)
			
			
			saveToMapSearchResult(respArticles, mapSearchResult)

		}
	}()

	<-done
	PREV_ARTICLES = mapSearchResult
	log.Printf("finished.")
	// repo.notification.Send(NOTI_COMPLETE_CRAWL_TITLE, NOTI_COMPLETE_CRAWL_TYPE, NOTI_COMPLETE_CRAWL_MESSAGE)
	// if len(PREV_ARTICLES) == 0 {
	// 	repo.notification.Send(NOTI_GOOGLE_CAPTCHA_TITLE, NOTI_GOOGLE_CAPTCHA_TYPE, NOTI_GOOGLE_CAPTCHA_MESSAGE)
	// }
}

func (repo *articleRepo) DeleteArticle(title string) error {
	
	req := esapi.DeleteRequest{
		Index:      ARTICLES_INDEX_NAME,
		DocumentID: strings.ToLower(title),
	}

	res, err := req.Do(context.Background(), repo.es)
	if err != nil {
		log.Errorf("Error getting response for delete request: %s", err)
		return fmt.Errorf("error getting response for delete request")
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Errorf("[%s] Error delete article document with ID='%s'", res.Status(), strings.ToLower(title))
		return fmt.Errorf(res.Status())
	} else {
		log.Printf("[%s] Deleted document with id: %s", res.Status(), strings.ToLower(title))
	}
	return nil
}
