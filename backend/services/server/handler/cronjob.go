package handler

import (
	"net/http"
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
	Cronjobs, err := h.handler.CronjobOnHour(time)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Get Cronjob failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "Cronjobs": Cronjobs})
}

func (h *CronjobHandler) CreateCronjobGetArticleFromGoogle() {
	h.handler.CreateCronjobGetArticleFromGoogle()
}

func (h *CronjobHandler) CreateCronjobRefreshCache() {
	h.handler.CreateCronjobRefreshCache()
}
