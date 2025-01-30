package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/domain/permission/model"
)

type Service interface {
	// CreatePermission creates a new permission and return the generated ID
	CreatePermission(ctx context.Context, name, description string) (res model.ServiceResponse, err error)

	// ListPermissions retrieves a list of all permissions
	ListPermissions(ctx context.Context) (res model.ServiceResponse, err error)

	// GetPermissionByID retrieves a single permission by ID
	GetPermissionByID(ctx context.Context, id uuid.UUID) (res model.ServiceResponse, err error)
}
