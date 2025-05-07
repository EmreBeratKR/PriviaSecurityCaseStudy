package models

import "privia-sec-case-study/shared"

type EmptyResponseModel struct {
	shared.StatusModel
	Message string `json:"message"`
}
