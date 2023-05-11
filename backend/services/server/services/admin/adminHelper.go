package adminservices

import (
	"errors"
	"regexp"
	"server/entities"
	serverhelper "server/helper"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v5"
)

type Admin struct {
	Username string `json:"username" validate:"required,min=8,max=16"`
	Password string `json:"password" validate:"required"`
}

type AdminWithConfirmPassword struct {
	Username             string `json:"username" validate:"required,min=8,max=16"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required"`
}

type JWTClaim struct {
	Username string `json:"username"`
	RandomString []byte `json:"random_string"`
	jwt.RegisteredClaims
}


func checkIsAdminCorrect(admin *Admin, adminJson entities.Admin) error {
	if admin.Username != adminJson.Username {
		return errors.New("username is incorrect")
	}
	if admin.Password != adminJson.Password {
		return errors.New("password is incorrect")
	}
	return nil
}

func validateAdmin(admin *Admin) error {
	validate := validator.New()
	match := checkRegexp(admin.Password)
	if !match {
		return errors.New("password must not contain special character")
	}
	match = checkRegexp(admin.Username)
	if !match {
		return errors.New("username must not contain special character")
	}
	err := validate.Struct(admin)
	if err != nil {
		return err
	}
	return nil
}

func checkRegexp(checkedString string) bool {
	match, _ := regexp.MatchString("^[a-zA-Z0-9_]*$", checkedString)
	return match
}

func generateToken(username string) (string, error) {
	expirationTime := time.Now().Add(TOKEN_LIFE)

	claims := &entities.JWTClaim{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(serverhelper.TOKEN_SERECT_KEY)
}
