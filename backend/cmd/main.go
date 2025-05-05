package main

import (
	"todo-backend-rest-api/internal/handler"
	"todo-backend-rest-api/internal/repository"
	"todo-backend-rest-api/internal/usercase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	userRepository := repository.NewMockUserRepository()
	userUsecase := usercase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	app.Get("/users/login", userHandler.LoginUser)

	app.Listen(":3000")
}
