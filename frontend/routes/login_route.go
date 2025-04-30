package routes

import (
	"todo-frontend-web-app/controllers"

	"github.com/gofiber/fiber/v2"
)

func MapLoginRoutes(app *fiber.App) {
	var path = "/login"
	app.Get(path, controllers.LoginControllerGet)
	app.Post(path, controllers.LoginControllerPost)
}
