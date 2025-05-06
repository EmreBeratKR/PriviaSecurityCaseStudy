package main

import (
	"os"
	"privia-sec-case-study/backend/internal/handler/concrete_handlers"
	"privia-sec-case-study/backend/internal/repository/concrete_repositories"
	"privia-sec-case-study/backend/internal/router"
	"privia-sec-case-study/backend/internal/usecase/concrete_usecases"
	"privia-sec-case-study/shared"

	"github.com/gofiber/fiber/v2"
)

func init() {
	shared.InitDotEnv()
}

func main() {
	app := fiber.New()

	userRepository := concrete_repositories.NewMockUserRepository()
	userUsecase := concrete_usecases.NewDefaultUserUsecase(userRepository)
	userHandler := concrete_handlers.NewDefaultUserHandler(userUsecase)

	todoListRepository := concrete_repositories.NewMockTodoListRepository()
	todoListUsecase := concrete_usecases.NewDefaultTodoListUsecase(todoListRepository)
	todoListHandler := concrete_handlers.NewDefaultTodoListHandler(todoListUsecase)

	todoTaskRepository := concrete_repositories.NewMockTodoTaskRepository()
	todoTaskUsecase := concrete_usecases.NewDefaultTodoTaskUsecase(todoListRepository, todoTaskRepository)
	todoTaskHandler := concrete_handlers.NewDefaultTodoTaskHandler(todoListUsecase, todoTaskUsecase)

	router.MapUserRoutes(app, userHandler)
	router.MapTodoListRouter(app, todoListHandler)
	router.MapTodoTaskRoutes(app, todoTaskHandler)

	port := os.Getenv("BACKEND_PORT")
	app.Listen(":" + port)
}
