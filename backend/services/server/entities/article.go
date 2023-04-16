package entities

import "time"

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Tags []string `json:"tags"`
}

type ElasticArticle struct {
	CreatedAt time.Time `json:"created_at"`
	Article
}