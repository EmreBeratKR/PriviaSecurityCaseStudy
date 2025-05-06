package common

import (
	"privia-sec-case-study/backend/internal/domain"
	"privia-sec-case-study/shared"

	"github.com/gofiber/fiber/v2"
)

func SendStatusOkWithValue(value any, context *fiber.Ctx) error {
	return context.Status(fiber.StatusOK).JSON(domain.ValueResponseOk(value))
}

func SendErrorStatus(status string, message string, context *fiber.Ctx) error {
	if status == shared.StatusBadRequest().Status {
		return SendStatusBadRequest(context, message)
	}

	if status == shared.StatusUnauthorized().Status {
		return SendStatusUnauthorized(context, message)
	}

	if status == shared.StatusForbidden().Status {
		return SendStatusForbidden(context, message)
	}

	if status == shared.StatusNotFound().Status {
		return SendStatusNotFound(context, message)
	}

	if status == shared.StatusInternalServerError().Status {
		return SendStatusInternalServerError(context, message)
	}

	return SendStatusInternalServerError(context, message)
}

func SendStatusBadRequest(context *fiber.Ctx, message string) error {
	return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": message,
	})
}

func SendStatusUnauthorized(context *fiber.Ctx, message string) error {
	return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": message,
	})
}

func SendStatusForbidden(context *fiber.Ctx, message string) error {
	return context.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"message": message,
	})
}

func SendStatusNotFound(context *fiber.Ctx, message string) error {
	return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": message,
	})
}

func SendStatusInternalServerError(context *fiber.Ctx, message string) error {
	return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": message,
	})
}
