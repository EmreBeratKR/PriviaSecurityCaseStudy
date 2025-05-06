package models

import "privia-sec-case-study/shared"

type TodoTaskGetResponseModel struct {
	shared.StatusModel
	Message  string        `json:"message"`
	TodoTask TodoTaskModel `json:"value"`
}
