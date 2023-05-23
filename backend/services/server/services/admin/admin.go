package adminservices

import (
	"fmt"
	"server/entities"
	"server/repository"
	"strings"

	log "github.com/sirupsen/logrus"

	"time"
)

const TOKEN_LIFE = 3 * time.Hour
const RANDOM_TOKEN_STRING_SIZE = 8

var NOTI_LOGIN_SUCCESS_TITLE = "Admin login success"
var NOTI_LOGIN_SUCCESS_TYPE = "INFO"
var NOTI_LOGIN_SUCCESS_MESSAGE = "Admin login sucess"

type AdminService struct {
	adminRepo repository.AdminRepository
}

func NewAdminService(admin repository.AdminRepository) *AdminService {
	adminService := &AdminService{
		adminRepo: admin,
	}
	return adminService
}

func (s *AdminService) GetAdminUsername(username string) (string, error) {
	err := s.CheckAdminUsernameToken(username)
	if err != nil {
		return "", fmt.Errorf("unauthorized access")
	}
	repoAdmin,err := s.adminRepo.Get(username)
	if err != nil {
		return "", err
	}
	return repoAdmin.Username, nil
}

func (s *AdminService) ChangePassword(admin *AdminWithConfirmPassword, usernameToken string) error {
	repoAdmin, err := s.adminRepo.Get(admin.Username)
	if err != nil {
		return err
	}
	if repoAdmin.Username != admin.Username {
		log.Errorf("error occrus when trying to change admin password: admin username input different from admin username\n")
		return fmt.Errorf("input invalid")
	}

	adminValidate := &Admin{
		Username: admin.Username,
		Password: admin.Password,
	}

	err = validateAdmin(adminValidate)
	if err != nil {
		log.Errorf("error occrus when trying to change admin password: %s\n", err)
		return fmt.Errorf("input invalid")
	}

	if admin.Password != admin.PasswordConfirmation {
		log.Errorf("error occrus when trying to change admin password: password different from password confirm\n")
		return fmt.Errorf("password different from password confirm")
	}

	entityAdmin := &entities.Admin{
		Username: admin.Username,
		Password: admin.Password,
	}

	err = s.adminRepo.UpdatePassword(entityAdmin)
	if err != nil {
		log.Errorf("error occrus when trying to change admin password: %s\n", err)
		return fmt.Errorf("internal server error")
	}

	return nil
}

func (s *AdminService) CheckAdminUsernameToken(username string) error {
	repoAdmin, err := s.adminRepo.Get(username)
	if err != nil {
		return err
	}
	if repoAdmin.Username != username {
		log.Errorf("Detect a strange token string: token username: %s\n", username)
		return fmt.Errorf("username input different from admin username")
	}
	return nil
}

func (s *AdminService) LoginWithUsername(admin *Admin) (string, error) {
	repoAdmin, err := s.adminRepo.Get(admin.Username)
	if err != nil {
		return "",err
	}
	err = validateAdmin(admin)
	if err != nil {
		log.Errorf("error occrus when a anonymous user try to login admin: %s\n", err)
		return "", fmt.Errorf("input invalid")
	}

	err = checkIsAdminCorrect(admin, *repoAdmin)
	if err != nil {
		log.Errorf("error occrus when a anonymous user try to login admin: %s\n", err)
		return "", fmt.Errorf("username or password is incorrect")
	}

	token, err := generateToken(admin.Username)
	if err != nil {
		log.Errorf("error occrus when a anonymous user try to login admin: %s\n", err)
		return "", fmt.Errorf("internal server error")
	}

	return token, nil
}

func (s *AdminService) LoginWithEmail(admin *Admin) (string, error) {
	err := validateAdminLoginWithEmail(admin)
	if err != nil {
		log.Errorf("error occrus when a anonymous user try to login admin: %s\n", err)
		return "", fmt.Errorf("input invalid")
	}

	repoAdmin, err := s.adminRepo.GetWithEmail(admin.Username)
	if err != nil {
		return "",err
	}

	err = checkIsAdminEmailCorrect(admin, *repoAdmin)
	if err != nil {
		log.Errorf("error occrus when a anonymous user try to login admin: %s\n", err)
		return "", fmt.Errorf("username or password is incorrect")
	}

	token, err := generateToken(repoAdmin.Username)
	if err != nil {
		log.Errorf("error occrus when a anonymous user try to login admin: %s\n", err)
		return "", fmt.Errorf("internal server error")
	}

	return token, nil
}

func (s *AdminService) Register(admin *RegisterUserInput) (error) {
	// validate
	err := validateRegisterAdmin(admin)
	if err != nil {
		return err
	}
	// create
	newAdmin := entities.Admin{
		Username: admin.Username,
		Password: admin.Password,
		Email: admin.Email,
	}
	err = s.adminRepo.Create(&newAdmin)
	if err != nil {
		log.Error(err)
		return fmt.Errorf("internal server error")
	}
	return nil
}

func (s *AdminService) GoogleOAuth(googleUser *GoogleUserResult) (string, error) {
	// upsert
	email := strings.ToLower(googleUser.Email)
	username := strings.Split(email, "@")
	newAdmin := &entities.Admin{
		Username: username[0],
		Email: email,
		Password: "",
	}
	err := s.adminRepo.Upsert(newAdmin)
	if err != nil {
		return "", fmt.Errorf("can not upsert user")
	}
	// select user
	admin, err := s.adminRepo.GetWithEmail(newAdmin.Email)
	if err != nil {
		return "", fmt.Errorf("user not found")
	} 

	token, err := generateToken(admin.Username)
	if err != nil {
		log.Errorf("error occrus when a anonymous user try to login admin: %s\n", err)
		return "", fmt.Errorf("internal server error")
	}

	return token, nil
}