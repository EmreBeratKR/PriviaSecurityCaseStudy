package controllers

import (
	"todo-frontend-web-app/common"
	"todo-frontend-web-app/services"

	"github.com/gofiber/fiber/v2"
)

type LogoutController struct {
	ServiceManager *services.ServiceManager
}

func (controller *LogoutController) LogoutControllerPost(context *fiber.Ctx) error {
	common.Logout(context)
	return common.RedirectToLoginPage(context)
}
