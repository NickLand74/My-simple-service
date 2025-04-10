package api

import (
	"my-todo-app/internal/service"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(service *service.Service) *fiber.App {
	app := fiber.New()

	apiGroup := app.Group("/api/v1")

	// Создание задачи
	apiGroup.Post("/tasks", service.CreateTask)

	// Получение всех задач
	apiGroup.Get("/tasks", service.GetAllTasks)

	// CRUD по ID
	taskGroup := apiGroup.Group("/tasks/:id")
	taskGroup.Get("", service.GetTask)
	taskGroup.Put("", service.UpdateTask)
	taskGroup.Delete("", service.DeleteTask)

	return app
}
