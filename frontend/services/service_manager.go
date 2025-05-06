package services

import (
	"os"
	"privia-sec-case-study/frontend/interfaces"

	"github.com/gofiber/fiber/v2"
)

type ServiceManager struct {
	UserService     interfaces.UserService
	TodoListService interfaces.TodoListService
	TodoTaskService interfaces.TodoTaskService
	Context         *fiber.Ctx
}

func CreateServiceManager() *ServiceManager {
	serviceManager := &ServiceManager{}

	/* userService := &MockUserService{
		ServiceManager: serviceManager,
	}
	userService.Init() */
	apiUrl := os.Getenv("API_URL")
	userService := NewApiUserService(apiUrl)

	todoListService := &MockTodoListService{
		ServiceManager: serviceManager,
	}
	todoListService.Init()

	todoTaskService := &MockTodoTaskService{
		ServiceManager:  serviceManager,
		TodoListService: todoListService,
	}
	todoTaskService.Init()

	serviceManager.UserService = userService
	serviceManager.TodoListService = todoListService
	serviceManager.TodoTaskService = todoTaskService

	return serviceManager
}

func (manager *ServiceManager) SetContext(context *fiber.Ctx) {
	manager.Context = context
}
