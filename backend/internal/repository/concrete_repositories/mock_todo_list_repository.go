package concrete_repositories

import (
	"privia-sec-case-study/backend/internal/domain"
	"privia-sec-case-study/backend/internal/repository/abstract_repositories"
	"privia-sec-case-study/shared"
	"strconv"
	"time"
)

type MockTodoListRepository struct {
	todoLists     []domain.TodoList
	todoListCount int
}

func NewMockTodoListRepository() *MockTodoListRepository {
	repository := &MockTodoListRepository{}

	repository.todoLists = []domain.TodoList{
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
	repository.todoListCount = len(repository.todoLists)
	return repository
}

func (repo *MockTodoListRepository) GetNonDeletedById(id string) *abstract_repositories.GetTodoListResponse {
	return repo.getByFilter(func(todoList *domain.TodoList) bool {
		return !todoList.IsDeleted() && todoList.Id == id
	})
}

func (repo *MockTodoListRepository) GetAllNonDeleted() *abstract_repositories.GetAllTodoListsResponse {
	return repo.getAllByFilter(func(todoList *domain.TodoList) bool {
		return !todoList.IsDeleted()
	})
}

func (repo *MockTodoListRepository) GetAllNonDeletedByUserId(userId string) *abstract_repositories.GetAllTodoListsResponse {
	return repo.getAllByFilter(func(todoList *domain.TodoList) bool {
		return todoList.UserId == userId && !todoList.IsDeleted()
	})
}

func (repo *MockTodoListRepository) UpdateNameById(id string, name string) *abstract_repositories.GetTodoListResponse {
	for i, _ := range repo.todoLists {
		if repo.todoLists[i].Id == id {
			repo.todoLists[i].Name = string([]byte(name)) // to fix fiber.Ctx.formValue bug
			repo.todoLists[i].ModifiedAt = time.Now()
			return &abstract_repositories.GetTodoListResponse{
				StatusModel: shared.StatusSuccess(),
				TodoList:    repo.todoLists[i],
			}
		}
	}

	return &abstract_repositories.GetTodoListResponse{
		StatusModel: shared.StatusNotFound(),
		Message:     "no todo list found",
	}
}

func (repo *MockTodoListRepository) AddWithUserIdAndName(userId string, name string) *abstract_repositories.GetTodoListResponse {
	return repo.add(func(todoList *domain.TodoList) {
		todoList.UserId = string([]byte(userId)) // to fix fiber.Ctx.formValue bug
		todoList.Name = string([]byte(name))     // to fix fiber.Ctx.formValue bug
	})
}

func (repo *MockTodoListRepository) DeleteById(id string) *abstract_repositories.GetTodoListResponse {
	for i, _ := range repo.todoLists {
		if repo.todoLists[i].Id == id {
			time := time.Now()
			repo.todoLists[i].ModifiedAt = time
			repo.todoLists[i].DeletedAt = &time
			return &abstract_repositories.GetTodoListResponse{
				StatusModel: shared.StatusSuccess(),
				TodoList:    repo.todoLists[i],
			}
		}
	}

	return &abstract_repositories.GetTodoListResponse{
		StatusModel: shared.StatusNotFound(),
		Message:     "no todo list found",
	}
}

func (repo *MockTodoListRepository) getByFilter(filter func(*domain.TodoList) bool) *abstract_repositories.GetTodoListResponse {
	for _, todoList := range repo.todoLists {
		if filter(&todoList) {
			return &abstract_repositories.GetTodoListResponse{
				StatusModel: shared.StatusSuccess(),
				TodoList:    todoList,
			}
		}
	}

	return &abstract_repositories.GetTodoListResponse{
		StatusModel: shared.StatusNotFound(),
		Message:     "Todo list not found",
	}
}

func (repo *MockTodoListRepository) getAllByFilter(filter func(*domain.TodoList) bool) *abstract_repositories.GetAllTodoListsResponse {
	var result = make([]domain.TodoList, 0)

	for _, todoList := range repo.todoLists {
		if filter(&todoList) {
			result = append(result, todoList)
		}
	}

	return &abstract_repositories.GetAllTodoListsResponse{
		StatusModel: shared.StatusSuccess(),
		TodoLists:   result,
	}
}

func (repo *MockTodoListRepository) add(modifier func(*domain.TodoList)) *abstract_repositories.GetTodoListResponse {
	todoList := domain.TodoList{
		Id:                strconv.Itoa(repo.todoListCount),
		Name:              "New todo list",
		CreatedAt:         time.Now(),
		ModifiedAt:        time.Now(),
		DeletedAt:         nil,
		CompletionPercent: 0,
		CompletedTasks:    0,
		TotalTasks:        0,
	}

	modifier(&todoList)
	repo.todoLists = append(repo.todoLists, todoList)
	repo.todoListCount += 1

	return &abstract_repositories.GetTodoListResponse{
		StatusModel: shared.StatusSuccess(),
		TodoList:    todoList,
	}
}

func (repo *MockTodoListRepository) modifyById(id string, modifier func(*domain.TodoList)) *abstract_repositories.GetTodoListResponse {
	for i, _ := range repo.todoLists {
		if repo.todoLists[i].Id == id {
			modifier(&repo.todoLists[i])
			repo.todoLists[i].ModifiedAt = time.Now()
			return &abstract_repositories.GetTodoListResponse{
				StatusModel: shared.StatusSuccess(),
				TodoList:    repo.todoLists[i],
			}
		}
	}

	return &abstract_repositories.GetTodoListResponse{
		StatusModel: shared.StatusNotFound(),
		Message:     "todo list not found",
	}
}
