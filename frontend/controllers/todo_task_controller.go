package controllers

import (
	"privia-sec-case-study/frontend/common"
	"privia-sec-case-study/frontend/services"

	"github.com/gofiber/fiber/v2"
)

type TodoTaskController struct {
	ServiceManager *services.ServiceManager
}

func (controller *TodoTaskController) TodoTaskControllerPost(context *fiber.Ctx) error {
	controller.ServiceManager.SetContext(context)

	listId := context.FormValue("list_id")
	if listId == "" {
		return common.SendStatusBadRequest(context)
	}

	content := context.FormValue("content")
	if content == "" {
		return common.SendStatusBadRequest(context)
	}

	response := controller.ServiceManager.TodoTaskService.AddWithListIdAndContent(listId, content)
	if response.IsNotSuccess() {
		return common.SendErrorStatus(response.Status, context)
	}

	return common.RedirectToTodoListPageById(context, listId)
}

func (controller *TodoTaskController) TodoTaskControllerPatch(context *fiber.Ctx) error {
	controller.ServiceManager.SetContext(context)

	id := context.FormValue("id")
	if id == "" {
		return common.SendStatusBadRequest(context)
	}

	action := context.FormValue("action")

	if action == "toggle" {
		return controller.sendPatchToggle(context, id)
	}

	if action == "edit" {
		return controller.sendPatchEdit(context, id)
	}

	return common.SendStatusBadRequest(context)
}

func (controller *TodoTaskController) TodoTaskControllerDelete(context *fiber.Ctx) error {
	controller.ServiceManager.SetContext(context)

	id := context.FormValue("id")
	if id == "" {
		return common.SendStatusBadRequest(context)
	}

	response := controller.ServiceManager.TodoTaskService.DeleteById(id)

	if response.IsNotSuccess() {
		return common.SendErrorStatus(response.Status, context)
	}

	return common.RedirectToTodoListPageById(context, response.TodoTask.TodoListId)
}

func (controller *TodoTaskController) sendPatchToggle(context *fiber.Ctx, id string) error {
	response := controller.ServiceManager.TodoTaskService.ToggleIsCompletedById(id)

	if response.IsNotSuccess() {
		return common.SendErrorStatus(response.Status, context)
	}

	return common.RedirectToTodoListPageById(context, response.TodoTask.TodoListId)
}

func (controller *TodoTaskController) sendPatchEdit(context *fiber.Ctx, id string) error {
	content := context.FormValue("content")
	if content == "" {
		return common.SendStatusBadRequest(context)
	}

	response := controller.ServiceManager.TodoTaskService.UpdateContentById(id, content)

	if response.IsNotSuccess() {
		return common.SendErrorStatus(response.Status, context)
	}

	return common.RedirectToTodoListPageById(context, response.TodoTask.TodoListId)
}
