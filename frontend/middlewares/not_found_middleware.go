package middlewares

import (
	"todo-frontend-web-app/common"

	"github.com/gofiber/fiber/v2"
)

func NotFoundMiddleware(context *fiber.Ctx) error {
	return common.SendStatusNotFound(context)
}
