package initializers

import (
	"todo-frontend-web-app/routes"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	routes.MapIndexRoutes(app)
	routes.MapLoginRoutes(app)
}
