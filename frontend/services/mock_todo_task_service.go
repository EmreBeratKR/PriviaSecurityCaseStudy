package services

import (
	"privia-sec-case-study/frontend/common"
	"privia-sec-case-study/frontend/models"
	"privia-sec-case-study/shared"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type MockTodoTaskService struct {
	ServiceManager  *ServiceManager
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

func (service *MockTodoTaskService) GetAllNonDeletedByTodoListId(context *fiber.Ctx, todoListId string) *models.TodoTaskGetAllResponseModel {
	todoListResponse := service.TodoListService.GetNonDeletedById(context, todoListId)

	if todoListResponse.IsNotSuccess() {
		return &models.TodoTaskGetAllResponseModel{
			StatusModel: todoListResponse.StatusModel,
		}
	}

	filtered := make([]models.TodoTaskModel, 0)

	for _, task := range service.TodoTasks {
		if task.TodoListId == todoListId && !task.IsDeleted() {
			filtered = append(filtered, task)
		}
	}

	return &models.TodoTaskGetAllResponseModel{
		StatusModel: shared.StatusSuccess(),
		Message:     "Tasks retrieved successfully",
		TodoTasks:   filtered,
	}
}

func (service *MockTodoTaskService) AddWithListIdAndContent(context *fiber.Ctx, todoListId string, content string) *models.TodoTaskGetResponseModel {
	todoListResponse := service.TodoListService.GetNonDeletedById(context, todoListId)
	if todoListResponse.IsNotSuccess() {
		return &models.TodoTaskGetResponseModel{
			StatusModel: todoListResponse.StatusModel,
		}
	}

	userId := common.GetAuthUserId(context)
	if userId != todoListResponse.TodoList.UserId {
		return &models.TodoTaskGetResponseModel{
			StatusModel: shared.StatusForbidden(),
		}
	}

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

	return &models.TodoTaskGetResponseModel{
		StatusModel: shared.StatusSuccess(),
		Message:     "todo task added",
		TodoTask:    service.TodoTasks[service.TodoTaskCount-1],
	}
}

func (service *MockTodoTaskService) DeleteById(context *fiber.Ctx, id string) *models.TodoTaskGetResponseModel {
	for i, todoTask := range service.TodoTasks {
		if todoTask.Id == id {
			todoListResponse := service.TodoListService.GetNonDeletedById(context, todoTask.TodoListId)
			if todoListResponse.IsNotSuccess() {
				return &models.TodoTaskGetResponseModel{
					StatusModel: todoListResponse.StatusModel,
				}
			}

			if todoTask.IsDeleted() {
				break
			}

			userId := common.GetAuthUserId(context)
			if userId != todoListResponse.TodoList.UserId {
				return &models.TodoTaskGetResponseModel{
					StatusModel: shared.StatusForbidden(),
				}
			}

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

			return &models.TodoTaskGetResponseModel{
				StatusModel: shared.StatusSuccess(),
				Message:     "todo task deleted",
				TodoTask:    service.TodoTasks[i],
			}
		}
	}

	return &models.TodoTaskGetResponseModel{
		StatusModel: shared.StatusNotFound(),
		Message:     "todo task not found or already deleted",
	}
}

func (service *MockTodoTaskService) ToggleIsCompletedById(context *fiber.Ctx, id string) *models.TodoTaskGetResponseModel {
	for i, todoTask := range service.TodoTasks {
		if todoTask.Id == id {
			todoListResponse := service.TodoListService.GetNonDeletedById(context, todoTask.TodoListId)
			if todoListResponse.IsNotSuccess() {
				return &models.TodoTaskGetResponseModel{
					StatusModel: todoListResponse.StatusModel,
				}
			}

			if todoTask.IsDeleted() {
				break
			}

			userId := common.GetAuthUserId(context)
			if userId != todoListResponse.TodoList.UserId {
				return &models.TodoTaskGetResponseModel{
					StatusModel: shared.StatusForbidden(),
				}
			}

			service.TodoTasks[i].ToggleIsCompleted()
			service.TodoTasks[i].UpdateModifiedAt()
			isCompleted := service.TodoTasks[i].IsCompleted
			todoListId := service.TodoTasks[i].TodoListId
			todoListService := service.TodoListService
			for i := range todoListService.TodoLists {
				if todoListService.TodoLists[i].Id == todoListId {
					if isCompleted {
						todoListService.TodoLists[i].CompletedTasks += 1
					} else {
						todoListService.TodoLists[i].CompletedTasks -= 1
					}
					todoListService.TodoLists[i].UpdateCompletionPercent()
					todoListService.TodoLists[i].UpdateModifiedAt()
				}
			}

			return &models.TodoTaskGetResponseModel{
				StatusModel: shared.StatusSuccess(),
				Message:     "todo task updated",
				TodoTask:    service.TodoTasks[i],
			}
		}
	}

	return &models.TodoTaskGetResponseModel{
		StatusModel: shared.StatusNotFound(),
		Message:     "todo task not found or already deleted",
	}
}

func (service *MockTodoTaskService) UpdateContentById(context *fiber.Ctx, id string, content string) *models.TodoTaskGetResponseModel {
	for i, todoTask := range service.TodoTasks {
		if todoTask.Id == id {
			todoListResponse := service.TodoListService.GetNonDeletedById(context, todoTask.TodoListId)
			if todoListResponse.IsNotSuccess() {
				return &models.TodoTaskGetResponseModel{
					StatusModel: todoListResponse.StatusModel,
				}
			}

			if todoTask.IsDeleted() {
				break
			}

			userId := common.GetAuthUserId(context)
			if userId != todoListResponse.TodoList.UserId {
				return &models.TodoTaskGetResponseModel{
					StatusModel: shared.StatusForbidden(),
				}
			}

			service.TodoTasks[i].Content = string([]byte(content)) // to fix fiber.Ctx.BodyParser bug

			return &models.TodoTaskGetResponseModel{
				StatusModel: shared.StatusSuccess(),
				Message:     "todo task content updated",
				TodoTask:    service.TodoTasks[i],
			}
		}
	}

	return &models.TodoTaskGetResponseModel{
		StatusModel: shared.StatusNotFound(),
		Message:     "todo task not found",
	}
}
