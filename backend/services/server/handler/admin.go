package handler

import (
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
	c.BindJSON(&inputAdmin)

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
	c.BindJSON(&inputAdmin)


	err := adminHandler.handler.CheckAdminUsernameToken(username.(string))
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

