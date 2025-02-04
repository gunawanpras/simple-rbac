package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/core/permission/domain"
)

type Service interface {
	// CreatePermission creates a new permission and return the generated ID
	CreatePermission(ctx context.Context, name, description string) (res domain.Permission, err error)

	// ListPermissions retrieves a list of all permissions
	ListPermissions(ctx context.Context) (res domain.ListPermissions, err error)

	// GetPermissionByID retrieves a single permission by ID
	GetPermissionByID(ctx context.Context, id uuid.UUID) (res domain.Permission, err error)
}
