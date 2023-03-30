package crawl

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type HtmlArticleClass struct {
	ArticleClass     string
	TitleClass       string
	DescriptionClass string
	LinkClass        string
}

type Article struct {
	Title       string
	Description string
	Link        string
}

func CrawlPage(url string, page int, htmlClasses HtmlArticleClass) ([]Article, error) {
	var articles []Article
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

	doc.Find(formatClassName(htmlClasses.ArticleClass)).Each(func(i int, s *goquery.Selection) {
		var article Article
		article.Title = s.Find(formatClassName(htmlClasses.TitleClass)).Text()
		article.Description = s.Find(formatClassName(htmlClasses.DescriptionClass)).Text()
		article.Link, _ = s.Find(formatClassName(htmlClasses.LinkClass)).Attr("href")
		articles = append(articles, article)
	})
	return articles, nil
}


func formatClassName(class string) string {
	var classes string
	hashParts := strings.Split(class, " ")
	for _, s := range hashParts {
		classes = classes + "." + s
	}
	return classes
}

func FormatKeywords(keyword string) string {
	var Regexp_A = `à|á|ạ|ã|ả|ă|ắ|ằ|ẳ|ẵ|ặ|â|ấ|ầ|ẩ|ẫ|ậ`
	var Regexp_E = `è|ẻ|ẽ|é|ẹ|ê|ề|ể|ễ|ế|ệ`
	var Regexp_I = `ì|ỉ|ĩ|í|ị`
	var Regexp_U = `ù|ủ|ũ|ú|ụ|ư|ừ|ử|ữ|ứ|ự`
	var Regexp_Y = `ỳ|ỷ|ỹ|ý|ỵ`
	var Regexp_O = `ò|ỏ|õ|ó|ọ|ô|ồ|ổ|ỗ|ố|ộ|ơ|ờ|ở|ỡ|ớ|ợ`
	var Regexp_D = `Đ|đ`
	reg_a := regexp.MustCompile(Regexp_A)
	reg_e := regexp.MustCompile(Regexp_E)
	reg_i := regexp.MustCompile(Regexp_I)
	reg_o := regexp.MustCompile(Regexp_O)
	reg_u := regexp.MustCompile(Regexp_U)
	reg_y := regexp.MustCompile(Regexp_Y)
	reg_d := regexp.MustCompile(Regexp_D)
	keyword = reg_a.ReplaceAllLiteralString(keyword, "a")
	keyword = reg_e.ReplaceAllLiteralString(keyword, "e")
	keyword = reg_i.ReplaceAllLiteralString(keyword, "i")
	keyword = reg_o.ReplaceAllLiteralString(keyword, "o")
	keyword = reg_u.ReplaceAllLiteralString(keyword, "u")
	keyword = reg_y.ReplaceAllLiteralString(keyword, "y")
	keyword = reg_d.ReplaceAllLiteralString(keyword, "d")

	// regexp remove charaters in ()
	var RegexpPara = `\(.*\)`
	reg_para := regexp.MustCompile(RegexpPara)
	keyword = reg_para.ReplaceAllLiteralString(keyword, "")

	keyword = strings.ToLower(keyword)
	return strings.Replace(keyword, " ", "+", -1)
}