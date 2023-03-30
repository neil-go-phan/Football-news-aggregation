package handler

import (
	"backend/services/server/services"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"log"
	"net/http"
)

// user search article, server query elastic search

type ArticleHandler struct {
	handler services.ArticleServices
}

func NewArticleHandler(handler services.ArticleServices) *ArticleHandler {
	userHandler := &ArticleHandler{
		handler: handler,
	}
	return userHandler
}

func (articleHandler *ArticleHandler) SearchWithIndexName(c *gin.Context) {
	keyword := c.Query("q")
	index := c.Query("index")

	// check if index is valid

	// request to elasticsearch

	articles, err := articleHandler.handler.FrontendSearchWithIndex(keyword, index)
	if err != nil {
		log.Printf("error occurred while services layer searching for keyword %s, with index: %s, err: %v\n", keyword, index, err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Server error"})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "articles": articles})
}

func (articleHandler *ArticleHandler) SignalToCrawler(cronjob *cron.Cron) {
	articleHandler.handler.GetArticles()
	_, err := cronjob.AddFunc("@every 0h01m", func() { articleHandler.handler.GetArticles() })
	if err != nil {
		log.Println("error occurred while seting up getArticle cronjob: ", err)
	}
}

// user get match result, server query cache (redis)
