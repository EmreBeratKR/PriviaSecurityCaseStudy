package abstract_handlers

import (
	"github.com/gofiber/fiber/v2"
)

type TodoListHandler interface {
	GetTodoLists(context *fiber.Ctx) error
	PostTodoLists(context *fiber.Ctx) error
	PatchTodoLists(context *fiber.Ctx) error
	DeleteTodoLists(context *fiber.Ctx) error
}
