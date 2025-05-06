package concrete_usecases

import (
	"privia-sec-case-study/backend/internal/repository/abstract_repositories"
	"privia-sec-case-study/shared"
)

type TodoListUsecase struct {
	repository abstract_repositories.TodoListRepository
}

func NewDefaultTodoListUsecase(repository abstract_repositories.TodoListRepository) *TodoListUsecase {
	return &TodoListUsecase{
		repository: repository,
	}
}

func (usecase *TodoListUsecase) GetNonDeletedById(id string) *abstract_repositories.GetTodoListResponse {
	response := usecase.repository.GetById(id)
	if response.IsNotSuccess() {
		return &abstract_repositories.GetTodoListResponse{
			StatusModel: response.StatusModel,
			Message:     response.Message,
		}
	}

	todoList := response.TodoList
	if todoList.IsDeleted() {
		return &abstract_repositories.GetTodoListResponse{
			StatusModel: shared.StatusNotFound(),
			Message:     "Todo list was deleted",
		}
	}

	return &abstract_repositories.GetTodoListResponse{
		StatusModel: shared.StatusSuccess(),
		TodoList:    todoList,
	}
}

func (usecase *TodoListUsecase) GetAllNonDeleted() *abstract_repositories.GetAllTodoListResponse {
	return usecase.repository.GetAllNonDeleted()
}

func (usecase *TodoListUsecase) GetAllNonDeletedByUserId(userId string) *abstract_repositories.GetAllTodoListResponse {
	return usecase.repository.GetAllNonDeletedByUserId(userId)
}

func (usecase *TodoListUsecase) AddWithUserIdAndName(userId string, name string) *abstract_repositories.GetTodoListResponse {
	return usecase.repository.AddWithUserIdAndName(userId, name)
}

func (usecase *TodoListUsecase) UpdateNameById(id string, name string) *abstract_repositories.GetTodoListResponse {
	return usecase.repository.UpdateNameById(id, name)
}

func (usecase *TodoListUsecase) DeleteById(id string) *abstract_repositories.GetTodoListResponse {
	return usecase.repository.DeleteById(id)
}
