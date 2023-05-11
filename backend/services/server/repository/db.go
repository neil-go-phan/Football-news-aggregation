package repository

import (
	log "github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDB(dbSource string) {
	var err error

	db, err := gorm.Open(postgres.Open(dbSource), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database : error=%v", err)
	}
	Db = db
}

func GetDB() *gorm.DB {
	return Db
}