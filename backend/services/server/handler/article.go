package handler

import (
	"backend/services/server/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// user search article, server query elastic search

type ArticleHandler struct {
	handler  services.ArticleServices
}

func NewArticleHandler(handler services.ArticleServices) *ArticleHandler {
	userHandler := &ArticleHandler{
		handler:  handler,
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
		log.Printf("error occurred while services layer searching for keyword %s, with index: %s, err: %v\n",keyword, index , err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message" : "Server error"})
	}
	 fmt.Println(articles)
	c.JSON(http.StatusOK, gin.H{"success" : true, "articles" : articles})
}


// user get match result, server query cache (redis)