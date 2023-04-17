package handler

import (
	"log"
	"net/http"
	"server/services"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	handler services.AdminServices
}

func NewAdminHandler(handler services.AdminServices) *AdminHandler {
	adminHandler := &AdminHandler{
		handler: handler,
	}
	return adminHandler;
}

func (adminHandler *AdminHandler) CheckAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Granted permission"})
}

func (adminHandler *AdminHandler) Get(c *gin.Context) {
	username, ok := c.Get("username")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Unauthorized access"})
		return
	}

	adminUsername, err := adminHandler.handler.GetAdminUsername(username.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Unauthorized access"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "username": adminUsername})
}

func (adminHandler *AdminHandler) Login(c *gin.Context) {
	var inputAdmin services.Admin
	err := c.BindJSON(&inputAdmin)
	if err != nil {
		log.Printf("error occrus: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "internal server error"})
		return
	}
	token, err := adminHandler.handler.Login(&inputAdmin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "token": token})
}

func (adminHandler *AdminHandler) ChangePassword(c *gin.Context) {
	username, ok := c.Get("username")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Unauthorized access"})
		return
	}

	var inputAdmin services.AdminWithConfirmPassword
	err := c.BindJSON(&inputAdmin)
	if err != nil {
		log.Printf("error occrus: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "internal server error"})
		return
	}

	err = adminHandler.handler.CheckAdminUsernameToken(username.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Unauthorized access"})
		return
	}

	err = adminHandler.handler.ChangePassword(&inputAdmin, username.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Change password success"})
}

