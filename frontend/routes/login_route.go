package routes

import (
	"privia-sec-case-study/frontend/controllers"
	"privia-sec-case-study/frontend/services"

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
