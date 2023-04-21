package notificationrepo

import (
	"server/entities"
	"time"
)

type NotificationRepo struct {
	Notification chan entities.Notification
}

func NewNotificationRepo(notification chan entities.Notification) *NotificationRepo {
	return &NotificationRepo{
		Notification: notification,
	}
}

func (repo *NotificationRepo) Send(title string, notiType string, message string) {
	repo.Notification <- entities.Notification{
		Time:     time.Now(),
		Title:    title,
		NotiType: notiType,
		Message:  message,
	}
}

func (repo *NotificationRepo) Listen() {

}