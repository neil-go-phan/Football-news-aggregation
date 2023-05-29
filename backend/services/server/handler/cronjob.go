package handler

import (
	"net/http"
	"server/handler/presenter"
	"server/services"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type CronjobHandler struct {
	handler services.CronjobServices
}

func NewCronjobHandler(handler services.CronjobServices) *CronjobHandler {
	CronjobHandler := &CronjobHandler{
		handler: handler,
	}
	return CronjobHandler
}

func (h *CronjobHandler) APIListCrawlerCronjob(c *gin.Context) {
	cronjobs := h.handler.GetCronjobRuntime()
	c.JSON(http.StatusOK, gin.H{"success": true, "cronjobs": cronjobs})
}

func (h *CronjobHandler) APIGetCronjobOnHour(c *gin.Context) {
	time := c.Query("time")
	cronjobs, err := h.handler.CronjobOnHour(time)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Get cronjob failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "cronjobs": cronjobs})
}

func (h *CronjobHandler) APIGetCronjobOnDay(c *gin.Context) {
	time := c.Query("time")
	cronjobs, err := h.handler.CronjobOnDay(time)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Get Cronjob failed"})
		return
	}

	reponse := parseChartDayToResponse(cronjobs)

	c.JSON(http.StatusOK, gin.H{"success": true, "cronjobs": reponse})
}

func parseChartDayToResponse(cronjobs *[24]services.ChartDay) [24]presenter.ChartDayResponse{
	response := [24]presenter.ChartDayResponse{}
	for index, cronjob := range cronjobs {
		response[index].Hour = cronjob.Hour
		response[index].AmountOfJob = cronjob.AmountOfJob
		cronjobResponese := make([]presenter.CronjobRunTimes, 0)
		for key, value := range cronjob.Cronjobs {
			cronjobResponese = append(cronjobResponese, presenter.CronjobRunTimes{
				Name: key,
				Times: value,
			})
		}
		response[index].Cronjobs = cronjobResponese
	}
	return response
}

func (h *CronjobHandler) CreateCronjobGetArticleFromGoogle() {
	h.handler.CreateCronjobGetArticleFromGoogle()
}

func (h *CronjobHandler) CreateCronjobRefreshCache() {
	h.handler.CreateCronjobRefreshCache()
}
