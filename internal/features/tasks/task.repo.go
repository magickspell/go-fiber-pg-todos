package tasks

import (
	"database/sql"
	"fmt"

	db "todo-go-fiber/internal/db"

	entities "todo-go-fiber/internal/db/entities"
)

func InsertTask(conn *sql.DB, query string, pi *int64) error {
	err := db.InsertTr(conn, query, pi)
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

	rows, err := conn.Query(query)
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

	return tasks, nil
}

func UpsertTask() {

}

func DeleteTask() {

}
