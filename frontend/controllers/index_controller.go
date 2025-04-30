package controllers

import (
	"todo-frontend-web-app/services"

	"github.com/gofiber/fiber/v2"
)

type IndexController struct {
	ServiceManager *services.ServiceManager
}

func (controller *IndexController) IndexControllerGet(context *fiber.Ctx) error {
	return context.SendString("hello from index")
}
