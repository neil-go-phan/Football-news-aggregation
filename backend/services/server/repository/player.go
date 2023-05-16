package repository

import (
	"gorm.io/gorm"
	"server/entities"
)

type PlayerRepo struct {
	DB *gorm.DB
}

func NewPlayerRepo(db *gorm.DB) *PlayerRepo {
	PlayerRepo := &PlayerRepo{
		DB: db,
	}
	return PlayerRepo
}

func (repo *PlayerRepo) FirstOrCreate(player *entities.Player) error {
	err := repo.DB.FirstOrCreate(player, entities.Player{MatchLineUpID: player.MatchLineUpID, PlayerName: player.PlayerName, PlayerNumber: player.PlayerNumber, MatchID: player.MatchID}).Error
	if err != nil {
		return err
	}
	return nil
}
