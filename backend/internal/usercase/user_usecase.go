package usercase

import (
	"privia-sec-case-study/backend/internal/domain"
	"privia-sec-case-study/shared"
)

type UserUsecase struct {
	repository domain.UserRepository
}

func NewUserUsecase(repository domain.UserRepository) *UserUsecase {
	return &UserUsecase{
		repository: repository,
	}
}

func (usecase *UserUsecase) GetUserWithUsernameAndPassword(username string, password string) *domain.GetUserResponse {
	response := usecase.repository.GetByUsername(username)
	if response.IsNotSuccess() {
		return &domain.GetUserResponse{
			StatusModel: response.StatusModel,
		}
	}

	user := response.User
	isValidPassword := shared.ComparePasswordAndHash(password, user.Hash)

	if isValidPassword {
		return &domain.GetUserResponse{
			StatusModel: shared.StatusSuccess(),
			User:        user,
		}
	}

	return &domain.GetUserResponse{
		StatusModel: shared.StatusUnauthorized(),
		Message:     "Wrong credentials",
	}
}
