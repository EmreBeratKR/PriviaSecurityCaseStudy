package interfaces

import "todo-frontend-web-app/models"

type TodoTaskService interface {
	GetAllNonDeletedByTodoListId(todoListId string) *models.TodoTaskGetAllResponseModel
}
