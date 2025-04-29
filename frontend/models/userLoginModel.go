package models

type UserLogin struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
