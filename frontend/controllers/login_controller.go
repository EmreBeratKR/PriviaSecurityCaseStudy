package controllers

import (
	"todo-frontend-web-app/models"

	"github.com/gofiber/fiber/v2"
)

func LoginControllerGet(context *fiber.Ctx) error {
	return context.Render("login", fiber.Map{})
}

func LoginControllerPost(context *fiber.Ctx) error {
	loginRequest, parseSuccess := tryParseLoginRequest(context)

	if !parseSuccess {
		return context.Status(fiber.StatusBadRequest).Render("bad_request", fiber.Map{})
	}

	loginSuccess, msg := loginRequest.TryLogin()

	if !loginSuccess {
		return context.Render("login", fiber.Map{
			"Error": msg,
		})
	}

	return context.SendString("Welcome!")
}

func tryParseLoginRequest(context *fiber.Ctx) (*models.LoginRequestModel, bool) {
	var loginRequest models.LoginRequestModel

	if err := context.BodyParser(&loginRequest); err != nil {
		return nil, false
	}

	return &loginRequest, true
}
