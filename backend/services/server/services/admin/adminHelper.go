package adminservices

import (
	"errors"
	"net/mail"
	"regexp"
	"server/entities"
	serverhelper "server/helper"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v5"
)

type Admin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AdminWithConfirmPassword struct {
	Username             string `json:"username" validate:"required"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required"`
}

type RegisterUserInput struct {
	Username             string `json:"username" validate:"required"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required"`
	Email                string `json:"email" validate:"required"`
}

type GoogleOauthToken struct {
	Access_token string
	Id_token     string
}

type GoogleUserResult struct {
	Id             string
	Email          string
	Verified_email bool
	Name           string
	Given_name     string
	Family_name    string
	Picture        string
	Locale         string
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

func checkIsAdminEmailCorrect(admin *Admin, adminJson entities.Admin) error {
	if admin.Username != adminJson.Email {
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

func validateRegisterAdmin(admin *RegisterUserInput) (error) {
	validate := validator.New()
	err := validate.Struct(admin)
	if err != nil {
		return err
	}
	match := checkRegexp(admin.Password)
	if !match {
		return errors.New("password must not contain special character")
	}
	match = checkRegexp(admin.Username)
	if !match {
		return errors.New("username must not contain special character")
	}
	match = validEmail(admin.Email)
	if !match {
		return errors.New("invalid email")
	}
	if admin.Password != admin.PasswordConfirmation {
		return errors.New("password confirm not match")
	}
	return nil
}

func validateAdminLoginWithEmail(admin *Admin) error {
	validate := validator.New()
	err := validate.Struct(admin)
	if err != nil {
		return err
	}
	match := checkRegexp(admin.Password)
	if !match {
		return errors.New("password must not contain special character")
	}
	match = validEmail(admin.Username)
	if !match {
		return errors.New("invalid email")
	}
	return nil
}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
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
