package handler

import (
	"net/http"
	"server/handler/presenter"
	"server/helper"
	"server/services"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	handler services.ArticleServices
}

type InputDeleteArticleID struct {
	ID uint `json:"id"`
}

func NewArticleHandler(handler services.ArticleServices) *ArticleHandler {
	userHandler := &ArticleHandler{
		handler: handler,
	}
	return userHandler
}

func (articleHandler *ArticleHandler) APISearchTagsAndKeyword(c *gin.Context) {
	keyword := c.Query("q")
	tags := c.Query("tags")
	fromString := c.Query("from")
	fromInt, err := strconv.Atoi(fromString)
	formatedTags := serverhelper.FortmatTagsFromRequest(tags)

	if err != nil {
		log.Printf("can not convert %s string to int err: %v\n", fromString, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "message": "Bad request"})
		return
	}

	keyword = strings.TrimSpace(keyword)
	articles, total, err := articleHandler.handler.SearchArticles(keyword, formatedTags, fromInt)
	if err != nil {
		log.Printf("error occurred while services layer searching for keyword %s, with index: %s, err: %v\n", keyword, "articles", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Server error"})
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

	c.JSON(http.StatusOK, gin.H{"success": true, "articles": resposeArticles, "total": total})
}

func (articleHandler *ArticleHandler) APICrawlArticleLeague(c *gin.Context) {
	leagueName := c.Query("league")

	league := []string{leagueName}

	articleHandler.handler.GetArticles(league)

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Signal crawl artilce success"})
}

func (articleHandler *ArticleHandler) APIUpdateNewTag(c *gin.Context) {
	tag := c.Query("tag")
	tagFormated := serverhelper.FormatVietnamese(tag)
	err := articleHandler.handler.AddTagForAllArticle(tagFormated)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Update tag failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Update tag successfull"})
}

func (articleHandler *ArticleHandler) APIDeleteArticle(c *gin.Context) {
	var inputArticle InputDeleteArticleID
	err := c.BindJSON(&inputArticle)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Delete article failed"})
		return
	}
	err = articleHandler.handler.DeleteArticle(inputArticle.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Delete article failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Delete article successfull"})
}

func (articleHandler *ArticleHandler) APIGetFirstPageOfLeagueRelatedArticle(c *gin.Context) {
	league := c.Query("league")

	articles, err := articleHandler.handler.GetFirstPageOfLeagueRelatedArticle(league)
	if err != nil {
		log.Printf("error occurred while services layer searching for keyword with index: %s, err: %v\n", "articles", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "articles": articles})
}

func (articleHandler *ArticleHandler) RefreshCache() {
	articleHandler.handler.RefreshCache()
}

func (articleHandler *ArticleHandler) GetArticleSeed() {
	articleHandler.handler.GetArticles(make([]string, 0))
}

func (articleHandler *ArticleHandler) APIGetArticleCount(c *gin.Context) {
	total, today, err := articleHandler.handler.GetArticleCount()
	if err != nil {
		log.Printf("error occrus when get article count %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Get article count success", "total": total, "today": today})
}
