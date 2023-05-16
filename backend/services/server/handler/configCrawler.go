package handler

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"server/services"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type ConfigCrawlerHandler struct {
	handler services.ConfigCrawlerServices
}

func NewConfigCrawlerHandler(handler services.ConfigCrawlerServices) *ConfigCrawlerHandler {
	configCrawler := &ConfigCrawlerHandler{
		handler: handler,
	}
	return configCrawler
}

func (configCrawler *ConfigCrawlerHandler) APIGetHtmlPage(c *gin.Context) {
	urlInput := c.Query("url")
	url, err := url.Parse(urlInput)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Bad request"})
		return
	}

	err = configCrawler.handler.GetHtmlPage(url)
	if err != nil {
		log.Errorln("error occurs when response page content to frontend: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Internal server error"})
		return
	}
	hostname := strings.TrimPrefix(url.Hostname(), "www.")
	filePath := fmt.Sprintf("page%s.html", hostname)

	file, err := os.Open(filePath)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()

	c.Header("Content-Type", "text/html")
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	err = os.Remove(filePath)
	if err != nil {
		log.Errorln("error occurs when delete html file: ", err)
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Get page success"})
}
