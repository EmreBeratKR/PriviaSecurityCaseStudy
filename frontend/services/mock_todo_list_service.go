package services

import (
	"strconv"
	"time"
	"todo-frontend-web-app/common"
	"todo-frontend-web-app/models"
)

type MockTodoListService struct {
	ServiceManager *ServiceManager
	TodoLists      []models.TodoListModel
	TodoListCount  int
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
	claims := common.GetUserClaims(service.ServiceManager.Context)
	for _, todoList := range service.TodoLists {
		if todoList.Id == id {
			if todoList.UserId != claims.Subject && !claims.IsAdmin() {
				return &models.TodoListGetResponseModel{
					StatusModel: models.StatusForbidden(),
					Message:     "You do not have access",
				}
			}

			return &models.TodoListGetResponseModel{
				StatusModel: models.StatusSuccess(),
				TodoList:    todoList,
				Message:     "todo list found",
			}
		}
	}
	return &models.TodoListGetResponseModel{
		StatusModel: models.StatusNotFound(),
		Message:     "todo list not found",
	}
}

func (service *MockTodoListService) GetAllNonDeleted() *models.TodoListGetAllResponseModel {
	isAdmin := common.IsAuthenticatedAsAdmin(service.ServiceManager.Context)

	if !isAdmin {
		return &models.TodoListGetAllResponseModel{
			StatusModel: models.StatusForbidden(),
			Message:     "You do not have access",
		}
	}

	return service.getAll(func(model *models.TodoListModel) bool {
		if model.IsDeleted() {
			return false
		}

		return true
	})
}

func (service *MockTodoListService) GetAllNonDeletedByUserId(userId string) *models.TodoListGetAllResponseModel {
	authUserId := common.GetAuthUserId(service.ServiceManager.Context)

	if userId != authUserId {
		return &models.TodoListGetAllResponseModel{
			StatusModel: models.StatusForbidden(),
			Message:     "You do not have access",
		}
	}

	return service.getAll(func(model *models.TodoListModel) bool {
		if model.IsDeleted() {
			return false
		}

		if model.UserId != userId {
			return false
		}

		return true
	})
}

func (service *MockTodoListService) AddWithUserIdAndName(userId string, name string) *models.EmptyResponseModel {
	authUserId := common.GetAuthUserId(service.ServiceManager.Context)

	if userId != authUserId {
		return &models.EmptyResponseModel{
			StatusModel: models.StatusForbidden(),
			Message:     "You do not have access",
		}
	}

	id := strconv.Itoa(service.TodoListCount)
	service.TodoLists = append(service.TodoLists, models.TodoListModel{
		Id:                id,
		UserId:            userId,
		Name:              string([]byte(name)), // to fix fiber.Ctx.FormValue bug
		CreatedAt:         time.Now(),
		ModifiedAt:        time.Now(),
		DeletedAt:         nil,
		CompletionPercent: 0,
		CompletedTasks:    0,
		TotalTasks:        0,
	})

	service.TodoListCount += 1

	return &models.EmptyResponseModel{
		StatusModel: models.StatusSuccess(),
		Message:     "todo list added",
	}
}

func (service *MockTodoListService) UpdateNameById(id string, name string) *models.EmptyResponseModel {
	authUserId := common.GetAuthUserId(service.ServiceManager.Context)
	for i, todoList := range service.TodoLists {
		if todoList.Id == id {
			if todoList.UserId != authUserId {
				return &models.EmptyResponseModel{
					StatusModel: models.StatusForbidden(),
					Message:     "You do not have access",
				}
			}

			service.TodoLists[i].Name = string([]byte(name)) // to fix fiber.Ctx.FormValue bug
			service.TodoLists[i].ModifiedAt = time.Now()
			return &models.EmptyResponseModel{
				StatusModel: models.StatusSuccess(),
				Message:     "todo list name updated",
			}
		}
	}

	return &models.EmptyResponseModel{
		StatusModel: models.StatusNotFound(),
		Message:     "todo list not found",
	}
}

func (service *MockTodoListService) DeleteById(id string) *models.EmptyResponseModel {
	authUserId := common.GetAuthUserId(service.ServiceManager.Context)
	for i, todoList := range service.TodoLists {
		if todoList.Id == id && todoList.DeletedAt == nil {
			if todoList.UserId != authUserId {
				return &models.EmptyResponseModel{
					StatusModel: models.StatusForbidden(),
					Message:     "You do not have access",
				}
			}
			now := time.Now()
			service.TodoLists[i].DeletedAt = &now
			return &models.EmptyResponseModel{
				StatusModel: models.StatusSuccess(),
				Message:     "todo list deleted",
			}
		}
	}

	return &models.EmptyResponseModel{
		StatusModel: models.StatusNotFound(),
		Message:     "todo list not found or already deleted",
	}
}

func (service *MockTodoListService) getAll(filter func(*models.TodoListModel) bool) *models.TodoListGetAllResponseModel {
	var filtered = make([]models.TodoListModel, 0)
	for _, todoList := range service.TodoLists {
		if filter(&todoList) {
			filtered = append(filtered, todoList)
		}
	}
	return &models.TodoListGetAllResponseModel{
		StatusModel: models.StatusSuccess(),
		TodoLists:   filtered,
		Message:     "todo lists are found",
	}
}
