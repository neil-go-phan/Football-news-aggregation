package services

import (
	"crawler/entities"
	"net/http"
	"time"

	"fmt"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

var GOAL_EVENT = []string{
	"goal",
	"ghi ban",
	"ghi-ban",
	"ghi_ban",
	"ghiban",
	"ghiBan",
	"GhiBan",
}

var YELLOW_CARD_EVENT = []string{
	"yellow card",
	"yellow-card",
	"yellow_card",
	"yellowCard",
	"YellowCard",
	"yellowcard",
	"the vang",
	"the-vang",
	"the-_vang",
	"thevang",
	"theVang",
	"TheVang",
}

var RED_CARD_EVENT = []string{
	"red card",
	"red-card",
	"red_card",
	"redCard",
	"RedCard",
	"redcard",
	"the do",
	"the-do",
	"the_do",
	"thedo",
	"theDo",
	"TheDo",
}

var SUBSTITUTION_EVENT = []string{
	"substitution",
	"ThayNguoi",
	"thayNguoi",
	"thaynguoi",
	"thay-nguoi",
	"thay_nguoi",
}
var TIME_OUT = 5 * time.Second

func CrawlMatchDetail(matchUrl string, xPath entities.XPathMatchDetail) (entities.MatchDetail, error) {
	var matchDetail entities.MatchDetail

	client := http.Client{
		Timeout: TIME_OUT,
	}
	req, err := http.NewRequest("GET", matchUrl, nil)
	if err != nil {
		log.Errorln("can not create when crawl match HTTP:", err)
		return matchDetail, err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Errorln("can not do http request:", err)
		return matchDetail, err
	}
	defer resp.Body.Close()

	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		log.Errorln("can not get match detail on url: ", matchUrl, " err: ", err)
		return matchDetail, err
	}

	// // crawl with xPath
	crawlErr := make(chan error, 5)
	var wg sync.WaitGroup

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		crawlErr <- findMatchTitle(doc, &matchDetail, xPath)
		wg.Done()
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		crawlErr <- findMatchOverview(doc, &matchDetail, xPath)
		wg.Done()
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		crawlErr <- findMatchStatistics(doc, &matchDetail, xPath)
		wg.Done()
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		crawlErr <- findMatchProgress(doc, &matchDetail, xPath)
		wg.Done()
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		crawlErr <- findMatchLineUp(doc, &matchDetail, xPath)
		wg.Done()
	}(&wg)

	// log crawl error
	done := make(chan bool)
	go func(done chan bool) {
		for err := range crawlErr {
			if err != nil {
				log.Printf("error occurs while crawl: %s\n", <-crawlErr)
			}
		}
		done <- true
	}(done)

	wg.Wait()
	close(crawlErr)
	<-done
	return matchDetail, nil
}

func findMatchTitle(doc *html.Node, matchDetail *entities.MatchDetail, xPath entities.XPathMatchDetail) error {
	matchScoreNode, err := htmlquery.Query(doc, xPath.MatchDetailTitle.MatchScore)
	if err == nil && matchScoreNode != nil {
		matchDetail.MatchDetailTitle.MatchScore = htmlquery.InnerText(matchScoreNode)
	}
	matchStatusNode, err := htmlquery.Query(doc, xPath.MatchDetailTitle.MatchStatus)
	if err == nil && matchStatusNode != nil {
		matchDetail.MatchDetailTitle.MatchStatus = htmlquery.InnerText(matchStatusNode)
	}

	matchClub1NameNode, err := htmlquery.Query(doc, xPath.MatchDetailTitle.Club1.Name)
	if err == nil && matchClub1NameNode != nil {
		matchDetail.MatchDetailTitle.Club1.Name = htmlquery.InnerText(matchClub1NameNode)
	}

	matchClub1LogoNode, err := htmlquery.Query(doc, xPath.MatchDetailTitle.Club1.Logo)
	if err == nil && matchClub1LogoNode != nil {
		matchDetail.MatchDetailTitle.Club1.Logo = htmlquery.SelectAttr(matchClub1LogoNode, "src")
	}

	matchClub2NameNode, err := htmlquery.Query(doc, xPath.MatchDetailTitle.Club2.Name)
	if err == nil && matchClub2NameNode != nil {
		matchDetail.MatchDetailTitle.Club2.Name = htmlquery.InnerText(matchClub2NameNode)
	}

	matchClub2LogoNode, err := htmlquery.Query(doc, xPath.MatchDetailTitle.Club2.Logo)
	if err == nil && matchClub2LogoNode != nil {
		matchDetail.MatchDetailTitle.Club2.Logo = htmlquery.SelectAttr(matchClub2LogoNode, "src")
	}

	return nil
}

