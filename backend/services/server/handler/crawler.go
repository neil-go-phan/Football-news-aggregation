package handler

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"server/handler/presenter"
	"server/services"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type ConfigCrawlerHandler struct {
	handler services.CrawlerServices
}

func NewCrawlerHandler(handler services.CrawlerServices) *ConfigCrawlerHandler {
	configCrawler := &ConfigCrawlerHandler{
		handler: handler,
	}
	return configCrawler
}

func (configCrawlerHandler *ConfigCrawlerHandler) APIGetHtmlPage(c *gin.Context) {
	urlInput := c.Query("url")
	url, err := url.Parse(urlInput)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Bad request"})
		return
	}

	err = configCrawlerHandler.handler.GetHtmlPage(url)
	if err != nil {
		log.Errorln("error occurs when response page content to frontend: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Internal server error"})
		return
	}
	hostname := strings.TrimPrefix(url.Hostname(), "www.")
	filePath := fmt.Sprintf("page%s.html", hostname)

	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()

	c.Header("Content-Type", "text/html")
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	err = os.Remove(filePath)
	if err != nil {
		log.Errorln("error occurs when delete html file: ", err)
	}
	// c.JSON(http.StatusOK, gin.H{"success": true, "message": "Get page success"})
}

func (configCrawlerHandler *ConfigCrawlerHandler) APIUpsertConfigCrawler(c *gin.Context) {
	var input services.Crawler
	err := c.BindJSON(&input)
	if err != nil {
		log.Printf("error occrus: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}
	err = configCrawlerHandler.handler.Upsert(&input)
	if err != nil {
		log.Printf("error occrus: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "upsert successfull"})
}

func (configCrawlerHandler *ConfigCrawlerHandler) APIListConfigCrawler(c *gin.Context) {
	configCrawlers, err := configCrawlerHandler.handler.List()
	if err != nil {
		log.Printf("error occrus: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	

	c.JSON(http.StatusOK, gin.H{"success": true, "config_crawlers": configCrawlers})
}

func (configCrawlerHandler *ConfigCrawlerHandler) APIGetConfigCrawler(c *gin.Context) {
	url := c.Query("url")
	configCrawler, err := configCrawlerHandler.handler.Get(url)
	if err != nil {
		log.Printf("error occrus: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "config_crawler": configCrawler})
}

func (configCrawlerHandler *ConfigCrawlerHandler) APIDeleteConfigCrawler(c *gin.Context) {
	url := c.Query("url")
	err := configCrawlerHandler.handler.Delete(url)
	if err != nil {
		log.Printf("error occrus: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "delete success"})
}

func (configCrawlerHandler *ConfigCrawlerHandler) APITestCrawler(c *gin.Context) {
	var input services.Crawler
	err := c.BindJSON(&input)
	if err != nil {
		log.Printf("error occrus: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "bad request"})
		return
	}
	articles,err := configCrawlerHandler.handler.TestCrawler(&input)
	if err != nil {
		log.Printf("error occrus: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}
		// parse to presenter
		resposeArticles := make([]presenter.Article, 0)

		for _, article := range articles {
			tags := make([]string, 0)
			for _, tag := range article.Tags {
				tags = append(tags, tag.TagName)
			}
			presenterArticle := presenter.Article{
				ID:          article.ID,
				Title:       article.Title,
				Description: article.Description,
				Link:        article.Link,
				Tags:        tags,
			}
			resposeArticles = append(resposeArticles, presenterArticle)
		}
		
	c.JSON(http.StatusOK, gin.H{"success": true, "articles": resposeArticles, "amount": len(resposeArticles)})
}

func (h *ConfigCrawlerHandler) CreateCustomCrawlerCronjob() {
	log.Println("start create custom crawler cronjob")
	err := h.handler.CreateCustomCrawlerCronjob()
	if err != nil {
		log.Error(err)
	}
	log.Println("end create custom crawler cronjob")
}

func (h *ConfigCrawlerHandler) APIChangeCronjobTime(c *gin.Context) {
	var inputCronjob services.CronjobChangeTimeRequestPayload
	err := c.BindJSON(&inputCronjob)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Change cronjob time failed"})
		return
	}
	err = h.handler.ChangeScheduleCronjob(inputCronjob)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Change cronjob time failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Change cronjob time success"})
}