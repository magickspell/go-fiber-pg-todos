package database

import (
	"database/sql"
	"fmt"
)

func InsertTr(db *sql.DB, query string, pi *int64) error {
	tran, err := db.Begin()
	if err != nil {
		return fmt.Errorf("cant start transaction: '%v'", err)
	}

	err = tran.QueryRow(query).Scan(pi)
	if err != nil {
		return fmt.Errorf("unable to QueryRow or Scan: '%v'", err)
	}

	err = tran.Commit()
	if err != nil {
		return fmt.Errorf("unable to commit: '%v'", err)
	}

	return nil
}
