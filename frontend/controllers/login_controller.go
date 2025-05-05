package controllers

import (
	"todo-frontend-web-app/common"
	"todo-frontend-web-app/models"
	"todo-frontend-web-app/services"

	"github.com/gofiber/fiber/v2"
)

type LoginController struct {
	ServiceManager *services.ServiceManager
}

func (controller *LoginController) LoginControllerGet(context *fiber.Ctx) error {
	controller.ServiceManager.SetContext(context)

	return context.Render("login", fiber.Map{
		"HideUserInfo": true,
	})
}

func (controller *LoginController) LoginControllerPost(context *fiber.Ctx) error {
	controller.ServiceManager.SetContext(context)

	loginRequest, parseSuccess := controller.tryParseLoginRequest(context)

	if !parseSuccess {
		return common.SendStatusBadRequest(context)
	}

	if !loginRequest.IsValid() {
		return common.SendStatusBadRequest(context)
	}

	response := controller.sendLoginRequest(loginRequest)

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

func (controller *LoginController) sendLoginRequest(request *models.LoginRequestModel) *models.LoginResponseModel {
	return controller.ServiceManager.UserService.Login(request)
}
