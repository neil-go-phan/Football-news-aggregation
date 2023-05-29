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
		routes.GET("get", middlewares.CheckToken(), configCrawlerRoute.configCrawlerHandler.APIGetConfigCrawler)
		routes.GET("delete", middlewares.CheckToken(), configCrawlerRoute.configCrawlerHandler.APIDeleteConfigCrawler)
		routes.GET("list", middlewares.CheckToken(), configCrawlerRoute.configCrawlerHandler.APIListConfigCrawler)
		routes.POST("upsert", middlewares.CheckToken(), configCrawlerRoute.configCrawlerHandler.APIUpsertConfigCrawler)
		routes.POST("test", middlewares.CheckToken(), configCrawlerRoute.configCrawlerHandler.APITestCrawler)
		routes.POST("change-time", middlewares.CheckToken(), configCrawlerRoute.configCrawlerHandler.APIChangeCronjobTime)
	}
}
