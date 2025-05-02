package common

import "github.com/gofiber/fiber/v2"

func RedirectToHomePage(context *fiber.Ctx) error {
	return context.Redirect("/")
}

func RedirectToLoginPage(context *fiber.Ctx) error {
	return context.Redirect("/login")
}

func SendStatusBadRequest(context *fiber.Ctx) error {
	return context.Status(fiber.StatusBadRequest).Render("bad_request", fiber.Map{})
}

func SendStatusNotFound(context *fiber.Ctx) error {
	return context.Status(fiber.StatusNotFound).Render("not_found", fiber.Map{})
}

func SendStatusInternalServerError(context *fiber.Ctx) error {
	return context.Status(fiber.StatusInternalServerError).Render("server_error", fiber.Map{})
}
