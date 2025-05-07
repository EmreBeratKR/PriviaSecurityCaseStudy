package abstract_usecases

import "privia-sec-case-study/backend/internal/repository/abstract_repositories"

type TodoTaskUsecase interface {
	GetNonDeletedById(id string) *abstract_repositories.GetTodoTaskResponse
	GetAllNonDeletedByListId(listId string) *abstract_repositories.GetAllTodoTasksResponse
	AddWithListIdAndContent(listId string, content string) *abstract_repositories.GetTodoTaskResponse
	ToggleIsCompletedById(id string) *abstract_repositories.GetTodoTaskResponse
	UpdateContentById(id string, content string) *abstract_repositories.GetTodoTaskResponse
	DeleteById(id string) *abstract_repositories.GetTodoTaskResponse
}
