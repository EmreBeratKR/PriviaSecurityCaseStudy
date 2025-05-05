package main

import (
	"os"
	"privia-sec-case-study/backend/internal/handler"
	"privia-sec-case-study/backend/internal/repository"
	"privia-sec-case-study/backend/internal/usercase"
	"privia-sec-case-study/shared"

	"github.com/gofiber/fiber/v2"
)

func init() {
	shared.InitDotEnv()
}

func main() {
	app := fiber.New()

	userRepository := repository.NewMockUserRepository()
	userUsecase := usercase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	app.Get("/users/login", userHandler.LoginUser)

	port := os.Getenv("BACKEND_PORT")
	app.Listen(":" + port)
}
