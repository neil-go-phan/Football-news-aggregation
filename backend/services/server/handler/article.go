package handler

import (
	"log"
	"net/http"
	"server/helper"
	"server/services"
	"strconv"
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
	fromString := c.Query("from")
	fromInt , err := strconv.Atoi(fromString)
	formatedTags := serverhelper.FortmatTagsFromRequest(tags)

	if err != nil {
		log.Printf("can not convert %s string to int err: %v\n",fromString, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "message": "Bad request"})
	}
	// request to elasticsearch
	keyword = strings.TrimSpace(keyword)
	articles, err := articleHandler.handler.SearchArticlesTagsAndKeyword(keyword, formatedTags, fromInt)
	if err != nil {
		log.Printf("error occurred while services layer searching for keyword %s, with index: %s, err: %v\n", keyword, "articles", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Server error"})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "articles": articles})
}



func (articleHandler *ArticleHandler) CrawlArticleLeague(c *gin.Context) {
	leagueName := c.Query("league")

	league := []string{leagueName}

	articleHandler.handler.GetArticles(league)

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Signal crawl artilce success"})
}

func (articleHandler *ArticleHandler) AddUpdateNewTag(c *gin.Context) {
	tag := c.Query("tag")

	err := articleHandler.handler.AddTagForAllArticle(tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Update tag failed"})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Update tag successfull"})
}

func (articleHandler *ArticleHandler) SignalToCrawlerAfter10Min(cronjob *cron.Cron) {
	_, err := cronjob.AddFunc("@every 0h10m", func() { articleHandler.handler.GetArticles(make([]string, 0)) })
	if err != nil {
		log.Println("error occurred while seting up getArticle cronjob: ", err)
	}
}
