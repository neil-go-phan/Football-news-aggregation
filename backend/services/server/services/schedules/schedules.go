package schedulesservices

import (
	"server/entities"
	pb "server/proto"
	"server/repository"
	"time"
)

type schedulesService struct {
	repo repository.SchedulesRepository
	matchDetailRepo repository.MatchDetailRepository
}

func NewSchedulesService(repo repository.SchedulesRepository, matchDetailRepo repository.MatchDetailRepository) *schedulesService {
	schedulesService := &schedulesService{
		repo: repo,
		matchDetailRepo: matchDetailRepo,
	}
	return schedulesService
}

func (s *schedulesService) SignalMatchDetailServiceToCrawl(matchURLs entities.MatchURLsOnDay){
	s.matchDetailRepo.GetMatchDetailsOnDayFromCrawler(matchURLs)
}

func (s *schedulesService) GetSchedules(date string) {
	in := &pb.Date{
		Date: date,
	}
	s.repo.GetSchedules(in)
}

func (s *schedulesService) GetAllScheduleLeagueOnDay(date time.Time) (entities.ScheduleOnDay, error) {
	return s.repo.GetAllScheduleLeagueOnDay(date)
}

func (s *schedulesService) GetScheduleLeagueOnDay(date time.Time, league string) (entities.ScheduleOnDay, error) {
	return s.repo.GetScheduleLeagueOnDay(date, league)
}

func (s *schedulesService) GetMatchURLsOnDay() entities.MatchURLsOnDay {
	return s.repo.GetMatchURLsOnDay()
}

func (s *schedulesService) ClearMatchURLsOnDay() {
	s.repo.ClearMatchURLsOnDay()
}

func (s *schedulesService)  GetMatchURLsOnTime() entities.MatchURLsWithTimeOnDay {
	return s.repo.GetMatchURLsOnTime()
}

func (s *schedulesService) ClearMatchURLsOnTime() {
	s.repo.ClearMatchURLsOnTime()
}