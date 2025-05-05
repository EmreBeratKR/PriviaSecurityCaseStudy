package controllers

import (
	"todo-frontend-web-app/common"
	"todo-frontend-web-app/services"

	"github.com/gofiber/fiber/v2"
)

type TodoListController struct {
	ServiceManager *services.ServiceManager
}

func (controller *TodoListController) TodoListControllerGet(context *fiber.Ctx) error {
	controller.ServiceManager.SetContext(context)

	todoListId := context.Query("id")
	if todoListId == "" {
		return common.SendStatusBadRequest(context)
	}

	todoListsResponse := controller.ServiceManager.TodoListService.GetNonDeletedById(todoListId)
	if todoListsResponse.IsNotSuccess() {
		return common.SendErrorStatus(todoListsResponse.Status, context)
	}

	todoTasksResponse := controller.ServiceManager.TodoTaskService.GetAllNonDeletedByTodoListId(todoListId)
	if todoTasksResponse.IsNotSuccess() {
		return common.SendErrorStatus(todoTasksResponse.Status, context)
	}

	userId := todoListsResponse.TodoList.UserId
	allowEditting := userId == common.GetAuthUserId(context)
	isEdittingListName := context.Query("edit_name") != "" && allowEditting
	editTodoTaskId := context.Query("edit_todo_task_id")
	editTaskContent := ""
	isEdittingTodoTask := editTodoTaskId != ""

	for i, todoTask := range todoTasksResponse.TodoTasks {
		todoTasksResponse.TodoTasks[i].AllowEditting = allowEditting
		if todoTask.Id == editTodoTaskId {
			editTaskContent = todoTask.Content
		}
	}

	return context.Render("todo_list", fiber.Map{
		"Username":              common.GetAuthUsername(context),
		"PageTitle":             todoListsResponse.TodoList.Name + " - Todo List",
		"TodoListId":            todoListsResponse.TodoList.Id,
		"Name":                  todoListsResponse.TodoList.Name,
		"CompletionPercent":     todoListsResponse.TodoList.CompletionPercent,
		"TodoTasks":             todoTasksResponse.TodoTasks,
		"IsTodoTasksEmpty":      len(todoTasksResponse.TodoTasks) <= 0,
		"EditTodoTaskId":        editTodoTaskId,
		"IsEdittingTodoTask":    isEdittingTodoTask,
		"IsNotEdittingTodoTask": !isEdittingTodoTask,
		"EditTaskContent":       editTaskContent,
		"AllowEditting":         allowEditting,
		"IsEdittingListName":    isEdittingListName,
	})
}

func (controller *TodoListController) TodoListControllerPost(context *fiber.Ctx) error {
	controller.ServiceManager.SetContext(context)

	userId := context.FormValue("user_id")
	if userId == "" {
		return common.SendStatusBadRequest(context)
	}

	name := context.FormValue("name")
	if name == "" {
		return common.SendStatusBadRequest(context)
	}

	if !controller.isValidTodoListName(name) {
		return common.SendStatusBadRequest(context)
	}

	response := controller.ServiceManager.TodoListService.AddWithUserIdAndName(userId, name)
	if response.IsNotSuccess() {
		return common.SendErrorStatus(response.Status, context)
	}

	return common.RedirectToTodoListPageById(context, response.TodoList.Id)
}

func (controller *TodoListController) TodoListControllerDelete(context *fiber.Ctx) error {
	controller.ServiceManager.SetContext(context)

	id := context.FormValue("id")
	if id == "" {
		return common.SendStatusBadRequest(context)
	}

	response := controller.ServiceManager.TodoListService.DeleteById(id)
	if response.IsNotSuccess() {
		return common.SendErrorStatus(response.Status, context)
	}

	return common.RedirectToHomePage(context)
}

func (controller *TodoListController) TodoListControllerPatch(context *fiber.Ctx) error {
	controller.ServiceManager.SetContext(context)

	id := context.FormValue("id")
	if id == "" {
		return common.SendStatusBadRequest(context)
	}

	name := context.FormValue("name")
	if name == "" {
		return common.SendStatusBadRequest(context)
	}

	if !controller.isValidTodoListName(name) {
		return common.SendStatusBadRequest(context)
	}

	response := controller.ServiceManager.TodoListService.UpdateNameById(id, name)
	if response.IsNotSuccess() {
		return common.SendErrorStatus(response.Status, context)
	}

	return common.RedirectToTodoListPageById(context, id)
}

func (controller *TodoListController) isValidTodoListName(name string) bool {
	return len(name) <= 20
}
