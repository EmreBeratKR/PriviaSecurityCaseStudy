package middlewares

import (
	"privia-sec-case-study/frontend/common"

	"github.com/gofiber/fiber/v2"
)

func AuthenticationMiddleware(context *fiber.Ctx) error {
	if context.Path() == "/health" {
		return context.Next()
	}

	isLoginPage := context.Path() == "/login"
	isAuthenticated := common.IsAuthenticated(context)

	if isAuthenticated {
		if isLoginPage {
			return common.RedirectToHomePage(context)
		}

		return context.Next()
	}

	if isLoginPage {
		return context.Next()
	}

	return common.RedirectToLoginPage(context)
}
