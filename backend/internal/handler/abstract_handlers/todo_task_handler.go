package abstract_handlers

import (
	"github.com/gofiber/fiber/v2"
)

type TodoTaskHandler interface {
	GetTodoTasks(context *fiber.Ctx) error
	PostTodoTasks(context *fiber.Ctx) error
	PatchTodoTasks(context *fiber.Ctx) error
	DeleteTodoTasks(context *fiber.Ctx) error
}
