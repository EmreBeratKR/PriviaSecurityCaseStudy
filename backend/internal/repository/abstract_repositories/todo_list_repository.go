package abstract_repositories

import (
	"privia-sec-case-study/backend/internal/domain"
	"privia-sec-case-study/shared"
)

type TodoListRepository interface {
	GetNonDeletedById(id string) *GetTodoListResponse
	GetAllNonDeleted() *GetAllTodoListsResponse
	GetAllNonDeletedByUserId(userId string) *GetAllTodoListsResponse
	AddWithUserIdAndName(userId string, name string) *GetTodoListResponse
	UpdateNameById(id string, name string) *GetTodoListResponse
	DeleteById(id string) *GetTodoListResponse
}

type GetTodoListResponse struct {
	shared.StatusModel
	Message  string          `json:"message"`
	TodoList domain.TodoList `json:"todo_list"`
}

type GetAllTodoListsResponse struct {
	shared.StatusModel
	Message   string            `json:"message"`
	TodoLists []domain.TodoList `json:"todo_lists"`
}
