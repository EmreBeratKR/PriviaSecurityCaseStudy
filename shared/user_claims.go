package shared

import "github.com/golang-jwt/jwt/v4"

type UserClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func (model *UserClaims) IsAdmin() bool {
	return model.Role == "admin"
}

func (model *UserClaims) IsNotAdmin() bool {
	return !model.IsAdmin()
}

func (model *UserClaims) IsAuthorizedForRead(userId string) bool {
	if model.IsAdmin() {
		return true
	}

	return model.Subject == userId
}

func (model *UserClaims) IsNotAuthorizedForRead(userId string) bool {
	return !model.IsAuthorizedForRead(userId)
}

func (model *UserClaims) IsAuthorizedForWrite(userId string) bool {
	return model.Subject == userId
}

func (model *UserClaims) IsNotAuthorizedForWrite(userId string) bool {
	return !model.IsAuthorizedForWrite(userId)
}
