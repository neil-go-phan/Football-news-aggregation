package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	serverhelper "server/helper"
	"server/services"
	adminservices "server/services/admin"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	handler services.AdminServices
}

func NewAdminHandler(handler services.AdminServices) *AdminHandler {
	adminHandler := &AdminHandler{
		handler: handler,
	}
	return adminHandler
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

func (adminHandler *AdminHandler) GoogleOAuth(c *gin.Context) {
	code := c.Query("code")
	var pathUrl string = "/"

	if c.Query("state") != "" {
		pathUrl = c.Query("state")
	}

	if code == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Authorization code not provided!"})
		return
	}

	tokenRes, err := GetGoogleOauthToken(code)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"success": false, "message": err.Error()})
		return
	}

	google_user, err := GetGoogleUser(tokenRes.Access_token, tokenRes.Id_token)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"success": false, "message": err.Error()})
		return
	}

	config, _ := serverhelper.LoadEnv(".")

	fmt.Println("google_user", google_user)

	token, err := adminHandler.handler.GoogleOAuth(google_user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.SetCookie("token", token, 24*60*60, "/", c.Request.Header.Get("Origin"), false, false)
	c.Redirect(http.StatusTemporaryRedirect, fmt.Sprint(config.FrontEndOrigin, pathUrl))
}

func (adminHandler *AdminHandler) Login(c *gin.Context) {
	var inputAdmin adminservices.Admin
	err := c.BindJSON(&inputAdmin)
	if err != nil {
		log.Printf("error occrus: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "internal server error"})
		return
	}
	var token string
	if !checkIsEmail(inputAdmin.Username) {
		log.Println("login with username")
		token, err = adminHandler.handler.LoginWithUsername(&inputAdmin)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
			return
		}
	} else {
		log.Println("login with email")
		token, err = adminHandler.handler.LoginWithEmail(&inputAdmin)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "token": token})
}

func checkIsEmail(checkedString string) bool {
	return strings.Contains(checkedString, "@")
}

func (adminHandler *AdminHandler) ChangePassword(c *gin.Context) {
	username, ok := c.Get("username")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Unauthorized access"})
		return
	}

	var inputAdmin adminservices.AdminWithConfirmPassword
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

func GetGoogleOauthToken(code string) (*adminservices.GoogleOauthToken, error) {
	const rootURl = "https://oauth2.googleapis.com/token"

	config, _ := serverhelper.LoadEnv(".")
	values := url.Values{}
	values.Add("grant_type", "authorization_code")
	values.Add("code", code)
	values.Add("client_id", config.GoogleClientID)
	values.Add("client_secret", config.GoogleClientSecret)
	values.Add("redirect_uri", config.GoogleOAuthRedirectUrl)

	query := values.Encode()

	req, err := http.NewRequest("POST", rootURl, bytes.NewBufferString(query))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve token")
	}

	var resBody bytes.Buffer
	_, err = io.Copy(&resBody, res.Body)
	if err != nil {
		return nil, err
	}

	var GoogleOauthTokenRes map[string]interface{}

	if err := json.Unmarshal(resBody.Bytes(), &GoogleOauthTokenRes); err != nil {
		return nil, err
	}

	tokenBody := &adminservices.GoogleOauthToken{
		Access_token: GoogleOauthTokenRes["access_token"].(string),
		Id_token:     GoogleOauthTokenRes["id_token"].(string),
	}

	return tokenBody, nil
}

func GetGoogleUser(access_token string, id_token string) (*adminservices.GoogleUserResult, error) {
	rootUrl := fmt.Sprintf("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=%s", access_token)

	req, err := http.NewRequest("GET", rootUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", id_token))

	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve user")
	}

	var resBody bytes.Buffer
	_, err = io.Copy(&resBody, res.Body)
	if err != nil {
		return nil, err
	}

	var GoogleUserRes map[string]interface{}

	if err := json.Unmarshal(resBody.Bytes(), &GoogleUserRes); err != nil {
		return nil, err
	}

	userBody := &adminservices.GoogleUserResult{
		Id:             GoogleUserRes["id"].(string),
		Email:          GoogleUserRes["email"].(string),
		Verified_email: GoogleUserRes["verified_email"].(bool),
		Name:           GoogleUserRes["name"].(string),
		Given_name:     GoogleUserRes["given_name"].(string),
		Picture:        GoogleUserRes["picture"].(string),
		Locale:         GoogleUserRes["locale"].(string),
	}

	return userBody, nil
}

func (adminHandler *AdminHandler) Register(c *gin.Context) {
	var inputUser adminservices.RegisterUserInput
	err := c.BindJSON(&inputUser)
	if err != nil {
		log.Printf("error occrus: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "input invalid"})
		return
	}

	err = adminHandler.handler.Register(&inputUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "register success"})
}
