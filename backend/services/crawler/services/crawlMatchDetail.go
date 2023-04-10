package services

import (
	"crawler/entities"
	"crawler/helper"
	"fmt"
	"log"
	"strings"
	"sync"

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

func CrawlMatchDetail(matchUrl string) (entities.MatchDetail, error) {
	var matchDetail entities.MatchDetail
	xPath, err := crawlerhelpers.ReadXPathClassMatchDetailJSON()
	if err != nil {
		log.Println("can not read file htmlSchedulesClass.json, err: ", err)
		return matchDetail, err
	}

	doc, err := htmlquery.LoadURL(fmt.Sprintf(`https://bongda24h.vn%s`, matchUrl))
	if err != nil {
		log.Println("can not get match detail on url: ", matchUrl, " err: ", err)
		return matchDetail, err
	}

	// crawl with xPath
	crawlErr := make(chan error, 5)
	var wg sync.WaitGroup

	wg.Add(1)
	go func (wg *sync.WaitGroup)  {
		crawlErr <- findMatchTitle(doc, &matchDetail, xPath)
		wg.Done()
	}(&wg)
	
	wg.Add(1)
	go func (wg *sync.WaitGroup)  {
		crawlErr <- findMatchOverview(doc, &matchDetail, xPath)
		wg.Done()
	}(&wg)

	wg.Add(1)
	go func (wg *sync.WaitGroup)  {
		crawlErr <- findMatchStatistics(doc, &matchDetail, xPath)
		wg.Done()
	}(&wg)

	wg.Add(1)
	go func (wg *sync.WaitGroup)  {
		crawlErr <- findMatchProgress(doc, &matchDetail, xPath)
		wg.Done()
	}(&wg)

	wg.Add(1)
	go func (wg *sync.WaitGroup)  {
		crawlErr <- findMatchLineUp(doc, &matchDetail, xPath)
		wg.Done()
	}(&wg)

	// log crawl error
	done := make(chan bool)
	go func (done chan bool)  {
		for err := range crawlErr {
			if err != nil {
				log.Printf("error occurs while crawl: %s\n", <-crawlErr)
			}
		}
		done <- true
	}(done)

	wg.Wait()
	close(crawlErr)
	<- done

	return matchDetail, nil
}

func findMatchTitle(doc *html.Node, matchDetail *entities.MatchDetail, xPath entities.XPathMatchDetail) error {
	matchScoreNode, err := htmlquery.Query(doc, xPath.MatchDetailTitle.MatchScore)
	if err != nil {
		return err
	}
	matchDetail.MatchDetailTitle.MatchScore = htmlquery.InnerText(matchScoreNode)

	matchClub1NameNode, err := htmlquery.Query(doc, xPath.MatchDetailTitle.Club1.Name)
	if err != nil {
		return err
	}
	matchDetail.MatchDetailTitle.Club1.Name = htmlquery.InnerText(matchClub1NameNode)

	matchClub1LogoNode, err := htmlquery.Query(doc, xPath.MatchDetailTitle.Club1.Logo)
	if err != nil {
		return err
	}
	matchDetail.MatchDetailTitle.Club1.Logo = htmlquery.SelectAttr(matchClub1LogoNode, "src")

	matchClub2NameNode, err := htmlquery.Query(doc, xPath.MatchDetailTitle.Club2.Name)
	if err != nil {
		return err
	}
	matchDetail.MatchDetailTitle.Club2.Name = htmlquery.InnerText(matchClub2NameNode)

	matchClub2LogoNode, err := htmlquery.Query(doc, xPath.MatchDetailTitle.Club2.Logo)
	if err != nil {
		return err
	}
	matchDetail.MatchDetailTitle.Club2.Logo = htmlquery.SelectAttr(matchClub2LogoNode, "src")

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
		if err != nil {
			log.Printf("cant find club1 overview event time\n")
		}
		overView.Time = htmlquery.InnerText(timeNode)

		imgNode, err := htmlquery.Query(node, xPath.MatchOverview.Club1OverviewClass.Img)
		if err != nil {
			log.Printf("cant find club1 overview event img\n")
		}
		overView.ImageType = getOverviewEventImgType(imgNode)

		infoNode, err := htmlquery.Query(node, xPath.MatchOverview.Club1OverviewClass.Info)
		if err != nil {
			log.Printf("cant find club1 overview event img\n")
		}
		overView.Info = htmlquery.InnerText(infoNode)

		matchDetail.MatchOverview.Club1Overview = append(matchDetail.MatchOverview.Club1Overview, overView)
	}

	club2OverviewList, err := htmlquery.QueryAll(doc, xPath.MatchOverview.Club2OverviewClass.List)
	if err != nil {
		return fmt.Errorf("cant find club 2 overview nodes")
	}
	for _, node := range club2OverviewList {
		var overView entities.OverviewItem
		timeNode, err := htmlquery.Query(node, xPath.MatchOverview.Club2OverviewClass.Time)
		if err != nil {
			log.Printf("cant find club2 overview event time\n")
		}
		overView.Time = htmlquery.InnerText(timeNode)

		imgNode, err := htmlquery.Query(node, xPath.MatchOverview.Club2OverviewClass.Img)
		if err != nil {
			log.Printf("cant find club2 overview event img\n")
		}
		overView.ImageType = getOverviewEventImgType(imgNode)

		infoNode, err := htmlquery.Query(node, xPath.MatchOverview.Club2OverviewClass.Info)
		if err != nil {
			log.Printf("cant find club2 overview event img\n")
		}
		overView.Info = htmlquery.InnerText(infoNode)

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
		if err != nil {
			log.Printf("cant find stat club 1 node\n")
		}
		stats.StatClub1 = htmlquery.InnerText(statClub1Node)

		statClub2Node, err := htmlquery.Query(node, xPath.MatchStatistics.StatisticsItem.StatClub2)
		if err != nil {
			log.Printf("cant find stat club 2 node\n")
		}
		stats.StatClub2 = htmlquery.InnerText(statClub2Node)

		statContentNode, err := htmlquery.Query(node, xPath.MatchStatistics.StatisticsItem.StatContent)
		if err != nil {
			log.Printf("cant find stat club 1 node\n")
		}
		stats.StatContent = htmlquery.InnerText(statContentNode)

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
		if err != nil {
			log.Printf("cant find event time node\n")
		}
		event.Time = htmlquery.InnerText(timeNode)

		contentNode, err := htmlquery.Query(node, xPath.MatchProgress.EventContent)
		if err != nil {
			log.Printf("cant find event content node\n")
		}
		event.Content = htmlquery.InnerText(contentNode)

		matchDetail.MatchProgress.Events = append(matchDetail.MatchProgress.Events, event)
	}
	return nil
}

