package schedulesservices

import (
	"server/entities"
	pb "server/proto"
	"server/repository"
	"time"
)

type schedulesService struct {
	repo repository.SchedulesRepository
}

func NewSchedulesService(repo repository.SchedulesRepository) *schedulesService {
	schedulesService := &schedulesService{
		repo: repo,
	}
	return schedulesService
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
