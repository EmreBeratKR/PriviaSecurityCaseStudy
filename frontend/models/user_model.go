package models

type UserModel struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Hash     string `json:"hash"`
	Role     string `json:"Role"`
}
