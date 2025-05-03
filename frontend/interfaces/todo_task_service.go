package interfaces

import "todo-frontend-web-app/models"

type TodoTaskService interface {
	GetById(id string) *models.TodoTaskGetResponseModel
	GetAllNonDeletedByTodoListId(todoListId string) *models.TodoTaskGetAllResponseModel
	AddWithListIdAndContent(todoListId string, content string) *models.EmptyResponseModel
	DeleteById(id string) *models.EmptyResponseModel
	ToggleIsCompletedById(id string) *models.EmptyResponseModel
	UpdateContentById(id string, content string) *models.EmptyResponseModel
}
