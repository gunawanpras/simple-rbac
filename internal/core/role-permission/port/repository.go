package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/core/role-permission/domain"
)

type Repository interface {
	// Create a new role permission
	CreateRolePermission(ctx context.Context, rolePermission domain.RolePermission) (roleID, permissionID uuid.UUID, err error)

	// Retrieve a list of all role permissions
	ListRolePermissions(ctx context.Context) (rolePermissions domain.ListRolePermissions, err error)

	// Retrieve a list of all role permissions by role ID
	ListRolePermissionsByRoleID(ctx context.Context, roleID uuid.UUID) (rolePermissions domain.ListRolePermissions, err error)

	// Retrieve a role permission by permission ID
	GetRolePermissionByPermissionID(ctx context.Context, roleID, permissionID uuid.UUID) (rolePermission domain.RolePermission, err error)
}
