package routes

import (
	"privia-sec-case-study/frontend/controllers"
	"privia-sec-case-study/frontend/services"

	"github.com/gofiber/fiber/v2"
)

func MapIndexRoutes(app *fiber.App, serviceManager *services.ServiceManager) {
	var path = "/"
	var controller = &controllers.IndexController{
		ServiceManager: serviceManager,
	}

	app.Get(path, controller.IndexControllerGet)
}
