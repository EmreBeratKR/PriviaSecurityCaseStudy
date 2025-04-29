package apis

import (
	"todo-frontend-web-app/models/responses"
)

type APIClient interface {
	Login(username string, password string) *responses.UserLoginResponse
}

func ClientInstance() APIClient {
	return &FakeAPIClient{}
}
