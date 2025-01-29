package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/domain/rbac/model"
)

type Repository interface {
	// Create a new user and return the generated ID
	CreateUser(ctx context.Context, user model.User) (id uuid.UUID, err error)

	// Retrieve a list of all users
	ListUsers(ctx context.Context) (users model.ListUsers, err error)

	// Retrieve a single user by ID
	GetUserByID(ctx context.Context, id uuid.UUID) (user model.User, err error)

	// GetUserByEmail retrieves a single user by email
	GetUserByEmail(ctx context.Context, email string) (user model.User, err error)

	// Create a new role and return the generated ID
	CreateRole(ctx context.Context, role model.Role) (id uuid.UUID, err error)

	// Retrieve a list of all roles
	ListRoles(ctx context.Context) (roles model.ListRoles, err error)

	// Retrieve a single role by ID
	GetRoleByID(ctx context.Context, id uuid.UUID) (role model.Role, err error)

	// Retrieve a single role by name
	GetRoleByName(ctx context.Context, name string) (role model.Role, err error)

	// Create a new permission and return the generated ID
	CreatePermission(ctx context.Context, permission model.Permission) (id uuid.UUID, err error)

	// Retrieve a list of all permissions
	ListPermissions(ctx context.Context) (permissions model.ListPermissions, err error)

	// Retrieve a single permission by ID
	GetPermissionByID(ctx context.Context, id uuid.UUID) (permission model.Permission, err error)

	// Retrieve a single permission by name
	GetPermissionByName(ctx context.Context, name string) (permission model.Permission, err error)

	// Create a new role permission
	CreateRolePermission(ctx context.Context, rolePermission model.RolePermission) (roleID, permissionID uuid.UUID, err error)

	// Retrieve a list of all role permissions
	ListRolePermissions(ctx context.Context) (rolePermissions model.ListRolePermissions, err error)

	// Retrieve a list of all role permissions by role ID
	ListRolePermissionsByRoleID(ctx context.Context, roleID uuid.UUID) (rolePermissions model.ListRolePermissions, err error)

	// Retrieve a role permission by permission ID
	GetRolePermissionByPermissionID(ctx context.Context, roleID, permissionID uuid.UUID) (rolePermission model.RolePermission, err error)
}
