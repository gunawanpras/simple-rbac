package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/core/role-permission/domain"
)

type Service interface {
	// CreateRolePermission creates a new role permission
	CreateRolePermission(ctx context.Context, roleID, permissionID uuid.UUID) (res domain.RolePermission, err error)

	// ListRolePermissions retrieves a list of all role permissions
	ListRolePermissions(ctx context.Context) (res domain.ListRolePermissions, err error)

	// ListRolePermissionsByRoleID retrieves a list of all role permissions by role ID
	ListRolePermissionsByRoleID(ctx context.Context, roleID uuid.UUID) (res domain.ListRolePermissions, err error)

	// GetRolePermissionByPermissionID retrieves a single role permission by permission ID
	GetRolePermissionByPermissionID(ctx context.Context, roleID, permissionID uuid.UUID) (res domain.RolePermission, err error)
}
