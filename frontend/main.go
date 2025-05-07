package main

import (
	"fmt"
	"os"
	"path/filepath"
	"privia-sec-case-study/frontend/initializers"
	"privia-sec-case-study/frontend/services"
	"privia-sec-case-study/shared"
	"strings"

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

	root := "."

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && strings.HasPrefix(path, ".git") {
			return filepath.SkipDir
		}
		fmt.Println(path)
		return nil
	})

	if err != nil {
		fmt.Println("Error walking the path:", err)
	}

	port := os.Getenv("FRONTEND_PORT")
	app.Listen(":" + port)
}
