package entities

import "gorm.io/gorm"

type Club struct {
	gorm.Model
	Name string `json:"name" gorm:"uniqueIndex:idx_club_name"`
	Logo string `json:"logo"`
}
