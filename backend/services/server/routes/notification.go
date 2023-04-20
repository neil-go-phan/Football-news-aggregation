package routes

import (
	"server/handler"

	"github.com/gin-gonic/gin"
)

type NotificationRoutes struct {
	notificationHandler handler.NotificationHandler
}

func NewNotificationRoutes(notificationHandler *handler.NotificationHandler) *NotificationRoutes {
	routes := &NotificationRoutes{
		notificationHandler: *notificationHandler,
	}
	return routes
}

func (notificationRoutes *NotificationRoutes) Setup(r *gin.Engine) {
	routes := r.Group("notification")
	{
		routes.GET("get", notificationRoutes.notificationHandler.Listen)
	}
}
