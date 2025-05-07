package router

import "github.com/gofiber/fiber/v2"

func MapHealthRouter(app *fiber.App) {
	path := "/health"
	app.Get(path, func(c *fiber.Ctx) error {
		return c.JSON("healthy")
	})
}
