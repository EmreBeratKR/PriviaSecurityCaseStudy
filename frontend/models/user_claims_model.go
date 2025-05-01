package models

import "github.com/golang-jwt/jwt/v4"

type UserClaims struct {
	Subject  string
	Username string
	Role     string
	jwt.RegisteredClaims
}
