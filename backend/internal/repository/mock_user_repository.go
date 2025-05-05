package repository

import "privia-sec-case-study/backend/internal/domain"

type MockUserRepository struct {
	users     []domain.User
	userCount int
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: []domain.User{
			{Id: "0", Username: "Emre"},
			{Id: "1", Username: "Berat"},
		},
		userCount: 2,
	}
}

func (repo *MockUserRepository) GetAll() *domain.GetAllUsersResponse {
	return &domain.GetAllUsersResponse{
		StatusModel: domain.StatusSuccess(),
		Users:       repo.users,
	}
}
