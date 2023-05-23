package entities

import (
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Username string 
	Password string 
	Email    string
}

type JWTClaim struct {
	Username string `json:"username"`
	RandomString []byte `json:"random_string"`
	jwt.RegisteredClaims
}

// type User struct {
// 	gorm.Model
// 	Username     string    `gorm:"type:varchar(100);not null"`

// 	Password string    `gorm:"not null"`

// 	Verified  bool      `gorm:"default:false;"`
// 	Provider  string    `gorm:"default:'local';"`
// }