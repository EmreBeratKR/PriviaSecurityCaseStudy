package routes

import (
	"todo-frontend-web-app/controllers"
	"todo-frontend-web-app/services"

	"github.com/gofiber/fiber/v2"
)

func MapTodoTaskRoutes(app *fiber.App, serviceManager *services.ServiceManager) {
	var path = "todo-task"
	var controller = &controllers.TodoTaskController{
		ServiceManager: serviceManager,
	}

	app.Post(path, controller.TodoTaskControllerPost)
	app.Post(path+"/delete", controller.TodoTaskControllerDelete)
	app.Post(path+"/toggle", controller.TodoTaskControllerToggle)
}
