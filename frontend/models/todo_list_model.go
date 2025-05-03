package models

import "time"

type TodoListModel struct {
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

func (model *TodoListModel) UpdateModifiedAt() {
	model.ModifiedAt = time.Now()
}

func (model *TodoListModel) UpdateCompletionPercent() {
	if model.TotalTasks <= 0 {
		model.CompletionPercent = 0
		return
	}
	model.CompletionPercent = (model.CompletedTasks * 100) / model.TotalTasks
}

func (model *TodoListModel) GetModifiedAtFormatted() string {
	return model.ModifiedAt.Format("02-01-2006 15:04")
}

func (model *TodoListModel) GetRemainingTasks() int {
	return model.TotalTasks - model.CompletedTasks
}

func (model *TodoListModel) IsDeleted() bool {
	return model.DeletedAt != nil
}
