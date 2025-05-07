package services

import (
	"os"
	"privia-sec-case-study/frontend/interfaces"
	"privia-sec-case-study/shared"
)

type ServiceManager struct {
	UserService     interfaces.UserService
	TodoListService interfaces.TodoListService
	TodoTaskService interfaces.TodoTaskService
}

func NewServiceManager() *ServiceManager {
	if shared.IsDevelopmentEnvironment() {
		return newServiceManagerWithMockServices()
	}

	return newServiceManagerWithApiServices()
}

func newServiceManagerWithMockServices() *ServiceManager {
	serviceManager := &ServiceManager{}

	userService := NewMockUserService(serviceManager)
	todoListService := NewMockTodoListService(serviceManager)
	todoTaskService := NewMockTodoTaskService(serviceManager, todoListService)

	serviceManager.UserService = userService
	serviceManager.TodoListService = todoListService
	serviceManager.TodoTaskService = todoTaskService

	return serviceManager
}

func newServiceManagerWithApiServices() *ServiceManager {
	serviceManager := &ServiceManager{}

	apiUrl := os.Getenv("API_URL")
	userService := NewApiUserService(apiUrl)
	todoListService := NewApiTodoListService(apiUrl)
	todoTaskService := NewApiTodoTaskService(apiUrl)

	serviceManager.UserService = userService
	serviceManager.TodoListService = todoListService
	serviceManager.TodoTaskService = todoTaskService

	return serviceManager
}
