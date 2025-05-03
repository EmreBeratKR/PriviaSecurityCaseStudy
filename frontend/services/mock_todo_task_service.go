package services

import (
	"strconv"
	"time"
	"todo-frontend-web-app/models"
)

type MockTodoTaskService struct {
	TodoListService *MockTodoListService
	TodoTasks       []models.TodoTaskModel
	TodoTaskCount   int
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
	service.TodoTaskCount = len(service.TodoTasks)
}

func (service *MockTodoTaskService) GetAllNonDeletedByTodoListId(todoListId string) *models.TodoTaskGetAllResponseModel {
	filtered := make([]models.TodoTaskModel, 0)

	for _, task := range service.TodoTasks {
		if task.TodoListId == todoListId && !task.IsDeleted() {
			filtered = append(filtered, task)
		}
	}

	return &models.TodoTaskGetAllResponseModel{
		Status:    "success",
		Message:   "Tasks retrieved successfully",
		TodoTasks: filtered,
	}
}

func (service *MockTodoTaskService) AddWithListIdAndContent(todoListId string, content string) *models.EmptyResponseModel {
	id := strconv.Itoa(service.TodoTaskCount)
	service.TodoTasks = append(service.TodoTasks, models.TodoTaskModel{
		Id:          id,
		TodoListId:  string([]byte(todoListId)), // to fix fiber.Ctx.BodyParser bug
		CreatedAt:   time.Now(),
		ModifiedAt:  time.Now(),
		DeletedAt:   nil,
		Content:     string([]byte(content)), // to fix fiber.Ctx.BodyParser bug
		IsCompleted: false,
	})

	service.TodoTaskCount += 1

	todoListService := service.TodoListService
	for i := range todoListService.TodoLists {
		if todoListService.TodoLists[i].Id == todoListId {
			todoListService.TodoLists[i].TotalTasks += 1
			todoListService.TodoLists[i].UpdateCompletionPercent()
			todoListService.TodoLists[i].UpdateModifiedAt()
		}
	}

	return &models.EmptyResponseModel{
		Status:  "success",
		Message: "todo task added",
	}
}

func (service *MockTodoTaskService) DeleteById(id string) *models.EmptyResponseModel {
	for i, todoList := range service.TodoTasks {
		if todoList.Id == id && todoList.DeletedAt == nil {
			now := time.Now()
			service.TodoTasks[i].DeletedAt = &now

			isCompleted := service.TodoTasks[i].IsCompleted
			todoListId := service.TodoTasks[i].TodoListId
			todoListService := service.TodoListService
			for i := range todoListService.TodoLists {
				if todoListService.TodoLists[i].Id == todoListId {
					todoListService.TodoLists[i].TotalTasks -= 1
					if isCompleted {
						todoListService.TodoLists[i].CompletedTasks -= 1
					}
					todoListService.TodoLists[i].UpdateCompletionPercent()
					todoListService.TodoLists[i].UpdateModifiedAt()
				}
			}

			return &models.EmptyResponseModel{
				Status:  "success",
				Message: "todo task deleted",
			}
		}
	}

	return &models.EmptyResponseModel{
		Status:  "not_found",
		Message: "todo task not found or already deleted",
	}
}
