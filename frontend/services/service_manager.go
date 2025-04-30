package services

import (
	"todo-frontend-web-app/interfaces"
)

type ServiceManager struct {
	UserService interfaces.UserService
}

func MockServiceManager() *ServiceManager {
	return &ServiceManager{
		UserService: &MockUserService{},
	}
}
