package models

import "privia-sec-case-study/shared"

type TodoListGetAllResponseModel struct {
	shared.StatusModel
	Message   string          `json:"message"`
	TodoLists []TodoListModel `json:"todo_lists"`
}

func (model *TodoListGetAllResponseModel) Filtered(filter func(*TodoListModel) bool) []TodoListModel {
	var filtered = make([]TodoListModel, 0)
	for _, todoList := range model.TodoLists {
		if filter(&todoList) {
			filtered = append(filtered, todoList)
		}
	}
	return filtered
}
