package crawler

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"server/entities"
	pb "server/proto"
	"server/repository"
	"server/services"
	"strings"
	"time"

	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

type CrawlerService struct {
	repo           repository.CrawlerRepository
	cronjobService services.CronjobServices
	grpcClient     pb.CrawlerServiceClient
	cron           *cron.Cron
	jobIDMap       map[string]cron.EntryID
}

func NewCrawlerService(repo repository.CrawlerRepository, cronjobService services.CronjobServices, grpcClient pb.CrawlerServiceClient, cron *cron.Cron, jobIDMap map[string]cron.EntryID) *CrawlerService {
	configCrawlerService := &CrawlerService{
		repo:           repo,
		cronjobService: cronjobService,
		grpcClient:     grpcClient,
		cron:           cron,
		jobIDMap:       jobIDMap,
	}
	return configCrawlerService
}

func (s *CrawlerService) GetHtmlPage(url *url.URL) error {
	const script = `(function(w, n, wn) {
		// Pass the Webdriver Test.
		Object.defineProperty(n, 'webdriver', {
			get: () => false,
		});
	
		// Pass the Plugins Length Test.
		// Overwrite the plugins property to use a custom getter.
		Object.defineProperty(n, 'plugins', {
			// This just needs to have length > 0 for the current test,
			// but we could mock the plugins too if necessary.
			get: () => [1, 2, 3, 4, 5],
		});
	
		// Pass the Languages Test.
		// Overwrite the plugins property to use a custom getter.
		Object.defineProperty(n, 'languages', {
			get: () => ['en-US', 'en'],
		});
	
		// Pass the Chrome Test.
		// We can mock this in as much depth as we need for the test.
		w.chrome = {
			app: {
				isInstalled: false,
			},
			webstore: {
				onInstallStageChanged: {},
				onDownloadProgress: {},
			},
			runtime: {
				PlatformOs: {
					MAC: 'mac',
					WIN: 'win',
					ANDROID: 'android',
					CROS: 'cros',
					LINUX: 'linux',
					OPENBSD: 'openbsd',
				},
				PlatformArch: {
					ARM: 'arm',
					X86_32: 'x86-32',
					X86_64: 'x86-64',
				},
				PlatformNaclArch: {
					ARM: 'arm',
					X86_32: 'x86-32',
					X86_64: 'x86-64',
				},
				RequestUpdateCheckStatus: {
					THROTTLED: 'throttled',
					NO_UPDATE: 'no_update',
					UPDATE_AVAILABLE: 'update_available',
				},
				OnInstalledReason: {
					INSTALL: 'install',
					UPDATE: 'update',
					CHROME_UPDATE: 'chrome_update',
					SHARED_MODULE_UPDATE: 'shared_module_update',
				},
				OnRestartRequiredReason: {
					APP_UPDATE: 'app_update',
					OS_UPDATE: 'os_update',
					PERIODIC: 'periodic',
				},
			},
		};
	
		// Pass the Permissions Test.
		const originalQuery = wn.permissions.query;
		return wn.permissions.query = (parameters) => (
			parameters.name === 'notifications' ?
				Promise.resolve({ state: Notification.permission }) :
				originalQuery(parameters)
		);
	
	})(window, navigator, window.navigator);`
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36"),
	)
	// var scriptID page.ScriptIdentifier
	// create context
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)

	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()
	var htmlContent string
	task := chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			_, err = page.AddScriptToEvaluateOnNewDocument(script).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
		chromedp.Navigate(url.String()),
		chromedp.Sleep(6 * time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			fmt.Print(node.NodeID)
			htmlContent, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)

			return err
		}),
	}

	if err := chromedp.Run(ctx, task); err != nil {
		fmt.Println(err)
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

func (s *CrawlerService) Upsert(configCrawler *services.Crawler) error {
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
	s.cronjobService.CreateCrawlerCronjob(newEntity)
	return nil
}

func (s *CrawlerService) List() ([]services.Crawler, error) {
	entites, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	configCrawlers := []services.Crawler{}
	for _, entity := range *entites {
		configCrawlers = append(configCrawlers, newConfigCrawler(&entity))
	}
	return configCrawlers, nil
}

func (s *CrawlerService) CreateCustomCrawlerCronjob() error {
	crawlers, err := s.repo.List()
	if err != nil {
		return err
	}
	for _, crawler := range *crawlers {
		s.cronjobService.CreateCrawlerCronjob(&crawler)
	}
	return nil
}

func (s *CrawlerService) Get(urlInput string) (*entities.Crawler, error) {
	configCrawler := &entities.Crawler{}
	_, err := url.ParseRequestURI(urlInput)
	if err != nil {
		return configCrawler, fmt.Errorf("url invalid")
	}

	entity, err := s.repo.Get(urlInput)
	if err != nil {
		return configCrawler, err
	}
	// configCrawler = newConfigCrawler(entity)
	return entity, nil
}

func (s *CrawlerService) Delete(urlInput string) error {
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

func (s *CrawlerService) UpdateRunEveryTime(crawler *entities.Crawler) error {
	return s.repo.UpdateRunEveryTime(crawler)
}

func (s *CrawlerService) TestCrawler(configCrawler *services.Crawler) ([]entities.Article, error) {
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

func (s *CrawlerService) GetArticles(configCrawler *services.Crawler) ([]entities.Article, error) {
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

func (s *CrawlerService) ChangeScheduleCronjob(cronjobIn services.CronjobChangeTimeRequestPayload) error {
	// search cronjob in map
	mapKey := newMapKey(cronjobIn.Url, cronjobIn.RunEveryMinOld)
	entryID, found := s.jobIDMap[mapKey]
	if !found {
		return fmt.Errorf("invalid request, not found cronjob")
	}
	// remove cronjob
	log.Println("remove cronjob", cronjobIn.Name)
	s.cron.Remove(entryID)
	delete(s.jobIDMap, mapKey)
	// add new cronjob
	// query db to get crawler
	crawler, err := s.Get(cronjobIn.Url)
	if err != nil {
		return err
	}
	crawler.RunEveryMin = cronjobIn.RunEveryMinNew
	s.cronjobService.CreateCrawlerCronjob(crawler)
	err = s.UpdateRunEveryTime(crawler)
	if err != nil {
		return err
	}
	return nil
}
