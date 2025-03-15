package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	config "todo-go-fiber/config"
	db "todo-go-fiber/internal/db"
	migrations "todo-go-fiber/internal/db/migrations"
	tasks "todo-go-fiber/internal/features/tasks"
	middlewares "todo-go-fiber/internal/middlewars"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	app := fiber.New()
	app.Use(recover.New())
	app.Use(middlewares.JsonMiddleware)

	config := config.GetConfig()

	conn := db.Connect(config)
	defer conn.Close()
	migrations.Migrate(conn)

	taskHandler := &tasks.TaskHandler{DB: conn}

	app.Get("/", func(c *fiber.Ctx) error {
		return fiber.NewError(782, "Custom error message")
	})

	app.Post("/tasks", func(c *fiber.Ctx) error {
		return taskHandler.CreateTask(c)
	})

	app.Get("/tasks", func(c *fiber.Ctx) error {
		return taskHandler.ReadTask(c)
	})

	app.Put("/tasks/:id", func(c *fiber.Ctx) error {
		return taskHandler.UpdateTask(c)
	})

	app.Delete("/tasks/:id", func(c *fiber.Ctx) error {
		return taskHandler.DeleteTask(c)
	})

	log.Fatal(app.Listen(":3000"))
}
