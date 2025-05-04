package initializers

import (
	"todo-frontend-web-app/routes"
	"todo-frontend-web-app/services"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App, serviceManager *services.ServiceManager) {
	routes.MapIndexRoutes(app, serviceManager)
	routes.MapLoginRoutes(app, serviceManager)
	routes.MapLogoutRoutes(app, serviceManager)
	routes.MapTodoListRoutes(app, serviceManager)
	routes.MapTodoTaskRoutes(app, serviceManager)
}
