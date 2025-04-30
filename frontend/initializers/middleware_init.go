package initializers

import (
	"todo-frontend-web-app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func InitMiddlewares(app *fiber.App) {
	app.Use(middlewares.NotFoundMiddleware)
}
