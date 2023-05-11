package entities

import "gorm.io/gorm"

type MatchLineUp struct {
	gorm.Model
	ClubName   string     
	Formation  string     
	ShirtColor string    
	Players  []Player `gorm:"foreignKey:MatchLineUpID"`
}