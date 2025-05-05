package initializers

import (
	"privia-sec-case-study/frontend/middlewares"
	"privia-sec-case-study/shared"

	"github.com/gofiber/fiber/v2"
)

func PreUseMiddlewares(app *fiber.App) {
	if shared.IsDevelopmentEnvironment() {
		app.Use(middlewares.LoggerMiddleware)
	}
	app.Use(middlewares.AuthenticationMiddleware)
}

func PostUseMiddlewares(app *fiber.App) {
	app.Use(middlewares.NotFoundMiddleware)
}
