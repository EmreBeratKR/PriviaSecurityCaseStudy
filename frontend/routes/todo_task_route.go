package routes

import (
	"privia-sec-case-study/frontend/controllers"
	"privia-sec-case-study/frontend/services"

	"github.com/gofiber/fiber/v2"
)

func MapTodoTaskRoutes(app *fiber.App, serviceManager *services.ServiceManager) {
	var path = "todo-task"
	var controller = &controllers.TodoTaskController{
		ServiceManager: serviceManager,
	}

	app.Post(path, controller.TodoTaskControllerPost)
	app.Post(path+"/patch", controller.TodoTaskControllerPatch)
	app.Post(path+"/delete", controller.TodoTaskControllerDelete)
}
