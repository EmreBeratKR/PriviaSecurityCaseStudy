package interfaces

import "todo-frontend-web-app/models"

type TodoListService interface {
	GetById(id string) *models.TodoListGetResponseModel
	GetAllNonDeletedByUserId(userId string) *models.TodoListGetAllResponseModel
	AddWithUserId(userId string) *models.EmptyResponseModel
	DeleteById(id string) *models.EmptyResponseModel
}
