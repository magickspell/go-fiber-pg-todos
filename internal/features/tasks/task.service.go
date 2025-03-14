package tasks

import (
	"database/sql"
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

	err = InsertTask(conn, strings.Join(queryStr, " "), &id)
	if err != nil {
		return 0, err
	}

	return int64(id), nil
}

func GetTask(c *fiber.Ctx, conn *sql.DB) ([]entities.Task, error) {
	queryId := c.Query("id")
	fmt.Println("[queryId]")
	fmt.Println(queryId)
	fmt.Println(len(queryId))
	var id int64 = 0
	if len(queryId) > 0 {
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
