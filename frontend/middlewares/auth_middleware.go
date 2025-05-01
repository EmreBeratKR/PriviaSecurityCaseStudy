package middlewares

import (
	"todo-frontend-web-app/common"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(context *fiber.Ctx) error {
	isLoginPage := context.Path() == "/login"
	isAuthenticated := common.IsAuthenticated(context)

	if isAuthenticated {
		if isLoginPage {
			return context.Redirect("/")
		}

		return context.Next()
	}

	if isLoginPage {
		return context.Next()
	}

	return context.Redirect("/login")
}
