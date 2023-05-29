package cronjob

import (
	"fmt"
	"server/entities"
	pb "server/proto"
	"server/services"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func newPbConfigCrawler(configCrawler *entities.Crawler) *pb.ConfigCrawler {
	return &pb.ConfigCrawler{
		Url:          configCrawler.Url,
		Div:          configCrawler.ArticleDiv,
		Title:        configCrawler.ArticleTitle,
		Description:  configCrawler.ArticleDescription,
		Link:         configCrawler.ArticleLink,
		NextPage:     configCrawler.NextPage,
		NetxPageType: configCrawler.NetxPageType,
	}
}

func readRunEveryMin(runEveryMin int) (string) {
	hours := runEveryMin / 60
	mins := runEveryMin % 60
	return fmt.Sprintf("@every %dh%02dm", hours, mins)
} 

func newCronjobResponse(mapKey string) services.CronjobResponse {
	if !strings.Contains(mapKey, "$") {
		return services.CronjobResponse {
			Name: mapKey,
			Url: "none",
			RunEveryMin: 1,
		}
	}
	part := strings.Split(mapKey, "$")
	cronjobResponse := services.CronjobResponse{
		Name: part[0],
		Url: part[1],
	}
	runEveryMin, err := strconv.Atoi(part[2])
	if err != nil {
		log.Error("can not get cronjob schedule")
		return cronjobResponse
	}
	cronjobResponse.RunEveryMin = runEveryMin
	return cronjobResponse
}

func newMapKey(url string, runEveryMin int) string {
	cronjobName := fmt.Sprintf("Crawl url :%s, every %v min", url, runEveryMin)
	return fmt.Sprintf("%s$%s$%v",cronjobName, url, runEveryMin)
}

func newCronjobInChart(cronjob entities.Cronjob) services.CronjobInChart {
	hour := cronjob.StartAt.Hour()
	min := cronjob.StartAt.Minute()
	startAt := fmt.Sprintf("%v:%v", addZeroWhenLowwerThanTen(hour), addZeroWhenLowwerThanTen(min))
	var endAt string
	if (cronjob.EndAt.Year() != 1) {
		hour = cronjob.EndAt.Hour()
		min = cronjob.EndAt.Minute()
		endAt = fmt.Sprintf("%v:%v", addZeroWhenLowwerThanTen(hour), addZeroWhenLowwerThanTen(min))
	} else {
		endAt = "runing"
	}

	return services.CronjobInChart{
		Name: cronjob.Name,
		StartAt: startAt,
		EndAt: endAt,
	}
} 

func addZeroWhenLowwerThanTen(time int) string {
	if time < 10 {
		return fmt.Sprintf("0%v", time)
	}
	return fmt.Sprintf("%v", time)
}