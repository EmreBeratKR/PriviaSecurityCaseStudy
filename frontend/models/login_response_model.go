package models

import (
	"privia-sec-case-study/shared"
)

type LoginResponseModel struct {
	shared.StatusModel
	Message string `json:"message"`
	Token   string `json:"value"`
}
