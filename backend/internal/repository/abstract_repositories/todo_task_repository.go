package abstract_repositories

import (
	"privia-sec-case-study/backend/internal/domain"
	"privia-sec-case-study/shared"
)

type TodoTaskRepository interface {
	GetNonDeletedById(id string) *GetTodoTaskResponse
	GetAllNonDeletedByListId(listId string) *GetAllTodoTasksResponse
	AddWithListIdAndContent(listId string, content string) *GetTodoTaskResponse
	ToggleIsCompletedById(id string) *GetTodoTaskResponse
	UpdateContentById(id string, content string) *GetTodoTaskResponse
	DeleteById(id string) *GetTodoTaskResponse
}

type GetTodoTaskResponse struct {
	shared.StatusModel
	Message  string          `json:"message"`
	TodoTask domain.TodoTask `json:"todo_task"`
}

type GetAllTodoTasksResponse struct {
	shared.StatusModel
	Message   string            `json:"message"`
	TodoTasks []domain.TodoTask `json:"todo_tasks"`
}
