package routes

// import "testing"

// func TestGetLeaguesName(t *testing.T) {
// 	mockResponse := `{"message":"Welcome to the Tech Company listing API with Golang"}`
// 	r := SetUpRouter()
// 	r.GET("/", HomepageHandler)
// 	req, _ := http.NewRequest("GET", "/", nil)
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	responseData, _ := ioutil.ReadAll(w.Body)
// 	assert.Equal(t, mockResponse, string(responseData))
// 	assert.Equal(t, http.StatusOK, w.Code)
// }