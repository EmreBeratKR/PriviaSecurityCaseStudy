package main

import (
	"todo-frontend-web-app/initializers"
	"todo-frontend-web-app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	serviceManager := services.MockServiceManager()
	engine := html.New("./views", ".tmpl")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	initializers.InitRoutes(app, serviceManager)
	initializers.InitMiddlewares(app)

	app.Listen(":3000")
}
