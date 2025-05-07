package domain

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Hash     string `json:"hash"`
	Role     string `json:"role"`
}
