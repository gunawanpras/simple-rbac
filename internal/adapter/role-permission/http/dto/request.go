package dto

import "github.com/google/uuid"

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
