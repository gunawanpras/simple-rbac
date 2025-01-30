package model

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

type GetRoleByIDRequest struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

type CreatePermissionRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"max=255"`
}

type GetPermissionByIDRequest struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

type CreateRolePermissionRequest struct {
	RoleID       uuid.UUID `json:"role_id" validate:"required,uuid"`
	PermissionID uuid.UUID `json:"permission_id" validate:"required,uuid"`
}

type ListRolePermissionsByRoleIDRequest struct {
	RoleID uuid.UUID `uri:"role_id" validate:"required,uuid"`
}

type GetRolePermissionByPermissionIDRequest struct {
	RoleID       uuid.UUID `uri:"role_id" validate:"required,uuid"`
	PermissionID uuid.UUID `uri:"permission_id" validate:"required,uuid"`
}
