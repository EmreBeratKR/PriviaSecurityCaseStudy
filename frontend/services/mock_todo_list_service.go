package services

import (
	"strconv"
	"time"
	"todo-frontend-web-app/models"
)

type MockTodoListService struct {
	TodoLists     []models.TodoListModel
	TodoListCount int
}

func (service *MockTodoListService) Init() {
	service.TodoLists = []models.TodoListModel{
		{
			Id:                "0",
			UserId:            "0",
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
			UserId:            "0",
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
			UserId:            "1",
			Name:              "Shopping List",
			CreatedAt:         time.Now(),
			ModifiedAt:        time.Now(),
			DeletedAt:         nil,
			CompletionPercent: 25,
			CompletedTasks:    1,
			TotalTasks:        4,
		},
	}
	service.TodoListCount = len(service.TodoLists)
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

func (service *MockTodoListService) GetAllNonDeletedByUserId(userId string) *models.TodoListGetAllResponseModel {
	var filtered = make([]models.TodoListModel, 0)
	for _, todo := range service.TodoLists {
		if todo.UserId == userId && !todo.IsDeleted() {
			filtered = append(filtered, todo)
		}
	}
	return &models.TodoListGetAllResponseModel{
		Status:    "success",
		TodoLists: filtered,
	}
}

func (service *MockTodoListService) AddWithUserId(userId string) *models.EmptyResponseModel {
	id := strconv.Itoa(service.TodoListCount)
	service.TodoLists = append(service.TodoLists, models.TodoListModel{
		Id:                id,
		UserId:            userId,
		Name:              "New Todo List #" + id,
		CreatedAt:         time.Now(),
		ModifiedAt:        time.Now(),
		DeletedAt:         nil,
		CompletionPercent: 0,
		CompletedTasks:    0,
		TotalTasks:        0,
	})

	service.TodoListCount += 1

	return &models.EmptyResponseModel{
		Status:  "success",
		Message: "todo list added",
	}
}

func (service *MockTodoListService) DeleteById(id string) *models.EmptyResponseModel {
	for i, todoList := range service.TodoLists {
		if todoList.Id == id && todoList.DeletedAt == nil {
			now := time.Now()
			service.TodoLists[i].DeletedAt = &now
			return &models.EmptyResponseModel{
				Status:  "success",
				Message: "todo list deleted",
			}
		}
	}

	return &models.EmptyResponseModel{
		Status:  "not_found",
		Message: "todo list not found or already deleted",
	}
}
