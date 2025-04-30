package routes

import (
	"todo-frontend-web-app/controllers"
	"todo-frontend-web-app/services"

	"github.com/gofiber/fiber/v2"
)

func MapLoginRoutes(app *fiber.App, serviceManger *services.ServiceManager) {
	var path = "/login"
	var controller = &controllers.LoginController{
		ServiceManager: serviceManger,
	}

	app.Get(path, controller.LoginControllerGet)
	app.Post(path, controller.LoginControllerPost)
}
