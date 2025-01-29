package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/domain/rbac/model"
)

type Service interface {
	// CreateUser creates a new user and return the generated ID
	CreateUser(ctx context.Context, name, email, password string, roleId uuid.UUID) (res model.ServiceResponse, err error)

	// ListUsers retrieves a list of all users
	ListUsers(ctx context.Context) (res model.ServiceResponse, err error)

	// GetUserByID retrieves a single user by ID
	GetUserByID(ctx context.Context, id uuid.UUID) (res model.ServiceResponse, err error)

	// CreateRole creates a new role and return the generated ID
	CreateRole(ctx context.Context, name string) (res model.ServiceResponse, err error)

	// ListRoles retrieves a list of all roles
	ListRoles(ctx context.Context) (res model.ServiceResponse, err error)

	// GetRoleByID retrieves a single role by ID
	GetRoleByID(ctx context.Context, id uuid.UUID) (res model.ServiceResponse, err error)

	// CreatePermission creates a new permission and return the generated ID
	CreatePermission(ctx context.Context, name, description string) (res model.ServiceResponse, err error)

	// ListPermissions retrieves a list of all permissions
	ListPermissions(ctx context.Context) (res model.ServiceResponse, err error)

	// GetPermissionByID retrieves a single permission by ID
	GetPermissionByID(ctx context.Context, id uuid.UUID) (res model.ServiceResponse, err error)

	// CreateRolePermission creates a new role permission
	CreateRolePermission(ctx context.Context, roleID, permissionID uuid.UUID) (res model.ServiceResponse, err error)

	// ListRolePermissions retrieves a list of all role permissions
	ListRolePermissions(ctx context.Context) (res model.ServiceResponse, err error)

	// ListRolePermissionsByRoleID retrieves a list of all role permissions by role ID
	ListRolePermissionsByRoleID(ctx context.Context, roleID uuid.UUID) (res model.ServiceResponse, err error)

	// GetRolePermissionByPermissionID retrieves a single role permission by permission ID
	GetRolePermissionByPermissionID(ctx context.Context, roleID, permissionID uuid.UUID) (res model.ServiceResponse, err error)
}
