package models

type LoginRequestModel struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func (request *LoginRequestModel) IsValid() bool {
	if request.Username == "" {
		return false
	}

	if request.Password == "" {
		return false
	}

	return true
}
