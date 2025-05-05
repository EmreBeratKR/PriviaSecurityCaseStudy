package handler

import (
	"privia-sec-case-study/backend/internal/common"
	"privia-sec-case-study/backend/internal/usercase"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	usecase *usercase.UserUsecase
}

func NewUserHandler(usecase *usercase.UserUsecase) *UserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}

func (handler *UserHandler) LoginUser(context *fiber.Ctx) error {
	authHeader := context.Get("Authorization")

	if authHeader == "" {
		return common.SendStatusUnauthorized(context)
	}

	response := handler.usecase.GetUserWithUsernameAndHash("0", "1")
	if response.IsNotSuccess() {
		return common.SendErrorStatus(response.Status, context)
	}

	return common.SendStatusOkWithValue(response.User, context)
}
