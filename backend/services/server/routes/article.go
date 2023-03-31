package routes

import (
	"server/handler"

	"github.com/gin-gonic/gin"
)

type ArticleRoutes struct {
	articleHandler handler.ArticleHandler
}

func NewArticleRoutes(articleHandler *handler.ArticleHandler) *ArticleRoutes{
	routes := &ArticleRoutes{
		articleHandler: *articleHandler,
	}
	return routes
}

func (articleRoutes *ArticleRoutes)Setup(r *gin.Engine) {
	routes := r.Group("article")
	{
		routes.GET("search-tag-keyword",  articleRoutes.articleHandler.SearchTagsAndKeyword)
		routes.GET("search-all",  articleRoutes.articleHandler.GetAllFromElastic)
	}
}