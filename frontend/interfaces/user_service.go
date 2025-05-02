package interfaces

import "todo-frontend-web-app/models"

type UserService interface {
	Login(request *models.LoginRequestModel) *models.LoginResponseModel
}
