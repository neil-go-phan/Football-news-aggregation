package repository

import (
	"server/entities"

	"gorm.io/gorm"
)

type AdminRepo struct {
	DB *gorm.DB
}

func NewAdminRepo(db *gorm.DB) *AdminRepo {
	return &AdminRepo{
		DB: db,
	}
}

func (repo *AdminRepo) Create(admin *entities.Admin) error {
	err := repo.DB.Create(admin).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *AdminRepo) Get(username string) (*entities.Admin, error) {
	admin := new(entities.Admin)
	err := repo.DB.Where(map[string]interface{}{"username": username}).Find(&admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (repo *AdminRepo) GetWithEmail(email string) (*entities.Admin, error) {
	admin := new(entities.Admin)
	err := repo.DB.Where(map[string]interface{}{"email": email}).Find(&admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (repo *AdminRepo) UpdatePassword(admin *entities.Admin) error  {
	err := repo.DB.Model(&admin).Where("username = ?", admin.Username).Updates(map[string]interface{}{"password": admin.Password}).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *AdminRepo) Upsert(admin *entities.Admin) (error) {
	if repo.DB.Model(&admin).Where("email = ?", admin.Email).Updates(&admin).RowsAffected == 0 {
		repo.DB.Create(&admin)
	}
	return nil
}