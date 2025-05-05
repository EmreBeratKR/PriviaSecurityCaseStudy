package common

import (
	"todo-backend-rest-api/internal/domain"

	"github.com/gofiber/fiber/v2"
)

func SendStatusOkWithValue(value any, context *fiber.Ctx) error {
	return context.Status(fiber.StatusBadRequest).JSON(domain.ValueResponseOk(value))
}

func SendErrorStatus(status string, context *fiber.Ctx) error {
	if status == domain.StatusBadRequest().Status {
		return SendStatusBadRequest(context)
	}

	if status == domain.StatusForbidden().Status {
		return SendStatusForbidden(context)
	}

	if status == domain.StatusNotFound().Status {
		return SendStatusNotFound(context)
	}

	if status == domain.StatusInternalServerError().Status {
		return SendStatusInternalServerError(context)
	}

	return SendStatusInternalServerError(context)
}

func SendStatusBadRequest(context *fiber.Ctx) error {
	return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": "400 - Bad Request",
	})
}

func SendStatusUnauthorized(context *fiber.Ctx) error {
	return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "401 - Unauthorized",
	})
}

func SendStatusForbidden(context *fiber.Ctx) error {
	return context.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"message": "403 - Forbidden",
	})
}

func SendStatusNotFound(context *fiber.Ctx) error {
	return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "404 - Not Found",
	})
}

func SendStatusInternalServerError(context *fiber.Ctx) error {
	return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "500 - Internal Server Error",
	})
}
