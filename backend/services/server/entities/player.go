package entities

import "gorm.io/gorm"

// this column can not scale to a big table like player info.
type Player struct {
	gorm.Model
	PlayerName    string
	PlayerNumber  string
	Row           int
	Col           int
	MatchLineUpID uint
	MatchID       uint
}
