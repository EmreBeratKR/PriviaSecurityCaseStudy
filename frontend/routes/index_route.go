package routes

import (
	"todo-frontend-web-app/controllers"
	"todo-frontend-web-app/services"

	"github.com/gofiber/fiber/v2"
)

func MapIndexRoutes(app *fiber.App, serviceManager *services.ServiceManager) {
	var path = "/"
	var controller = &controllers.IndexController{
		ServiceManager: serviceManager,
	}

	app.Get(path, controller.IndexControllerGet)
}