func findMatchOverview(doc *html.Node, matchDetail *entities.MatchDetail, xPath entities.XPathMatchDetail) error {
	club1OverviewList, err := htmlquery.QueryAll(doc, xPath.MatchOverview.Club1OverviewClass.List)
	if err != nil {
		return fmt.Errorf("cant find club 1 overview nodes")
	}
	for _, node := range club1OverviewList {
		var overView entities.OverviewItem

		timeNode, err := htmlquery.Query(node, xPath.MatchOverview.Club1OverviewClass.Time)

		if err == nil && timeNode != nil {
			overView.Time = htmlquery.InnerText(timeNode)
		} else {
			log.Printf("cant find club1 overview event time\n")
		}

		imgNode, err := htmlquery.Query(node, xPath.MatchOverview.Club1OverviewClass.Img)

		if err == nil && imgNode != nil {
			overView.ImageType = getOverviewEventImgType(imgNode)
		} else {
			log.Printf("cant find club1 overview event img\n")
		}

		infoNode, err := htmlquery.Query(node, xPath.MatchOverview.Club1OverviewClass.Info)

		if err == nil && infoNode != nil {
			overView.Info = htmlquery.InnerText(infoNode)
		} else {
			log.Printf("cant find club1 overview event img\n")
		}

		matchDetail.MatchOverview.Club1Overview = append(matchDetail.MatchOverview.Club1Overview, overView)
	}

	club2OverviewList, err := htmlquery.QueryAll(doc, xPath.MatchOverview.Club2OverviewClass.List)
	if err != nil {
		return fmt.Errorf("cant find club 2 overview nodes")
	}
	for _, node := range club2OverviewList {
		var overView entities.OverviewItem

		timeNode, err := htmlquery.Query(node, xPath.MatchOverview.Club2OverviewClass.Time)

		if err == nil && timeNode != nil {
			overView.Time = htmlquery.InnerText(timeNode)
		} else {
			log.Printf("cant find club2 overview event time\n")
		}

		imgNode, err := htmlquery.Query(node, xPath.MatchOverview.Club2OverviewClass.Img)

		if err == nil && imgNode != nil {
			overView.ImageType = getOverviewEventImgType(imgNode)
		} else {
			log.Printf("cant find club2 overview event img\n")
		}

		infoNode, err := htmlquery.Query(node, xPath.MatchOverview.Club2OverviewClass.Info)

		if err == nil && infoNode != nil {
			overView.Info = htmlquery.InnerText(infoNode)
		} else {
			log.Printf("cant find club1 overview event img\n")
		}

		matchDetail.MatchOverview.Club2Overview = append(matchDetail.MatchOverview.Club2Overview, overView)
	}
	return nil
}

func getOverviewEventImgType(node *html.Node) string {
	nodeHtml := htmlquery.OutputHTML(node, true)
	// is goal
	if stringInSlice(nodeHtml, GOAL_EVENT) {
		return "goal"
	}
	if stringInSlice(nodeHtml, YELLOW_CARD_EVENT) {
		return "yellow-card"
	}
	if stringInSlice(nodeHtml, RED_CARD_EVENT) {
		return "red-card"
	}
	if stringInSlice(nodeHtml, SUBSTITUTION_EVENT) {
		return "substitution"
	}
	return nodeHtml
}

func stringInSlice(nodeHtml string, list []string) bool {
	for _, subString := range list {
		if strings.Contains(strings.ToLower(nodeHtml), subString) {
			return true
		}
	}
	return false
}

