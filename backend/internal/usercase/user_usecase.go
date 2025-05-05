package usercase

import "todo-backend-rest-api/internal/domain"

type UserUsecase struct {
	repository domain.UserRepository
}

func NewUserUsecase(repository domain.UserRepository) *UserUsecase {
	return &UserUsecase{
		repository: repository,
	}
}

func (usecase *UserUsecase) GetUserWithUsernameAndHash(username string, hash string) *domain.GetUserResponse {
	response := usecase.repository.GetAll()
	if response.IsNotSuccess() {
		return &domain.GetUserResponse{
			StatusModel: response.StatusModel,
		}
	}

	users := response.Users
	for _, user := range users {
		if user.Username != username {
			continue
		}

		if user.Hash != hash {
			continue
		}

		return &domain.GetUserResponse{
			StatusModel: domain.StatusSuccess(),
			User:        user,
		}
	}

	return &domain.GetUserResponse{
		StatusModel: domain.StatusNotFound(),
	}
}
