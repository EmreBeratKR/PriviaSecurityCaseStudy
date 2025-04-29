package models

import (
	"todo-frontend-web-app/apis"
	"todo-frontend-web-app/models/responses"
)

func LoginUser(username string, password string) *responses.UserLoginResponse {
	client := apis.ClientInstance()
	response := client.Login(username, password)
	if response == nil {
		return nil
	}

	return response
}
