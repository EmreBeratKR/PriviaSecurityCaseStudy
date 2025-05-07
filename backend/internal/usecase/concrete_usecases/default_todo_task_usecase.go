package concrete_usecases

import "privia-sec-case-study/backend/internal/repository/abstract_repositories"

type DefaultTodoTaskUsecase struct {
	listRepository abstract_repositories.TodoListRepository
	taskRepository abstract_repositories.TodoTaskRepository
}

func NewDefaultTodoTaskUsecase(listRepository abstract_repositories.TodoListRepository, taskRepository abstract_repositories.TodoTaskRepository) *DefaultTodoTaskUsecase {
	return &DefaultTodoTaskUsecase{
		listRepository: listRepository,
		taskRepository: taskRepository,
	}
}

func (usecase *DefaultTodoTaskUsecase) GetNonDeletedById(id string) *abstract_repositories.GetTodoTaskResponse {
	return usecase.taskRepository.GetNonDeletedById(id)
}

func (usecase *DefaultTodoTaskUsecase) GetAllNonDeletedByListId(listId string) *abstract_repositories.GetAllTodoTasksResponse {
	return usecase.taskRepository.GetAllNonDeletedByListId(listId)
}

func (usecase *DefaultTodoTaskUsecase) AddWithListIdAndContent(listId string, content string) *abstract_repositories.GetTodoTaskResponse {
	taskResponse := usecase.taskRepository.AddWithListIdAndContent(listId, content)
	if taskResponse.IsNotSuccess() {
		return taskResponse
	}

	isCompleted := taskResponse.TodoTask.IsCompleted
	listResponse := usecase.listRepository.IncrementTaskCountById(listId, isCompleted)
	if listResponse.IsNotSuccess() {
		return &abstract_repositories.GetTodoTaskResponse{
			StatusModel: listResponse.StatusModel,
			Message:     listResponse.Message,
		}
	}

	return taskResponse
}

func (usecase *DefaultTodoTaskUsecase) ToggleIsCompletedById(id string) *abstract_repositories.GetTodoTaskResponse {
	taskResponse := usecase.taskRepository.ToggleIsCompletedById(id)
	if taskResponse.IsNotSuccess() {
		return taskResponse
	}

	listId := taskResponse.TodoTask.TodoListId

	if taskResponse.TodoTask.IsCompleted {
		listResponse := usecase.listRepository.IncrementCompletedTaskCountById(listId)
		if listResponse.IsNotSuccess() {
			return &abstract_repositories.GetTodoTaskResponse{
				StatusModel: listResponse.StatusModel,
				Message:     listResponse.Message,
			}
		}
	} else {
		listResponse := usecase.listRepository.DecrementCompletedTaskCountById(listId)
		if listResponse.IsNotSuccess() {
			return &abstract_repositories.GetTodoTaskResponse{
				StatusModel: listResponse.StatusModel,
				Message:     listResponse.Message,
			}
		}
	}

	return taskResponse
}

func (usecase *DefaultTodoTaskUsecase) UpdateContentById(id string, content string) *abstract_repositories.GetTodoTaskResponse {
	taskResponse := usecase.taskRepository.UpdateContentById(id, content)
	if taskResponse.IsNotSuccess() {
		return taskResponse
	}

	listId := taskResponse.TodoTask.TodoListId
	listResponse := usecase.listRepository.UpdateModifiedAtById(listId)
	if listResponse.IsNotSuccess() {
		return &abstract_repositories.GetTodoTaskResponse{
			StatusModel: listResponse.StatusModel,
			Message:     listResponse.Message,
		}
	}

	return taskResponse
}

func (usecase *DefaultTodoTaskUsecase) DeleteById(id string) *abstract_repositories.GetTodoTaskResponse {
	taskResponse := usecase.taskRepository.DeleteById(id)
	if taskResponse.IsNotSuccess() {
		return taskResponse
	}

	listId := taskResponse.TodoTask.TodoListId
	isCompleted := taskResponse.TodoTask.IsCompleted
	listResponse := usecase.listRepository.DecrementTaskCountById(listId, isCompleted)
	if listResponse.IsNotSuccess() {
		return &abstract_repositories.GetTodoTaskResponse{
			StatusModel: listResponse.StatusModel,
			Message:     listResponse.Message,
		}
	}

	return taskResponse
}
