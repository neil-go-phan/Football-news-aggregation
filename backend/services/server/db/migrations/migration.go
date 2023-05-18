package migrations

import (
	log "github.com/sirupsen/logrus"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)


func RunDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("cant create new migrate instance: ", err)
	}
	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("fail to run migrate up: ", err)
	}
	log.Println("db migrate successfully")
}