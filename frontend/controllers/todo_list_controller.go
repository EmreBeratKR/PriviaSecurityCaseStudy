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
		return common.SendStatusNotFound(context)
	}

	userId := todoList.UserId
	if !common.IsAuthorizedForUserId(context, userId) {
		return common.SendStatusForbidden(context)
	}

	todoTasks := controller.getTodoTasksByTodoListId(todoList.Id)

	if todoTasks == nil {
		return common.SendStatusInternalServerError(context)
	}

	allowEditting := userId == common.GetAuthUserId(context)
	editTodoTaskId := context.Query("edit_todo_task_id")
	isEdittingTodoTask := editTodoTaskId != ""

	for i, _ := range todoTasks {
		todoTasks[i].AllowEditting = allowEditting
	}

	if !isEdittingTodoTask {
		return context.Render("todo_list", fiber.Map{
			"TodoListId":            todoList.Id,
			"Name":                  todoList.Name,
			"CompletionPercent":     todoList.CompletionPercent,
			"TodoTasks":             todoTasks,
			"EditTodoTaskId":        editTodoTaskId,
			"IsNotEdittingTodoTask": !isEdittingTodoTask,
			"IsEdittingTodoTask":    isEdittingTodoTask,
			"AllowEditting":         allowEditting,
		})
	}

	todoTask := controller.getTodoTaskById(editTodoTaskId)

	if todoTask == nil {
		return common.SendStatusInternalServerError(context)
	}

	return context.Render("todo_list", fiber.Map{
		"TodoListId":            todoList.Id,
		"Name":                  todoList.Name,
		"CompletionPercent":     todoList.CompletionPercent,
		"TodoTasks":             todoTasks,
		"EditTodoTaskId":        editTodoTaskId,
		"IsEdittingTodoTask":    isEdittingTodoTask,
		"IsNotEdittingTodoTask": !isEdittingTodoTask,
		"EditTaskContent":       todoTask.Content,
		"AllowEditting":         allowEditting,
	})
}

func (controller *TodoListController) TodoListControllerPost(context *fiber.Ctx) error {
	success := controller.tryAddTodoListForAuthenticatedUser(context)

	if !success {
		return common.SendStatusInternalServerError(context)
	}

	return common.RedirectToHomePage(context)
}

func (controller *TodoListController) TodoListControllerDelete(context *fiber.Ctx) error {
	if controller.tryDeleteTodoListByQueryParams(context) {
		return common.RedirectToHomePage(context)
	}

	return common.SendStatusInternalServerError(context)
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
	response := controller.ServiceManager.TodoTaskService.GetAllNonDeletedByTodoListId(todoListId)

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

	name := context.FormValue("content")

	if name == "" {
		return false
	}

	response := controller.ServiceManager.TodoListService.AddWithUserIdAndName(userId, name)

	return response.IsSuccess()
}

func (controller *TodoListController) tryDeleteTodoListByQueryParams(context *fiber.Ctx) bool {
	id := context.Query("id")

	if id == "" {
		return false
	}

	response := controller.ServiceManager.TodoListService.DeleteById(id)

	return response.IsSuccess()
}

func (controller *TodoListController) getTodoTaskById(id string) *models.TodoTaskModel {
	response := controller.ServiceManager.TodoTaskService.GetById(id)

	if response.IsSuccess() {
		return &response.TodoTask
	}

	return nil
}
