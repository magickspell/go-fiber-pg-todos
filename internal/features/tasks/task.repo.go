package tasks

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	entities "todo-go-fiber/internal/db/entities"
)

func UpdateTask(conn *sql.DB, query string, pi *int64) error {
	err := taskTransaction(conn, query, pi)
	if err != nil {
		return err
	}
	return nil
}

func SelectTask(conn *sql.DB, id int64) ([]entities.Task, error) {
	var query string = "SELECT * FROM tasks"
	if id > 0 {
		query = query + " WHERE id = " + fmt.Sprint(id)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	tran, err := conn.Begin()
	if err != nil {
		return nil, fmt.Errorf("cant start transaction: '%v'", err)
	}

	rows, err := tran.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []entities.Task

	for rows.Next() {
		var task entities.Task
		if err := rows.Scan(
			&task.Id,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	err = tran.Commit()
	if err != nil {
		return nil, fmt.Errorf("unable to commit: '%v'", err)
	}

	return tasks, nil
}

func taskTransaction(db *sql.DB, query string, pi *int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	tran, err := db.Begin()
	if err != nil {
		return fmt.Errorf("cant start transaction: '%w'", err)
	}

	err = tran.QueryRowContext(ctx, query).Scan(pi)
	if err != nil {
		return fmt.Errorf("unable to QueryRow or Scan: '%w'", err)
	}

	err = tran.Commit()
	if err != nil {
		return fmt.Errorf("unable to commit: '%w'", err)
	}

	return nil
}
