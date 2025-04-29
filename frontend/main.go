package main

import (
	"todo-frontend-web-app/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".tmpl")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", controllers.Hello)
	app.Get("/login", controllers.Login)
	app.Post("/login", controllers.SubmitLogin)

	app.Listen(":3000")
}
