package abstract_usecases

import "privia-sec-case-study/backend/internal/repository/abstract_repositories"

type TodoListUsecase interface {
	GetNonDeletedById(id string) *abstract_repositories.GetTodoListResponse
	GetAllNonDeleted() *abstract_repositories.GetAllTodoListsResponse
	GetAllNonDeletedByUserId(userId string) *abstract_repositories.GetAllTodoListsResponse
	AddWithUserIdAndName(userId string, name string) *abstract_repositories.GetTodoListResponse
	UpdateNameById(id string, name string) *abstract_repositories.GetTodoListResponse
	DeleteById(id string) *abstract_repositories.GetTodoListResponse
}
