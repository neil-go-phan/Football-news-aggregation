package routes

import (
	"server/handler"
	"server/middlewares"

	"github.com/gin-gonic/gin"
)

type AdminRoutes struct {
	adminHandler handler.AdminHandler
}

func NewAdminRoutes(adminHandler *handler.AdminHandler) *AdminRoutes{
	routes := &AdminRoutes{
		adminHandler: *adminHandler,
	}
	return routes
}

func (adminRoute *AdminRoutes)Setup(r *gin.Engine) {
	routes := r.Group("admin")
	{
		routes.POST("login", adminRoute.adminHandler.Login)
		routes.GET("check-auth", middlewares.CheckToken(), adminRoute.adminHandler.CheckAuth)
		routes.GET("get", middlewares.CheckToken(), adminRoute.adminHandler.Get)
		routes.POST("change-password", middlewares.CheckToken(), adminRoute.adminHandler.ChangePassword)
	}
}