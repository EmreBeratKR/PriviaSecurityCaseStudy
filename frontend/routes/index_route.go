package routes

import (
	"todo-frontend-web-app/controllers"

	"github.com/gofiber/fiber/v2"
)

func MapIndexRoutes(app *fiber.App) {
	var path = "/"
	app.Get(path, controllers.IndexControllerGet)
}
