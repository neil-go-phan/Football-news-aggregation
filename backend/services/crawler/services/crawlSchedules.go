package services

import (
	"crawler/entities"
	"crawler/helpers"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func CrawlSchedules(date string) (entities.ScheduleOnDay, error) {
	// date format: dd-mm-yyyy
	var schedules entities.ScheduleOnDay
	htmlClasses, err := helpers.ReadHtmlClassScheduleJSON()
	if err != nil {
		log.Println("can not read file htmlSchedulesClass.json, err: ", err)
		return schedules, err
	}

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

	doc.Find(helpers.FormatClassName(htmlClasses.LeagueClass)).Each(func(i int, s *goquery.Selection) {
		var schedule entities.ScheduleOnLeague

		schedule.LeagueName = s.Find("h3").Text()

		// Tìm các match thuộc về league hiện tại
		matches := s.NextUntil(helpers.FormatClassName(htmlClasses.LeagueClass))
		matches.Each(func(j int, m *goquery.Selection) {
			var match entities.Match
			match.Time = m.Find(helpers.FormatClassName(htmlClasses.HtmlMatchClass.Time)).Text()
			match.Round = m.Find(helpers.FormatClassName(htmlClasses.HtmlMatchClass.Round)).Text()
			match.Time = strings.TrimSpace(strings.Replace(match.Time, match.Round, "", -1))
			match.Scores = m.Find(helpers.FormatClassName(htmlClasses.HtmlMatchClass.Scores)).Text()
			m.Find(fmt.Sprintf("%s a", helpers.FormatClassName(htmlClasses.HtmlMatchClass.MatchDetailLink))).Each(func(i int, s *goquery.Selection) {
				href, _ := s.Attr("href")
				match.MatchDetailLink = href
			})
			match.Club1.Name = m.Find(helpers.FormatClassName(htmlClasses.HtmlMatchClass.Club1.Name)).Text()
			match.Club2.Name = m.Find(helpers.FormatClassName(htmlClasses.HtmlMatchClass.Club2.Name)).Text()
			m.Find(helpers.FormatClassName(htmlClasses.HtmlMatchClass.Club2.Name)).Children()
			m.Find(helpers.FormatClassName(htmlClasses.HtmlMatchClass.Club1.Name)).Each(func(i int, s *goquery.Selection) {
				attr, _ := s.Find(helpers.FormatClassName(htmlClasses.HtmlMatchClass.Club2.Logo)).Attr("src")
				match.Club1.Logo = attr
			})
			m.Find(helpers.FormatClassName(htmlClasses.HtmlMatchClass.Club2.Name)).Each(func(i int, s *goquery.Selection) {
				attr, _ := s.Find(helpers.FormatClassName(htmlClasses.HtmlMatchClass.Club2.Logo)).Attr("src")
				match.Club2.Logo = attr
			})
			schedule.Matchs = append(schedule.Matchs, match)
		})

		schedules.ScheduleOnLeagues = append(schedules.ScheduleOnLeagues, schedule)
	})

	return schedules, nil
}
