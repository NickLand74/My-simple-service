package dto

import "github.com/gofiber/fiber/v2"

type TaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Status      string `json:"status" validate:"oneof=pending completed"`
}

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

// Обработка ошибок
func BadRequest(ctx *fiber.Ctx, msg string) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(Response{
		Error: struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		}{
			Code:    "BAD_REQUEST",
			Message: msg,
		},
	})
}

func NotFound(ctx *fiber.Ctx, msg string) error {
	return ctx.Status(fiber.StatusNotFound).JSON(Response{
		Error: struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		}{
			Code:    "NOT_FOUND",
			Message: msg,
		},
	})
}

func InternalServerError(ctx *fiber.Ctx, msg string) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(Response{
		Error: struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		}{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: msg,
		},
	})
}
