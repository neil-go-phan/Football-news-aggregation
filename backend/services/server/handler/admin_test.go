package handler

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	mock "server/handler/mocks"
// 	adminservices "server/services/admin"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )

// func TestCheckAuth(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	mockAdminServices := new(mock.MockAdminServices)
// 	adminHandler := NewAdminHandler(mockAdminServices)
// 	assert := assert.New(t)

// 	adminHandler.CheckAuth(c)

// 	assert.Equal(http.StatusOK, w.Code, "Expected HTTP status code OK")
// 	assert.JSONEq(`{"success":true,"message":"Granted permission"}`, w.Body.String(), "Expected 'Granted permission' message to be returned")
// }

// func TestGetFail(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	mockAdminServices := new(mock.MockAdminServices)
// 	adminHandler := NewAdminHandler(mockAdminServices)
// 	assert := assert.New(t)

// 	adminHandler.Get(c)

// 	assert.Equal(http.StatusUnauthorized, w.Code, "Expected HTTP status code 401")
// 	assert.JSONEq(`{"success":false,"message":"Unauthorized access"}`, w.Body.String(), "Expected 'Unauthorized access' message to be returned")
// }

// func TestGetAdminFail(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	c.Set("username", "admin2023")

// 	mockAdminServices := new(mock.MockAdminServices)
// 	adminHandler := NewAdminHandler(mockAdminServices)
// 	assert := assert.New(t)

// 	mockAdminServices.On("GetAdminUsername", "admin2023").Return("no admin", fmt.Errorf("cant get admin username"))
// 	adminHandler.Get(c)

// 	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
// 	assert.JSONEq(`{"success":false,"message":"Unauthorized access"}`, w.Body.String(), "Expected 'Unauthorized access' message to be returned")
// }

// func TestGetAdminSuccess(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	c.Set("username", "admin2023")

// 	mockAdminServices := new(mock.MockAdminServices)
// 	adminHandler := NewAdminHandler(mockAdminServices)
// 	assert := assert.New(t)

// 	mockAdminServices.On("GetAdminUsername", "admin2023").Return("admin2023", nil)
// 	adminHandler.Get(c)

// 	assert.Equal(http.StatusOK, w.Code, "Expected HTTP status code 200")
// 	assert.JSONEq(`{"success":true,"username":"admin2023"}`, w.Body.String(), "Expected admin username to be returned")
// }

// func mockJsonPost(c *gin.Context, content interface{}) {
// 	c.Request.Method = "POST"
// 	c.Request.Header.Set("Content-Type", "application/json")

// 	jsonbytes, err := json.Marshal(content)
// 	if err != nil {
// 		panic(err)
// 	}

// 	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
// }

// func TestLoginBindJSONFail(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)
// 	c.Request = &http.Request{
// 		Header: make(http.Header),
// 	}
// 	invalidAdminJson := "invalid"

// 	mockJsonPost(c, invalidAdminJson)

// 	mockAdminServices := new(mock.MockAdminServices)
// 	adminHandler := NewAdminHandler(mockAdminServices)
// 	assert := assert.New(t)

// 	adminHandler.Login(c)

// 	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
// 	assert.JSONEq(`{"success":false,"message":"internal server error"}`, w.Body.String(), "Expected 'internal server error' to be returned")
// }

// func TestLoginFail(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)
// 	c.Request = &http.Request{
// 		Header: make(http.Header),
// 	}
// 	adminJson := map[string]interface{}{"username": "admintest", "password": "passwordtest"}

// 	mockJsonPost(c, adminJson)
// 	admin := adminservices.Admin{
// 		Username: "admintest",
// 		Password: "passwordtest",
// 	}

// 	mockAdminServices := new(mock.MockAdminServices)
// 	adminHandler := NewAdminHandler(mockAdminServices)
// 	assert := assert.New(t)

// 	mockAdminServices.On("Login", &admin).Return("token string", fmt.Errorf("login failed"))
// 	adminHandler.Login(c)

// 	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
// 	assert.JSONEq(`{"success":false,"message":"login failed"}`, w.Body.String(), "Expected 'login failed' to be returned")
// }

// func TestLoginSuccess(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)
// 	c.Request = &http.Request{
// 		Header: make(http.Header),
// 	}
// 	adminJson := map[string]interface{}{"username": "admintest", "password": "passwordtest"}

// 	mockJsonPost(c, adminJson)
// 	admin := adminservices.Admin{
// 		Username: "admintest",
// 		Password: "passwordtest",
// 	}

// 	mockAdminServices := new(mock.MockAdminServices)
// 	adminHandler := NewAdminHandler(mockAdminServices)
// 	assert := assert.New(t)

// 	mockAdminServices.On("Login", &admin).Return("token string",  nil)
// 	adminHandler.Login(c)

