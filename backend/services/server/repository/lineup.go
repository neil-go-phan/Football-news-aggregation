package repository

import (
	"gorm.io/gorm"
	"server/entities"
)

type LineupRepo struct {
	DB *gorm.DB
}

func NewLineupRepo(db *gorm.DB) *LineupRepo {
	LineupRepo := &LineupRepo{
		DB: db,
	}
	return LineupRepo
}

func (repo *LineupRepo) FirstOrCreate(lineup *entities.MatchLineUp) (*entities.MatchLineUp, error) {
	err := repo.DB.FirstOrCreate(lineup, entities.MatchLineUp{ClubName: lineup.ClubName, Formation: lineup.Formation, ShirtColor: lineup.ShirtColor}).Error
	if err != nil {
		return lineup, err
	}
	return lineup, nil
}

func (repo *LineupRepo) Get(id uint) (*entities.MatchLineUp, error) {
	lineup := new(entities.MatchLineUp)
	err := repo.DB.
	Preload("Players").
	Where("id = ?", id).
	Find(lineup).Error
	if err != nil {
		return nil, err
	}
	return lineup, nil
}
