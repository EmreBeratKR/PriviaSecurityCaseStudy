package concrete_repositories

import (
	"privia-sec-case-study/backend/internal/domain"
	"privia-sec-case-study/backend/internal/repository/abstract_repositories"
	"privia-sec-case-study/shared"
	"strconv"
)

type MockUserRepository struct {
	users     []domain.User
	userCount int
}

func NewMockUserRepository() *MockUserRepository {
	repo := &MockUserRepository{}
	repo.createUser("user1", "1234", "user")
	repo.createUser("user2", "1234", "user")
	repo.createUser("user3", "1234", "user")
	repo.createUser("user4", "1234", "user")
	repo.createUser("user5", "1234", "user")
	repo.createUser("user6", "1234", "user")
	repo.createUser("Emre", "admin1234", "admin")
	repo.createUser("Berat", "admin1234", "admin")
	repo.createUser("Simge", "admin1234", "admin")
	return repo
}

func (repo *MockUserRepository) GetByUsername(username string) *abstract_repositories.GetUserResponse {
	for _, user := range repo.users {
		if user.Username == username {
			return &abstract_repositories.GetUserResponse{
				StatusModel: shared.StatusSuccess(),
				User:        user,
			}
		}
	}

	return &abstract_repositories.GetUserResponse{
		StatusModel: shared.StatusNotFound(),
		Message:     "User does not exist",
	}
}

func (repo *MockUserRepository) createUser(username string, password string, role string) {
	id := strconv.Itoa(repo.userCount)
	repo.users = append(repo.users, domain.User{
		Id:       id,
		Username: username,
		Hash:     shared.GeneratePasswordHash(password),
		Role:     role,
	})
	repo.userCount += 1
}
