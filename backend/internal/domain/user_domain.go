package domain

import "privia-sec-case-study/shared"

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Hash     string `json:"hash"`
	Role     string `json:"role"`
}

type GetUserResponse struct {
	shared.StatusModel
	Message string `json:"message"`
	User    User
}

type GetAllUsersResponse struct {
	shared.StatusModel
	Message string `json:"message"`
	Users   []User
}

type UserRepository interface {
	GetByUsername(username string) *GetUserResponse
}
