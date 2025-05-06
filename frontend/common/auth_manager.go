package common

import (
	"privia-sec-case-study/frontend/models"
	"privia-sec-case-study/shared"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Login(context *fiber.Ctx, response *models.LoginResponseModel) {
	context.Cookie(&fiber.Cookie{
		Name:     getAuthCookieName(),
		Value:    response.Token,
		Expires:  shared.CalculateJWTExpireTime(),
		HTTPOnly: true,
		Secure:   shared.IsProductionEnvironment(),
		SameSite: "Lax",
	})
}

func Logout(context *fiber.Ctx) {
	context.Cookie(&fiber.Cookie{
		Name:     getAuthCookieName(),
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
		Secure:   shared.IsProductionEnvironment(),
		SameSite: "Lax",
	})
}

func IsAuthenticated(context *fiber.Ctx) bool {
	return GetUserClaims(context) != nil
}

func IsAuthenticatedAsAdmin(context *fiber.Ctx) bool {
	claims := GetUserClaims(context)
	if claims == nil {
		return false
	}

	return claims.IsAdmin()
}

func IsAuthorizedForUserId(context *fiber.Ctx, userId string) bool {
	claims := GetUserClaims(context)
	if claims == nil {
		return false
	}

	if claims.IsAdmin() {
		return true
	}

	return claims.Subject == userId
}

func GetAuthUserId(context *fiber.Ctx) string {
	claims := GetUserClaims(context)
	if claims == nil {
		return ""
	}

	return claims.Subject
}

func GetAuthUsername(context *fiber.Ctx) string {
	claims := GetUserClaims(context)
	if claims == nil {
		return ""
	}

	return claims.Username
}

func GetUserClaims(context *fiber.Ctx) *shared.UserClaims {
	tokenStr := context.Cookies(getAuthCookieName())
	if tokenStr == "" {
		return nil
	}

	return shared.GetUserClaims(tokenStr)
}

func getAuthCookieName() string {
	return "auth_token"
}
