package initializers

import (
	"privia-sec-case-study/frontend/routes"
	"privia-sec-case-study/frontend/services"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App, serviceManager *services.ServiceManager) {
	routes.MapHealthRoutes(app)
	routes.MapIndexRoutes(app, serviceManager)
	routes.MapLoginRoutes(app, serviceManager)
	routes.MapLogoutRoutes(app, serviceManager)
	routes.MapTodoListRoutes(app, serviceManager)
	routes.MapTodoTaskRoutes(app, serviceManager)
}
