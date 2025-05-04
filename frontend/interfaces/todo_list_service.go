package interfaces

import "todo-frontend-web-app/models"

type TodoListService interface {
	GetById(id string) *models.TodoListGetResponseModel
	GetAllNonDeletedByUserId(userId string) *models.TodoListGetAllResponseModel
	GetAllNonDeletedWithoutUserId(userId string) *models.TodoListGetAllResponseModel
	AddWithUserIdAndName(userId string, name string) *models.EmptyResponseModel
	DeleteById(id string) *models.EmptyResponseModel
}
