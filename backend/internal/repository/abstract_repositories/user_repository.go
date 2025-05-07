package abstract_repositories

import (
	"privia-sec-case-study/backend/internal/domain"
	"privia-sec-case-study/shared"
)

type UserRepository interface {
	GetByUsername(username string) *GetUserResponse
}

type GetUserResponse struct {
	shared.StatusModel
	Message string      `json:"message"`
	User    domain.User `json:"user"`
}
