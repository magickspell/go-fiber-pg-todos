package migrations

import (
	"database/sql"
	"log"

	"github.com/pressly/goose"
)

func RunMigrations(dbconn *sql.DB) {
	log.Println("transaction started")

	if err := goose.SetDialect("postgres"); err != nil {
		log.Println(err)
	}

	if err := goose.Up(dbconn, "internal/db/migrations/sql"); err != nil {
		log.Println(err)
	}

	log.Println("transaction passed")
}
