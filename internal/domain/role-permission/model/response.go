package model

import "github.com/google/uuid"

type CreateRolePermissionResponse struct {
	RoleID       uuid.UUID `json:"role_id"`
	PermissionID uuid.UUID `json:"permission_id"`
}

type GetRolePermissionResponse struct {
	RoleID       uuid.UUID `json:"role_id"`
	PermissionID uuid.UUID `json:"permission_id"`
	CreatedBy    string    `json:"created_by"`
	CreatedAt    string    `json:"created_at"`
}

type ListRolePermissionsResponse []GetRolePermissionResponse

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ServiceResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