func findMatchLineUp(doc *html.Node, matchDetail *entities.MatchDetail, xPath entities.XPathMatchDetail) error {
	lineUpNode, err := htmlquery.Query(doc, xPath.MatchLineup.Lineup)
	if err != nil {
		return fmt.Errorf("cant find line up node")
	}
	// club name and formation
	club1NameNode, err := htmlquery.Query(lineUpNode, xPath.MatchLineup.Club1.ClubName)
	if err != nil {
		log.Printf("cant find club 1 name %s\n", err)
	}
	matchDetail.MatchLineup.LineupClub1.ClubName = htmlquery.InnerText(club1NameNode)

	club1FormationNode, err := htmlquery.Query(lineUpNode, xPath.MatchLineup.Club1.Formation)
	if err != nil {
		log.Printf("cant find club 1 formation %s\n", err)
	}
	matchDetail.MatchLineup.LineupClub1.Formation = htmlquery.InnerText(club1FormationNode)

	club2NameNode, err := htmlquery.Query(lineUpNode, xPath.MatchLineup.Club2.ClubName)
	if err != nil {
		log.Printf("cant find club 2 name %s\n", err)
	}
	matchDetail.MatchLineup.LineupClub2.ClubName = htmlquery.InnerText(club2NameNode)

	club2FormationNode, err := htmlquery.Query(lineUpNode, xPath.MatchLineup.Club2.Formation)
	if err != nil {
		log.Printf("cant find club 2 formation %s\n", err)
	}
	matchDetail.MatchLineup.LineupClub2.Formation = htmlquery.InnerText(club2FormationNode)

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
			if err != nil {
				log.Printf("cant find club 1 pitch row player name %s\n", err)
			}
			pitchRowDetail.PlayerName = htmlquery.InnerText(playerNameNode)

			playerNumberNode, err := htmlquery.Query(playerNode, xPath.MatchLineup.Club1.PitchRows.ListPlayer.PlayerNumber)
			if err != nil {
				log.Printf("cant find club 1 pitch row player number %s\n", err)
			}
			pitchRowDetail.PlayerNumber = htmlquery.InnerText(playerNumberNode)

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
			if err != nil {
				log.Printf("cant find club 2 pitch row player name %s\n", err)
			}
			pitchRowDetail.PlayerName = htmlquery.InnerText(playerNameNode)

			playerNumberNode, err := htmlquery.Query(playerNode, xPath.MatchLineup.Club2.PitchRows.ListPlayer.PlayerNumber)
			if err != nil {
				log.Printf("cant find club 2 pitch row player number %s\n", err)
			}
			pitchRowDetail.PlayerNumber = htmlquery.InnerText(playerNumberNode)

			pitchRow.PitchRowsDetail = append(pitchRow.PitchRowsDetail, pitchRowDetail)
		}

		matchDetail.MatchLineup.LineupClub2.PitchRows = append(matchDetail.MatchLineup.LineupClub1.PitchRows, pitchRow)
	}
	return nil
}
