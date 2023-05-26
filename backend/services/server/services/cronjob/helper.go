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

// func newChartHourResponse(entityChart entities.Cronjob) services.ChartHourResponse {
// 	cronjobName := make([]string, 0)
// 	for _, cronjob := range entityChart.Cronjobs {
// 		cronjobName = append(cronjobName, cronjob.Name)
// 	}
// 	return services.ChartHourResponse{
// 		Time: entityChart.Time,
// 		AmountOfJob: entityChart.AmountOfJob,
// 		MemoryUsage: entityChart.MemoryUsage,
// 		CronjobNames: cronjobName,
// 	}
// }