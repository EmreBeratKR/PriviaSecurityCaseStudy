package abstract_handlers

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	LoginUser(context *fiber.Ctx) error
}
