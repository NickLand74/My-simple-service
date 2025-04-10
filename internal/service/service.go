package service

import (
	"my-todo-app/dto"
	"my-todo-app/internal/repo"

	"github.com/gofiber/fiber/v2"
)

type Service struct {
	repo *repo.Repository
}

func NewService(repo *repo.Repository) *Service {
	return &Service{repo: repo}
}

// CreateTask - обработчик создания задачи
func (s *Service) CreateTask(ctx *fiber.Ctx) error {
	var req dto.TaskRequest
	if err := ctx.BodyParser(&req); err != nil {
		return dto.BadRequest(ctx, "Invalid request body")
	}

	task := repo.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}
	id, err := s.repo.CreateTask(task)
	if err != nil {
		return dto.InternalServerError(ctx, "Failed to create task")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   fiber.Map{"task_id": id},
	})
}

// GetAllTasks - обработчик получения списка задач
func (s *Service) GetAllTasks(ctx *fiber.Ctx) error {
	tasks, err := s.repo.GetAllTasks()
	if err != nil {
		return dto.InternalServerError(ctx, "Failed to get tasks")
	}
	return ctx.JSON(tasks)
}

// GetTask - обработчик получения задачи по ID
func (s *Service) GetTask(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return dto.BadRequest(ctx, "Invalid task ID")
	}

	task, err := s.repo.GetTask(id)
	if err != nil {
		return dto.NotFound(ctx, "Task not found")
	}
	return ctx.JSON(task)
}

// UpdateTask - обработчик обновления задачи
func (s *Service) UpdateTask(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return dto.BadRequest(ctx, "Invalid task ID")
	}

	var req dto.TaskRequest
	if err := ctx.BodyParser(&req); err != nil {
		return dto.BadRequest(ctx, "Invalid request body")
	}

	err = s.repo.UpdateTask(id, repo.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	})
	if err != nil {
		return dto.NotFound(ctx, "Task not found")
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   fiber.Map{"task_id": id},
	})
}

// DeleteTask - обработчик удаления задачи
func (s *Service) DeleteTask(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return dto.BadRequest(ctx, "Invalid task ID")
	}

	err = s.repo.DeleteTask(id)
	if err != nil {
		return dto.NotFound(ctx, "Task not found")
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   fiber.Map{"deleted_id": id},
	})
}
