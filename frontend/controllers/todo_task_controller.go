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
