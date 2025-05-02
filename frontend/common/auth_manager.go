package common

import (
	"os"
	"todo-frontend-web-app/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Login(context *fiber.Ctx, response *models.LoginResponseModel) {
	context.Cookie(&fiber.Cookie{
		Name:     getAuthCookieName(),
		Value:    response.Token,
		Expires:  response.ExpiresAt,
		HTTPOnly: true,
		Secure:   os.Getenv("ENVIRONMENT") == "prod",
		SameSite: "Lax",
	})
}

func IsAuthenticated(context *fiber.Ctx) bool {
	return getUserClaims(context) != nil
}

func GetAuthUserId(context *fiber.Ctx) string {
	claims := getUserClaims(context)
	if claims == nil {
		return ""
	}

	return claims.Subject
}

func GetAuthUsername(context *fiber.Ctx) string {
	claims := getUserClaims(context)
	if claims == nil {
		return ""
	}

	return claims.Username
}

func getAuthCookieName() string {
	return "auth_token"
}

func getUserClaims(context *fiber.Ctx) *models.UserClaims {
	tokenStr := context.Cookies(getAuthCookieName())
	if tokenStr == "" {
		return nil
	}

	token, err := jwt.ParseWithClaims(tokenStr, &models.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return nil
	}

	claims, ok := token.Claims.(*models.UserClaims)
	if !ok {
		return nil
	}

	return claims
}
