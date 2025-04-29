package controllers

import (
	"todo-frontend-web-app/models"

	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	// Call the model to get the message
	message := models.Hello()

	// Render the view with the message
	return c.Render("hello", fiber.Map{
		"Message": message,
	})
}
