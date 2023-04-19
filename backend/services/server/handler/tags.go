package handler

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"server/services"

	"github.com/gin-gonic/gin"
)

// user search article, server query elastic search

type TagsHandler struct {
	handler services.TagsServices
}

func NewTagsHandler(handler services.TagsServices) *TagsHandler {
	tagsHandler := &TagsHandler{
		handler: handler,
	}
	return tagsHandler;
}

func (tagsHandler *TagsHandler) ListTags(c *gin.Context) {
	tags := tagsHandler.handler.ListTags()
	c.JSON(http.StatusOK, gin.H{"success": true, "tags": tags})
}

func (tagsHandler *TagsHandler) DeleteTag(c *gin.Context) {
	tag := c.Query("tag")
	err := tagsHandler.handler.DeleteTag(tag)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Delete successfull"})
}

func (tagsHandler *TagsHandler) AddTag(c *gin.Context) {
	tag := c.Query("tag")
	err := tagsHandler.handler.AddTag(tag)
	if err != nil {
		log.Printf("error occus when add tag: %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Add tag failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Add tag successfull"})
}