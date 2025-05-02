package controllers

import (
	"todo-frontend-web-app/common"
	"todo-frontend-web-app/models"
	"todo-frontend-web-app/services"

	"github.com/gofiber/fiber/v2"
)

type IndexController struct {
	ServiceManager *services.ServiceManager
}

func (controller *IndexController) IndexControllerGet(context *fiber.Ctx) error {
	todoLists := controller.getAllTodoListsForAuthenticatedUser(context)

	if todoLists == nil {
		return common.SendStatusInternalServerError(context)
	}

	return context.Render("index", fiber.Map{
		"Username":  common.GetAuthUsername(context),
		"TodoLists": todoLists,
	})
}

func (controller *IndexController) getAllTodoListsForAuthenticatedUser(context *fiber.Ctx) []models.TodoListModel {
	userId := common.GetAuthUserId(context)

	if userId == "" {
		return nil
	}

	response := controller.ServiceManager.TodoListService.GetAllNonDeletedByUserId(userId)

	if !response.IsSuccess() {
		return nil
	}

	return response.TodoLists
}
