package handler

import (
	"encoding/base64"
	"privia-sec-case-study/backend/internal/common"
	"privia-sec-case-study/backend/internal/usercase"
	"privia-sec-case-study/shared"
	"strings"

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
		return common.SendStatusUnauthorized(context, "Auth header missing")
	}

	authHeaderSplits := strings.Split(authHeader, " ")

	authType := authHeaderSplits[0]
	if authType != "Basic" {
		return common.SendStatusUnauthorized(context, "Basic auth header missing")
	}

	encodedCredentials := authHeaderSplits[1]
	if encodedCredentials == "" {
		return common.SendStatusUnauthorized(context, "Credentials missing")
	}

	decodedBytes, err := base64.StdEncoding.DecodeString(encodedCredentials)
	if err != nil {
		return common.SendStatusInternalServerError(context, "Can't decode credentials")
	}

	credentials := string(decodedBytes)
	credentialsSplits := strings.Split(credentials, ":")

	username := credentialsSplits[0]
	if username == "" {
		return common.SendStatusBadRequest(context, "Username is required")
	}

	password := credentialsSplits[1]
	if password == "" {
		return common.SendStatusBadRequest(context, "Password is required")
	}

	response := handler.usecase.GetUserWithUsernameAndPassword(username, password)
	if response.IsNotSuccess() {
		return common.SendErrorStatus(response.Status, response.Message, context)
	}

	user := response.User
	jwtToken := shared.CreateJWT(shared.UserClaims{
		Username: user.Username,
		Role:     user.Role,
	}, user.Id)

	return common.SendStatusOkWithValue(jwtToken, context)
}
