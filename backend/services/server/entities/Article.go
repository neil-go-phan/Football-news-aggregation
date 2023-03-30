package entities

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Tags []string `json:"tags"`
}