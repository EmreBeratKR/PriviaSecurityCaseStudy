package routes

import (
	"todo-frontend-web-app/controllers"
	"todo-frontend-web-app/services"

	"github.com/gofiber/fiber/v2"
)

func MapLogoutRoutes(app *fiber.App, serviceManager *services.ServiceManager) {
	var path = "/logout"
	var controller = &controllers.LogoutController{
		ServiceManager: serviceManager,
	}

	app.Post(path, controller.LogoutControllerPost)
}
