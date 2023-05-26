package services

import (
	"context"
	"crawler/entities"
	"sync"

	crawlerhelpers "crawler/helper"
	pb "crawler/proto"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"
)

var AMOUNT_PAGE_CRAWL = 3

func CrawlArticleAddedCrawler(configCrawler *pb.ConfigCrawler) ([]entities.Article, error) {
	var articles []entities.Article
	var cacheCrawlerArticle = make(map[string]bool)

	articles, err := crawlWithGoQuery(configCrawler, cacheCrawlerArticle)
	if err != nil {
		if err.Error() == "maybe this is a javascript render button" {
			articlesDP, err := CrawlWithChromedp(configCrawler, cacheCrawlerArticle)
			if err != nil {
				log.Println(err)
				return articles, err
			}
			articles = articlesDP
			return articles, nil
		}
		return articles, err
	}

	if checkIfEmpty(articles) {
		log.Println("not found article try crawl with chromedp")
		articlesDP, err := CrawlWithChromedp(configCrawler, cacheCrawlerArticle)
		articles = articlesDP
		if err != nil {
			log.Println(err)
			return articles, err
		}

	}
	return articles, nil
}

func checkIfEmpty(articles []entities.Article) bool {
	var count int
	for _, article := range articles {
		if article.Title == "" && article.Description == "" && article.Link == "" {
			count++
		}
	}
	return count == len(articles)
}

func CrawlWithChromedp(configCrawler *pb.ConfigCrawler, cacheCrawlerArticle map[string]bool) ([]entities.Article, error) {
	var articles []entities.Article
	var nextPageCount int
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
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)

	ctx, cancelch := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancelch()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()
	task := chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			_, err = page.AddScriptToEvaluateOnNewDocument(script).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
		chromedp.Navigate(configCrawler.Url),
		chromedp.Sleep(6 * time.Second),
	}

	if err := chromedp.Run(ctx, task); err != nil {
		return articles, err
	}

	for nextPageCount < AMOUNT_PAGE_CRAWL {
		articleList, err := getArticleList(ctx, configCrawler, cacheCrawlerArticle)
		if err != nil {
			return articles, err
		}
		articles = append(articles, articleList...)

		var articleNodes []*cdp.Node

		if configCrawler.NetxPageType == "none" {
			return articles, nil
		}

		// nextPageCtx, cancelNextPage:= context.WithTimeout(ctx, 2*time.Second)
		// defer cancelNextPage()

		err = chromedp.Run(ctx, chromedp.Nodes(fmt.Sprintf(`//*[@class='%s']`, configCrawler.NextPage), &articleNodes))
		if err != nil {
			return articles, err
		}
		nodeIDs := []cdp.NodeID{articleNodes[0].NodeID}
		err = chromedp.Run(ctx,
			chromedp.Click(nodeIDs, chromedp.ByNodeID),
			chromedp.Sleep(3*time.Second),
		)
		if err != nil {
			return articles, err
		}
		nextPageCount++
	}

	return articles, nil
}

func getArticleList(ctx context.Context, configCrawler *pb.ConfigCrawler, cacheCrawlerArticle map[string]bool) ([]entities.Article, error) {
	var articleNodes []*cdp.Node
	var articles []entities.Article

	err := chromedp.Run(ctx, chromedp.Nodes(fmt.Sprintf(`//*[@class='%s']`, configCrawler.Div), &articleNodes))

	if err != nil {
		return articles, err
	}

	for _, node := range articleNodes {
		var article entities.Article
		// title
		crawlErr := make(chan error, 3)
		var wg sync.WaitGroup

		wg.Add(1)
		go func(article *entities.Article) {
			crawlErr <- crawlTitle(article, ctx, configCrawler, node)
			wg.Done()
		}(&article)

		// Description

		wg.Add(1)
		go func(article *entities.Article) {
			crawlErr <- crawlDescription(article, ctx, configCrawler, node)
			defer wg.Done()
		}(&article)

		// link
		// there are only one node
		wg.Add(1)
		go func(article *entities.Article) {
			defer wg.Done()
			crawlErr <- crawlLink(article, ctx, configCrawler, node)
		}(&article)

		done := make(chan bool)
		go func(done chan bool) {
			for err := range crawlErr {
				if err != nil {
					log.Printf("error occurs while crawl: %s", <-crawlErr)
				}
			}
			done <- true
		}(done)

		wg.Wait()
		close(crawlErr)
		<-done
		close(done)
		mapKey := fmt.Sprintf("%s-%s", article.Title, article.Link)
		_, ok := cacheCrawlerArticle[mapKey]
		if !ok {
			articles = append(articles, article)
			cacheCrawlerArticle[mapKey] = true
		}

	}

	return articles, nil
}

func crawlTitle(article *entities.Article, ctx context.Context, configCrawler *pb.ConfigCrawler, node *cdp.Node) error {
	titleCtx, cancelTitle := context.WithTimeout(ctx, 1*time.Second)
	defer cancelTitle()
	titleQuery := fmt.Sprintf(`%s//*[@class='%s']`, node.FullXPath(), configCrawler.Title)
	err := chromedp.Run(titleCtx,
		chromedp.Text(titleQuery, &article.Title))
	if err != nil {
		return err
	}
	return nil
}

func crawlDescription(article *entities.Article, ctx context.Context, configCrawler *pb.ConfigCrawler, node *cdp.Node) error {
	descriptionCtx, cancelDescription := context.WithTimeout(ctx, 1*time.Second)
	defer cancelDescription()
	descriptionQuery := fmt.Sprintf(`%s//*[@class='%s']`, node.FullXPath(), configCrawler.Description)
	err := chromedp.Run(descriptionCtx, chromedp.Text(descriptionQuery, &article.Description))
	if err != nil {
		return err
	}
	return nil
}

