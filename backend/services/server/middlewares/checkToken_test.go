package middlewares

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"server/entities"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCheckToken(t *testing.T) {
	r := gin.New()
	r.Use(CheckToken())
	req1 := httptest.NewRequest(http.MethodGet, "/check", nil)
	w1 := httptest.NewRecorder()
	r.ServeHTTP(w1, req1)

	// Case no token
	if w1.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %v but got %v", http.StatusUnauthorized, w1.Code)
	}
	expectedMsg1 := `{"message":"Unauthorized access","success":false}`
	if strings.TrimSpace(w1.Body.String()) != expectedMsg1 {
		t.Errorf("Expected response body %v but got %v", expectedMsg1, w1.Body.String())
	}

	// Case invalid token
	tokenString := "invalid-token"
	req2 := httptest.NewRequest(http.MethodGet, "/check", nil)

	req2.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tokenString))
	w2 := httptest.NewRecorder()

	r.ServeHTTP(w2, req2)

	if w1.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %v but got %v", http.StatusUnauthorized, w1.Code)
	}
	expectedMsg2 := `{"message":"Unauthorized access","success":false}`
	if strings.TrimSpace(w1.Body.String()) != expectedMsg2 {
		t.Errorf("Expected response body %v but got %v", expectedMsg2, w1.Body.String())
	}

	// case OK
	tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluMjAyMyIsInJhbmRvbV9zdHJpbmciOm51bGwsImV4cCI6OTk5OTk5OTk5OX0.4kZY-hdYr_7TFdnFZ-gw0pcglG_MXAjly61HIEdc9mQ"

	req3 := httptest.NewRequest(http.MethodGet, "/check", nil)
	req3.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tokenString))

	w3 := httptest.NewRecorder()
	r.ServeHTTP(w3, req3)

	if w3.Code != http.StatusNotFound {
		t.Errorf("Expected status code %v but got %v", http.StatusUnauthorized, w1.Code)
	}
}

func TestValidateToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluMjAyMyIsInJhbmRvbV9zdHJpbmciOm51bGwsImV4cCI6OTk5OTk5OTk5OX0.4kZY-hdYr_7TFdnFZ-gw0pcglG_MXAjly61HIEdc9mQ"

	// Test case 1: valid token
	want := &entities.JWTClaim{
		Username: "admin2023",
	}
	claims, err := validateToken(tokenString)
	assert.NoError(t, err, "Expected no error for a valid token")
	assert.Equal(t, want.Username, claims.Username, "Expected username to be set in the claims")

	// Test case 2: invalid token
	tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluMjAyMyIsInJhbmRvbV9zdHJpbmciOm51bGwsImV4cCI6MTY4MzAyMjg5M30.nkGigx0qCuN5lAovGqNpuGFBkZjvG_Wc9PsIibMYPjc"
	claims, err = validateToken(tokenString)

	assert.Error(t, err, "Expected an error for an invalid token")
	assert.Nil(t, claims, "Expected no claims to be set for an invalid token")

	// Test case 3: unexpected signing method"

	tokenString = "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJFUzI1NmluT1RBIiwibmFtZSI6IkpvaG4gRG9lIn0.MEQCICRphRrc0GWowZgJAy0gL6At628Kw8YPE22iD-aKIi4PAiA0JWU-qFNL8I0tP0ws3Bbmg0FfVMn4_yk2lGGquAGOXA"
	claims, err = validateToken(tokenString)

	assert.Error(t, err, "Expected an error for an invalid token")
	assert.Nil(t, claims, "Expected no claims to be set for an invalid token")
}
