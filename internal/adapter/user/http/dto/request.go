package dto

import "github.com/google/uuid"

type CreateUserRequest struct {
	Name     string    `json:"name" validate:"required,min=3,max=100"`
	Email    string    `json:"email" validate:"required,email"`
	Password string    `json:"password" validate:"required,min=8"`
	RoleID   uuid.UUID `json:"role_id" validate:"required,uuid"`
}

type GetUserByIDRequest struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

type CreateRoleRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}
