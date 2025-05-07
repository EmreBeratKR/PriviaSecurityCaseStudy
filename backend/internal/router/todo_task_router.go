package router

import (
	"privia-sec-case-study/backend/internal/handler/abstract_handlers"

	"github.com/gofiber/fiber/v2"
)

func MapTodoTaskRoutes(app *fiber.App, handler abstract_handlers.TodoTaskHandler) {
	path := "/todo-tasks"
	app.Get(path, handler.GetTodoTasks)
	app.Post(path, handler.PostTodoTasks)
	app.Patch(path, handler.PatchTodoTasks)
	app.Delete(path, handler.DeleteTodoTasks)
}
