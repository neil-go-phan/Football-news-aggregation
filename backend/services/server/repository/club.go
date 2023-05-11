package repository

import (
	"server/entities"

	"gorm.io/gorm"
)

type clubRepo struct {
	DB *gorm.DB
}

func NewClubRepo(db *gorm.DB) *clubRepo {
	return &clubRepo{
		DB: db,
	}
}

func (repo *clubRepo) FirstOrCreate(clubName string, logo string) (*entities.Club,error) {
	club := &entities.Club{
		Name: clubName,
		Logo: logo,
	}
	err := repo.DB.Where("name = ?", clubName).FirstOrCreate(&club).Error
	if err != nil {
		return club, err
	}
	return club, nil
}

func (repo *clubRepo) GetByName(clubName string) (*entities.Club,error) {
	club := new(entities.Club)
	err := repo.DB.Where("name = ?", clubName).Find(&club).Error
	if err != nil {
		return club, err
	}
	return club, nil
}