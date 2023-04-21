package handler

import (
	notificationservices "server/services/notification"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	handler *notificationservices.NotificationService
	// notificationChan chan entities.Notification
}

func NewNotificationHandler(handler *notificationservices.NotificationService) *NotificationHandler {
	notificationHandler := &NotificationHandler{
		handler: handler,
	}
	return notificationHandler
}

func (notificationHandler *NotificationHandler) Listen(c *gin.Context) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	defer close(notificationHandler.handler.Repo.Notification)

	go func() {
			for {
					message := <-notificationHandler.handler.Repo.Notification
					c.SSEvent("message", message)
			}
	}()

	// block the request until the connection is closed
	<-c.Writer.CloseNotify()
}