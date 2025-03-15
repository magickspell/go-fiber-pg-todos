package tasks

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

// todo
/*
новые запросы в горутинах?
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

type TaskhandlerI interface {
	CreateTask(ctx *fiber.Ctx) error
}

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

func (h *TaskHandler) UpdateTask(ctx *fiber.Ctx) error {
	taskId, err := PutTask(ctx, h.DB)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"message": "error", "err": err.Error()})
	}

	return ctx.Status(200).JSON(fiber.Map{"message": "success", "id": taskId})
}

func (h *TaskHandler) DeleteTask(ctx *fiber.Ctx) error {
	taskId, err := DeleteTask(ctx, h.DB)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"message": "error", "err": err.Error()})
	}

	return ctx.Status(200).JSON(fiber.Map{"message": "success", "id": taskId})
}
