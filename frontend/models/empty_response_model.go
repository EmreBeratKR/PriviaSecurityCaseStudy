package models

type EmptyResponseModel struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (response *EmptyResponseModel) IsSuccess() bool {
	return response.Status == "success"
}
