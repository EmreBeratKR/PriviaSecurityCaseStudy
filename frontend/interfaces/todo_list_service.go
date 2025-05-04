package interfaces

import "todo-frontend-web-app/models"

type TodoListService interface {
	GetNonDeletedById(id string) *models.TodoListGetResponseModel
	GetAllNonDeleted() *models.TodoListGetAllResponseModel
	GetAllNonDeletedByUserId(userId string) *models.TodoListGetAllResponseModel
	AddWithUserIdAndName(userId string, name string) *models.EmptyResponseModel
	UpdateNameById(id string, name string) *models.EmptyResponseModel
	DeleteById(id string) *models.EmptyResponseModel
}
