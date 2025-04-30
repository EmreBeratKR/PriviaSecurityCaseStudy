package services

import "todo-frontend-web-app/models"

type MockUserService struct{}

func (service *MockUserService) Login(request *models.LoginRequestModel) *models.LoginResponseModel {
	if request.Username == "Emre" && request.Password == "1234" {
		return &models.LoginResponseModel{
			Status:  "success",
			Message: "Welcome back, Emre",
		}
	}

	return &models.LoginResponseModel{
		Status:  "error",
		Message: "Wrong credentials",
	}
}
