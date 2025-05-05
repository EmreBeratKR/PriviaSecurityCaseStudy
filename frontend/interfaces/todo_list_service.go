package interfaces

import "privia-sec-case-study/frontend/models"

type TodoListService interface {
	GetNonDeletedById(id string) *models.TodoListGetResponseModel
	GetAllNonDeleted() *models.TodoListGetAllResponseModel
	GetAllNonDeletedByUserId(userId string) *models.TodoListGetAllResponseModel
	AddWithUserIdAndName(userId string, name string) *models.TodoListGetResponseModel
	UpdateNameById(id string, name string) *models.TodoListGetResponseModel
	DeleteById(id string) *models.TodoListGetResponseModel
}
