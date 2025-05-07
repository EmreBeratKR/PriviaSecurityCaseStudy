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
	viewsPath := os.Getenv("FRONTEND_VIEWS_PATH")
	serviceManager := services.NewServiceManager()
	engine := html.New(viewsPath, ".tmpl")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	publicPath := os.Getenv("FRONTEND_PUBLIC_PATH")
	app.Static("/", publicPath)

	initializers.PreUseMiddlewares(app)
	initializers.InitRoutes(app, serviceManager)
	initializers.PostUseMiddlewares(app)

	port := os.Getenv("FRONTEND_PORT")
	app.Listen(":" + port)
}