// 	assert.Equal(http.StatusOK, w.Code, "Expected HTTP status code 200")
// 	assert.JSONEq(`{"success":true,"token":"token string"}`, w.Body.String(), "Expected a token string to be returned")
// }

// func TestChangePasswordFailToGetContext(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	c.Set("there is no username", "admin2023")

// 	mockAdminServices := new(mock.MockAdminServices)
// 	adminHandler := NewAdminHandler(mockAdminServices)
// 	assert := assert.New(t)

// 	adminHandler.ChangePassword(c)

// 	assert.Equal(http.StatusUnauthorized, w.Code, "Expected HTTP status code 401")
// 	assert.JSONEq(`{"success":false,"message":"Unauthorized access"}`, w.Body.String(), "Expected 'Unauthorized access' message to be returned")
// }

// func TestChangePasswordFailToBindJSON(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	c.Set("username", "admin2023")

// 	mockAdminServices := new(mock.MockAdminServices)
// 	adminHandler := NewAdminHandler(mockAdminServices)
// 	assert := assert.New(t)

// 	adminHandler.ChangePassword(c)

// 	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
// 	assert.JSONEq(`{"success":false,"message":"internal server error"}`, w.Body.String(), "Expected 'internal server error' message to be returned")
// }

// func TestChangePasswordFailCheckUsernameToken(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	c.Set("username", "admintest")

// 	c.Request = &http.Request{
// 		Header: make(http.Header),
// 	}
// 	adminJson := map[string]interface{}{"username": "admintest", "password": "passwordtest", "password_confirmation":"passwordtest"}
// 	mockJsonPost(c, adminJson)

// 	// admin := adminservices.AdminWithConfirmPassword{
// 	// 	Username: "admintest",
// 	// 	Password: "passwordtest",
// 	// 	PasswordConfirmation: "passwordtest",
// 	// }

// 	mockAdminServices := new(mock.MockAdminServices)
// 	adminHandler := NewAdminHandler(mockAdminServices)
// 	assert := assert.New(t)

// 	mockAdminServices.On("CheckAdminUsernameToken", "admintest").Return(fmt.Errorf("cant check admin username token"))
// 	adminHandler.ChangePassword(c)

// 	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
// 	assert.JSONEq(`{"success":false,"message":"Unauthorized access"}`, w.Body.String(), "Expected 'Unauthorized access' message to be returned")
// }

// func TestChangePasswordFail(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	c.Set("username", "admintest")

// 	c.Request = &http.Request{
// 		Header: make(http.Header),
// 	}
// 	adminJson := map[string]interface{}{"username": "admintest", "password": "passwordtest", "password_confirmation":"passwordtest"}
// 	mockJsonPost(c, adminJson)

// 	admin := adminservices.AdminWithConfirmPassword{
// 		Username: "admintest",
// 		Password: "passwordtest",
// 		PasswordConfirmation: "passwordtest",
// 	}

// 	mockAdminServices := new(mock.MockAdminServices)
// 	adminHandler := NewAdminHandler(mockAdminServices)
// 	assert := assert.New(t)

// 	mockAdminServices.On("CheckAdminUsernameToken", "admintest").Return(nil)
// 	mockAdminServices.On("ChangePassword", &admin, "admintest").Return(fmt.Errorf("change password fail"))
// 	adminHandler.ChangePassword(c)

// 	assert.Equal(http.StatusBadRequest, w.Code, "Expected HTTP status code 400")
// 	assert.JSONEq(`{"success":false,"message":"change password fail"}`, w.Body.String(), "Expected 'change password fail' message to be returned")
// }

// func TestChangePasswordSuccess(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	c.Set("username", "admintest")

// 	c.Request = &http.Request{
// 		Header: make(http.Header),
// 	}
// 	adminJson := map[string]interface{}{"username": "admintest", "password": "passwordtest", "password_confirmation":"passwordtest"}
// 	mockJsonPost(c, adminJson)

// 	admin := adminservices.AdminWithConfirmPassword{
// 		Username: "admintest",
// 		Password: "passwordtest",
// 		PasswordConfirmation: "passwordtest",
// 	}

// 	mockAdminServices := new(mock.MockAdminServices)
// 	adminHandler := NewAdminHandler(mockAdminServices)
// 	assert := assert.New(t)

// 	mockAdminServices.On("CheckAdminUsernameToken", "admintest").Return(nil)
// 	mockAdminServices.On("ChangePassword", &admin, "admintest").Return(nil)
// 	adminHandler.ChangePassword(c)

// 	assert.Equal(http.StatusOK, w.Code, "Expected HTTP status code 200")
// 	assert.JSONEq(`{"success":true,"message":"Change password success"}`, w.Body.String(), "Expected 'Change password success' message to be returned")
// }

