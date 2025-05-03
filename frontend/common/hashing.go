package common

import (
	"os"

	"golang.org/x/crypto/bcrypt"
)

var passwordPepper = os.Getenv("PASSWORD_PEPPER")

func GeneratePasswordHash(password string) string {
	peppered := password + passwordPepper
	hashed, _ := bcrypt.GenerateFromPassword([]byte(peppered), bcrypt.DefaultCost)
	return string(hashed)
}

func ComparePasswordAndHash(password string, hash string) bool {
	peppered := password + passwordPepper
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(peppered))
	return err == nil
}
