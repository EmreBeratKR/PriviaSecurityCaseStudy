package models

type TodoTaskGetResponseModel struct {
	Status   string        `json:"status"`
	Message  string        `json:"message"`
	TodoTask TodoTaskModel `json:"todo_task"`
}

func (response *TodoTaskGetResponseModel) IsSuccess() bool {
	return response.Status == "success"
}
