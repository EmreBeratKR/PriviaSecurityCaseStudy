package models

type TodoTaskGetResponseModel struct {
	StatusModel
	Message  string        `json:"message"`
	TodoTask TodoTaskModel `json:"todo_task"`
}
