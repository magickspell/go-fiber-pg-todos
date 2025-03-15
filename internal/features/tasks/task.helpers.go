package tasks

import (
	"fmt"
	"strconv"
	"strings"

	entities "todo-go-fiber/internal/db/entities"
)

func validateId(idStr string) (int64, error) {
	if len(idStr) > 0 {
		idInt, err := strconv.Atoi(idStr)
		if err != nil {
			return 0, err
		}
		return int64(idInt), nil
	}
	return 0, fmt.Errorf("id was not provided")
}

func getCreateTaskQuery(task *entities.Task) (string, error) {
	counter := 0
	queryStr := []string{"INSERT INTO tasks (\n", "", "\n)\nVALUES (", "", ")\nRETURNING id;"}
	if len(task.Title) != 0 {
		queryStr[1] = queryStr[1] + "title"
		queryStr[3] = queryStr[3] + fmt.Sprintf("'%v'", task.Title)
		counter++
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
		counter++
	}
	if counter == 0 {
		return "", fmt.Errorf("params was not provided")
	}
	return strings.Join(queryStr, " "), nil
}

func getUpdateTaskQuery(task *entities.Task, id int64) (string, error) {
	paramsCount := 0
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
		return "", fmt.Errorf("params was not provided")
	}
	queryStr = queryStr + fmt.Sprintf(",updated_at = now()\nwhere id = '%v'\nRETURNING id;", id)
	return queryStr, nil
}
