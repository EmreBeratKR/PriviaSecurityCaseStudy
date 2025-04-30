package models

type LoginRequestModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (request *LoginRequestModel) TryLogin() (bool, string) {
	if request.Username == "Emre" && request.Password == "1234" {
		return true, ""
	}

	return false, "Wrong credentials"
}
