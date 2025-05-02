package models

type TodoListGetResponseModel struct {
	Status   string        `json:"status"`
	Message  string        `json:"message"`
	TodoList TodoListModel `json:"todo_list"`
}

func (response *TodoListGetResponseModel) IsSuccess() bool {
	return response.Status == "success"
}
