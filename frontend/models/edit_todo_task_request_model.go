package models

type EditTodoTaskRequestModel struct {
	Id      string `form:"id"`
	Content string `form:"content"`
}
