package entities

import "gorm.io/gorm"

type StatisticsItem struct {
	gorm.Model
	StatClub1     string 
	StatContent   string
	StatClub2     string 
	MatchID uint  
}
