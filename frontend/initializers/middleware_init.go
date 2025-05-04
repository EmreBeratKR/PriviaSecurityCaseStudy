package initializers

import (
	"todo-frontend-web-app/common"
	"todo-frontend-web-app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func PreUseMiddlewares(app *fiber.App) {
	if common.IsDevelopmentEnvironment() {
		app.Use(middlewares.LoggerMiddleware)
	}
	app.Use(middlewares.AuthenticationMiddleware)
}

func PostUseMiddlewares(app *fiber.App) {
	app.Use(middlewares.NotFoundMiddleware)
}
