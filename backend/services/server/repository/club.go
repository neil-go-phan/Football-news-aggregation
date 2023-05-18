package repository

import (
	"server/entities"

	"gorm.io/gorm"
)

type ClubRepo struct {
	DB *gorm.DB
}

func NewClubRepo(db *gorm.DB) *ClubRepo {
	return &ClubRepo{
		DB: db,
	}
}

func (repo *ClubRepo) FirstOrCreate(clubName string, logo string) (*entities.Club,error) {
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

func (repo *ClubRepo) GetByName(clubName string) (*entities.Club,error) {
	club := new(entities.Club)
	err := repo.DB.Where("name = ?", clubName).Find(&club).Error
	if err != nil {
		return club, err
	}
	return club, nil
}