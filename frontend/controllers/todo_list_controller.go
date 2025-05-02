package controllers

import (
	"todo-frontend-web-app/common"
	"todo-frontend-web-app/models"
	"todo-frontend-web-app/services"

	"github.com/gofiber/fiber/v2"
)

type TodoListController struct {
	ServiceManager *services.ServiceManager
}

func (controller *TodoListController) TodoListControllerGet(context *fiber.Ctx) error {
	todoList := controller.getTodoListByQueryParams(context)

	if todoList == nil {
		return common.SendStatusInternalServerError(context)
	}

	todoTasks := controller.getTodoTasksByTodoListId(todoList.Id)

	if todoTasks == nil {
		return common.SendStatusInternalServerError(context)
	}

	return context.Render("todo_list", fiber.Map{
		"Name":              todoList.Name,
		"CompletionPercent": todoList.CompletionPercent,
		"TodoTasks":         todoTasks,
	})
}

func (controller *TodoListController) TodoListControllerPost(context *fiber.Ctx) error {
	success := controller.tryAddTodoListForAuthenticatedUser(context)

	if !success {
		return common.SendStatusInternalServerError(context)
	}

	return common.RedirectToHomePage(context)
}

func (controller *TodoListController) getTodoListByQueryParams(context *fiber.Ctx) *models.TodoListModel {
	todoListId := context.Query("id")

	if todoListId == "" {
		return nil
	}

	response := controller.ServiceManager.TodoListService.GetById(todoListId)

	if !response.IsSuccess() {
		return nil
	}

	return &response.TodoList
}

func (controller *TodoListController) getTodoTasksByTodoListId(todoListId string) []models.TodoTaskModel {
	response := controller.ServiceManager.TodoTaskService.GetAllByTodoListId(todoListId)

	if !response.IsSuccess() {
		return nil
	}

	return response.TodoTasks
}

func (controller *TodoListController) tryAddTodoListForAuthenticatedUser(context *fiber.Ctx) bool {
	userId := common.GetAuthUserId(context)

	if userId == "" {
		return false
	}

	response := controller.ServiceManager.TodoListService.AddWithUserId(userId)

	return response.IsSuccess()
}
