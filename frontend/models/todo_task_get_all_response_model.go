package models

type TodoTaskGetAllResponseModel struct {
	Status    string          `json:"status"`
	Message   string          `json:"message"`
	TodoTasks []TodoTaskModel `json:"todo_tasks"`
}

func (response *TodoTaskGetAllResponseModel) IsSuccess() bool {
	return response.Status == "success"
}
