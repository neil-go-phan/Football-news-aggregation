package articlesservices

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/url"
	"server/entities"
	serverhelper "server/helper"
	"server/repository"
	"server/services"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	pb "server/proto"

	"github.com/elastic/go-elasticsearch/v7"
	jsoniter "github.com/json-iterator/go"
	"github.com/patrickmn/go-cache"

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

type work = []entities.ElasticArticle
type result struct {
	ids             []uint
	bulkRequestBody []byte
}

type ArticleService struct {
	grpcClient     pb.CrawlerServiceClient
	es             *elasticsearch.Client
	leaguesService services.LeaguesServices
	tagsService    services.TagsServices
	configCrawler  services.ConfigCrawlerServices
	repo           repository.ArticleRepository
}

func NewArticleService(leaguesService services.LeaguesServices, tagsService services.TagsServices, grpcClient pb.CrawlerServiceClient, es *elasticsearch.Client, repo repository.ArticleRepository, configCrawler services.ConfigCrawlerServices) *ArticleService {
	articleService := &ArticleService{
		grpcClient:     grpcClient,
		es:             es,
		leaguesService: leaguesService,
		tagsService:    tagsService,
		configCrawler:  configCrawler,
		repo:           repo,
	}
	return articleService
}

// GetArticles request crawler to scrapt data and sync elastic search
func (s *ArticleService) GetArticles(keywords []string) {

	leagues, err := s.leaguesService.GetLeaguesNameActive()
	if err != nil {
		log.Error(err)
	}
	// TODO: refactor this
	leagues = append(leagues, "Tin tuc bong da")

	in := &pb.KeywordToSearch{
		Keyword: leagues,
	}

	if len(keywords) != 0 {
		in.Keyword = keywords
	}

	// send gRPC request to crawler
	stream, err := s.grpcClient.GetArticles(context.Background(), in)
	if err != nil {
		log.Errorln("error occurred while openning stream error ", err)
		return
	}

	done := make(chan bool)
	var mapSearchResult = make(map[string]bool)
	log.Printf("Start get stream of article...\n")
	// recieve stream of article from crawler
	go func() {
		for {
			resp, err := stream.Recv()
			if err != nil {
				status, _ := status.FromError(err)
				// this error mean crawler is shutdown
				if status.Code().String() == "Unavailable" || err == io.EOF {
					done <- true
					return
				}
			}

			league := resp.GetLeague()

			respArticles := resp.GetArticles()
			s.StoreArticles(respArticles, league)

			saveToMapSearchResult(respArticles, mapSearchResult)
		}
	}()

	<-done
	PREV_ARTICLES = mapSearchResult

	// configCrawler crawl
	
	crawlers, err := s.configCrawler.List()
	if err != nil {
		log.Error(err)
	}

	for _, crawler := range crawlers {
		articles := make([]*pb.Article, 0)
		entitiesArticle, err, _ := s.configCrawler.GetArticles(&crawler)
		if err != nil {
			log.Error(err)
		}
		for _, entity := range entitiesArticle {
			pbArticle := newPbArticle(entity)
			articles = append(articles, pbArticle)
		}
		url, err := url.ParseRequestURI(crawler.Url)
		if err != nil {
			log.Error(err)
		}
		s.StoreArticles(articles, url.Hostname())
	}

		

	log.Printf("finished.")
}

func (s *ArticleService) StoreArticles(respArticles []*pb.Article, league string) {
	tags, err := s.tagsService.ListTagsName()
	if err != nil {
		log.Error(err)
	}

	for _, respArticle := range respArticles {
		ok := validateArticles(respArticle)
		if !ok {
			continue
		}

		// tagged articles
		tagNames := taggedWhenCrawl(respArticle, tags, league)
		articleTag, err := s.tagsService.GetTagsByTagNames(tagNames)
		if err != nil {
			log.Error(err)
		}

		// store article to database
		newArticle := newEntitiesArticle(respArticle, *articleTag)
		err = s.repo.FirstOrCreate(&newArticle)
		if err != nil {
			log.Error(err)
		}
		// store article to elasticsearch
		storeArticleToElasticsearch(newArticle, s.es)
	}
}

func (s *ArticleService) GetArticleCount() (total int64, today int64, err error) {
	today, err = s.repo.GetCrawledArticleToday()
	if err != nil {
		return total, today, err
	}
	total, err = s.repo.GetTotalCrawledArticle()
	if err != nil {
		return total, today, err
	}
	return total, today, nil
}

func (s *ArticleService) SearchArticles(keyword string, formatedTags []string, from int) ([]entities.Article, int64, error) {
	articles := []entities.Article{}
	var total int64
	ids, total, err := s.SearchArticlesOnElasticSearch(keyword, formatedTags, from)
	if err != nil {
		return articles, total, err
	}
	articles, err = s.repo.GetArticlesByIDs(ids)
	if err != nil {
		return articles, total, err
	}
	return articles, total - 1, nil
}

// return IDs of articles that match the condition
func (s *ArticleService) SearchArticlesOnElasticSearch(keyword string, formatedTags []string, from int) ([]uint, int64, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	ids := make([]uint, 0)
	var buffer bytes.Buffer
	query := querySearchArticle(keyword, formatedTags, from)
	var total float64

	err := json.NewEncoder(&buffer).Encode(query)
	if err != nil {
		return ids, int64(total), fmt.Errorf("encode query failed")
	}

	resp, err := s.es.Search(s.es.Search.WithIndex(ARTICLES_INDEX_NAME), s.es.Search.WithBody(&buffer))
	if err != nil {
		return ids, int64(total), fmt.Errorf("request to elastic search fail")
	}

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return ids, int64(total), fmt.Errorf("decode respose from elastic search failed")
	}

	value := result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"]
	total = value.(float64)
	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		newArticle := hit.(map[string]interface{})["_source"].(map[string]interface{})
		article := newElasticArticleFromMap(newArticle)
		ids = append(ids, article.ID)
	}

	return ids, int64(total), nil
}

