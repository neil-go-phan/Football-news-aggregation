package repository

import (
	"sync"

	log "github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var once sync.Once

func ConnectDB(dbSource string) *gorm.DB {
	once.Do(func() {
		db, err := gorm.Open(postgres.Open(dbSource), &gorm.Config{})
		if err != nil {
			log.Fatalln("error connecting to database : error=", err)
		}
		Db = db
	})
	return Db
}