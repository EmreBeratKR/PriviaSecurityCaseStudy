package concrete_usecases

import "privia-sec-case-study/backend/internal/repository/abstract_repositories"

type DefaultTodoTaskUsecase struct {
	repository abstract_repositories.TodoTaskRepository
}

func NewDefaultTodoTaskUsecase(repository abstract_repositories.TodoTaskRepository) *DefaultTodoTaskUsecase {
	return &DefaultTodoTaskUsecase{
		repository: repository,
	}
}

func (usecase *DefaultTodoTaskUsecase) GetNonDeletedById(id string) *abstract_repositories.GetTodoTaskResponse {
	return usecase.repository.GetNonDeletedById(id)
}

func (usecase *DefaultTodoTaskUsecase) GetAllNonDeletedByListId(listId string) *abstract_repositories.GetAllTodoTasksResponse {
	return usecase.repository.GetAllNonDeletedByListId(listId)
}

func (usecase *DefaultTodoTaskUsecase) AddWithListIdAndContent(listId string, content string) *abstract_repositories.GetTodoTaskResponse {
	return usecase.repository.AddWithListIdAndContent(listId, content)
}

func (usecase *DefaultTodoTaskUsecase) ToggleIsCompletedById(id string) *abstract_repositories.GetTodoTaskResponse {
	return usecase.repository.ToggleIsCompletedById(id)
}

func (usecase *DefaultTodoTaskUsecase) UpdateContentById(id string, content string) *abstract_repositories.GetTodoTaskResponse {
	return usecase.repository.UpdateContentById(id, content)
}

func (usecase *DefaultTodoTaskUsecase) DeleteById(id string) *abstract_repositories.GetTodoTaskResponse {
	return usecase.repository.DeleteById(id)
}
