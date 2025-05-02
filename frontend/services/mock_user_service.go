package services

import (
	"os"
	"time"
	"todo-frontend-web-app/models"

	"github.com/golang-jwt/jwt/v4"
)

type MockUserService struct{}

func (service *MockUserService) Login(request *models.LoginRequestModel) *models.LoginResponseModel {
	if request.Username == "Emre" && request.Password == "1234" {
		userId := "1234567890"
		expireAt := calculateJWTExpireTime()
		return &models.LoginResponseModel{
			Status:  "success",
			Message: "Welcome back, Emre",
			Token: createJWT(expireAt, models.UserClaims{
				Username: request.Username,
				Role:     "user",
			}, userId),
			ExpiresAt: expireAt,
		}
	}

	return &models.LoginResponseModel{
		Status:  "error",
		Message: "Wrong credentials",
	}
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
