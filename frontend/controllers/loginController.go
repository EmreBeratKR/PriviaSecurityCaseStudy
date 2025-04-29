package controllers

import (
	"todo-frontend-web-app/models"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{})
}

func SubmitLogin(c *fiber.Ctx) error {
	var userLogin models.UserLogin
	err := c.BodyParser(&userLogin)

	if err != nil {
		return c.Render("login", fiber.Map{
			"Error": "Failed to parse form data!",
		})
	}

	username := userLogin.Username
	password := userLogin.Password
	user := models.LoginUser(username, password)

	if user == nil {
		return c.Render("login", fiber.Map{
			"Username": userLogin.Username,
			"Error":    "Wrong username or password!",
		})
	}

	return c.SendString("Welcome " + username)
}
