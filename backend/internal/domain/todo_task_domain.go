package domain

import "time"

type TodoTask struct {
	Id          string     `json:"id"`
	TodoListId  string     `json:"todo_list_id"`
	CreatedAt   time.Time  `json:"created_at"`
	ModifiedAt  time.Time  `json:"modified_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	Content     string     `json:"content"`
	IsCompleted bool       `json:"is_completed"`
}

func (model *TodoTask) IsDeleted() bool {
	return model.DeletedAt != nil
}
