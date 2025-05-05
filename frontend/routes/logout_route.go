package routes

import (
	"privia-sec-case-study/frontend/controllers"
	"privia-sec-case-study/frontend/services"

	"github.com/gofiber/fiber/v2"
)

func MapLogoutRoutes(app *fiber.App, serviceManager *services.ServiceManager) {
	var path = "/logout"
	var controller = &controllers.LogoutController{
		ServiceManager: serviceManager,
	}

	app.Post(path, controller.LogoutControllerPost)
}
