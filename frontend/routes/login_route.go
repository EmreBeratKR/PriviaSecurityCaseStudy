package routes

import (
	"todo-frontend-web-app/controllers"
	"todo-frontend-web-app/services"

	"github.com/gofiber/fiber/v2"
)

func MapLoginRoutes(app *fiber.App, serviceManager *services.ServiceManager) {
	var path = "/login"
	var controller = &controllers.LoginController{
		ServiceManager: serviceManager,
	}

	app.Get(path, controller.LoginControllerGet)
	app.Post(path, controller.LoginControllerPost)
}
