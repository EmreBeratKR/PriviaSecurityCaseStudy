package models

type LoginRequestModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
