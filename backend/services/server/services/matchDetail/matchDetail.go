package matchdetailservices

import (
	"server/entities"
	"server/repository"
	"time"
)

type matchDetailService struct {
	repo repository.MatchDetailRepository
}

func NewMatchDetailervice(repo repository.MatchDetailRepository) *matchDetailService {
	matchDetailService := &matchDetailService{
		repo: repo,
	}
	return matchDetailService
}

func (s *matchDetailService) GetMatchDetailsOnDayFromCrawler(matchURLs entities.MatchURLsOnDay) {
	s.repo.GetMatchDetailsOnDayFromCrawler(matchURLs)
}

func (s *matchDetailService) GetMatchDetail(date time.Time, club1Name string, club2Name string) (entities.MatchDetail, error) {
	return s.repo.GetMatchDetail(date, club1Name, club2Name)
}
