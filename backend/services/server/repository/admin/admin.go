package repository

import (
	"fmt"
	"io"
	"os"
	"server/entities"

	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
)


type adminRepo struct {
	admin entities.Admin
	path  string
}

func NewAdminRepo(admin entities.Admin, path string) *adminRepo {
	adminRepo := &adminRepo{
		admin: admin,
		path:  path,
	}
	return adminRepo
}

func (repo *adminRepo)GetAdmin() (entities.Admin) {
	return repo.admin
}

func (repo *adminRepo)SetAdmin(newAdmin entities.Admin) () {
	repo.admin = newAdmin
}

func (repo *adminRepo)ReadAdminJSON() (entities.Admin, error) {
	var adminConfig entities.Admin
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	adminConfigJson, err := os.Open(fmt.Sprintf("%sadminConfig.json", repo.path))
	if err != nil {
		log.Errorln(err)
		return adminConfig, err
	}
	defer adminConfigJson.Close()

	adminConfigByte, err := io.ReadAll(adminConfigJson)
	if err != nil {
		log.Errorln(err)
		return adminConfig, err
	}

	err = json.Unmarshal(adminConfigByte, &adminConfig)
	if err != nil {
		log.Println(err)
		return adminConfig, err
	}
	repo.admin = adminConfig
	return adminConfig, nil
}

func (repo *adminRepo) WriteAdminJSON(admin *entities.Admin) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	file, err := os.Create(fmt.Sprintf("%sadminConfig.json", repo.path))
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