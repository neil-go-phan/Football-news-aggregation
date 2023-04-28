package articlesservices

import (
	"server/entities"
	"server/repository"
)

type articleService struct {
	repo repository.ArticleRepository
}

func NewArticleService(repo repository.ArticleRepository) *articleService {
	articleService := &articleService{
		repo: repo,
	}
	return articleService
}

func (s *articleService) SearchArticlesTagsAndKeyword(keyword string, formatedTags []string, from int) ([]entities.Article, float64, error) {
	return s.repo.SearchArticlesTagsAndKeyword(keyword, formatedTags, from)
}

func (s *articleService) GetFirstPageOfLeagueRelatedArticle(leagueName string) ([]entities.Article, error) {
	return s.repo.GetFirstPageOfLeagueRelatedArticle(leagueName)
}

func (s *articleService) RefreshCache() {
	s.repo.RefreshCache()
}

func (s *articleService) GetArticleCount() (total float64, today float64, err error) {
	return s.repo.GetArticleCount()
}

func (s *articleService) AddTagForAllArticle(tag string) error {
	return s.repo.AddTagForAllArticle(tag)
}

func (s *articleService) GetArticles(keywords []string) {
	s.repo.GetArticles(keywords)
}

func (s *articleService) DeleteArticle(title string) error {
	return s.repo.DeleteArticle(title)
}
