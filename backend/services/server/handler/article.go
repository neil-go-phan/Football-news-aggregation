package handler

import (
	"server/helper"
	"server/services"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
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

func (articleHandler *ArticleHandler) SearchTagsAndKeyword(c *gin.Context) {
	keyword := c.Query("q")
	tags := c.Query("tags")
	formatedTags := serverhelper.FortmatTagsFromRequest(tags)

	// request to elasticsearch
	keyword= strings.TrimSpace(keyword)
	articles, err := articleHandler.handler.APISearchArticlesTagsAndKeyword(keyword, formatedTags)
	if err != nil {
		log.Printf("error occurred while services layer searching for keyword %s, with index: %s, err: %v\n", keyword, "articles", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Server error"})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "articles": articles})
}

func (articleHandler *ArticleHandler) GetAllFromElastic(c *gin.Context) {
	search_type := c.Query("search_type")
	scroll := c.Query("scroll")
	size := c.Query("size")

	articles, err := articleHandler.handler.APISearchAll(search_type, scroll, size)
	if err != nil {
		log.Printf("error occurred while services layer try to get all article from eslaticsearch err: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Server error"})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "articles": articles})
}

func (articleHandler *ArticleHandler) SignalToCrawlerAfter10Min(cronjob *cron.Cron) {
	_, err := cronjob.AddFunc("@every 0h10m", func() { articleHandler.handler.GetArticles(make([]string, 0)) })
	if err != nil {
		log.Println("error occurred while seting up getArticle cronjob: ", err)
	}
}
