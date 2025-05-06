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
