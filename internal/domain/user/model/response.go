package model

import "github.com/google/uuid"

type CreateUserResponse struct {
	ID uuid.UUID `json:"id"`
}

type GetUserResponse struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	RoleID    *uuid.UUID `json:"role_id"`
	CreatedBy string     `json:"created_by"`
	CreatedAt string     `json:"created_at"`
}

type ListUsersResponse []GetUserResponse

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ServiceResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
