package schedulesservices

import (
	"context"
	"server/entities"
	pb "server/proto"
	"server/repository"
	"server/services"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/elastic/go-elasticsearch/v7"
)

var SCHEDULE_INDEX_NAME = "schedules"

var NOTI_COMPLETE_CRAWL_TITLE = "Crawl schedule success"
var NOTI_COMPLETE_CRAWL_TYPE = "INFO"

type schedulesService struct {
	grpcClient             pb.CrawlerServiceClient
	es                     *elasticsearch.Client
	leagueServices         services.LeaguesServices
	tagsServices           services.TagsServices
	matchService           services.MatchServices
	repo                   repository.SchedulesRepository
	matchURLsWithTimeOnDay repository.MatchURLsWithTimeOnDay
	allMatchURLsOnDay      repository.AllMatchURLsOnDay
}

func NewSchedulesService(leagueServices services.LeaguesServices, tagsServices services.TagsServices, grpcClient pb.CrawlerServiceClient, es *elasticsearch.Client, repo repository.SchedulesRepository, matchService services.MatchServices) *schedulesService {
	schedulesService := &schedulesService{
		grpcClient:         grpcClient,
		es:                 es,
		leagueServices:     leagueServices,
		tagsServices:       tagsServices,
		matchService: matchService,
		repo:               repo,
	}
	return schedulesService
}

func (s *schedulesService) GetSchedules(dateString string) {
	pbdate := &pb.Date{
		Date: dateString,
	}
	// send gRPC request to crawler
	pbSchedules, err := s.grpcClient.GetSchedulesOnDay(context.Background(), pbdate)
	if err != nil {
		log.Errorf("error occurred while get schedule on day from crawler error %v \n", err)
		return
	}

	dateCasted, err := time.Parse("02-01-2006", pbSchedules.GetDateFormated())
	if err != nil {
		log.Errorf("error when parse date: %v", err)
	}

	// crawler return all schedule on that day, we need to crack down it and save each schedule on db
	for _, pbSchedule := range pbSchedules.GetScheduleOnLeagues() {
		schedule := s.storeSchedule(pbSchedule, dateCasted)
		s.createMatchUrl(schedule, dateCasted)
		s.autoStoreNewLeague(schedule)
	}
}

func (s *schedulesService) storeSchedule(pbSchedule *pb.ScheduleOnLeague, dateCasted time.Time) entities.Schedule {
	schedule := newEntitiesSchedule(pbSchedule, dateCasted)
	err := s.repo.FirstOrCreate(&schedule)
	if err != nil {
		log.Error(err)
	}

	for _, match := range schedule.Matches {
		err := s.matchService.StoreMatch_ScheduleCrawl(match, schedule.ID, schedule.Date)
		if err != nil {
			log.Error(err)
		}
	}

	return schedule
}

// get the link of the matches taking place during the day, used to crawl the match details
func (s *schedulesService) createMatchUrl(schedule entities.Schedule, dateCasted time.Time) {
	for _, match := range schedule.Matches {
		exactTime, err := readTime(match, schedule.Date)
		if err != nil {
			log.Error(err)
			continue
		}
		addMatchUrl(exactTime, match.MatchDetailLink, &s.matchURLsWithTimeOnDay)
	}
	allMatchUrl := getMatchUrl(schedule)
	s.allMatchURLsOnDay.Date = dateCasted
	s.allMatchURLsOnDay.Urls = append(s.allMatchURLsOnDay.Urls, allMatchUrl...)
}

func (s *schedulesService) autoStoreNewLeague(schedule entities.Schedule) {
	leagues, err := s.leagueServices.ListLeagues()
	if err != nil {
		log.Error(err)
	}
	isNewLeague := isNewLeague(*leagues, schedule.LeagueName)
	if isNewLeague {
		log.Println("detect a new league: ", schedule.LeagueName)
		err = s.leagueServices.CreateLeague(schedule.LeagueName)
		if err != nil {
			log.Error(err)
		}
	}
}

func (s *schedulesService) SignalMatchDetailServiceToCrawl(matchURLs repository.AllMatchURLsOnDay) []*pb.MatchDetail {
	return s.matchService.GetMatchDetailsOnDayFromCrawler(matchURLs)
}

func (s *schedulesService) GetAllScheduleLeagueOnDay(date time.Time) (repository.ScheduleOnDay, error) {
	var scheduleOnDay repository.ScheduleOnDay

	scheduleOnDay.Date = date
	scheduleOnDay.DateWithWeekday = date.Weekday().String()

	entitySchedules, err := s.repo.GetScheduleOnDay(date)
	if err != nil {
		return scheduleOnDay, err
	}

	activeLeague, err := s.leagueServices.GetLeaguesNameActive()
	if err != nil {
		log.Error(err)
	}

	for _, entitySchedule := range *entitySchedules {
		if checkIsScheduleOnActiveLeague(activeLeague, entitySchedule.LeagueName) {
			scheduleOnDay.ScheduleOnLeagues = append(scheduleOnDay.ScheduleOnLeagues, newScheduleOnLeague(entitySchedule))
		}
	}

	return scheduleOnDay, nil
}

func (s *schedulesService) GetScheduleLeagueOnDay(date time.Time, league string) (repository.ScheduleOnDay, error) {
	var scheduleOnDay repository.ScheduleOnDay

	scheduleOnDay.Date = date
	scheduleOnDay.DateWithWeekday = date.Weekday().String()

	entitySchedule, err := s.repo.GetScheduleOnLeague(league, date)
	if err != nil {
		return scheduleOnDay, err
	}

	if len(entitySchedule.Matches) == 0 {
		return scheduleOnDay, nil
	}

	scheduleOnDay.ScheduleOnLeagues = append(scheduleOnDay.ScheduleOnLeagues, newScheduleOnLeague(*entitySchedule))

	return scheduleOnDay, nil
}

func (s *schedulesService) GetMatchURLsOnTime() repository.MatchURLsWithTimeOnDay {
	return s.matchURLsWithTimeOnDay
}

func (s *schedulesService) ClearMatchURLsOnTime() {
	s.matchURLsWithTimeOnDay = repository.MatchURLsWithTimeOnDay{}
}

func (s *schedulesService) GetAllMatchURLs() repository.AllMatchURLsOnDay {
	return s.allMatchURLsOnDay
}

func (s *schedulesService) ClearAllMatchURLs() {
	s.allMatchURLsOnDay = repository.AllMatchURLsOnDay{}
}
