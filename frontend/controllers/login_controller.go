package controllers

import (
	"privia-sec-case-study/frontend/common"
	"privia-sec-case-study/frontend/models"
	"privia-sec-case-study/frontend/services"

	"github.com/gofiber/fiber/v2"
)

type LoginController struct {
	ServiceManager *services.ServiceManager
}

func (controller *LoginController) LoginControllerGet(context *fiber.Ctx) error {
	return context.Render("login", fiber.Map{
		"HideUserInfo": true,
	})
}

func (controller *LoginController) LoginControllerPost(context *fiber.Ctx) error {
	loginRequest, parseSuccess := controller.tryParseLoginRequest(context)

	if !parseSuccess {
		return common.SendStatusBadRequest(context)
	}

	if !loginRequest.IsValid() {
		return common.SendStatusBadRequest(context)
	}

	response := controller.sendLoginRequest(context, loginRequest)

	if response.IsNotSuccess() {
		return context.Render("login", fiber.Map{
			"HideUserInfo": true,
			"Username":     loginRequest.Username,
			"Error":        response.Message,
		})
	}

	common.Login(context, response)

	return common.RedirectToHomePage(context)
}

func (controller *LoginController) tryParseLoginRequest(context *fiber.Ctx) (*models.LoginRequestModel, bool) {
	var loginRequest models.LoginRequestModel

	if err := context.BodyParser(&loginRequest); err != nil {
		return nil, false
	}

	return &loginRequest, true
}

func (controller *LoginController) sendLoginRequest(context *fiber.Ctx, request *models.LoginRequestModel) *models.LoginResponseModel {
	return controller.ServiceManager.UserService.Login(context, request)
}
