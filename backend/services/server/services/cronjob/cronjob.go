package cronjob

import (
	"context"
	"fmt"
	"server/entities"
	pb "server/proto"
	"server/repository"
	"server/services"
	"time"

	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

// each crawler will have only one cronjob in runtime
// jobIDMap: map[crawler_url]cron.EntryID use to track cronjob with crawler url
type CronjobService struct {
	repo           repository.CronjobRepository
	articleService services.ArticleServices

	cron       *cron.Cron
	grpcClient pb.CrawlerServiceClient
	jobIDMap   map[string]cron.EntryID
}

func NewCronjobService(repo repository.CronjobRepository, articleService services.ArticleServices, cronjob *cron.Cron, grpcClient pb.CrawlerServiceClient, jobIDMap map[string]cron.EntryID) *CronjobService {
	CronjobService := &CronjobService{
		repo:           repo,
		articleService: articleService,
		cron:           cronjob,
		grpcClient:     grpcClient,
		jobIDMap:       jobIDMap,
	}
	return CronjobService
}

func (s *CronjobService) CreateCrawlerCronjob(crawler *entities.Crawler) {
	cronjobSchedule := readRunEveryMin(crawler.RunEveryMin)
	cronjobName := fmt.Sprintf("Crawl url :%s, every %v min", crawler.Url, crawler.RunEveryMin)

	entryID, err := s.cron.AddFunc(cronjobSchedule, func() { s.cronjobCrawlerFunction(crawler, cronjobName) })
	if err != nil {
		log.Error("error occurred while seting up cronjob: ", err)
	}

	log.Printf("create cronjob crawl url %s every %v minute success", crawler.Url, crawler.RunEveryMin)
	mapKey := newMapKey(crawler.Url, crawler.RunEveryMin)
	s.jobIDMap[mapKey] = entryID
}

func (s *CronjobService) cronjobCrawlerFunction(crawler *entities.Crawler, cronjobName string) {
	log.Println("create a new record on cronjob table")
	cronjobDB, err := s.createCronjobCrawlerDB(crawler, cronjobName)
	if err != nil {
		log.Error(err)
	}
	log.Println("start crawl articles")
	pbArticles, err := s.getArticles(crawler)
	if err != nil {
		log.Error(err)
	}
	log.Println("start store articles")
	s.articleService.StoreArticles(pbArticles, crawler.Url)
	log.Println("update end_at collumn in cronjob table")
	err = s.updateEndAt(cronjobDB)
	if err != nil {
		log.Error(err)
	}
}

func (s *CronjobService) getArticles(crawler *entities.Crawler) ([]*pb.Article, error) {
	pbArticles := []*pb.Article{}
	in := newPbConfigCrawler(crawler)
	pbAllarticles, err := s.grpcClient.GetArticlesFromAddedCrawler(context.Background(), in)
	if err != nil {
		log.Errorf("error occurred while get schedule on day from crawler error %v \n", err)
		return pbArticles, err
	}
	pbArticles = pbAllarticles.Articles

	return pbArticles, nil
}

func (s *CronjobService) createCronjobCrawlerDB(crawler *entities.Crawler, cronjobName string) (*entities.Cronjob, error) {
	cronjobDB := &entities.Cronjob{
		StartAt:     time.Now(),
		CrawlerID:   crawler.ID,
		Crawler:     *crawler,
		RunEveryMin: crawler.RunEveryMin,
		Name:        cronjobName,
	}
	cronjobDB, err := s.repo.Create(cronjobDB)
	if err != nil {
		return nil, err
	}
	return cronjobDB, nil
}

func (s *CronjobService) updateEndAt(cronjob *entities.Cronjob) error {
	cronjob.EndAt = time.Now()
	err := s.repo.UpdateEndAt(cronjob)
	if err != nil {
		return err
	}
	return nil
}

func (s *CronjobService) CreateCronjobGetArticleFromGoogle() {
	_, err := s.cron.AddFunc("@every 0h10m", func() { s.articleService.GetArticles(make([]string, 0)) })
	if err != nil {
		log.Println("error occurred while seting up cronjob Get article from google: ", err)
	}
	// s.jobIDMap["Get article from google"] = entryID
	log.Println("create cronjob 'Get article from google' every 10 minute success")
}

func (s *CronjobService) CreateCronjobRefreshCache() {
	_, err := s.cron.AddFunc("@every 0h5m", func() { s.articleService.RefreshCache() })
	if err != nil {
		log.Println("error occurred while seting up RefreshCacheAfter5Min cronjob: ", err)
	}
	// s.jobIDMap["Refresh cache"] = entryID
	log.Println("create cronjob 'Refresh cache' every 5 minute success")
}

func (s *CronjobService) GetCronjobRuntime() []services.CronjobResponse {
	cronjobResponses := make([]services.CronjobResponse, 0)
	for key := range s.jobIDMap {
		log.Println(key)
		cronjobResponses = append(cronjobResponses, newCronjobResponse(key))
	}
	return cronjobResponses
}

func (s *CronjobService) CronjobOnHour(timeString string) (*[60]services.ChartHour, error) {
	charts := [60]services.ChartHour{}
	for index := range charts {
		charts[index].Minute = index
		charts[index].AmountOfJob = 0
	}

	hour, err := time.Parse("2006-01-02 15", timeString)
	if err != nil {
		return &charts, err
	}
	
	endOfHour := hour.Add(time.Duration(59) * time.Minute)

	entities, err := s.repo.Get(hour, endOfHour)
	if err != nil {
		return &charts, err
	}
	minuteNow := time.Now().Minute()
	for _, entityChart := range *entities {
		minuteStart := entityChart.StartAt.Minute()
		if entityChart.EndAt.Year() != 1 {
			minuteEnd := entityChart.EndAt.Minute()
			for i := minuteStart; i <= minuteEnd; i++ {
				charts[i].AmountOfJob += 1
				charts[i].Cronjobs = append(charts[i].Cronjobs, newCronjobInChart(entityChart))
			}
		} else {
			minuteEnd := minuteNow
			for i := minuteStart; i <= minuteEnd; i++ {
				charts[i].AmountOfJob += 1
				charts[i].Cronjobs = append(charts[i].Cronjobs, newCronjobInChart(entityChart))
			}
		}

	}

	// if now.Hour() == hour.Hour() {
	// 	runningCronjob, err := s.repo.GetRuning()
	// 	if err != nil {
	// 		return &charts, err
	// 	}

	// }

	return &charts, nil
}

func (s *CronjobService) CronjobOnDay(timeString string) (*[24]services.ChartDay, error) {
	charts := [24]services.ChartDay{}
	for index := range charts {
		charts[index].Hour = index
		charts[index].AmountOfJob = 0
	}
	day, err := time.Parse("2006-01-02", timeString)
	if err != nil {
		return &charts, err
	}

	endOfDay := day.Add(time.Duration(23)*time.Hour + time.Duration(59)*time.Minute)
	entities, err := s.repo.Get(day, endOfDay)
	if err != nil {
		return &charts, err
	}
	for _, entityChart := range *entities {
		hour := entityChart.StartAt.Hour()
		charts[hour].AmountOfJob += 1
		if charts[hour].Cronjobs == nil {
			charts[hour].Cronjobs = map[string]int{}
		}
		charts[hour].Cronjobs[entityChart.Name] += 1
	}
	return &charts, nil
}
