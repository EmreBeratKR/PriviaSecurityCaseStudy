package initializers

import (
	"todo-frontend-web-app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func PreUseMiddlewares(app *fiber.App) {
	app.Use(middlewares.AuthMiddleware)
}

func PostUseMiddlewares(app *fiber.App) {
	app.Use(middlewares.NotFoundMiddleware)
}
