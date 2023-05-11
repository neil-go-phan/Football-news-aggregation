package services

import (
	"crawler/entities"
	crawlerhelpers "crawler/helper"
	"fmt"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/PuerkitoBio/goquery"
)

func CrawlSchedules(date string, htmlClasses entities.HtmlSchedulesClass) (entities.ScheduleOnDay, error) {
	// date format: dd-mm-yyyy
	var schedules entities.ScheduleOnDay

	req, err := http.NewRequest("GET", fmt.Sprintf(`https://bongda24h.vn/LiveScore/AjaxLivescore?date=%s&leagueId=0`, date), nil)
	if err != nil {
		log.Println("can not get schedule on day: ", date, " err: ", err)
		return schedules, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("can not set header, err:", err)
		return schedules, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println("can not create goquery document, err:", err)
		return schedules, err
	}

	schedules.Date = date

	doc.Find(crawlerhelpers.FormatClassName(htmlClasses.LeagueClass)).Each(func(i int, s *goquery.Selection) {
		var schedule entities.ScheduleOnLeague

		schedule.LeagueName = s.Find("h3").Text()

		// Tìm các match thuộc về league hiện tại
		matches := s.NextUntil(crawlerhelpers.FormatClassName(htmlClasses.LeagueClass))
		matches.Each(func(j int, m *goquery.Selection) {
			var match entities.Match
			match.Time = m.Find(crawlerhelpers.FormatClassName(htmlClasses.HtmlMatchClass.Time)).Text()
			match.Round = m.Find(crawlerhelpers.FormatClassName(htmlClasses.HtmlMatchClass.Round)).Text()
			match.Round = strings.TrimSpace(match.Round)
			match.Time = strings.TrimSpace(strings.Replace(match.Time, match.Round, "", -1))
			match.Scores = m.Find(crawlerhelpers.FormatClassName(htmlClasses.HtmlMatchClass.Scores)).Text()
			match.Scores = strings.TrimSpace(match.Scores)
			m.Find(fmt.Sprintf("%s a", crawlerhelpers.FormatClassName(htmlClasses.HtmlMatchClass.MatchDetailLink))).Each(func(i int, s *goquery.Selection) {
				href, _ := s.Attr("href")
				match.MatchDetailLink = href
			})
			match.Club1.Name = m.Find(crawlerhelpers.FormatClassName(htmlClasses.HtmlMatchClass.Club1.Name)).First().Clone().Children().Remove().End().Text()
			match.Club1.Name = strings.TrimSpace(match.Club1.Name)
			match.Club2.Name = m.Find(crawlerhelpers.FormatClassName(htmlClasses.HtmlMatchClass.Club2.Name)).First().Clone().Children().Remove().End().Text()
			match.Club2.Name = strings.TrimSpace(match.Club2.Name)
			m.Find(crawlerhelpers.FormatClassName(htmlClasses.HtmlMatchClass.Club2.Name)).Children()
			m.Find(crawlerhelpers.FormatClassName(htmlClasses.HtmlMatchClass.Club1.Name)).Each(func(i int, s *goquery.Selection) {
				attr, _ := s.Find(crawlerhelpers.FormatClassName(htmlClasses.HtmlMatchClass.Club2.Logo)).Attr("src")
				match.Club1.Logo = attr
			})
			m.Find(crawlerhelpers.FormatClassName(htmlClasses.HtmlMatchClass.Club2.Name)).Each(func(i int, s *goquery.Selection) {
				attr, _ := s.Find(crawlerhelpers.FormatClassName(htmlClasses.HtmlMatchClass.Club2.Logo)).Attr("src")
				match.Club2.Logo = attr
			})
			schedule.Matches = append(schedule.Matches, match)
		})

		schedules.ScheduleOnLeagues = append(schedules.ScheduleOnLeagues, schedule)
	})

	return schedules, nil
}
