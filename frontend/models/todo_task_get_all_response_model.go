package models

import "privia-sec-case-study/shared"

type TodoTaskGetAllResponseModel struct {
	shared.StatusModel
	Message   string          `json:"message"`
	TodoTasks []TodoTaskModel `json:"todo_tasks"`
}

func (response *TodoTaskGetAllResponseModel) IsSuccess() bool {
	return response.Status == "success"
}
