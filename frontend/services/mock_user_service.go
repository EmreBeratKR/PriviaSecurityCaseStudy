package services

import (
	"privia-sec-case-study/frontend/models"
	"privia-sec-case-study/shared"
	"strconv"
)

type MockUserService struct {
	ServiceManager *ServiceManager
	Users          []models.UserModel
	UserCount      int
}

func (service *MockUserService) Init() {
	service.UserCount = 0
	service.createUser("Emre", "1234", "user")
	service.createUser("Berat", "1234", "admin")
}

func (service *MockUserService) Login(request *models.LoginRequestModel) *models.LoginResponseModel {
	for _, user := range service.Users {
		if user.Username != request.Username {
			continue
		}

		if shared.ComparePasswordAndHash(request.Password, user.Hash) {
			return &models.LoginResponseModel{
				StatusModel: shared.StatusSuccess(),
				Token: shared.CreateJWT(shared.UserClaims{
					Username: request.Username,
					Role:     user.Role,
				}, user.Id),
			}
		}
	}

	return &models.LoginResponseModel{
		StatusModel: shared.StatusUnauthorized(),
		Message:     "Wrong credentials",
	}
}

func (service *MockUserService) createUser(username string, password string, role string) {
	id := strconv.Itoa(service.UserCount)
	service.Users = append(service.Users, models.UserModel{
		Id:       id,
		Username: username,
		Hash:     shared.GeneratePasswordHash(password),
		Role:     role,
	})
	service.UserCount += 1
}
