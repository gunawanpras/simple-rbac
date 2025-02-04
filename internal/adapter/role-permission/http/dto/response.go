package dto

import (
	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/core/role-permission/domain"
)

type CreateRolePermissionResponse struct {
	RoleID       uuid.UUID `json:"role_id"`
	PermissionID uuid.UUID `json:"permission_id"`
}

func (r *CreateRolePermissionResponse) ToResponse(rolePermission domain.RolePermission) {
	*r = CreateRolePermissionResponse{
		RoleID:       rolePermission.RoleID,
		PermissionID: rolePermission.PermissionID,
	}
}

type GetRolePermissionResponse struct {
	RoleID       uuid.UUID `json:"role_id"`
	PermissionID uuid.UUID `json:"permission_id"`
	CreatedBy    string    `json:"created_by"`
	CreatedAt    string    `json:"created_at"`
}

func (r *GetRolePermissionResponse) ToResponse(rolePermission domain.RolePermission) {
	*r = GetRolePermissionResponse{
		RoleID:       rolePermission.RoleID,
		PermissionID: rolePermission.PermissionID,
		CreatedBy:    rolePermission.CreatedBy,
		CreatedAt:    rolePermission.CreatedAt.String(),
	}
}

type ListRolePermissionsResponse []GetRolePermissionResponse

func (r *ListRolePermissionsResponse) ToResponse(rolePermissions domain.ListRolePermissions) {
	for _, rolePermission := range rolePermissions {
		*r = append(*r, GetRolePermissionResponse{
			RoleID:       rolePermission.RoleID,
			PermissionID: rolePermission.PermissionID,
			CreatedBy:    rolePermission.CreatedBy,
			CreatedAt:    rolePermission.CreatedAt.String(),
		})
	}
}
