package services

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"server/entities"
	serverhelper "server/helper"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v5"
	jsoniter "github.com/json-iterator/go"
)

const TOKEN_LIFE = 1 * time.Hour // 1 hour
const RANDOM_TOKEN_STRING_SIZE = 8

type adminService struct {
	admin entities.Admin
	path  string
}

func NewAdminService(admin entities.Admin, path string) *adminService {
	adminService := &adminService{
		admin: admin,
		path:  path,
	}
	return adminService
}

func (s *adminService) GetAdminUsername(username string) (string, error) {
	err := s.CheckAdminUsernameToken(username)
	if err != nil {
		return "", fmt.Errorf("unauthorized access")
	}

	return s.admin.Username, nil

}

func (s *adminService) ChangePassword(admin *AdminWithConfirmPassword, usernameToken string) error {

	if s.admin.Username != admin.Username {
		log.Printf("error occrus when trying to change admin password: admin username input different from admin username\n")
		return fmt.Errorf("input invalid")
	}

	adminValidate := &Admin{
		Username: admin.Username,
		Password: admin.Password,
	}

	err := validateAdmin(adminValidate)
	if err != nil {
		log.Printf("error occrus when trying to change admin password: %s\n", err)
		return fmt.Errorf("input invalid")
	}

	if admin.Password != admin.PasswordConfirmation {
		log.Printf("error occrus when trying to change admin password: password different from password confirm\n")
		return fmt.Errorf("password different from password confirm")
	}

	entityAdmin := &entities.Admin{
		Username: admin.Username,
		Password: admin.Password,
	}
	
	err = s.WriteAdminJSON(entityAdmin)
	if err != nil {
		log.Printf("error occrus when trying to change admin password: %s\n", err)
		return fmt.Errorf("internal server error")
	}

	newAdmin, err := ReadAdminJSON(s.path)
if err != nil {
	log.Printf("error occrus when trying to change admin password: %s\n", err)
	err := s.WriteAdminJSON(&s.admin)
	if err != nil {
		log.Printf("error occurs: %s", err)
		return fmt.Errorf("internal server error")
	}
	return fmt.Errorf("internal server error")
}

	s.admin = newAdmin

	return nil
}

func (s *adminService) CheckAdminUsernameToken(username string) error {
	if s.admin.Username != username {
		log.Printf("Detect a strange token string: token username: %s\n", username)
		return fmt.Errorf("username input different from admin username")
	}
	return nil
}

func (s *adminService) Login(admin *Admin) (string, error) {
	err := validateAdmin(admin)
	if err != nil {
		log.Printf("error occrus when a anonymous user try to login admin: %s\n", err)
		return "", fmt.Errorf("input invalid")
	}

	err = checkIsAdminCorrect(admin, s.admin)
	if err != nil {
		log.Printf("error occrus when a anonymous user try to login admin: %s\n", err)
		return "", fmt.Errorf("username or password is incorrect")
	}

	token, err := generateToken(admin.Username)
	if err != nil {
		log.Printf("error occrus when a anonymous user try to login admin: %s\n", err)
		return "", fmt.Errorf("internal server error")
	}
	return token, nil
}

func ReadAdminJSON(jsonPath string) (entities.Admin, error) {
	var adminConfig entities.Admin
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	adminConfigJson, err := os.Open(fmt.Sprintf("%sadminConfig.json", jsonPath))
	if err != nil {
		log.Println(err)
		return adminConfig, err
	}
	defer adminConfigJson.Close()

	adminConfigByte, err := io.ReadAll(adminConfigJson)
	if err != nil {
		log.Println(err)
		return adminConfig, err
	}

	err = json.Unmarshal(adminConfigByte, &adminConfig)
	if err != nil {
		log.Println(err)
		return adminConfig, err
	}
	return adminConfig, nil
}

func (s *adminService) WriteAdminJSON(admin *entities.Admin) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	file, err := os.Create(fmt.Sprintf("%sadminConfig.json", s.path))
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(admin)
	if err != nil {
		return err
	}

	return nil
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
