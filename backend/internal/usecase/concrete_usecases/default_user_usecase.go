package concrete_usecases

import (
	"privia-sec-case-study/backend/internal/repository/abstract_repositories"
	"privia-sec-case-study/shared"
)

type DefaultUserUsecase struct {
	repository abstract_repositories.UserRepository
}

func NewDefaultUserUsecase(repository abstract_repositories.UserRepository) *DefaultUserUsecase {
	return &DefaultUserUsecase{
		repository: repository,
	}
}

func (usecase *DefaultUserUsecase) GetUserWithUsernameAndPassword(username string, password string) *abstract_repositories.GetUserResponse {
	response := usecase.repository.GetByUsername(username)
	if response.IsNotSuccess() {
		return &abstract_repositories.GetUserResponse{
			StatusModel: response.StatusModel,
			Message:     response.Message,
		}
	}

	user := response.User
	isValidPassword := shared.ComparePasswordAndHash(password, user.Hash)

	if isValidPassword {
		return &abstract_repositories.GetUserResponse{
			StatusModel: shared.StatusSuccess(),
			User:        user,
		}
	}

	return &abstract_repositories.GetUserResponse{
		StatusModel: shared.StatusUnauthorized(),
		Message:     "Wrong credentials",
	}
}
