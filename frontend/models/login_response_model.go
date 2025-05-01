package models

import "time"

type LoginResponseModel struct {
	Status    string    `json:"status"`
	Message   string    `json:"message"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`
}

func (response *LoginResponseModel) IsSuccess() bool {
	return response.Status == "success"
}
