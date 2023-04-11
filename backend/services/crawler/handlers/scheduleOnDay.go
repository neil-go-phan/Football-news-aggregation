package handlers

import (
	"context"
	"crawler/services"
	"fmt"
	"log"
	"time"

	"crawler/entities"
	pb "crawler/proto"

	jsoniter "github.com/json-iterator/go"
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
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	pbSchedules := &pb.SchedulesReponse{}
  scheduleByte, err := json.Marshal(crawledSchedules)
	if err != nil {
		log.Printf("error occrus when marshal crawled schedules: %s", err)
	}
  err = json.Unmarshal(scheduleByte, pbSchedules)
	if err != nil {
		log.Printf("error occrus when unmarshal crawled schedules to proto.Schedules: %s", err)
	}
	return pbSchedules
}
