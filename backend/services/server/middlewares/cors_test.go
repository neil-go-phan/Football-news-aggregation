package middlewares

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCors(t *testing.T) {
	r := gin.New()
	r.Use(Cors())

	req := httptest.NewRequest(http.MethodGet, "/cors", nil)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	fmt.Println("header", w.Header())
	if w.Header().Get("Access-Control-Allow-Credentials") != "true" {
		t.Errorf("Expected header Access-Control-Allow-Credentials is %v but got %v", "true", w.Header().Get("Access-Control-Allow-Credentials"))
	}

	if w.Header().Get("Access-Control-Allow-Methods") != "POST, OPTIONS, GET, PUT, DELETE" {
		t.Errorf("Expected header Access-Control-Allow-Methods is %v but got %v", "POST, OPTIONS, GET, PUT, DELETE", w.Header().Get("Access-Control-Allow-Methods"))
	}
	
}