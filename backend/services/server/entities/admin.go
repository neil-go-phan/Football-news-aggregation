package entities

import "github.com/golang-jwt/jwt/v5"

type Admin struct {
	Username string `json:"username" validate:"required,min=8,max=16"`
	Password string `json:"password" validate:"required"`
}

type JWTClaim struct {
	Username string `json:"username"`
	RandomString []byte `json:"random_string"`
	jwt.RegisteredClaims
}
