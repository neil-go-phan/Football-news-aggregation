package routes

import (
	"server/handler"
	"server/middlewares"

	"github.com/gin-gonic/gin"
)

type ArticleRoutes struct {
	articleHandler handler.ArticleHandler
}

func NewArticleRoutes(articleHandler *handler.ArticleHandler) *ArticleRoutes {
	routes := &ArticleRoutes{
		articleHandler: *articleHandler,
	}
	return routes
}

func (articleRoutes *ArticleRoutes) Setup(r *gin.Engine) {
	routes := r.Group("article")
	{
		routes.GET("search-tag-keyword", articleRoutes.articleHandler.APISearchTagsAndKeyword)
		routes.GET("signal-crawler", middlewares.CheckToken(), articleRoutes.articleHandler.APICrawlArticleLeague)
		routes.GET("update-tag", middlewares.CheckToken(), articleRoutes.articleHandler.APIAddUpdateNewTag)
		routes.GET("get-first-page", articleRoutes.articleHandler.APIGetFirstPageOfLeagueRelatedArticle)
		routes.GET("count", middlewares.CheckToken(), articleRoutes.articleHandler.APIGetArticleCount)
		routes.POST("delete", middlewares.CheckToken(), articleRoutes.articleHandler.APIDeleteArticle)
	}
}
