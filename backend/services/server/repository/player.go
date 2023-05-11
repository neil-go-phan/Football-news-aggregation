package repository

import (
	"gorm.io/gorm"
	"server/entities"
)

type playerRepo struct {
	DB *gorm.DB
}

func NewPlayerRepo(db *gorm.DB) *playerRepo {
	playerRepo := &playerRepo{
		DB: db,
	}
	return playerRepo
}

func (repo *playerRepo) FirstOrCreate(player *entities.Player) error {
	err := repo.DB.FirstOrCreate(player, entities.Player{MatchLineUpID: player.MatchLineUpID, PlayerName: player.PlayerName, PlayerNumber: player.PlayerNumber, MatchID: player.MatchID}).Error
	if err != nil {
		return err
	}
	return nil
}
