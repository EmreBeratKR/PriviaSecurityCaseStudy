package router

import (
	"privia-sec-case-study/backend/internal/handler/abstract_handlers"

	"github.com/gofiber/fiber/v2"
)

func MapUserRoutes(app *fiber.App, handler abstract_handlers.UserHandler) {
	app.Get("/users/login", handler.LoginUser)
}
