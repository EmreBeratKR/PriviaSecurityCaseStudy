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
	todoListService := &MockTodoListService{}
	todoListService.Init()
	todoTaskService := &MockTodoTaskService{
		TodoListService: todoListService,
	}
	todoTaskService.Init()
	return &ServiceManager{
		UserService:     &MockUserService{},
		TodoListService: todoListService,
		TodoTaskService: todoTaskService,
	}
}
