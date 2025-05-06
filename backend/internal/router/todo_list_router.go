package router

import (
	"privia-sec-case-study/backend/internal/handler/abstract_handlers"

	"github.com/gofiber/fiber/v2"
)

func MapTodoListRouter(app *fiber.App, handler abstract_handlers.TodoListHandler) {
	path := "/todo-lists"
	app.Get(path, handler.GetTodoLists)
	app.Post(path, handler.PostTodoLists)
	app.Patch(path, handler.PatchTodoLists)
	app.Delete(path, handler.DeleteTodoLists)
}
