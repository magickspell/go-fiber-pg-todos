package tasks

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

// todo
/*
валидаторы
миграции
тразакции
документация
грейсфулшотдаун
докер
горм?
демка скрин
vendor
air
*/

type TaskHandler struct {
	DB *sql.DB
}

func (h *TaskHandler) CreateTask(ctx *fiber.Ctx) error {
	id, err := CreateTask(ctx, h.DB)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"message": "error", "err": err.Error()})
	}

	return ctx.Status(200).JSON(fiber.Map{"message": "success", "id": id})
}

func (h *TaskHandler) ReadTask(ctx *fiber.Ctx) error {
	tasks, err := GetTask(ctx, h.DB)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"message": "error", "err": err.Error()})
	}

	return ctx.Status(200).JSON(fiber.Map{"message": "success", "tasks": tasks})
}

func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
	// id := c.Params("id")
	return nil
}

func (h *TaskHandler) DeleteTask(c *fiber.Ctx) error {
	return nil
}
