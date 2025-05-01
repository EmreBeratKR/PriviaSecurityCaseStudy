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
	return context.Render("login", fiber.Map{})
}

func (controller *LoginController) LoginControllerPost(context *fiber.Ctx) error {
	loginRequest, parseSuccess := controller.tryParseLoginRequest(context)

	if !parseSuccess {
		return context.Status(fiber.StatusBadRequest).Render("bad_request", fiber.Map{})
	}

	response := controller.sendLoginRequest(loginRequest)

	if !response.IsSuccess() {
		return context.Render("login", fiber.Map{
			"Error": response.Message,
		})
	}

	common.Login(context, response)

	return context.Redirect("/")
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
