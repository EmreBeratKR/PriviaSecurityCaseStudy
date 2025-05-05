package services

import (
	"todo-frontend-web-app/interfaces"

	"github.com/gofiber/fiber/v2"
)

type ServiceManager struct {
	UserService     interfaces.UserService
	TodoListService interfaces.TodoListService
	TodoTaskService interfaces.TodoTaskService
	Context         *fiber.Ctx
}

func MockServiceManager() *ServiceManager {
	serviceManager := &ServiceManager{}

	userService := &MockUserService{
		ServiceManager: serviceManager,
	}
	userService.Init()

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