func findMatchStatistics(doc *html.Node, matchDetail *entities.MatchDetail, xPath entities.XPathMatchDetail) error {
	statsList, err := htmlquery.QueryAll(doc, xPath.MatchStatistics.MatchStatisticsListItem)
	if err != nil {
		return fmt.Errorf("cant find match statistics list nodes")
	}

	for _, node := range statsList {
		var stats entities.StatisticsItem

		statClub1Node, err := htmlquery.Query(node, xPath.MatchStatistics.StatisticsItem.StatClub1)

		if err == nil && statClub1Node != nil {
			stats.StatClub1 = htmlquery.InnerText(statClub1Node)
		} else {
			log.Printf("cant find stat club 1 node\n")
		}

		statClub2Node, err := htmlquery.Query(node, xPath.MatchStatistics.StatisticsItem.StatClub2)

		if err == nil && statClub2Node != nil {
			stats.StatClub2 = htmlquery.InnerText(statClub2Node)
		} else {
			log.Printf("cant find stat club 2 node\n")
		}

		statContentNode, err := htmlquery.Query(node, xPath.MatchStatistics.StatisticsItem.StatContent)

		if err == nil && statContentNode != nil {
			stats.StatContent = htmlquery.InnerText(statContentNode)
		} else {
			log.Printf("cant find stat content node\n")
		}

		matchDetail.MatchStatistics.Statistics = append(matchDetail.MatchStatistics.Statistics, stats)
	}
	return nil
}

func findMatchProgress(doc *html.Node, matchDetail *entities.MatchDetail, xPath entities.XPathMatchDetail) error {
	eventList, err := htmlquery.QueryAll(doc, xPath.MatchProgress.Events)
	if err != nil {
		return fmt.Errorf("cant find match process event list nodes")
	}
	for _, node := range eventList {
		var event entities.MatchEvent

		timeNode, err := htmlquery.Query(node, xPath.MatchProgress.EventTime)

		if err == nil && timeNode != nil {
			event.Time = htmlquery.InnerText(timeNode)
		} else {
			log.Printf("cant find event time node\n")
		}

		contentNode, err := htmlquery.Query(node, xPath.MatchProgress.EventContent)

		if err == nil && contentNode != nil {
			event.Content = htmlquery.InnerText(contentNode)
		} else {
			log.Printf("cant find event content node\n")
		}

		matchDetail.MatchProgress.Events = append(matchDetail.MatchProgress.Events, event)
	}
	return nil
}

