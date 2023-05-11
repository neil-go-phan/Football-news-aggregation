package entities

import "gorm.io/gorm"

type League struct {
	gorm.Model
	LeagueName string `gorm:"uniqueIndex:idx_league_name"`
	Active bool 
}

