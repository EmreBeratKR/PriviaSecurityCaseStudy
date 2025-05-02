package interfaces

import "todo-frontend-web-app/models"

type TodoTaskService interface {
	GetAllByTodoListId(todoListId string) *models.TodoTaskGetAllResponseModel
}
