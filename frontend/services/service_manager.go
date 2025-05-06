package services

import (
	"os"
	"privia-sec-case-study/frontend/interfaces"
)

type ServiceManager struct {
	UserService     interfaces.UserService
	TodoListService interfaces.TodoListService
	TodoTaskService interfaces.TodoTaskService
}

func CreateServiceManager() *ServiceManager {
	serviceManager := &ServiceManager{}

	/* userService := &MockUserService{
		ServiceManager: serviceManager,
	}
	userService.Init() */
	apiUrl := os.Getenv("API_URL")
	userService := NewApiUserService(apiUrl)
	todoListService := NewApiTodoListService(apiUrl)
	todoTaskService := NewApiTodoTaskService(apiUrl)

	/* todoListService := &MockTodoListService{
		ServiceManager: serviceManager,
	}
	todoListService.Init() */

	/* todoTaskService := &MockTodoTaskService{
		ServiceManager:  serviceManager,
		TodoListService: todoListService,
	}
	todoTaskService.Init() */

	serviceManager.UserService = userService
	serviceManager.TodoListService = todoListService
	serviceManager.TodoTaskService = todoTaskService

	return serviceManager
}
