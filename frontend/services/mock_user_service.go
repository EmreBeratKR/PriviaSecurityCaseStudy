package services

import (
	"os"
	"strconv"
	"time"
	"todo-frontend-web-app/common"
	"todo-frontend-web-app/models"

	"github.com/golang-jwt/jwt/v4"
)

type MockUserService struct {
	Users     []models.UserModel
	UserCount int
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

		if common.ComparePasswordAndHash(request.Password, user.Hash) {
			expireAt := calculateJWTExpireTime()
			return &models.LoginResponseModel{
				Status:  "success",
				Message: "Welcome back, " + user.Username,
				Token: createJWT(expireAt, models.UserClaims{
					Username: request.Username,
					Role:     user.Role,
				}, user.Id),
				ExpiresAt: expireAt,
			}
		}
	}

	return &models.LoginResponseModel{
		Status:  "error",
		Message: "Wrong credentials",
	}
}

func (service *MockUserService) createUser(username string, password string, role string) {
	id := strconv.Itoa(service.UserCount)
	service.Users = append(service.Users, models.UserModel{
		Id:       id,
		Username: username,
		Hash:     common.GeneratePasswordHash(password),
		Role:     role,
	})
	service.UserCount += 1
}

func calculateJWTExpireTime() time.Time {
	return time.Now().Add(24 * time.Hour)
}

func getJWTSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

func createJWT(expireAt time.Time, claims models.UserClaims, subject string) string {
	secret := getJWTSecret()
	claims.RegisteredClaims = jwt.RegisteredClaims{
		Subject:   subject,
		Issuer:    "Mock User Service",
		ExpiresAt: jwt.NewNumericDate(expireAt),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secret)

	if err != nil {
		return ""
	}

	return signedToken
}
