package tasks

import (
	"database/sql"
	"errors"
	"fmt"

	entities "todo-go-fiber/internal/db/entities"

	"github.com/gofiber/fiber/v2"
)

func CreateTask(c *fiber.Ctx, conn *sql.DB) (int64, error) {
	var task entities.Task
	err := c.BodyParser(&task)
	if err != nil {
		return 0, fmt.Errorf("cannot parse request body: %v", err)
	}

	var id int64
	queryStr, err := getCreateTaskQuery(&task)
	if err != nil {
		return 0, err
	}

	err = UpdateTask(conn, queryStr, &id)
	if err != nil {
		return 0, err
	}

	return int64(id), nil
}

func GetTask(c *fiber.Ctx, conn *sql.DB) ([]entities.Task, error) {
	queryId := c.Query("id")

	var id int64 = 0
	if len(queryId) > 0 {
		idInt, err := validateId(queryId)
		if err != nil {
			return []entities.Task{}, err
		}
		id = int64(idInt)
	}

	tasks, err := SelectTask(conn, id)
	if err != nil {
		return []entities.Task{}, err
	}

	return tasks, nil
}

func PutTask(c *fiber.Ctx, conn *sql.DB) (int64, error) {
	idInt, err := validateId(c.Params("id"))
	if err != nil {
		return 0, err
	}

	var task entities.Task
	err = c.BodyParser(&task)
	if err != nil {
		return 0, fmt.Errorf("cannot parse request body: %v", err)
	}

	query, err := getUpdateTaskQuery(&task, idInt)
	if err != nil {
		return 0, err
	}

	var id int64 = 0
	err = UpdateTask(conn, query, &id)
	if err != nil {
		return 0, err
	}

	return int64(id), nil
}

func DeleteTask(c *fiber.Ctx, conn *sql.DB) (int64, error) {
	idInt, err := validateId(c.Params("id"))
	if err != nil {
		return 0, err
	}
	queryStr := fmt.Sprintf("DELETE FROM tasks where id = %v RETURNING id;", idInt)

	var id int64 = 0
	err = UpdateTask(conn, queryStr, &id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, fmt.Errorf("cant delete")
	}

	return int64(id), nil
}
