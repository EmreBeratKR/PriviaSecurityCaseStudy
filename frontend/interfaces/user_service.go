package interfaces

import "privia-sec-case-study/frontend/models"

type UserService interface {
	Login(request *models.LoginRequestModel) *models.LoginResponseModel
}
