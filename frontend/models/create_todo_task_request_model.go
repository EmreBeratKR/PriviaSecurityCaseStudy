package models

type CreateTodoTaskRequestModel struct {
	ListId  string `form:"list_id"`
	Content string `form:"content"`
}
