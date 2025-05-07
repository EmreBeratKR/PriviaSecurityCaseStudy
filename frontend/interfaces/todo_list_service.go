package interfaces

import (
	"privia-sec-case-study/frontend/models"

	"github.com/gofiber/fiber/v2"
)

type TodoListService interface {
	GetNonDeletedById(context *fiber.Ctx, id string) *models.TodoListGetResponseModel
	GetAllNonDeleted(context *fiber.Ctx) *models.TodoListGetAllResponseModel
	GetAllNonDeletedByUserId(context *fiber.Ctx, userId string) *models.TodoListGetAllResponseModel
	AddWithUserIdAndName(context *fiber.Ctx, userId string, name string) *models.TodoListGetResponseModel
	UpdateNameById(context *fiber.Ctx, id string, name string) *models.TodoListGetResponseModel
	DeleteById(context *fiber.Ctx, id string) *models.TodoListGetResponseModel
}
