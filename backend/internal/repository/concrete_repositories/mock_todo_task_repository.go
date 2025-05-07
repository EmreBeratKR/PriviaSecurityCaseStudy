package concrete_repositories

import (
	"privia-sec-case-study/backend/internal/domain"
	"privia-sec-case-study/backend/internal/repository/abstract_repositories"
	"privia-sec-case-study/shared"
	"strconv"
	"time"
)

type MockTodoTaskRepository struct {
	todoTasks     []domain.TodoTask
	todoTaskCount int
}

func NewMockTodoTaskRepository() *MockTodoTaskRepository {
	repository := &MockTodoTaskRepository{}

	repository.todoTasks = []domain.TodoTask{
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
	repository.todoTaskCount = len(repository.todoTasks)
	return repository
}

func (repo *MockTodoTaskRepository) GetNonDeletedById(id string) *abstract_repositories.GetTodoTaskResponse {
	return repo.getByFilter(func(todoTask *domain.TodoTask) bool {
		return !todoTask.IsDeleted() && todoTask.Id == id
	})
}

func (repo *MockTodoTaskRepository) GetAllNonDeletedByListId(listId string) *abstract_repositories.GetAllTodoTasksResponse {
	return repo.getAllByFilter(func(todoTask *domain.TodoTask) bool {
		return !todoTask.IsDeleted() && todoTask.TodoListId == listId
	})
}

func (repo *MockTodoTaskRepository) AddWithListIdAndContent(listId string, content string) *abstract_repositories.GetTodoTaskResponse {
	return repo.add(func(todoTask *domain.TodoTask) {
		todoTask.TodoListId = string([]byte(listId)) // to fix fiber.Ctx.formValue bug
		todoTask.Content = string([]byte(content))   // to fix fiber.Ctx.formValue bug
	})
}

func (repo *MockTodoTaskRepository) ToggleIsCompletedById(id string) *abstract_repositories.GetTodoTaskResponse {
	return repo.modifyById(id, func(todoTask *domain.TodoTask) {
		todoTask.IsCompleted = !todoTask.IsCompleted
	})
}

func (repo *MockTodoTaskRepository) UpdateContentById(id string, content string) *abstract_repositories.GetTodoTaskResponse {
	return repo.modifyById(id, func(todoTask *domain.TodoTask) {
		todoTask.Content = string([]byte(content)) // to fix fiber.Ctx.formValue bug
	})
}

func (repo *MockTodoTaskRepository) DeleteById(id string) *abstract_repositories.GetTodoTaskResponse {
	deleteTime := time.Now()
	return repo.modifyById(id, func(todoTask *domain.TodoTask) {
		todoTask.DeletedAt = &deleteTime
	})
}

func (repo *MockTodoTaskRepository) getByFilter(filter func(*domain.TodoTask) bool) *abstract_repositories.GetTodoTaskResponse {
	for _, todoTask := range repo.todoTasks {
		if filter(&todoTask) {
			return &abstract_repositories.GetTodoTaskResponse{
				StatusModel: shared.StatusSuccess(),
				TodoTask:    todoTask,
			}
		}
	}

	return &abstract_repositories.GetTodoTaskResponse{
		StatusModel: shared.StatusNotFound(),
		Message:     "todo task not found",
	}
}

func (repo *MockTodoTaskRepository) getAllByFilter(filter func(*domain.TodoTask) bool) *abstract_repositories.GetAllTodoTasksResponse {
	var result = make([]domain.TodoTask, 0)

	for _, todoTask := range repo.todoTasks {
		if filter(&todoTask) {
			result = append(result, todoTask)
		}
	}

	return &abstract_repositories.GetAllTodoTasksResponse{
		StatusModel: shared.StatusSuccess(),
		TodoTasks:   result,
	}
}

func (repo *MockTodoTaskRepository) add(modifier func(*domain.TodoTask)) *abstract_repositories.GetTodoTaskResponse {
	todoTask := domain.TodoTask{
		Id:          strconv.Itoa(repo.todoTaskCount),
		CreatedAt:   time.Now(),
		ModifiedAt:  time.Now(),
		DeletedAt:   nil,
		Content:     "new todo task",
		IsCompleted: false,
	}

	modifier(&todoTask)
	repo.todoTasks = append(repo.todoTasks, todoTask)
	repo.todoTaskCount += 1

	return &abstract_repositories.GetTodoTaskResponse{
		StatusModel: shared.StatusSuccess(),
		TodoTask:    todoTask,
	}
}

func (repo *MockTodoTaskRepository) modifyById(id string, modifier func(*domain.TodoTask)) *abstract_repositories.GetTodoTaskResponse {
	for i, _ := range repo.todoTasks {
		if repo.todoTasks[i].Id == id {
			modifier(&repo.todoTasks[i])
			repo.todoTasks[i].ModifiedAt = time.Now()
			return &abstract_repositories.GetTodoTaskResponse{
				StatusModel: shared.StatusSuccess(),
				TodoTask:    repo.todoTasks[i],
			}
		}
	}

	return &abstract_repositories.GetTodoTaskResponse{
		StatusModel: shared.StatusNotFound(),
		Message:     "todo task not found",
	}
}
