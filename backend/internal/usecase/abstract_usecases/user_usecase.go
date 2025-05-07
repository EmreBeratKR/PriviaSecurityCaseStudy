package abstract_usecases

import (
	"privia-sec-case-study/backend/internal/repository/abstract_repositories"
)

type UserUsecase interface {
	GetUserWithUsernameAndPassword(username string, password string) *abstract_repositories.GetUserResponse
}
