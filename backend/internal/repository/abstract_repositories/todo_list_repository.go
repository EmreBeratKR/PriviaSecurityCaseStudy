package abstract_repositories

import (
	"privia-sec-case-study/backend/internal/domain"
	"privia-sec-case-study/shared"
)

type TodoListRepository interface {
	GetById(id string) *GetTodoListResponse
	GetAllNonDeleted() *GetAllTodoListResponse
	GetAllNonDeletedByUserId(userId string) *GetAllTodoListResponse
	AddWithUserIdAndName(userId string, name string) *GetTodoListResponse
	UpdateNameById(id string, name string) *GetTodoListResponse
	DeleteById(id string) *GetTodoListResponse
}

type GetTodoListResponse struct {
	shared.StatusModel
	Message  string          `json:"message"`
	TodoList domain.TodoList `json:"todo_list"`
}

type GetAllTodoListResponse struct {
	shared.StatusModel
	Message   string            `json:"message"`
	TodoLists []domain.TodoList `json:"todo_lists"`
}
