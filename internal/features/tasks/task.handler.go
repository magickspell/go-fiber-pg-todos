package tasks

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type TaskHandlerI interface {
	CreateTask(ctx *fiber.Ctx) error
	ReadTask(ctx *fiber.Ctx) error
	UpdateTask(ctx *fiber.Ctx) error
	DeleteTask(ctx *fiber.Ctx) error
}

type TaskHandler struct {
	DB *sql.DB
}

func (h *TaskHandler) CreateTask(ctx *fiber.Ctx) error {
	taskId, err := CreateTask(ctx, h.DB)
	return handleResponse(ctx, taskId, err)
}

func (h *TaskHandler) ReadTask(ctx *fiber.Ctx) error {
	tasks, err := GetTask(ctx, h.DB)
	return handleResponse(ctx, tasks, err)
}

func (h *TaskHandler) UpdateTask(ctx *fiber.Ctx) error {
	taskId, err := PutTask(ctx, h.DB)
	return handleResponse(ctx, taskId, err)
}

func (h *TaskHandler) DeleteTask(ctx *fiber.Ctx) error {
	taskId, err := DeleteTask(ctx, h.DB)
	return handleResponse(ctx, taskId, err)
}

func handleResponse(ctx *fiber.Ctx, result any, err error) error {
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"message": "error", "err": err.Error()})
	}

	return ctx.Status(200).JSON(fiber.Map{"message": "success", "result": result})
}
