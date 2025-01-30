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

type CreateRoleResponse struct {
	ID uuid.UUID `json:"id"`
}

type GetRoleResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedBy string    `json:"created_by"`
	CreatedAt string    `json:"created_at"`
}

type ListRolesResponse []GetRoleResponse

type CreatePermissionResponse struct {
	ID uuid.UUID `json:"id"`
}

type GetPermissionResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedBy string    `json:"created_by"`
	CreatedAt string    `json:"created_at"`
}

type ListPermissionsResponse []GetPermissionResponse

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
