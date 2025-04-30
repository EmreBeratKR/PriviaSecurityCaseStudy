package middlewares

import "github.com/gofiber/fiber/v2"

func NotFoundMiddleware(context *fiber.Ctx) error {
	return context.Status(fiber.StatusNotFound).Render("not_found", fiber.Map{})
}
