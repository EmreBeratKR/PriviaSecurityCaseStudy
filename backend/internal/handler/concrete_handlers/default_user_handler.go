package concrete_handlers

import (
	"encoding/base64"
	"privia-sec-case-study/backend/internal/common"
	"privia-sec-case-study/backend/internal/usecase/abstract_usecases"
	"privia-sec-case-study/shared"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type DefaultUserHandler struct {
	usecase abstract_usecases.UserUsecase
}

func NewDefaultUserHandler(usecase abstract_usecases.UserUsecase) *DefaultUserHandler {
	return &DefaultUserHandler{
		usecase: usecase,
	}
}

func (handler *DefaultUserHandler) LoginUser(context *fiber.Ctx) error {
	authHeader := context.Get("Authorization")

	if authHeader == "" {
		return common.SendStatusUnauthorized(context, "Authorization header is required")
	}

	authHeaderSplits := strings.Split(authHeader, " ")
	if len(authHeaderSplits) < 2 {
		return common.SendStatusUnauthorized(context, "Unproccessable authorization header")
	}

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
