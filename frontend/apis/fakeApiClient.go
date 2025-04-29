package apis

import (
	"todo-frontend-web-app/models/responses"
)

type FakeAPIClient struct{}

func (f *FakeAPIClient) Login(username string, password string) *responses.UserLoginResponse {
	if username == "emre" && password == "1234" {
		return &responses.UserLoginResponse{}
	}

	return nil
}
