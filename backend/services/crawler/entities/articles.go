package entities

type HtmlArticleClass struct {
	ArticleClass            string`json:"article_class"`
	ArticleTitleClass       string`json:"article_title_class"`
	ArticleDescriptionClass string`json:"article_description_class"`
	ArticleLinkClass        string`json:"article_link_class"`
}

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
}