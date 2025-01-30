package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/domain/role-permission/model"
)

type Service interface {
	// CreateRolePermission creates a new role permission
	CreateRolePermission(ctx context.Context, roleID, permissionID uuid.UUID) (res model.ServiceResponse, err error)

	// ListRolePermissions retrieves a list of all role permissions
	ListRolePermissions(ctx context.Context) (res model.ServiceResponse, err error)

	// ListRolePermissionsByRoleID retrieves a list of all role permissions by role ID
	ListRolePermissionsByRoleID(ctx context.Context, roleID uuid.UUID) (res model.ServiceResponse, err error)

	// GetRolePermissionByPermissionID retrieves a single role permission by permission ID
	GetRolePermissionByPermissionID(ctx context.Context, roleID, permissionID uuid.UUID) (res model.ServiceResponse, err error)
}
