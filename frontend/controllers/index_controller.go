package controllers

import (
	"privia-sec-case-study/frontend/common"
	"privia-sec-case-study/frontend/models"
	"privia-sec-case-study/frontend/services"

	"github.com/gofiber/fiber/v2"
)

var pageTitle = "Emrello"

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
	todoListsResponse := controller.ServiceManager.TodoListService.GetAllNonDeleted(context)

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

	for i := range ownedTodoLists {
		ownedTodoLists[i].ShowDeleteButton = true
	}

	for i := range notOwnedTodoLists {
		notOwnedTodoLists[i].ShowDeleteButton = false
	}

	return context.Render("index", fiber.Map{
		"PageTitle":             pageTitle,
		"UserId":                userId,
		"Username":              common.GetAuthUsername(context),
		"TodoLists":             ownedTodoLists,
		"IsTodoListsEmpty":      len(ownedTodoLists) <= 0,
		"OthersTodoLists":       notOwnedTodoLists,
		"IsOtherTodoListsEmpty": len(notOwnedTodoLists) <= 0,
		"IsAdmin":               true,
		"IsCreatingNewList":     controller.isCreatingNewList(context),
	})
}

func (controller *IndexController) sendUserPage(context *fiber.Ctx) error {
	userId := common.GetAuthUserId(context)
	todoListsResponse := controller.ServiceManager.TodoListService.GetAllNonDeletedByUserId(context, userId)

	if todoListsResponse.IsNotSuccess() {
		return common.SendErrorStatus(todoListsResponse.Status, context)
	}

	for i := range todoListsResponse.TodoLists {
		todoListsResponse.TodoLists[i].ShowDeleteButton = true
	}

	return context.Render("index", fiber.Map{
		"UserId":            userId,
		"Username":          common.GetAuthUsername(context),
		"TodoLists":         todoListsResponse.TodoLists,
		"IsTodoListsEmpty":  len(todoListsResponse.TodoLists) <= 0,
		"IsAdmin":           false,
		"IsCreatingNewList": controller.isCreatingNewList(context),
	})
}

func (controller *IndexController) isCreatingNewList(context *fiber.Ctx) bool {
	return context.Query("create") != ""
}
