package adminservices

import (
	"fmt"
	"server/entities"
	"server/repository"

	log "github.com/sirupsen/logrus"

	"time"
)

const TOKEN_LIFE = 1 * time.Hour
const RANDOM_TOKEN_STRING_SIZE = 8

var NOTI_LOGIN_SUCCESS_TITLE = "Admin login success"
var NOTI_LOGIN_SUCCESS_TYPE = "INFO"
var NOTI_LOGIN_SUCCESS_MESSAGE = "Admin login sucess"

type adminService struct {
	adminRepo repository.AdminRepository
}

func NewAdminService(admin repository.AdminRepository) *adminService {
	adminService := &adminService{
		adminRepo: admin,
	}
	return adminService
}

func (s *adminService) GetAdminUsername(username string) (string, error) {
	err := s.CheckAdminUsernameToken(username)
	if err != nil {
		return "", fmt.Errorf("unauthorized access")
	}
	repoAdmin := s.adminRepo.GetAdmin()
	return repoAdmin.Username, nil
}

func (s *adminService) ChangePassword(admin *AdminWithConfirmPassword, usernameToken string) error {
	repoAdmin := s.adminRepo.GetAdmin()
	if repoAdmin.Username != admin.Username {
		log.Errorf("error occrus when trying to change admin password: admin username input different from admin username\n")
		return fmt.Errorf("input invalid")
	}

	adminValidate := &Admin{
		Username: admin.Username,
		Password: admin.Password,
	}

	err := validateAdmin(adminValidate)
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

	err = s.adminRepo.WriteAdminJSON(entityAdmin)
	if err != nil {
		log.Errorf("error occrus when trying to change admin password: %s\n", err)
		return fmt.Errorf("internal server error")
	}

	newAdmin, err := s.adminRepo.ReadAdminJSON()
	if err != nil {
		log.Errorf("error occrus when trying to change admin password: %s\n", err)
		err := s.adminRepo.WriteAdminJSON(&repoAdmin)
		if err != nil {
			log.Errorf("error occurs: %s", err)
			return fmt.Errorf("internal server error")
		}
		return fmt.Errorf("internal server error")
	}

	s.adminRepo.SetAdmin(newAdmin)

	return nil
}

func (s *adminService) CheckAdminUsernameToken(username string) error {
	repoAdmin := s.adminRepo.GetAdmin()
	if repoAdmin.Username != username {
		log.Errorf("Detect a strange token string: token username: %s\n", username)
		return fmt.Errorf("username input different from admin username")
	}
	return nil
}

func (s *adminService) Login(admin *Admin) (string, error) {
	repoAdmin := s.adminRepo.GetAdmin()
	err := validateAdmin(admin)
	if err != nil {
		log.Errorf("error occrus when a anonymous user try to login admin: %s\n", err)
		return "", fmt.Errorf("input invalid")
	}

	err = checkIsAdminCorrect(admin, repoAdmin)
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
