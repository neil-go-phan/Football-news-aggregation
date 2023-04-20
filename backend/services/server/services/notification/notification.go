package notificationservices

import (
	notificationrepo "server/repository/notification"
)

type NotificationService struct {
	Repo *notificationrepo.NotificationRepo
}

func NewNotificationService(repo *notificationrepo.NotificationRepo) *NotificationService {
	notificationService := &NotificationService{
		Repo: repo,
	}
	return notificationService
}

func (s *NotificationService)Send(title string, notiType string, message string)  {
	s.Repo.Send(title, notiType, message)
}