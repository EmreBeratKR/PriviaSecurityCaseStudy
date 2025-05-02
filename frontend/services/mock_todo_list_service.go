package services

import (
	"time"
	"todo-frontend-web-app/models"
)

type MockTodoListService struct {
	TodoLists []models.TodoListModel
}

func (service *MockTodoListService) Init() {
	service.TodoLists = []models.TodoListModel{
		{
			Id:                "0",
			UserId:            "1234567890",
			Name:              "Work Tasks",
			CreatedAt:         time.Now(),
			ModifiedAt:        time.Now(),
			DeletedAt:         nil,
			CompletionPercent: 67,
			CompletedTasks:    2,
			TotalTasks:        3,
		},
		{
			Id:                "1",
			UserId:            "1234567890",
			Name:              "Personal Goals",
			CreatedAt:         time.Now(),
			ModifiedAt:        time.Now(),
			DeletedAt:         nil,
			CompletionPercent: 100,
			CompletedTasks:    2,
			TotalTasks:        2,
		},
		{
			Id:                "2",
			UserId:            "1234567890",
			Name:              "Shopping List",
			CreatedAt:         time.Now(),
			ModifiedAt:        time.Now(),
			DeletedAt:         nil,
			CompletionPercent: 25,
			CompletedTasks:    1,
			TotalTasks:        4,
		},
	}
}

func (service *MockTodoListService) GetById(id string) *models.TodoListGetResponseModel {
	for _, todoList := range service.TodoLists {
		if todoList.Id == id {
			return &models.TodoListGetResponseModel{
				Status:   "success",
				TodoList: todoList,
			}
		}
	}
	return &models.TodoListGetResponseModel{
		Status: "not_found",
	}
}

func (service *MockTodoListService) GetAllByUserId(userId string) *models.TodoListGetAllResponseModel {
	var filtered []models.TodoListModel
	for _, todo := range service.TodoLists {
		if todo.UserId == userId {
			filtered = append(filtered, todo)
		}
	}
	return &models.TodoListGetAllResponseModel{
		Status:    "success",
		TodoLists: filtered,
	}
}
