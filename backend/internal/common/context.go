package common

import (
	"privia-sec-case-study/shared"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetClaimsFromHeaders(context *fiber.Ctx) (*shared.UserClaims, error) {
	authHeader := context.Get("Authorization")
	if authHeader == "" {
		return nil, SendStatusUnauthorized(context, "Authorization header is required")
	}

	authHeaderSplits := strings.Split(authHeader, " ")
	if len(authHeaderSplits) < 2 {
		return nil, SendStatusUnauthorized(context, "Unproccessable authorization header")
	}

	authType := authHeaderSplits[0]
	if authType != "Bearer" {
		return nil, SendStatusUnauthorized(context, "Bearer authorization header required")
	}

	jwtToken := authHeaderSplits[1]
	if jwtToken == "" {
		return nil, SendStatusUnauthorized(context, "JWT token is required")
	}

	claims := shared.GetUserClaims(jwtToken)
	if claims == nil {
		return nil, SendStatusUnauthorized(context, "Invalid JWT token")
	}

	return claims, nil
}

func SendStatusOkWithValue(value any, context *fiber.Ctx) error {
	return context.Status(fiber.StatusOK).JSON(ValueResponseOk(value))
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
