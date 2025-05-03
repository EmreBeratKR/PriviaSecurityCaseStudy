package services

import (
	"todo-frontend-web-app/interfaces"
)

type ServiceManager struct {
	UserService     interfaces.UserService
	TodoListService interfaces.TodoListService
	TodoTaskService interfaces.TodoTaskService
}

func MockServiceManager() *ServiceManager {
	userService := &MockUserService{}
	userService.Init()
	todoListService := &MockTodoListService{}
	todoListService.Init()
	todoTaskService := &MockTodoTaskService{
		TodoListService: todoListService,
	}
	todoTaskService.Init()
	return &ServiceManager{
		UserService:     userService,
		TodoListService: todoListService,
		TodoTaskService: todoTaskService,
	}
}
