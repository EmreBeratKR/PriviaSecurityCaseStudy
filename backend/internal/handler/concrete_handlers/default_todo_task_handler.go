package concrete_handlers

import (
	"privia-sec-case-study/backend/internal/common"
	"privia-sec-case-study/backend/internal/usecase/abstract_usecases"

	"github.com/gofiber/fiber/v2"
)

type DefaultTodoTaskHandler struct {
	todoListUsecase abstract_usecases.TodoListUsecase
	todoTaskUsecase abstract_usecases.TodoTaskUsecase
}

func NewDefaultTodoTaskHandler(todoListUsecase abstract_usecases.TodoListUsecase, todoTaskUsecase abstract_usecases.TodoTaskUsecase) *DefaultTodoTaskHandler {
	return &DefaultTodoTaskHandler{
		todoListUsecase: todoListUsecase,
		todoTaskUsecase: todoTaskUsecase,
	}
}

func (handler *DefaultTodoTaskHandler) GetTodoTasks(context *fiber.Ctx) error {
	claims, errCallback := common.GetClaimsFromHeaders(context)
	if claims == nil {
		return errCallback()
	}

	listId := context.Query("list_id")
	if listId == "" {
		return common.SendStatusBadRequest(context, "list_id is required")
	}

	todoListResponse := handler.todoListUsecase.GetNonDeletedById(listId)
	if todoListResponse.IsNotSuccess() {
		return common.SendErrorStatus(todoListResponse.Status, todoListResponse.Message, context)
	}

	userId := todoListResponse.TodoList.UserId
	if claims.IsNotAuthorizedForRead(userId) {
		return common.SendStatusForbidden(context, "You don't have access")
	}

	todoTasksResponse := handler.todoTaskUsecase.GetAllNonDeletedByListId(listId)
	if todoTasksResponse.IsNotSuccess() {
		return common.SendErrorStatus(todoTasksResponse.Status, todoTasksResponse.Message, context)
	}

	return common.SendStatusOkWithValue(todoTasksResponse.TodoTasks, context)
}

func (handler *DefaultTodoTaskHandler) PostTodoTasks(context *fiber.Ctx) error {
	claims, errCallback := common.GetClaimsFromHeaders(context)
	if claims == nil {
		return errCallback()
	}

	listId := context.FormValue("list_id")
	if listId == "" {
		return common.SendStatusBadRequest(context, "list_id is required")
	}

	content := context.FormValue("content")
	if content == "" {
		return common.SendStatusBadRequest(context, "content is required")
	}

	todoListResponse := handler.todoListUsecase.GetNonDeletedById(listId)
	if todoListResponse.IsNotSuccess() {
		return common.SendErrorStatus(todoListResponse.Status, todoListResponse.Message, context)
	}

	userId := todoListResponse.TodoList.UserId
	if claims.IsNotAuthorizedForRead(userId) {
		return common.SendStatusForbidden(context, "You don't have access")
	}

	todoTaskResponse := handler.todoTaskUsecase.AddWithListIdAndContent(listId, content)
	if todoTaskResponse.IsNotSuccess() {
		return common.SendErrorStatus(todoTaskResponse.Status, todoTaskResponse.Message, context)
	}

	return common.SendStatusOkWithValue(todoTaskResponse.TodoTask, context)
}

func (handler *DefaultTodoTaskHandler) PatchTodoTasks(context *fiber.Ctx) error {
	claims, errCallback := common.GetClaimsFromHeaders(context)
	if claims == nil {
		return errCallback()
	}

	id := context.FormValue("id")
	if id == "" {
		return common.SendStatusBadRequest(context, "id is required")
	}

	todoTaskResponse := handler.todoTaskUsecase.GetNonDeletedById(id)
	if todoTaskResponse.IsNotSuccess() {
		return common.SendErrorStatus(todoTaskResponse.Status, todoTaskResponse.Message, context)
	}

	listId := todoTaskResponse.TodoTask.TodoListId
	todoListResponse := handler.todoListUsecase.GetNonDeletedById(listId)
	if todoListResponse.IsNotSuccess() {
		return common.SendErrorStatus(todoListResponse.Status, todoListResponse.Message, context)
	}

	userId := todoListResponse.TodoList.UserId
	if claims.IsNotAuthorizedForWrite(userId) {
		return common.SendStatusForbidden(context, "You don't have access")
	}

	action := context.FormValue("action")
	if action == "" {
		return common.SendStatusBadRequest(context, "action is required")
	}

	if action == "toggle" {
		return handler.toggleIsCompletedById(context, id)
	}

	if action == "edit" {
		return handler.updateContentById(context, id)
	}

	return common.SendStatusBadRequest(context, "invalid action provided")
}

func (handler *DefaultTodoTaskHandler) DeleteTodoTasks(context *fiber.Ctx) error {
	claims, errCallback := common.GetClaimsFromHeaders(context)
	if claims == nil {
		return errCallback()
	}

	id := context.FormValue("id")
	if id == "" {
		return common.SendStatusBadRequest(context, "id is required")
	}

	todoTaskResponse := handler.todoTaskUsecase.GetNonDeletedById(id)
	if todoTaskResponse.IsNotSuccess() {
		return common.SendErrorStatus(todoTaskResponse.Status, todoTaskResponse.Message, context)
	}

	listId := todoTaskResponse.TodoTask.TodoListId
	todoListResponse := handler.todoListUsecase.GetNonDeletedById(listId)
	if todoListResponse.IsNotSuccess() {
		return common.SendErrorStatus(todoListResponse.Status, todoListResponse.Message, context)
	}

	userId := todoListResponse.TodoList.UserId
	if claims.IsNotAuthorizedForWrite(userId) {
		return common.SendStatusForbidden(context, "You don't have access")
	}

	deleteResponse := handler.todoTaskUsecase.DeleteById(id)
	if deleteResponse.IsNotSuccess() {
		return common.SendErrorStatus(deleteResponse.Status, deleteResponse.Message, context)
	}

	return common.SendStatusOkWithValue(deleteResponse.TodoTask, context)
}

func (handler *DefaultTodoTaskHandler) toggleIsCompletedById(context *fiber.Ctx, id string) error {
	response := handler.todoTaskUsecase.ToggleIsCompletedById(id)
	if response.IsNotSuccess() {
		return common.SendErrorStatus(response.Status, response.Message, context)
	}

	return common.SendStatusOkWithValue(response.TodoTask, context)
}

func (handler *DefaultTodoTaskHandler) updateContentById(context *fiber.Ctx, id string) error {
	content := context.FormValue("content")
	if content == "" {
		return common.SendStatusBadRequest(context, "content is required")
	}

	response := handler.todoTaskUsecase.UpdateContentById(id, content)
	if response.IsNotSuccess() {
		return common.SendErrorStatus(response.Status, response.Message, context)
	}

	return common.SendStatusOkWithValue(response.TodoTask, context)
}
