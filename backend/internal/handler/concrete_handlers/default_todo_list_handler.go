package concrete_handlers

import (
	"privia-sec-case-study/backend/internal/common"
	"privia-sec-case-study/backend/internal/usecase/abstract_usecases"
	"privia-sec-case-study/shared"

	"github.com/gofiber/fiber/v2"
)

type DefaultTodoListHandler struct {
	usecase abstract_usecases.TodoListUsecase
}

func NewDefaultTodoListHandler(usecase abstract_usecases.TodoListUsecase) *DefaultTodoListHandler {
	return &DefaultTodoListHandler{
		usecase: usecase,
	}
}

func (handler *DefaultTodoListHandler) GetTodoLists(context *fiber.Ctx) error {
	claims, errCallback := common.GetClaimsFromHeaders(context)
	if claims == nil {
		return errCallback()
	}

	id := context.Query("id")
	if id != "" {
		return handler.getNonDeletedById(context, id, claims)
	}

	userId := context.Query("user_id")
	if userId != "" {
		return handler.getAllNonDeletedByUserId(context, userId, claims)
	}

	return handler.getAllNonDeleted(context, claims)
}

func (handler *DefaultTodoListHandler) PostTodoLists(context *fiber.Ctx) error {
	claims, errCallback := common.GetClaimsFromHeaders(context)
	if claims == nil {
		return errCallback()
	}

	userId := context.FormValue("user_id")
	if userId == "" {
		return common.SendStatusBadRequest(context, "user_id is required")
	}

	name := context.FormValue("name")
	if name == "" {
		return common.SendStatusBadRequest(context, "name is required")
	}

	if claims.IsNotAuthorizedForWrite(userId) {
		return common.SendStatusForbidden(context, "You don't have permission")
	}

	response := handler.usecase.AddWithUserIdAndName(userId, name)
	if response.IsNotSuccess() {
		return common.SendErrorStatus(response.Status, response.Message, context)
	}

	todoList := response.TodoList
	return common.SendStatusOkWithValue(todoList, context)
}

func (handler *DefaultTodoListHandler) PatchTodoLists(context *fiber.Ctx) error {
	claims, errCallback := common.GetClaimsFromHeaders(context)
	if claims == nil {
		return errCallback()
	}

	id := context.FormValue("id")
	if id == "" {
		return common.SendStatusBadRequest(context, "id is required")
	}

	response := handler.usecase.GetNonDeletedById(id)
	if response.IsNotSuccess() {
		return common.SendErrorStatus(response.Status, response.Message, context)
	}

	todoList := response.TodoList
	userId := todoList.UserId
	if claims.IsNotAuthorizedForWrite(userId) {
		return common.SendStatusForbidden(context, "You don't have permission")
	}

	name := context.FormValue("name")
	if name == "" {
		return common.SendStatusBadRequest(context, "name is required")
	}

	updataResponse := handler.usecase.UpdateNameById(id, name)
	if updataResponse.IsNotSuccess() {
		return common.SendErrorStatus(updataResponse.Status, updataResponse.Message, context)
	}

	return common.SendStatusOkWithValue(updataResponse.TodoList, context)
}

func (handler *DefaultTodoListHandler) DeleteTodoLists(context *fiber.Ctx) error {
	claims, err := common.GetClaimsFromHeaders(context)
	if err != nil {
		return nil
	}

	id := context.FormValue("id")
	if id == "" {
		return common.SendStatusBadRequest(context, "id is required")
	}

	response := handler.usecase.GetNonDeletedById(id)
	if response.IsNotSuccess() {
		return common.SendErrorStatus(response.Status, response.Message, context)
	}

	todoList := response.TodoList
	userId := todoList.UserId
	if claims.IsNotAuthorizedForWrite(userId) {
		return common.SendStatusForbidden(context, "You don't have permission")
	}

	deleteResponse := handler.usecase.DeleteById(id)
	if deleteResponse.IsNotSuccess() {
		return common.SendErrorStatus(deleteResponse.Status, deleteResponse.Message, context)
	}

	return common.SendStatusOkWithValue(deleteResponse.TodoList, context)
}

func (handler *DefaultTodoListHandler) getAllNonDeleted(context *fiber.Ctx, claims *shared.UserClaims) error {
	if claims.IsNotAdmin() {
		return common.SendStatusForbidden(context, "You don't have access")
	}

	response := handler.usecase.GetAllNonDeleted()
	if response.IsNotSuccess() {
		return common.SendErrorStatus(response.Status, response.Message, context)
	}

	todoLists := response.TodoLists
	return common.SendStatusOkWithValue(todoLists, context)
}

func (handler *DefaultTodoListHandler) getAllNonDeletedByUserId(context *fiber.Ctx, userId string, claims *shared.UserClaims) error {
	if claims.IsNotAuthorizedForRead(userId) {
		return common.SendStatusForbidden(context, "You don't have access")
	}

	response := handler.usecase.GetAllNonDeletedByUserId(userId)
	if response.IsNotSuccess() {
		return common.SendErrorStatus(response.Status, response.Message, context)
	}

	todoLists := response.TodoLists
	return common.SendStatusOkWithValue(todoLists, context)
}

func (handler *DefaultTodoListHandler) getNonDeletedById(context *fiber.Ctx, id string, claims *shared.UserClaims) error {
	response := handler.usecase.GetNonDeletedById(id)
	if response.IsNotSuccess() {
		return common.SendErrorStatus(response.Status, response.Message, context)
	}

	todoList := response.TodoList
	userId := todoList.UserId

	if claims.IsNotAuthorizedForRead(userId) {
		return common.SendStatusForbidden(context, "You don't have access")
	}

	return common.SendStatusOkWithValue(todoList, context)
}
