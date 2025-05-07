package common

import (
	"privia-sec-case-study/shared"

	"github.com/gofiber/fiber/v2"
)

func RedirectToHomePage(context *fiber.Ctx) error {
	return context.Redirect("/")
}

func RedirectToLoginPage(context *fiber.Ctx) error {
	return context.Redirect("/login")
}

func RedirectToTodoListPageById(context *fiber.Ctx, id string) error {
	return context.Redirect("/todo-list?id=" + id)
}

func SendErrorStatus(status string, context *fiber.Ctx) error {
	if status == shared.StatusBadRequest().Status {
		return SendStatusBadRequest(context)
	}

	if status == shared.StatusForbidden().Status {
		return SendStatusForbidden(context)
	}

	if status == shared.StatusForbidden().Status {
		return SendStatusForbidden(context)
	}

	if status == shared.StatusNotFound().Status {
		return SendStatusNotFound(context)
	}

	if status == shared.StatusInternalServerError().Status {
		return SendStatusInternalServerError(context)
	}

	return SendStatusInternalServerError(context)
}

func SendStatusBadRequest(context *fiber.Ctx) error {
	return context.Status(fiber.StatusBadRequest).Render("bad_request", fiber.Map{})
}

func SendStatusUnauthorized(context *fiber.Ctx) error {
	return context.Status(fiber.StatusForbidden).Render("unauthorized", fiber.Map{})
}

func SendStatusForbidden(context *fiber.Ctx) error {
	return context.Status(fiber.StatusForbidden).Render("forbidden", fiber.Map{})
}

func SendStatusNotFound(context *fiber.Ctx) error {
	return context.Status(fiber.StatusNotFound).Render("not_found", fiber.Map{})
}

func SendStatusInternalServerError(context *fiber.Ctx) error {
	return context.Status(fiber.StatusInternalServerError).Render("server_error", fiber.Map{})
}
