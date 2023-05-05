package services

import (
	"crawler/entities"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrawlArticles_Success(t *testing.T) {
	htmlDoc := `
		<html>
			<body>
				<div class="article">
					<h2 class="article-title">Article 1</h2>
					<p class="article-description">Description 1</p>
					<a class="article-link" href="https://example.com/article1">Read more</a>
				</div>
				<div class="article">
					<h2 class="article-title">Article 2</h2>
					<p class="article-description">Description 2</p>
					<a class="article-link" href="https://example.com/article2">Read more</a>
				</div>
			</body>
		</html>`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(htmlDoc))
	}))
	defer server.Close()

	assert := assert.New(t)

	searchUrl := fmt.Sprintf("%s/", server.URL)
	page := 0
	htmlClasses := entities.HtmlArticleClass{
		ArticleClass:            "article",
		ArticleTitleClass:       "article-title",
		ArticleDescriptionClass: "article-description",
		ArticleLinkClass:        "article-link",
	}

	articles, err := CrawlArticles(searchUrl, page, htmlClasses)

	assert.Nil(err)

	assert.Len(articles, 2, "Expected 2 articles, but got %d", len(articles))

	expectedTitles := []string{"Article 1", "Article 2"}
	for i, article := range articles {
		assert.Equal(article.Title, expectedTitles[i], "Expected title '%s', but got '%s'", expectedTitles[i], article.Title)
	}
}

func TestCrawlArticles_ReCaptcha(t *testing.T) {
	htmlDoc := `
			<html>
					<body>
							<div id="recaptcha">
							</div>
					</body>
			</html>`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(htmlDoc))
	}))
	defer server.Close()

	assert := assert.New(t)

	searchUrl := fmt.Sprintf("%s/", server.URL)
	page := 0
	htmlClasses := entities.HtmlArticleClass{
		ArticleClass:            "article",
		ArticleTitleClass:       "article-title",
		ArticleDescriptionClass: "article-description",
		ArticleLinkClass:        "article-link",
	}

	articles, err := CrawlArticles(searchUrl, page, htmlClasses)

	assert.Nil(err)

	assert.Len(articles, 0, "Expected 0 articles, but got %d", len(articles))
}
