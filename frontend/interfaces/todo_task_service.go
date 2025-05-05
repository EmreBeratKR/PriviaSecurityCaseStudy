package interfaces

import "privia-sec-case-study/frontend/models"

type TodoTaskService interface {
	GetAllNonDeletedByTodoListId(todoListId string) *models.TodoTaskGetAllResponseModel
	AddWithListIdAndContent(todoListId string, content string) *models.TodoTaskGetResponseModel
	DeleteById(id string) *models.TodoTaskGetResponseModel
	ToggleIsCompletedById(id string) *models.TodoTaskGetResponseModel
	UpdateContentById(id string, content string) *models.TodoTaskGetResponseModel
}
