package matchservices

import (
	"context"
	"fmt"
	"io"
	"server/entities"
	pb "server/proto"
	"server/repository"
	"server/services"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

type MatchService struct {
	grpcClient          pb.CrawlerServiceClient
	repo                repository.MatchRepository
	clubService         services.ClubServices
	statsItemService    services.StatsItemServices
	eventService        services.EventServices
	overviewItemService services.OverviewItemServices
	lineupService       services.LineUpServices
	playerService       services.PlayerServices
}

var DEFAULT_LEAGUE_NAME = "Tin tức bóng đá"

func NewMatchService(grpcClient pb.CrawlerServiceClient, repo repository.MatchRepository, clubService services.ClubServices, statsItemService services.StatsItemServices, eventService services.EventServices, overviewItemService services.OverviewItemServices, lineupService services.LineUpServices, playerService services.PlayerServices) *MatchService {
	matchService := &MatchService{
		grpcClient:          grpcClient,
		repo:                repo,
		clubService:         clubService,
		statsItemService:    statsItemService,
		eventService:        eventService,
		overviewItemService: overviewItemService,
		lineupService:       lineupService,
		playerService:       playerService,
	}
	return matchService
}

func (s *MatchService) GetMatchDetailsOnDayFromCrawler(matchURLs repository.AllMatchURLsOnDay) []*pb.MatchDetail {
	in := &pb.MatchURLs{
		Url: matchURLs.Urls,
	}

	matches := make([]*pb.MatchDetail, 0)

	// send gRPC request to crawler
	stream, err := s.grpcClient.GetMatchDetail(context.Background(), in)
	if err != nil {
		log.Errorf("error occurred while openning stream error %v ", err)
		return matches
	}

	done := make(chan bool)
	log.Printf("Start get stream of match detail...")

	go func(date time.Time) {
		for {
			resp, err := stream.Recv()
			if err != nil {
				status, _ := status.FromError(err)
				if status.Code().String() == "Unavailable" || err == io.EOF {
					done <- true
					return
				}
			}

			err = s.StoreMatch_MatchDetailCrawl(resp, date)
			if err != nil {
				log.Error(err)
			} else {
				log.Println("Update match success")
			}

			matches = append(matches, resp)
		}
	}(matchURLs.Date)

	<-done

	log.Printf("finished crawl match detail")
	return matches
}

func (s *MatchService) StoreMatch_MatchDetailCrawl(resp *pb.MatchDetail, date time.Time) error {
	// get match
	if resp.MatchDetailTitle.Club_1.Name == "" && resp.MatchDetailTitle.Club_2.Name == "" {
		return fmt.Errorf("crawler return a empty match (timeout)")
	}
	// TODO: handle error match not found
	match, err := s.repo.GetIDWithDateAndClubName(date, resp.MatchDetailTitle.Club_1.Name, resp.MatchDetailTitle.Club_2.Name)
	if err != nil {
		return err
	}

	// create statis (statstic only render on the crawled website when match ended. so I use FirstOrCreate method in case it already end, same with oveview and event)
	for _, item := range resp.MatchStatistics.GetStatistics() {
		statsItem := &entities.StatisticsItem{
			MatchID:     match.ID,
			StatClub1:   item.StatClub_1,
			StatContent: item.StatContent,
			StatClub2:   item.StatClub_2,
		}
		err = s.statsItemService.FirstOrCreate(statsItem)
		if err != nil {
			log.Error(err)
		}
	}

	// create event
	for _, respEvent := range resp.MatchProgress.Events {
		exactTime, err := readTimeEvent(respEvent.Time, match.TimeStart)
		if err != nil {
			log.Error(exactTime)
		}
		event := &entities.MatchEvent{
			TimeString: respEvent.Time,
			MatchID:    match.ID,
			Content:    respEvent.Content,
			Time:       exactTime,
		}
		err = s.eventService.FirstOrCreate(event)
		if err != nil {
			log.Error(err)
		}
	}

	// create overview
	// club 1
	for _, overviewItem := range resp.MatchOverview.Club_1Overview {
		if overviewItem.Info == "" {
			continue
		}
		exactTime, err := readTimeEvent(overviewItem.Time, match.TimeStart)
		if err != nil {
			log.Error(err)
		}
		overview := &entities.OverviewItem{
			TimeString: overviewItem.Time,
			MatchID:    match.ID,
			ClubID:     match.Club1ID,
			Time:       exactTime,
			Info:       overviewItem.Info,
			ImageType:  overviewItem.ImageType,
		}
		err = s.overviewItemService.FirstOrCreate(overview)
		if err != nil {
			log.Error(err)
		}
	}
	// club 2
	for _, overviewItem := range resp.MatchOverview.Club_2Overview {
		if overviewItem.Info == "" {
			continue
		}
		exactTime, err := readTimeEvent(overviewItem.Time, match.TimeStart)
		if err != nil {
			log.Error(err)
		}
		overview := &entities.OverviewItem{
			TimeString: overviewItem.Time,
			MatchID:    match.ID,
			ClubID:     match.Club2ID,
			Time:       exactTime,
			Info:       overviewItem.Info,
			ImageType:  overviewItem.ImageType,
		}
		err = s.overviewItemService.FirstOrCreate(overview)
		if err != nil {
			log.Error(err)
		}
	}

	// create lineup
	// club 1
	newlineup1 := &entities.MatchLineUp{
		ClubName:   resp.MatchLineup.LineupClub_1.ClubName,
		Formation:  resp.MatchLineup.LineupClub_1.Formation,
		ShirtColor: resp.MatchLineup.LineupClub_1.ShirtColor,
	}

	lineup1, err := s.lineupService.GetOrCreate(newlineup1)
	if err != nil {
		log.Error(err)
	}
	// create player position
	for row, rowDetail := range resp.MatchLineup.LineupClub_1.PitchRow {
		for column, playerOnPosition := range rowDetail.GetPitchRowsDetail() {
			player := &entities.Player{
				PlayerName:    playerOnPosition.PlayerName,
				PlayerNumber:  playerOnPosition.PlayerNumber,
				MatchLineUpID: lineup1.ID,
				MatchID:       match.ID,
				Col:           column,
				Row:           row,
			}
			err := s.playerService.FirstOrCreate(player)
			if err != nil {
				log.Error(err)
			}
		}
	}

	// club 2
	newlineup2 := &entities.MatchLineUp{
		ClubName:   resp.MatchLineup.LineupClub_2.ClubName,
		Formation:  resp.MatchLineup.LineupClub_2.Formation,
		ShirtColor: resp.MatchLineup.LineupClub_2.ShirtColor,
	}

	lineup2, err := s.lineupService.GetOrCreate(newlineup2)
	if err != nil {
		log.Error(err)
	}
	// create player position
	for row, rowDetail := range resp.MatchLineup.LineupClub_2.PitchRow {
		for column, playerOnPosition := range rowDetail.GetPitchRowsDetail() {
			player := &entities.Player{
				PlayerName:    playerOnPosition.PlayerName,
				PlayerNumber:  playerOnPosition.PlayerNumber,
				MatchLineUpID: lineup2.ID,
				MatchID:       match.ID,
				Col:           column,
				Row:           row,
			}
			err := s.playerService.FirstOrCreate(player)
			if err != nil {
				log.Error(err)
			}
		}
	}

	match.MatchStatus = resp.MatchDetailTitle.MatchStatus
	match.Scores = resp.MatchDetailTitle.MatchScore
	match.LineupClub1ID = lineup1.ID
	match.LineupClub2ID = lineup2.ID
	// update score, match status
	err = s.repo.UpdateWhenMatchDetailCrawl(match)
	if err != nil {
		log.Error(err)
	}

	return nil
}

func (s *MatchService) GetMatch(date time.Time, club1Name string, club2Name string) (*entities.Match, error) {
	// get match
	match, err := s.repo.GetIDWithDateAndClubName(date, club1Name, club2Name)
	if err != nil {
		return nil, err
	}

	match, err = s.repo.GetMatch(match)
	if err != nil {
		return nil, err
	}

	return match, nil
}

func (s *MatchService)GetLineUps(id1 uint, id2 uint) (*entities.MatchLineUp, *entities.MatchLineUp, error) {
	return s.lineupService.GetLineUps(id1, id2)
}

func (s *MatchService) StoreMatch_ScheduleCrawl(match entities.Match, ID uint, date time.Time) error {
	// check if match exist
	timeStart, err := readTimeStart(match, date)
	if err != nil {
		log.Error(err)
	}
	match.TimeStart = timeStart
	existMatch, err := s.repo.GetIDWithDateAndClubName(date, match.Club1.Name, match.Club2.Name)
	if err != nil {
		// match is not exist
		log.Println("Match not exist, create a new one")
		match.ScheduleID = ID
		err = s.createMatch(&match)
		if err != nil {
			return err
		}
	} else {
		// match already exist, then we will update some collumn
		log.Println("Match update id =", existMatch.ID)
		existMatch.Time = match.Time
		existMatch.Round = match.Round
		existMatch.Scores = match.Scores
		existMatch.MatchDetailLink = match.MatchDetailLink
		err = s.repo.UpdateWhenScheduleCrawl(existMatch)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *MatchService) createMatch(match *entities.Match) error {
	club1, err := s.clubService.GetOrCreate(match.Club1.Name, match.Club1.Logo)
	if err != nil {
		log.Error(err)
	}

	match.Club1 = *club1

	club2, err := s.clubService.GetOrCreate(match.Club2.Name, match.Club2.Logo)
	if err != nil {
		log.Error(err)
	}
	match.Club2 = *club2
	err = s.repo.Create(match)
	if err != nil {
		return err
	}
	return nil
}
