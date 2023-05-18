package configcrawler

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"server/entities"
	pb "server/proto"
	"server/repository"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

type ConfigCrawlerService struct {
	repo       repository.ConfigCrawlerRepository
	grpcClient pb.CrawlerServiceClient
}

func NewConfigCrawlerService(repo repository.ConfigCrawlerRepository, grpcClient pb.CrawlerServiceClient) *ConfigCrawlerService {
	configCrawlerService := &ConfigCrawlerService{
		repo:       repo,
		grpcClient: grpcClient,
	}
	return configCrawlerService
}

func (s *ConfigCrawlerService) GetHtmlPage(url *url.URL) error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 300*time.Second)
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate(url.String()),
		chromedp.Sleep(3 * time.Second),
	)
	if err != nil {
		return err
	}
	var htmlContent string

	err = chromedp.Run(ctx,
		chromedp.OuterHTML("html", &htmlContent),
		chromedp.Sleep(3 * time.Second),
	)
	if err != nil {
		return err
	}

	hostname := strings.TrimPrefix(url.Hostname(), "www.")
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
	}

	removeScriptTags(doc)

	htmlWithoutScript := renderNode(doc)
	err = os.WriteFile(fmt.Sprintf("page%s.html", hostname), []byte(htmlWithoutScript), 0644)
	if err != nil {
		return err
	}

	return nil
}

func (s *ConfigCrawlerService) Upsert(configCrawler *ConfigCrawler) error {
	configCrawler = trimConfigCrawler(configCrawler)
	err := validateConfigCrawler(configCrawler)
	if err != nil {
		return err
	}
	newEntity := newEntityConfigCrawler(configCrawler)
	err = s.repo.Upsert(newEntity)
	if err != nil {
		return err
	}
	return nil
}

func (s *ConfigCrawlerService) List() ([]ConfigCrawler, error) {
	entites, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	configCrawlers := []ConfigCrawler{}
	for _, entity := range *entites {
		configCrawlers = append(configCrawlers, newConfigCrawler(&entity))
	}
	return configCrawlers, nil
}

func (s *ConfigCrawlerService) Get(urlInput string) (ConfigCrawler, error) {
	configCrawler := ConfigCrawler{}
	_, err := url.ParseRequestURI(urlInput)
	if err != nil {
		return configCrawler, fmt.Errorf("url invalid")
	}
	
	entity, err := s.repo.Get(urlInput)
	if err != nil {
		return configCrawler, err
	}
	configCrawler = newConfigCrawler(entity)
	return configCrawler, nil
}

func (s *ConfigCrawlerService) Delete(urlInput string) error {
	_, err := url.ParseRequestURI(urlInput)
	if err != nil {
		return fmt.Errorf("url invalid")
	}
	err = s.repo.Delete(urlInput)
	if err != nil {
		return err
	}
	return nil
}

func (s *ConfigCrawlerService) TestCrawler(configCrawler *ConfigCrawler) ([]entities.Article, error) {
	articles := []entities.Article{}
	configCrawler = trimConfigCrawler(configCrawler)
	err := validateConfigCrawler(configCrawler)
	if err != nil {
		return articles, err
	}
	articles, err = s.GetArticles(configCrawler)
	if err != nil {
		return articles, err
	}
	return articles, nil
}

func (s *ConfigCrawlerService) GetArticles(configCrawler *ConfigCrawler) ([]entities.Article, error) {
	articles := []entities.Article{}
	in := newPbConfigCrawler(configCrawler)
	pbAllarticles, err := s.grpcClient.GetArticlesFromAddedCrawler(context.Background(), in)
	if err != nil {
		log.Errorf("error occurred while get schedule on day from crawler error %v \n", err)
		return articles, err
	}
	pbArticles := pbAllarticles.Articles
	for _, pbArticle := range pbArticles {
		article := newEntitiesArticle(pbArticle)
		articles = append(articles, article)
	}
	return articles, nil
}