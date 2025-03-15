package tasks

import (
	"database/sql"
	"fmt"

	db "todo-go-fiber/internal/db"

	entities "todo-go-fiber/internal/db/entities"
)

func UpdateTask(conn *sql.DB, query string, pi *int64) error {
	err := db.TaskTransaction(conn, query, pi)
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

	tran, err := conn.Begin()
	if err != nil {
		return nil, fmt.Errorf("cant start transaction: '%v'", err)
	}

	rows, err := tran.Query(query)
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
