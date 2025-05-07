package interfaces

import (
	"privia-sec-case-study/frontend/models"

	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	Login(context *fiber.Ctx, request *models.LoginRequestModel) *models.LoginResponseModel
}