func (s *ArticleService) GetFirstPageOfLeagueRelatedArticle(leagueName string) ([]services.ArticleCache, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	articles := []services.ArticleCache{}

	key := fmt.Sprintf("%s_%s", ARTICLE_CACHE_KEY_PREFIX, serverhelper.FormatCacheKey(leagueName))
	articleInterface, found := ARTICLE_CACHE.Get(key)
	if found {
		log.Println("FOUND CACHE: ", key)
		articleByte, err := json.Marshal(articleInterface)
		if err != nil {
			log.Errorf("error occrus when marshal cache articles: %s\n", err)
		}

		err = json.Unmarshal(articleByte, &articles)
		if err != nil {
			log.Errorf("error occrus when unmarshal cache articles to entity articles: %s\n", err)
		}

		return articles, nil
	}
	log.Println("MISS CACHE: ", key)
	// Cache miss, fetch article from database
	formatedTags := serverhelper.FortmatTagsFromRequest(leagueName)
	articlesFromDB, _, err := s.SearchArticles("", formatedTags, 0)
	if err != nil {
		return nil, err
	}

	articleCaches := make([]services.ArticleCache, 0)

	for _, article := range articlesFromDB {
		tags := make([]string, 0)
		for _, tag := range article.Tags {
			tags = append(tags, tag.TagName)
		}
		articleCache := services.ArticleCache{
			ID:          article.ID,
			Title:       article.Title,
			Description: article.Description,
			Link:        article.Link,
			Tags:        tags,
		}
		articleCaches = append(articleCaches, articleCache)
	}

	return articleCaches, nil
}

func (s *ArticleService) RefreshCache() {
	log.Printf("refresh cache...")
	var wg sync.WaitGroup
	leagues, err := s.leaguesService.GetLeaguesNameActive()
	if err != nil {
		log.Error(err)
	}
	for _, league := range leagues {
		wg.Add(1)
		go func(league string) {
			defer wg.Done()

			tagFromLeague := serverhelper.FormatVietnamese(league)

			firstPageArticleLeague, _, err := s.SearchArticles("", []string{tagFromLeague}, 0)
			if err != nil {
				log.Errorf("can not reset cache: %s\n", err)
			}

			key := fmt.Sprintf("%s_%s", ARTICLE_CACHE_KEY_PREFIX, serverhelper.FormatCacheKey(league))

			// parse to presenter
			articleCaches := make([]services.ArticleCache, 0)

			for _, article := range firstPageArticleLeague {
				tags := make([]string, 0)
				for _, tag := range article.Tags {
					tags = append(tags, tag.TagName)
				}
				articleCache := services.ArticleCache{
					ID:          article.ID,
					Title:       article.Title,
					Description: article.Description,
					Link:        article.Link,
					Tags:        tags,
				}
				articleCaches = append(articleCaches, articleCache)
			}
			ARTICLE_CACHE.Set(key, articleCaches, cache.DefaultExpiration)
			log.Printf("refresh cache key %s\n", key)
		}(league)
	}
	wg.Wait()
	log.Printf("refresh cache end.")
}

func (s *ArticleService) DeleteArticle(id uint) error {
	err := deleteArticleFromElasticSearch(id, s.es)
	if err != nil {
		return err
	}

	err = s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

// implement search_after query, // implement worker pool
// Add tag for article in elasticsearch and database. find articles potentially tagged from elasticsearch. then check it. Then modify the database and elasticsearch. apply elasticsearch's search_query. Get 10 results each time
func (s *ArticleService) AddTagForAllArticle(tag string) error {
	tagFormated := serverhelper.FormatVietnamese(tag)

	entityTag, err := s.tagsService.Get(tagFormated)
	if err != nil {
		return err
	}

	pitID, err := requestOpenPointInTime(s.es)
	if err != nil {
		return err
	}
	firstPageArticles, newPitID, searchAfterQuery, err := requestFirstPageOfArticleSearchWithTagAsKeyword(tagFormated, pitID, s.es)
	if err != nil {
		return err
	}

	bulkRequestBody, ids, err := createBulkRequestBodyAddTag(firstPageArticles, tagFormated)
	if err != nil {
		log.Errorf("error when worker create bulk request body %s\n", err)
	}
	if len(bulkRequestBody) != 0 {
		requestUpdateTagsArticle(bulkRequestBody, s.es)
		err := s.repo.AddTag(ids, entityTag)
		if err != nil {
			log.Error(err)
		}
	}

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
		// worker: take job from jobs chanel -> Do something -> add result to results chanel
		go workerAddTagBulkRequest(jobs, results, tagFormated, &wg1, i)
	}
	// producer: create job and add to jobs chanel
	go producerRequestArticle(jobs, s.es, tagFormated, newPitID, searchAfterQuery)

	wg2.Add(1)
	// analyzer: summarizing results (log error)
	go analyzedResult(results, s, &wg2, entityTag)

	wg1.Wait()

	close(results)

	wg2.Wait()
	log.Printf("add tag success")

	return nil
}
