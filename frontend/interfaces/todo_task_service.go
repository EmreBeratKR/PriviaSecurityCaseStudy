package interfaces

import (
	"privia-sec-case-study/frontend/models"

	"github.com/gofiber/fiber/v2"
)

type TodoTaskService interface {
	GetAllNonDeletedByTodoListId(context *fiber.Ctx, todoListId string) *models.TodoTaskGetAllResponseModel
	AddWithListIdAndContent(context *fiber.Ctx, todoListId string, content string) *models.TodoTaskGetResponseModel
	ToggleIsCompletedById(context *fiber.Ctx, id string) *models.TodoTaskGetResponseModel
	UpdateContentById(context *fiber.Ctx, id string, content string) *models.TodoTaskGetResponseModel
	DeleteById(context *fiber.Ctx, id string) *models.TodoTaskGetResponseModel
}
