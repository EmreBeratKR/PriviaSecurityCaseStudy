package controllers

import (
	"privia-sec-case-study/frontend/common"
	"privia-sec-case-study/frontend/services"

	"github.com/gofiber/fiber/v2"
)

type LogoutController struct {
	ServiceManager *services.ServiceManager
}

func (controller *LogoutController) LogoutControllerPost(context *fiber.Ctx) error {
	common.Logout(context)
	return common.RedirectToLoginPage(context)
}
