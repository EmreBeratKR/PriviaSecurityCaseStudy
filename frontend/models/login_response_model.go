package models

type LoginResponseModel struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (response *LoginResponseModel) IsSuccess() bool {
	return response.Status == "success"
}
