package handler

import (
	"net/http"
	"server/handler/presenter"
	serverhelper "server/helper"
	"server/services"

	log "github.com/sirupsen/logrus"

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
	tags, err := tagsHandler.handler.ListTagsName()
	if err != nil {
		log.Errorln("error occurs when response list tag to frontend: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Internal server error"})
		return
	}
	resposeTags := presenter.Tags{
		Tags: tags,
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "tags": resposeTags})
}

func (tagsHandler *TagsHandler) DeleteTag(c *gin.Context) {
	tag := c.Query("tag")
	tagFormated := serverhelper.FormatVietnamese(tag)
	err := tagsHandler.handler.DeleteTag(tagFormated)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Delete successfull"})
}

func (tagsHandler *TagsHandler) AddTag(c *gin.Context) {
	tag := c.Query("tag")
	tagFormated := serverhelper.FormatVietnamese(tag)
	err := tagsHandler.handler.AddTag(tagFormated)
	if err != nil {
		log.Printf("error occus when add tag: %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Add tag failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Add tag successfull"})
}