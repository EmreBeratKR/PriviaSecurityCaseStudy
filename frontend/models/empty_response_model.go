package models

type EmptyResponseModel struct {
	StatusModel
	Message string `json:"message"`
}
