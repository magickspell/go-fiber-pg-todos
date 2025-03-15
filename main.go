package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	db "todo-go-fiber/internal/db"
	migrations "todo-go-fiber/internal/db/migrations"
	tasks "todo-go-fiber/internal/features/tasks"
	middlewares "todo-go-fiber/internal/middlewars"
	config "todo-go-fiber/pkg/config"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	app := fiber.New(fiber.Config{
		StrictRouting: true,
	})
	app.Use(middlewares.JsonMiddleware)
	app.Use(recover.New())

	config := config.GetConfig()

	conn := db.Connect(config)
	defer conn.Close()
	migrations.RunMigrations(conn)

	taskHandler := &tasks.TaskHandler{DB: conn}

	app.Get("/panic", func(c *fiber.Ctx) error {
		panic("panic")
	})

	app.Get("/long", func(c *fiber.Ctx) error {
		time.Sleep(5 * time.Second)
		return fiber.NewError(500, "request overtime passed wrong")
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

	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-c
	log.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	log.Println("Running cleanup tasks...")
	conn.Close()
	log.Println("Fiber was successful shutdown.")
}
