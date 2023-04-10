package handlers

import (
	"context"
	"crawler/services"
	"fmt"
	"log"
	"time"

	"crawler/entities"
	pb "crawler/proto"
)

func (s *gRPCServer) GetSchedulesOnDay(ctx context.Context, date *pb.Date) (*pb.SchedulesReponse, error) {

	log.Println("Start scrapt schedule on date", date.GetDate())

	schedules, err := crawlSchedulesAndParse(date)
	fmt.Println(schedules.DateFormated)
	if err != nil {
		log.Printf("error occurred while get schedule for day: %s, err: %v \n", date.GetDate(), err)
	}

	log.Println("Finish scrapt schedule")
	return schedules, nil
}

func crawlSchedulesAndParse(date *pb.Date) (*pb.SchedulesReponse, error) {
	dateIn := date.GetDate()
	ok := checkDateFormat(dateIn)
	if !ok {
		log.Println("Date is invalid")
		return nil, fmt.Errorf("date %s is invalid", dateIn)
	}

	schedulesCrawl, err := services.CrawlSchedules(dateIn)
	if err != nil {
		return nil, fmt.Errorf("error occurred during crawl schedule on day %s, err: %v", dateIn, err)
	}

	schedules := crawledSchedulesToPbSchedules(schedulesCrawl)
	if err != nil {
		return nil, fmt.Errorf("error occurred while sending response to client: %v", err)
	}

	log.Printf("Date: %s crawl successfully", dateIn)
	return schedules, nil
}

func checkDateFormat(dateStr string) bool {
	_, err := time.Parse("02-01-2006", dateStr)
	return err == nil
}

func crawledSchedulesToPbSchedules(crawledSchedules entities.ScheduleOnDay) *pb.SchedulesReponse {
	pbSchedules := &pb.SchedulesReponse{}
	for _, scheduleOnLeague := range crawledSchedules.ScheduleOnLeagues {
		pbSchedule := &pb.ScheduleOnLeague{
			LeagueName: scheduleOnLeague.LeagueName,
		}
		for _, match := range scheduleOnLeague.Matchs {
			pgMatch := &pb.Match{
				Time:            match.Time,
				Round:           match.Round,
				Scores:          match.Scores,
				MatchDetailLink: match.MatchDetailLink,
				Club1: &pb.Club{
					Name: match.Club1.Name,
					Logo: match.Club1.Logo,
				},
				Club2: &pb.Club{
					Name: match.Club2.Name,
					Logo: match.Club2.Logo,
				},
			}

			pbSchedule.Matchs = append(pbSchedule.Matchs, pgMatch)
		}
		pbSchedules.ScheduleOnLeagues = append(pbSchedules.ScheduleOnLeagues, pbSchedule)
	}
	pbSchedules.DateFormated = crawledSchedules.Date
	return pbSchedules
}
