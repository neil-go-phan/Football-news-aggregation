package services

import (
	"context"
	"fmt"
	"log"
	"server/entities"
	pb "server/proto"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"google.golang.org/grpc"
)

var SCHEDULE_ON_DAY_INDEX_NAME = "scheduleonday"

type scheduleOnDayService struct {
	conn *grpc.ClientConn
	es   *elasticsearch.Client
}

func NewScheduleOnDayService(conn *grpc.ClientConn, es *elasticsearch.Client) *scheduleOnDayService {
	scheduleOnDayService := &scheduleOnDayService{
		conn: conn,
		es:   es,
	}
	return scheduleOnDayService
}

func (s *scheduleOnDayService) GetSchedules() {
	client := pb.NewCrawlerServiceClient(s.conn)

	date := time.Now().Format("02-01-2006")

	in := &pb.Date{
		Date: date,
	}
	// send gRPC request to crawler
	pbSchedules, err := client.GetSchedulesOnDay(context.Background(), in)
	if err != nil {
		log.Printf("error occurred while get schedule on day from crawler error %v \n", err)
		return
	}
	fmt.Println(pbSchedules)
}

func PbSchedulesToEntitySchedule(pbSchedule *pb.SchedulesReponse) entities.ScheduleOnDay {
	entitySchedules := entities.ScheduleOnDay{
		DateWithWeekday: pbSchedule.DateWithWeekday,
		Date:            time.Now(),
	}

	for _, scheduleOnLeague := range pbSchedule.GetScheduleOnLeagues() {
		entitySchedule := entities.ScheduleOnLeague{
			LeagueName: scheduleOnLeague.LeagueName,
		}
		for _, match := range scheduleOnLeague.Matchs {
			entityMatch := entities.Match{
				Time:       match.Time,
				Round:      match.Round,
				Scores:     match.Scores,
				MatchDetailLink: match.MatchDetailLink,
				Club1: entities.Club{
					Name: match.Club1.Name,
					Logo: match.Club1.Logo,
				},
				Club2: entities.Club{
					Name: match.Club2.Name,
					Logo: match.Club2.Logo,
				},
			}

			entitySchedule.Matchs = append(entitySchedule.Matchs, entityMatch)
		}
		entitySchedules.ScheduleOnLeagues = append(entitySchedules.ScheduleOnLeagues, entitySchedule)
	}

	return entitySchedules
}
