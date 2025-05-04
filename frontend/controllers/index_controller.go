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
	isAdmin := common.IsAuthenticatedAsAdmin(context)

	if isAdmin {
		return controller.sendAdminPage(context)
	}

	return controller.sendUserPage(context)
}

func (controller *IndexController) sendAdminPage(context *fiber.Ctx) error {
	todoListsResponse := controller.ServiceManager.TodoListService.GetAllNonDeleted()

	if todoListsResponse.IsNotSuccess() {
		return common.SendErrorStatus(todoListsResponse.Status, context)
	}

	userId := common.GetAuthUserId(context)
	ownedTodoLists := todoListsResponse.Filtered(func(model *models.TodoListModel) bool {
		return model.UserId == userId
	})

	notOwnedTodoLists := todoListsResponse.Filtered(func(model *models.TodoListModel) bool {
		return model.UserId != userId
	})

	return context.Render("index", fiber.Map{
		"Username":          common.GetAuthUsername(context),
		"TodoLists":         ownedTodoLists,
		"OthersTodoLists":   notOwnedTodoLists,
		"IsAdmin":           true,
		"IsCreatingNewList": controller.isCreatingNewList(context),
	})
}

func (controller *IndexController) sendUserPage(context *fiber.Ctx) error {
	userId := common.GetAuthUserId(context)
	todoListsResponse := controller.ServiceManager.TodoListService.GetAllNonDeletedByUserId(userId)

	if todoListsResponse.IsNotSuccess() {
		return common.SendErrorStatus(todoListsResponse.Status, context)
	}

	return context.Render("index", fiber.Map{
		"Username":          common.GetAuthUsername(context),
		"TodoLists":         todoListsResponse.TodoLists,
		"IsAdmin":           false,
		"IsCreatingNewList": controller.isCreatingNewList(context),
	})
}

func (controller *IndexController) isCreatingNewList(context *fiber.Ctx) bool {
	return context.Query("create") != ""
}
