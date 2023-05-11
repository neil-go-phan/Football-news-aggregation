package entities

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title       string `gorm:"uniqueIndex:idx_title_link"`
	Description string 
	Link        string `gorm:"uniqueIndex:idx_title_link"`
	Tags []Tag `gorm:"many2many:article_tag;constraint:OnDelete:CASCADE;"`
}

type ElasticArticle struct {
	ID uint `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Tags []string `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
}