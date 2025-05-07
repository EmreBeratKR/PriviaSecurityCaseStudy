package routes

import "github.com/gofiber/fiber/v2"

func MapHealthRoutes(app *fiber.App) {
	var path = "/health"

	app.Get(path, func(c *fiber.Ctx) error {
		return c.JSON("healthy")
	})
}
