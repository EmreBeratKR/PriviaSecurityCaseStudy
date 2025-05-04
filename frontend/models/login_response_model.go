package models

import "time"

type LoginResponseModel struct {
	StatusModel
	Message   string    `json:"message"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`
}
