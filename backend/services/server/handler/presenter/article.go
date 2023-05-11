package presenter

type Article struct {
	ID uint `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Tags []string `json:"tags"`
}