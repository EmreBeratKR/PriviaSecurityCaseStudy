package main

import (
	"os"
	"privia-sec-case-study/frontend/initializers"
	"privia-sec-case-study/frontend/services"
	"privia-sec-case-study/shared"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func init() {
	shared.InitDotEnv()
}

func main() {
	serviceManager := services.MockServiceManager()
	engine := html.New("./frontend/views", ".tmpl")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./frontend/public")

	initializers.PreUseMiddlewares(app)
	initializers.InitRoutes(app, serviceManager)
	initializers.PostUseMiddlewares(app)

	port := os.Getenv("FRONTEND_PORT")
	app.Listen(":" + port)
}
