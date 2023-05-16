package handler

import (
	"net/http"
	"server/entities"
	"server/handler/presenter"
	"server/repository"
	"server/services"
	"sort"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type MatchDetailHandler struct {
	handler services.MatchServices
}

func NewMatchDetailHandler(handler services.MatchServices) *MatchDetailHandler {
	matchDetailHandler := &MatchDetailHandler{
		handler: handler,
	}
	return matchDetailHandler
}

func (matchDetailHandler *MatchDetailHandler) SignalToCrawler(matchUrl repository.AllMatchURLsOnDay) {
	matchDetailHandler.handler.GetMatchDetailsOnDayFromCrawler(matchUrl)
}

func (matchDetailHandler *MatchDetailHandler) APIGetMatchDetail(c *gin.Context) {
	dateString := c.Query("date")
	club1Name := c.Query("club_1")
	club2Name := c.Query("club_2")

	date, err := time.Parse(DATE_LAYOUT, dateString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "message": "Date invalid"})
		return
	}

	match, err := matchDetailHandler.handler.GetMatch(date, club1Name, club2Name)
	if err != nil {
		log.Printf("error occurred while services layer get match: %s\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Server error"})
		return
	}

	lineup1, lineup2, err := matchDetailHandler.handler.GetLineUps(match.LineupClub1ID, match.LineupClub2ID)
	if err != nil {
		log.Printf("error occurred while services layer get lineup: %s\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Server error"})
		return
	}

	matchResponse := castToPresenterMatch(*match, *lineup1, *lineup2)
	c.JSON(http.StatusOK, gin.H{"success": true, "match_detail": matchResponse})
}

func castToPresenterMatch(matchDetail entities.Match, lineup1 entities.MatchLineUp, lineup2 entities.MatchLineUp) presenter.MatchDetail {
	respMatchDetail := presenter.MatchDetail{}

	// cast title
	respMatchDetail.MatchDetailTitle.Club1.Name = matchDetail.Club1.Name
	respMatchDetail.MatchDetailTitle.Club1.Logo = matchDetail.Club1.Logo
	respMatchDetail.MatchDetailTitle.Club2.Name = matchDetail.Club2.Name
	respMatchDetail.MatchDetailTitle.Club2.Logo = matchDetail.Club2.Logo
	respMatchDetail.MatchDetailTitle.MatchScore = matchDetail.Scores
	respMatchDetail.MatchDetailTitle.MatchStatus = matchDetail.MatchStatus

	// cast overview
	club1Overview := make([]presenter.OverviewItem, 0)
	for _, item := range matchDetail.Club1Overview {
			overviewItem := presenter.OverviewItem{
				Time:      item.TimeString,
				Info:      item.Info,
				ImageType: item.ImageType,
			}
			club1Overview = append(club1Overview, overviewItem)
	}
	club2Overview := make([]presenter.OverviewItem, 0)
	for _, item := range matchDetail.Club2Overview {
		overviewItem := presenter.OverviewItem{
			Time:      item.TimeString,
			Info:      item.Info,
			ImageType: item.ImageType,
		}
		club2Overview = append(club2Overview, overviewItem)

	}
	if len(club1Overview) != 0 {
		respMatchDetail.MatchOverview.Club1Overview = club1Overview
	}
	if len(club2Overview) != 0 {
		respMatchDetail.MatchOverview.Club2Overview = club2Overview
	}

	// cast MatchStatistics
	stats := make([]presenter.StatisticsItem, 0)
	for _, item := range matchDetail.Statistics {
		stat := presenter.StatisticsItem{
			StatClub1:   item.StatClub1,
			StatContent: item.StatContent,
			StatClub2:   item.StatClub2,
		}
		stats = append(stats, stat)
	}
	if len(stats) != 0 {
		respMatchDetail.MatchStatistics.Statistics = stats
	}

	// cast MatchProgress
	events := make([]presenter.MatchEvent, 0)
	for _, item := range matchDetail.Events {
		event := presenter.MatchEvent{
			Time:    item.TimeString,
			Content: item.Content,
		}
		events = append(events, event)
	}
	if len(events) != 0 {
		respMatchDetail.MatchProgress.Events = events
	}

	// cast MatchLineup
	// sắp xếp theo thứ tự bé đến lớn. loop qua nếu cùng 1 hàng thì append vào. 
	sort.Slice(lineup1.Players, func(i, j int) bool {
		return lineup1.Players[i].Row < lineup1.Players[j].Row
	})

	rowPlayers := make([]presenter.PitchRows, 0)
	currentRow := 0
	var currentColumns []presenter.PitchRowsDetail

	for _, player := range lineup1.Players {
		if player.Row != currentRow {
			if len(currentColumns) > 0 {
				rowPlayers = append(rowPlayers, presenter.PitchRows{PitchRowsDetail: currentColumns})
			}
			currentRow += 1
			currentColumns = []presenter.PitchRowsDetail{}
		}

		currentColumns = append(currentColumns, presenter.PitchRowsDetail{PlayerName: player.PlayerName, PlayerNumber: player.PlayerNumber})
	}

	// Thêm row cuối cùng vào slice rows
	if len(currentColumns) > 0 {
		rowPlayers = append(rowPlayers, presenter.PitchRows{PitchRowsDetail: currentColumns})
	}

	lineupClub1 := presenter.MatchLineUpDetail{
		ClubName:   lineup1.ClubName,
		Formation:  lineup1.Formation,
		ShirtColor: lineup1.ShirtColor,
		PitchRows:  rowPlayers,
	}

	sort.Slice(lineup1.Players, func(i, j int) bool {
		return lineup1.Players[i].Row < lineup1.Players[j].Row
	})

	rowPlayers2 := make([]presenter.PitchRows, 0)
	currentRow = 0
	currentColumns = []presenter.PitchRowsDetail{}

	for _, player := range lineup2.Players {
		if player.Row != currentRow {
			if len(currentColumns) > 0 {
				rowPlayers2 = append(rowPlayers2, presenter.PitchRows{PitchRowsDetail: currentColumns})
			}
			currentRow += 1
			currentColumns = []presenter.PitchRowsDetail{}
		}

		currentColumns = append(currentColumns, presenter.PitchRowsDetail{PlayerName: player.PlayerName, PlayerNumber: player.PlayerNumber})
	}

	// Thêm row cuối cùng vào slice rows
	if len(currentColumns) > 0 {
		rowPlayers2 = append(rowPlayers2, presenter.PitchRows{PitchRowsDetail: currentColumns})
	}

	lineupClub2 := presenter.MatchLineUpDetail{
		ClubName:   lineup2.ClubName,
		Formation:  lineup2.Formation,
		ShirtColor: lineup2.ShirtColor,
		PitchRows:  rowPlayers2,
	}
	if len(lineupClub1.PitchRows) != 0 {
		respMatchDetail.MatchLineup.LineupClub1 = lineupClub1
	}
	if len(lineupClub2.PitchRows) != 0 {
		respMatchDetail.MatchLineup.LineupClub2 = lineupClub2
	}

	return respMatchDetail
}
