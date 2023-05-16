package repository

import (
	"server/entities"

	"gorm.io/gorm"
)

type LeaguesRepo struct {
	DB *gorm.DB
}

func NewLeaguesRepo(db *gorm.DB) *LeaguesRepo {
	return &LeaguesRepo{
		DB: db,
	}
}

func (repo *LeaguesRepo) Create(league *entities.League) error {
	err := repo.DB.Create(league).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *LeaguesRepo) GetLeaguesName() (*[]entities.League, error) {
	leagues := make([]entities.League, 10)
	err := repo.DB.Select("league_name", "id").Find(&leagues).Error
	if err != nil {
		return nil, err
	}
	return &leagues, nil
}

func (repo *LeaguesRepo) GetLeaguesNameActive() (*[]entities.League, error) {
	leagues := make([]entities.League, 0)
	err := repo.DB.Where(map[string]interface{}{"active": true}).Find(&leagues).Error
	if err != nil {
		return nil, err
	}
	return &leagues, nil
}

func (repo *LeaguesRepo) List() (*[]entities.League, error) {
	leagues := make([]entities.League, 0)
	err := repo.DB.Order("active desc").Find(&leagues).Error
	if err != nil {
		return nil, err
	}
	return &leagues, nil
}

func (repo *LeaguesRepo) GetByName(leagueName string) (*entities.League, error) {
	league := new(entities.League)
	err := repo.DB.Where(map[string]interface{}{"league_name": leagueName}).Find(&league).Error
	if err != nil {
		return nil, err
	}
	return league, nil
}

func (repo *LeaguesRepo) Update(league *entities.League) error {
	err := repo.DB.Model(&league).Where("id = ?", league.ID).Updates(map[string]interface{}{"active": league.Active}).Error
	if err != nil {
		return err
	}
	return nil
}
