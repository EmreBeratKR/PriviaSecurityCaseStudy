package controllers

import "github.com/gofiber/fiber/v2"

func IndexControllerGet(context *fiber.Ctx) error {
	return context.SendString("hello from index")
}
