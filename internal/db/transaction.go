package database

import (
	"database/sql"
	"fmt"
)

func TaskTransaction(db *sql.DB, query string, pi *int64) error {
	tran, err := db.Begin()
	if err != nil {
		return fmt.Errorf("cant start transaction: '%w'", err)
	}

	err = tran.QueryRow(query).Scan(pi)
	if err != nil {
		return fmt.Errorf("unable to QueryRow or Scan: '%w'", err)
	}

	err = tran.Commit()
	if err != nil {
		return fmt.Errorf("unable to commit: '%w'", err)
	}

	return nil
}
