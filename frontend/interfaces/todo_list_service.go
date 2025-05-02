package interfaces

import "todo-frontend-web-app/models"

type TodoListService interface {
	GetById(id string) *models.TodoListGetResponseModel
	GetAllByUserId(userId string) *models.TodoListGetAllResponseModel
}
