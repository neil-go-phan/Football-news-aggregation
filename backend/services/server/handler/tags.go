package handler

import (
	"backend/services/server/services"
	"net/http"

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

