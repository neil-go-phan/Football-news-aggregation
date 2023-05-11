package repository

import (
	"server/entities"

	"gorm.io/gorm"
)

type adminRepo struct {
	DB *gorm.DB
}

func NewAdminRepo(db *gorm.DB) *adminRepo {
	return &adminRepo{
		DB: db,
	}
}

func (repo *adminRepo) Create(admin *entities.Admin) error {
	err := repo.DB.Create(admin).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *adminRepo) Get(username string) (*entities.Admin, error) {
	admin := new(entities.Admin)
	err := repo.DB.Where(map[string]interface{}{"username": username}).Find(&admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (repo *adminRepo) UpdatePassword(admin *entities.Admin) error  {
	err := repo.DB.Model(&admin).Where("username = ?", admin.Username).Updates(map[string]interface{}{"password": admin.Password}).Error
	if err != nil {
		return err
	}
	return nil
}
