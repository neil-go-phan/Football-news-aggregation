package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

type HtmlArticleClass struct {
	ArticleClass     string
	TitleClass       string
	DescriptionClass string
	ThumbnailClass   string
	LinkClass        string
}

var htmlArticleClass = HtmlArticleClass{
	ArticleClass:     `SoaBEf`,
	TitleClass:       `mCBkyc ynAwRc MBeuO nDgy9d`,
	DescriptionClass: `GI74Re nDgy9d`,
	ThumbnailClass:   `uhHOwf BYbUcd`,
	LinkClass:        `WlydOe`,
}

var PAGES int = 10

type Article struct {
	Title       string
	Description string
	Thumbnail   string
	Link        string
}

func main() {
	articles, _ := searchKeyWord("ngoai hang anh")
	for index, a:= range articles {
		fmt.Println("index: ", index, " title: ", a.Title)
	}
	
}

func searchKeyWord(keyword string) ([]Article, error) {

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	ctx2, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	err := chromedp.Run(ctx2, chromedp.Navigate(fmt.Sprintf("https://www.google.com/search?q=%s", formatKeywords(keyword))))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = chromedp.Run(ctx2, chromedp.WaitVisible(`#search`, chromedp.ByID))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var nodes []*cdp.Node
	err = chromedp.Run(ctx2, chromedp.Nodes(`//div[@class="hdtb-mitem"]//a[text()="Tin tức"]`, &nodes))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var newURL string
	for _, node := range nodes {
		newURL, _ = node.Attribute("href")
	}

	newURL = `https://www.google.com` + newURL

	// articles, _ := searchInMultiPages(ctx2, newURL, PAGES)
	articles := multiPage(newURL)

	return articles, nil
}

func multiPage(newURL string) ([]Article) {
	var articles []Article
	var wg sync.WaitGroup
	for i := 0; i < PAGES; i ++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			news, err := fetchHtml(newURL, index)
			if err != nil {
				log.Fatal(err)
			}
			articles = append(articles, news...)
			
		} (i)
		
	}
	wg.Wait()
	return articles
}

func fetchHtml(url string, page int) ([]Article, error) {
	var articles []Article
	//
	req, err := http.NewRequest("GET", fmt.Sprintf(`%s&start=%d0`, url, page), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(formatClassName(htmlArticleClass.ArticleClass)).Each(func(i int, s *goquery.Selection) {
		var article Article
		article.Title = s.Find(formatClassName(htmlArticleClass.TitleClass)).Text()
		article.Description = s.Find(formatClassName(htmlArticleClass.DescriptionClass)).Text()
		article.Thumbnail, _ = s.Find(formatClassName(htmlArticleClass.ThumbnailClass)).Find("img").Attr("src")
		article.Link, _ = s.Find(formatClassName(htmlArticleClass.LinkClass)).Attr("href")
		articles = append(articles, article)
	})
	return articles, nil
}

func formatClassName(class string) string {
	var classes string
	hashParts := strings.Split(class, " ")
	for _,s := range hashParts {
		classes = classes + "." + s
	}
	return classes
}

func formatKeywords(keyword string) string {
	return strings.Replace(keyword, " ", "+", -1)
}

// Use chromedp

// func searchInMultiPages(ctx context.Context, url string, pages int) ([]Article, error) {
// 	var articles []Article

// 	var wg sync.WaitGroup

// 	for i := 0; i < pages; i++ {
// 		wg.Add(1)

// 		go func(ctx context.Context, index int) {
// 			ctx, cancel := chromedp.NewContext(ctx)
// 			defer cancel()
// 			ctx2, cancel := context.WithTimeout(ctx, 3*time.Second)
// 			defer cancel()
// 			defer wg.Done()

// 			err := chromedp.Run(ctx2, chromedp.Navigate(fmt.Sprintf(`%s&start=%d0`, url, index)))
// 			if err != nil {
// 				log.Println("err: ", err)
// 			}
// 			err = chromedp.Run(ctx2, chromedp.WaitVisible(`#search`, chromedp.ByID))
// 			if err != nil {
// 				log.Println(err)
// 			}
// 			articles = append(articles, searchResultInOnePage(ctx)...)
// 		}(ctx, i)

// 	}

// 	wg.Wait()

// 	return articles, nil
// }

// func searchResultInOnePage(ctx context.Context) []Article {
// 	var nodes []*cdp.Node
// 	var articles []Article
// 	err := chromedp.Run(ctx, chromedp.Nodes(fmt.Sprintf(`//div[@class='%s']`, htmlArticleClass.ArticleClass), &nodes))
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	for _, node := range nodes {
// 		var article Article
// 		// title
// 		err = chromedp.Run(ctx, chromedp.Text(fmt.Sprintf(`%s//div[@class='%s']`, node.FullXPath(), htmlArticleClass.TitleClass), &article.Title))
// 		if err != nil {
// 			log.Println(err)
// 		}

// 		// Description
// 		err = chromedp.Run(ctx, chromedp.Text(fmt.Sprintf(`%s//div[@class='%s']`, node.FullXPath(), htmlArticleClass.DescriptionClass), &article.Description))
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		// Link

// 		err = chromedp.Run(ctx, chromedp.AttributeValue(fmt.Sprintf(`%s//a[@class='%s']`, node.FullXPath(), htmlArticleClass.LinkClass), "href", &article.Link, nil))
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		// Thumbnail
// 		err = chromedp.Run(ctx, chromedp.AttributeValue(fmt.Sprintf(`%s//div[@class='%s']//img`, node.FullXPath(), htmlArticleClass.ThumbnailClass), "src", &article.Thumbnail, nil))
// 		if err != nil {
// 			log.Println(err)
// 		}

// 		articles = append(articles, article)
// 	}
// 	return articles
// }