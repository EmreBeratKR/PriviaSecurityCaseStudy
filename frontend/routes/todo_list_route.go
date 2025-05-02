package routes

import (
	"todo-frontend-web-app/controllers"
	"todo-frontend-web-app/services"

	"github.com/gofiber/fiber/v2"
)

func MapTodoListRoutes(app *fiber.App, serviceManager *services.ServiceManager) {
	var path = "/todo-list"
	var controller = &controllers.TodoListController{
		ServiceManager: serviceManager,
	}

	app.Get(path, controller.TodoListControllerGet)
}
