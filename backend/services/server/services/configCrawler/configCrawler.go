package configcrawler

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	// log "github.com/sirupsen/logrus"
)

type ConfigCrawlerService struct {
}

func NewConfigCrawlerService() *ConfigCrawlerService {
	configCrawlerService := &ConfigCrawlerService{}
	return configCrawlerService
}

func (s *ConfigCrawlerService) GetHtmlPage(url *url.URL) error {

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate(url.String()),
	)
	if err != nil {
		return err
	}
	var htmlContent string

	err = chromedp.Run(ctx,
		chromedp.OuterHTML("html", &htmlContent),
	)
	if err != nil {
		return err
	}

	hostname := strings.TrimPrefix(url.Hostname(), "www.")

	err = os.WriteFile(fmt.Sprintf("page%s.html", hostname), []byte(htmlContent), 0644)
	if err != nil {
		return err
	}

	return nil
}
