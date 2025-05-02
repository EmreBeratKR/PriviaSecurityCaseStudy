package services

import (
	"time"
	"todo-frontend-web-app/models"
)

type MockTodoTaskService struct {
	TodoTasks []models.TodoTaskModel
}

func (service *MockTodoTaskService) Init() {
	service.TodoTasks = []models.TodoTaskModel{
		{
			Id:          "0",
			TodoListId:  "0",
			CreatedAt:   time.Now(),
			ModifiedAt:  time.Now(),
			DeletedAt:   nil,
			Content:     "Finish the UI layout",
			IsCompleted: true,
		},
		{
			Id:          "1",
			TodoListId:  "0",
			CreatedAt:   time.Now(),
			ModifiedAt:  time.Now(),
			DeletedAt:   nil,
			Content:     "Write the backend logic",
			IsCompleted: false,
		},
		{
			Id:          "2",
			TodoListId:  "0",
			CreatedAt:   time.Now(),
			ModifiedAt:  time.Now(),
			DeletedAt:   nil,
			Content:     "Connect frontend to backend",
			IsCompleted: true,
		},
		{
			Id:          "3",
			TodoListId:  "1",
			CreatedAt:   time.Now(),
			ModifiedAt:  time.Now(),
			DeletedAt:   nil,
			Content:     "Wake up early",
			IsCompleted: true,
		},
		{
			Id:          "4",
			TodoListId:  "1",
			CreatedAt:   time.Now(),
			ModifiedAt:  time.Now(),
			DeletedAt:   nil,
			Content:     "Go to gym",
			IsCompleted: true,
		},
		{
			Id:          "5",
			TodoListId:  "2",
			CreatedAt:   time.Now(),
			ModifiedAt:  time.Now(),
			DeletedAt:   nil,
			Content:     "Bread",
			IsCompleted: false,
		},
		{
			Id:          "6",
			TodoListId:  "2",
			CreatedAt:   time.Now(),
			ModifiedAt:  time.Now(),
			DeletedAt:   nil,
			Content:     "Tomato",
			IsCompleted: false,
		},
		{
			Id:          "7",
			TodoListId:  "2",
			CreatedAt:   time.Now(),
			ModifiedAt:  time.Now(),
			DeletedAt:   nil,
			Content:     "Eggs",
			IsCompleted: true,
		},
		{
			Id:          "8",
			TodoListId:  "2",
			CreatedAt:   time.Now(),
			ModifiedAt:  time.Now(),
			DeletedAt:   nil,
			Content:     "Oranges",
			IsCompleted: false,
		},
	}
}

func (service *MockTodoTaskService) GetAllByTodoListId(todoListId string) *models.TodoTaskGetAllResponseModel {
	filtered := make([]models.TodoTaskModel, 0)

	for _, task := range service.TodoTasks {
		if task.TodoListId == todoListId && task.DeletedAt == nil {
			filtered = append(filtered, task)
		}
	}

	return &models.TodoTaskGetAllResponseModel{
		Status:    "success",
		Message:   "Tasks retrieved successfully",
		TodoTasks: filtered,
	}
}
