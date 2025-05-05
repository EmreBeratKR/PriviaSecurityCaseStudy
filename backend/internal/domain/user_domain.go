package domain

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Hash     string `json:"hash"`
}

type GetUserResponse struct {
	StatusModel
	User User
}

type GetAllUsersResponse struct {
	StatusModel
	Users []User
}

type UserRepository interface {
	GetAll() *GetAllUsersResponse
}
