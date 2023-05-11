package entities

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	TagName string `gorm:"uniqueIndex"`
}