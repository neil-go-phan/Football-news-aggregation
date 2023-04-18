package routes

import (
	"server/handler"

	"github.com/gin-gonic/gin"
)

type TagsRoutes struct {
	tagsHandler handler.TagsHandler
}

func NewTagsRoutes(tagsHandler *handler.TagsHandler) *TagsRoutes{
	routes := &TagsRoutes{
		tagsHandler: *tagsHandler,
	}
	return routes
}

func (tagsRoutes *TagsRoutes)Setup(r *gin.Engine) {
	routes := r.Group("tags")
	{
		routes.GET("list", tagsRoutes.tagsHandler.ListTags)
		routes.GET("delete", tagsRoutes.tagsHandler.DeleteTag)
		routes.GET("add", tagsRoutes.tagsHandler.AddTag)
	}
}