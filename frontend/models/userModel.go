package models

type User struct {
	Username string
	Password string
}

func ValidateUser(username, password string) *User {
	if username == "emre" && password == "1234" {
		return &User{
			Username: username,
			Password: password,
		}
	}
	return nil
}
