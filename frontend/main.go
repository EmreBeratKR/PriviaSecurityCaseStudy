package main

import (
	"os"
	"todo-frontend-web-app/initializers"
	"todo-frontend-web-app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func init() {
	initializers.InitDotEnv()
}

func main() {
	serviceManager := services.MockServiceManager()
	engine := html.New("./views", ".tmpl")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	initializers.PreUseMiddlewares(app)
	initializers.InitRoutes(app, serviceManager)
	initializers.PostUseMiddlewares(app)

	port := os.Getenv("PORT")
	app.Listen(":" + port)
}
