package database

import (
	"database/sql"
	"fmt"
	"os"

	config "todo-go-fiber/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(c *config.Config) *sql.DB {
	dbconn, err := sql.Open("pgx", c.DbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return dbconn
}
