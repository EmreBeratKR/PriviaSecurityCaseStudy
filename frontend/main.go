package main

import (
	"todo-frontend-web-app/initializers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".tmpl")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	initializers.InitRoutes(app)
	initializers.InitMiddlewares(app)

	app.Listen(":3000")
}