func findMatchLineUp(doc *html.Node, matchDetail *entities.MatchDetail, xPath entities.XPathMatchDetail) error {
	lineUpNode, err := htmlquery.Query(doc, xPath.MatchLineup.Lineup)
	if err != nil || lineUpNode == nil {
		return fmt.Errorf("cant find line up node")
	}

	// club name and formation
	club1NameNode, err := htmlquery.Query(lineUpNode, xPath.MatchLineup.Club1.ClubName)

	if err == nil && club1NameNode != nil {
		matchDetail.MatchLineup.LineupClub1.ClubName = htmlquery.InnerText(club1NameNode)
	} else {
		log.Printf("cant find club 1 name %s\n", err)
	}

	club1FormationNode, err := htmlquery.Query(lineUpNode, xPath.MatchLineup.Club1.Formation)

	if err == nil && club1FormationNode != nil {
		matchDetail.MatchLineup.LineupClub1.Formation = htmlquery.InnerText(club1FormationNode)
	} else {
		log.Printf("cant find club 1 formation %s\n", err)
	}

	club1ShirtColorNode, err := htmlquery.Query(lineUpNode, xPath.MatchLineup.Club1.ShirtColor)

	if err == nil && club1ShirtColorNode != nil {
		styleAttr := htmlquery.SelectAttr(club1ShirtColorNode, "style")
		_, shirtColor1, _ := strings.Cut(styleAttr, "fill:")
		matchDetail.MatchLineup.LineupClub1.ShirtColor = shirtColor1
	} else {
		log.Printf("cant find club 1 shirt color %s\n", err)
	}

	club2NameNode, err := htmlquery.Query(lineUpNode, xPath.MatchLineup.Club2.ClubName)

	if err == nil && club2NameNode != nil {
		matchDetail.MatchLineup.LineupClub2.ClubName = htmlquery.InnerText(club2NameNode)
	} else {
		log.Printf("cant find club 2 name %s\n", err)
	}

	club2FormationNode, err := htmlquery.Query(lineUpNode, xPath.MatchLineup.Club2.Formation)

	if err == nil && club2FormationNode != nil {
		matchDetail.MatchLineup.LineupClub2.Formation = htmlquery.InnerText(club2FormationNode)
	} else {
		log.Printf("cant find club 2 formation %s\n", err)
	}

	club2ShirtColorNode, err := htmlquery.Query(lineUpNode, xPath.MatchLineup.Club2.ShirtColor)

	if err == nil && club2ShirtColorNode != nil {
		styleAttr := htmlquery.SelectAttr(club2ShirtColorNode, "style")
		_, shirtColor2, _ := strings.Cut(styleAttr, "fill:")
		matchDetail.MatchLineup.LineupClub2.ShirtColor = shirtColor2
	} else {
		log.Printf("cant find club 2 shirt color %s\n", err)
	}

	//club1 pitch row
	club1PitchRowNodes, err := htmlquery.QueryAll(lineUpNode, xPath.MatchLineup.Club1.PitchRows.List)
	if err != nil {
		return fmt.Errorf("cant find club 1 pitch row nodes %s", err)
	}
	for _, node := range club1PitchRowNodes {
		var pitchRow entities.PitchRows

		playerList, err := htmlquery.QueryAll(node, xPath.MatchLineup.Club1.PitchRows.ListPlayer.List)

		if err != nil {
			log.Printf("cant find club 1 pitch row player list %s\n", err)
			continue
		}
		for _, playerNode := range playerList {
			var pitchRowDetail entities.PitchRowsDetail

			playerNameNode, err := htmlquery.Query(playerNode, xPath.MatchLineup.Club1.PitchRows.ListPlayer.PlayerName)

			if err == nil && playerNameNode != nil {
				pitchRowDetail.PlayerName = htmlquery.InnerText(playerNameNode)
			} else {
				log.Printf("cant find club 1 pitch row player name %s\n", err)
			}

			playerNumberNode, err := htmlquery.Query(playerNode, xPath.MatchLineup.Club1.PitchRows.ListPlayer.PlayerNumber)

			if err == nil && playerNumberNode != nil {
				pitchRowDetail.PlayerNumber = htmlquery.InnerText(playerNumberNode)
			} else {
				log.Printf("cant find club 1 pitch row player number %s\n", err)
			}

			pitchRow.PitchRowsDetail = append(pitchRow.PitchRowsDetail, pitchRowDetail)
		}

		matchDetail.MatchLineup.LineupClub1.PitchRows = append(matchDetail.MatchLineup.LineupClub1.PitchRows, pitchRow)
	}

	//club2 pitch row
	club2PitchRowNodes, err := htmlquery.QueryAll(lineUpNode, xPath.MatchLineup.Club2.PitchRows.List)
	if err != nil {
		return fmt.Errorf("cant find club 1 pitch row nodes %s", err)
	}
	for _, node := range club2PitchRowNodes {
		var pitchRow entities.PitchRows
		playerList, err := htmlquery.QueryAll(node, xPath.MatchLineup.Club2.PitchRows.ListPlayer.List)
		if err != nil {
			log.Printf("cant find club 2 pitch row player list %s\n", err)
			continue
		}
		for _, playerNode := range playerList {
			var pitchRowDetail entities.PitchRowsDetail

			playerNameNode, err := htmlquery.Query(playerNode, xPath.MatchLineup.Club2.PitchRows.ListPlayer.PlayerName)

			if err == nil && playerNameNode != nil {
				pitchRowDetail.PlayerName = htmlquery.InnerText(playerNameNode)
			} else {
				log.Printf("cant find club 2 pitch row player name %s\n", err)
			}

			playerNumberNode, err := htmlquery.Query(playerNode, xPath.MatchLineup.Club2.PitchRows.ListPlayer.PlayerNumber)

			if err == nil && playerNumberNode != nil {
				pitchRowDetail.PlayerNumber = htmlquery.InnerText(playerNumberNode)
			} else {
				log.Printf("cant find club 2 pitch row player number %s\n", err)
			}

			pitchRow.PitchRowsDetail = append(pitchRow.PitchRowsDetail, pitchRowDetail)
		}

		matchDetail.MatchLineup.LineupClub2.PitchRows = append(matchDetail.MatchLineup.LineupClub2.PitchRows, pitchRow)
	}
	return nil
}
