package domain

import (
	"time"
)

type TodoList struct {
	Id                string     `json:"id"`
	UserId            string     `json:"user_id"`
	Name              string     `json:"name"`
	CreatedAt         time.Time  `json:"created_at"`
	ModifiedAt        time.Time  `json:"modified_at"`
	DeletedAt         *time.Time `json:"deleted_at"`
	CompletionPercent int        `json:"completion_percent"`
	CompletedTasks    int        `json:"completed_tasks"`
	TotalTasks        int        `json:"total_tasks"`
}

func (todoList *TodoList) IsDeleted() bool {
	return todoList.DeletedAt != nil
}

func (model *TodoList) UpdateModifiedAt() {
	model.ModifiedAt = time.Now()
}

func (model *TodoList) UpdateCompletionPercent() {
	if model.TotalTasks <= 0 {
		model.CompletionPercent = 0
		return
	}
	model.CompletionPercent = (model.CompletedTasks * 100) / model.TotalTasks
}
