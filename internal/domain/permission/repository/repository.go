package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/domain/permission/model"
)

type Repository interface {
	// Create a new permission and return the generated ID
	CreatePermission(ctx context.Context, permission model.Permission) (id uuid.UUID, err error)

	// Retrieve a list of all permissions
	ListPermissions(ctx context.Context) (permissions model.ListPermissions, err error)

	// Retrieve a single permission by ID
	GetPermissionByID(ctx context.Context, id uuid.UUID) (permission model.Permission, err error)

	// Retrieve a single permission by name
	GetPermissionByName(ctx context.Context, name string) (permission model.Permission, err error)
}
