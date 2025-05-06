package models

import "privia-sec-case-study/shared"

type TodoListGetResponseModel struct {
	shared.StatusModel
	Message  string        `json:"message"`
	TodoList TodoListModel `json:"todo_list"`
}
