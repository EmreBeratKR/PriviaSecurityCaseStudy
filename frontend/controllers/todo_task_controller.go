package controllers

import (
	"todo-frontend-web-app/common"
	"todo-frontend-web-app/models"
	"todo-frontend-web-app/services"

	"github.com/gofiber/fiber/v2"
)

type TodoTaskController struct {
	ServiceManager *services.ServiceManager
}

func (controller *TodoTaskController) TodoTaskControllerPost(context *fiber.Ctx) error {
	request := controller.parseTodoTaskRequest(context)

	if request == nil {
		return common.SendStatusBadRequest(context)
	}

	if controller.tryAddTodoTaskByRequest(request) {
		return common.RedirectToTodoListPageById(context, request.ListId)
	}

	return common.SendStatusInternalServerError(context)
}

func (controller *TodoTaskController) TodoTaskControllerDelete(context *fiber.Ctx) error {
	if controller.tryDeleteTodoTaskByQueryParams(context) {
		return controller.redirectBackByQueryParams(context)
	}

	return common.SendStatusInternalServerError(context)
}

func (controller *TodoTaskController) TodoTaskControllerToggle(context *fiber.Ctx) error {
	if controller.tryToggleTodoTaskByQueryParams(context) {
		return controller.redirectBackByQueryParams(context)
	}

	return common.SendStatusInternalServerError(context)
}

func (controller *TodoTaskController) parseTodoTaskRequest(context *fiber.Ctx) *models.CreateTodoTaskRequestModel {
	var request models.CreateTodoTaskRequestModel

	if err := context.BodyParser(&request); err != nil {
		return nil
	}

	return &request
}

func (controller *TodoTaskController) tryAddTodoTaskByRequest(request *models.CreateTodoTaskRequestModel) bool {
	response := controller.ServiceManager.TodoTaskService.AddWithListIdAndContent(request.ListId, request.Content)

	return response.IsSuccess()
}

func (controller *TodoTaskController) tryDeleteTodoTaskByQueryParams(context *fiber.Ctx) bool {
	id := context.Query("id")

	if id == "" {
		return false
	}

	response := controller.ServiceManager.TodoTaskService.DeleteById(id)

	return response.IsSuccess()
}

func (controller *TodoTaskController) tryToggleTodoTaskByQueryParams(context *fiber.Ctx) bool {
	id := context.Query("id")

	if id == "" {
		return false
	}

	response := controller.ServiceManager.TodoTaskService.ToggleIsCompletedById(id)

	return response.IsSuccess()
}

func (controller *TodoTaskController) redirectBackByQueryParams(context *fiber.Ctx) error {
	redirectId := context.Query("redirect_id")

	if redirectId == "" {
		return common.SendStatusBadRequest(context)
	}

	return common.RedirectToTodoListPageById(context, redirectId)
}
