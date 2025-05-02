package models

type TodoListGetAllResponseModel struct {
	Status    string          `json:"status"`
	Message   string          `json:"message"`
	TodoLists []TodoListModel `json:"todo_lists"`
}

func (response *TodoListGetAllResponseModel) IsSuccess() bool {
	return response.Status == "success"
}
