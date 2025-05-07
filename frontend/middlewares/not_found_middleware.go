package middlewares

import (
	"privia-sec-case-study/frontend/common"

	"github.com/gofiber/fiber/v2"
)

func NotFoundMiddleware(context *fiber.Ctx) error {
	return common.SendStatusNotFound(context)
}