func crawlLink(article *entities.Article, ctx context.Context, configCrawler *pb.ConfigCrawler, node *cdp.Node) error {
	linkCtx, cancelLink := context.WithTimeout(ctx, 1*time.Second)
	defer cancelLink()
	var linkNodes []*cdp.Node
	linkNodesQuery := fmt.Sprintf(`%s//*[@class='%s']`, node.FullXPath(), configCrawler.Link)
	err := chromedp.Run(linkCtx, chromedp.Nodes(linkNodesQuery, &linkNodes))
	if err != nil {
		return err
	}

	for _, linkNode := range linkNodes {
		link, ok := linkNode.Attribute("href")
		if ok {
			article.Link = formatLink(link, configCrawler.Url)
		} else {
			var linkChilds []*cdp.Node
			err = chromedp.Run(linkCtx, chromedp.Nodes(fmt.Sprintf(`%s//*`, linkNode.FullXPath()), &linkChilds))
			if err != nil {
				continue
			}
			for _, child := range linkChilds {
				linkChild, ok := child.Attribute("href")
				if ok {
					article.Link = formatLink(linkChild, configCrawler.Url)
					break
				}
			}
		}
	}
	if err != nil {
		return err
	}
	return nil
}

// use goquery, only apply with web with next page type is button and
func crawlWithGoQuery(configCrawler *pb.ConfigCrawler, cacheCrawlerArticle map[string]bool) ([]entities.Article, error) {
	var articles []entities.Article

	doc, err := getGoqueryDoc(configCrawler.Url)
	if err != nil {
		return articles, nil
	}

	if configCrawler.NetxPageType == "none" {
		crawledArticles, _ := crawlOnePageWithGoQuery(doc, configCrawler, cacheCrawlerArticle)
		articles = append(articles, crawledArticles...)
		return articles, err
	}

	for i := 0; i < AMOUNT_PAGE_CRAWL; i++ {
		crawledArticles, nextPageLink := crawlOnePageWithGoQuery(doc, configCrawler, cacheCrawlerArticle)
		articles = append(articles, crawledArticles...)

		var nextPageDoc *goquery.Document
		// maybe this is a javascript render button
		if nextPageLink == "" {
			return []entities.Article{}, fmt.Errorf("maybe this is a javascript render button")
		}
		nextPageDoc, err = getGoqueryDoc(nextPageLink)
		if err != nil {
			return articles, err
		}

		doc = nextPageDoc
	}

	return articles, nil
}

func getGoqueryDoc(url string) (*goquery.Document, error) {
	client := http.Client{}
	doc := new(goquery.Document)
	log.Println(" url", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorln("can not create when crawl HTTP:", err)
		return doc, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Errorln("can not do http request:", err)
		return doc, err
	}
	defer resp.Body.Close()

	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("can not create goquery document, err:%s\n", err)
		return doc, err
	}
	return doc, nil
}

func crawlOnePageWithGoQuery(doc *goquery.Document, configCrawler *pb.ConfigCrawler, cacheCrawlerArticle map[string]bool) ([]entities.Article, string) {
	var articles []entities.Article
	doc.Find(crawlerhelpers.FormatClassName(configCrawler.Div)).Each(func(i int, s *goquery.Selection) {
		var article entities.Article
		article.Title = s.Find(crawlerhelpers.FormatClassName(configCrawler.Title)).Text()
		article.Description = s.Find(crawlerhelpers.FormatClassName(configCrawler.Description)).Text()
		link, ok := s.Find(crawlerhelpers.FormatClassName(configCrawler.Link)).Attr("href")
		if ok {
			article.Link = formatLink(link, configCrawler.Url)
		} else {
			s.Find(crawlerhelpers.FormatClassName(configCrawler.Link)).Each(func(i int, s *goquery.Selection) {
				link, ok := s.Find("a").Attr("href")
				if ok {
					article.Link = formatLink(link, configCrawler.Url)
					return
				}
				article.Link = ""
			})
		}

		mapKey := fmt.Sprintf("%s-%s", article.Title, article.Link)
		_, ok = cacheCrawlerArticle[mapKey]
		if !ok {
			articles = append(articles, article)
			cacheCrawlerArticle[mapKey] = true
		}
	})
	if configCrawler.NetxPageType == "none" {
		return articles, ""
	}
	nextPage, ok := doc.Find(crawlerhelpers.FormatClassName(configCrawler.NextPage)).Attr("href")
	if ok {
		return articles, formatLink(nextPage, configCrawler.Url)
	} else {
		doc.Find(crawlerhelpers.FormatClassName(configCrawler.Link)).Each(func(i int, s *goquery.Selection) {
			next, ok := s.Find("a").Attr("href")
			if ok {
				nextPage = next
				return
			}
		})

	}
	return articles, formatLink(nextPage, configCrawler.Url)
}

func formatLink(link string, urlIn string) string {
	// xử lý trường hợp web để link kiểu /hello/halu thay vì example.com/hello/halu
	linkFormated := link
	if link != "" {
		ok := strings.Contains(linkFormated, urlIn)
		if !ok {
			tempUrl, err := url.ParseRequestURI(urlIn)
			if err != nil {
				return fmt.Sprintf("%s%s", urlIn, linkFormated)
			}
			linkFormated = fmt.Sprintf("https://%s%s", tempUrl.Hostname(), linkFormated)
		}
	}

	return linkFormated
}
