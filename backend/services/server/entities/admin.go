package entities

import (
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Username string `gorm:"uniqueIndex"`
	Password string 
}

type JWTClaim struct {
	Username string `json:"username"`
	RandomString []byte `json:"random_string"`
	jwt.RegisteredClaims
}
