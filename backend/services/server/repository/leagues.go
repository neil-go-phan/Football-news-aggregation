package repository

import (
	"server/entities"

	"gorm.io/gorm"
)

type leaguesRepo struct {
	DB *gorm.DB
}

func NewLeaguesRepo(db *gorm.DB) *leaguesRepo {
	return &leaguesRepo{
		DB: db,
	}
}

func (repo *leaguesRepo) Create(league *entities.League) error {
	err := repo.DB.Create(league).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *leaguesRepo) GetLeaguesName() (*[]entities.League, error) {
	leagues := make([]entities.League, 10)
	err := repo.DB.Select("league_name", "id").Find(&leagues).Error
	if err != nil {
		return nil, err
	}
	return &leagues, nil
}

func (repo *leaguesRepo) GetLeaguesNameActive() (*[]entities.League, error) {
	leagues := make([]entities.League, 0)
	err := repo.DB.Where(map[string]interface{}{"active": true}).Find(&leagues).Error
	if err != nil {
		return nil, err
	}
	return &leagues, nil
}

func (repo *leaguesRepo) List() (*[]entities.League, error) {
	leagues := make([]entities.League, 0)
	err := repo.DB.Order("active desc").Find(&leagues).Error
	if err != nil {
		return nil, err
	}
	return &leagues, nil
}

func (repo *leaguesRepo) GetByName(leagueName string) (*entities.League, error) {
	league := new(entities.League)
	err := repo.DB.Where(map[string]interface{}{"league_name": leagueName}).Find(&league).Error
	if err != nil {
		return nil, err
	}
	return league, nil
}

func (repo *leaguesRepo) Update(league *entities.League) error {
	err := repo.DB.Model(&league).Where("id = ?", league.ID).Updates(map[string]interface{}{"active": league.Active}).Error
	if err != nil {
		return err
	}
	return nil
}
