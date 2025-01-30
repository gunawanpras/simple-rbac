package model

import (
	"time"

	"github.com/google/uuid"
)

type RolePermission struct {
	RoleID       uuid.UUID `db:"role_id"`
	PermissionID uuid.UUID `db:"permission_id"`
	CreatedAt    time.Time `db:"created_at"`
	CreatedBy    string    `db:"created_by"`
}

func (rolePermission RolePermission) ToResponse() GetRolePermissionResponse {
	return GetRolePermissionResponse{
		RoleID:       rolePermission.RoleID,
		PermissionID: rolePermission.PermissionID,
		CreatedBy:    rolePermission.CreatedBy,
		CreatedAt:    rolePermission.CreatedAt.String(),
	}
}

type ListRolePermissions []RolePermission

func (rolePermissions ListRolePermissions) ToResponse() ListRolePermissionsResponse {
	var res ListRolePermissionsResponse
	for _, rolePermission := range rolePermissions {
		res = append(res, GetRolePermissionResponse{
			RoleID:       rolePermission.RoleID,
			PermissionID: rolePermission.PermissionID,
			CreatedBy:    rolePermission.CreatedBy,
			CreatedAt:    rolePermission.CreatedAt.String(),
		})
	}
	return res
}
