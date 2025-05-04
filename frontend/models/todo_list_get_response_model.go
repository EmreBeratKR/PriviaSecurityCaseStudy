package models

type TodoListGetResponseModel struct {
	StatusModel
	Message  string        `json:"message"`
	TodoList TodoListModel `json:"todo_list"`
}
