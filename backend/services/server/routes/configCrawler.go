package routes

import (
	"server/handler"
	"server/middlewares"

	"github.com/gin-gonic/gin"
)

type ConfigCrawlerRoutes struct {
	configCrawlerHandler handler.ConfigCrawlerHandler
}

func NewConfigCrawlerRoutes(configCrawler *handler.ConfigCrawlerHandler) *ConfigCrawlerRoutes {
	routes := &ConfigCrawlerRoutes{
		configCrawlerHandler: *configCrawler,
	}
	return routes
}

func (configCrawlerRoute *ConfigCrawlerRoutes) Setup(r *gin.Engine) {
	routes := r.Group("crawler")
	{
		routes.GET("get-html-page", middlewares.CheckToken(), configCrawlerRoute.configCrawlerHandler.APIGetHtmlPage)
	}
}
