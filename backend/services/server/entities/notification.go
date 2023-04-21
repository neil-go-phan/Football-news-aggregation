package entities

import "time"

type Notification struct {
	Time time.Time `json:"time"`
	NotiType string `json:"noti_type"` // include: warning, error, info
	Title string `json:"title"`
	Message string `json:"message"`
}