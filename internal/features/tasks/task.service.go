package tasks

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	entities "todo-go-fiber/internal/db/entities"

	"github.com/gofiber/fiber/v2"
)

func CreateTask(c *fiber.Ctx, conn *sql.DB) (int64, error) {
	var task entities.Task
	err := c.BodyParser(&task)
	if err != nil {
		return 0, fmt.Errorf("cannot parse request body: %v", err)
	}

	// todo Валидация
	var id int64
	queryStr := []string{"INSERT INTO tasks (\n", "", "\n)\nVALUES (", "", ")\nRETURNING id;"}
	if len(task.Title) != 0 {
		queryStr[1] = queryStr[1] + "title"
		queryStr[3] = queryStr[3] + fmt.Sprintf("'%v'", task.Title)
	}
	if task.Description != nil && len(*task.Description) > 0 {
		if len(queryStr[1]) > 0 {
			queryStr[1] = queryStr[1] + ", "
		}
		if len(queryStr[3]) > 0 {
			queryStr[3] = queryStr[3] + ", "
		}
		queryStr[1] = queryStr[1] + "description"
		queryStr[3] = queryStr[3] + fmt.Sprintf("'%v'", *task.Description)
	}

	err = UpdateTask(conn, strings.Join(queryStr, " "), &id)
	if err != nil {
		return 0, err
	}

	return int64(id), nil
}

func GetTask(c *fiber.Ctx, conn *sql.DB) ([]entities.Task, error) {
	queryId := c.Query("id")

	var id int64 = 0
	if len(queryId) > 0 { // todo func
		idInt, err := strconv.Atoi(queryId)
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
	pId := c.Params("id")
	if pId == "" {
		return 0, fmt.Errorf("id was not provided")
	}
	var idInt int64 = 0
	if len(pId) > 0 { // todo func
		id, err := strconv.Atoi(pId)
		if err != nil {
			return 0, err
		}
		idInt = int64(id)
	}

	var task entities.Task
	err := c.BodyParser(&task)
	if err != nil {
		return 0, fmt.Errorf("cannot parse request body: %v", err)
	}

	paramsCount := 0
	var id int64 = 0
	queryStr := "UPDATE tasks set "
	if len(task.Title) != 0 {
		queryStr = queryStr + fmt.Sprintf("title = '%v'", task.Title)
		paramsCount++
	}
	if task.Description != nil && len(*task.Description) > 0 {
		if paramsCount > 0 {
			queryStr = queryStr + ", "
		}
		queryStr = queryStr + fmt.Sprintf("description = '%v'", *task.Description)
		paramsCount++
	}
	if len(task.Status) != 0 {
		if paramsCount > 0 {
			queryStr = queryStr + ", "
		}
		queryStr = queryStr + fmt.Sprintf("status = '%v'", task.Status)
		paramsCount++
	}
	if paramsCount == 0 {
		return 0, fmt.Errorf("params was not provided")
	}
	queryStr = queryStr + fmt.Sprintf(",updated_at = now()\nwhere id = '%v'\nRETURNING id;", idInt)

	err = UpdateTask(conn, queryStr, &id)
	if err != nil {
		return 0, err
	}

	return int64(id), nil
}

func DeleteTask(c *fiber.Ctx, conn *sql.DB) (int64, error) {
	pId := c.Params("id")
	if pId == "" {
		return 0, fmt.Errorf("id was not provided")
	}

	var id int64 = 0

	var idInt int64 = 0
	if len(pId) > 0 { // todo func
		id, err := strconv.Atoi(pId)
		if err != nil {
			return 0, err
		}
		idInt = int64(id)
	}
	queryStr := fmt.Sprintf("DELETE FROM tasks where id = %v RETURNING id;", idInt)

	err := UpdateTask(conn, queryStr, &id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, fmt.Errorf("cant delete")
	}

	return int64(id), nil
}
