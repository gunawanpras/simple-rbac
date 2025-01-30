package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/domain/role/model"
)

type Service interface {
	// CreateRole creates a new role and return the generated ID
	CreateRole(ctx context.Context, name string) (res model.ServiceResponse, err error)

	// ListRoles retrieves a list of all roles
	ListRoles(ctx context.Context) (res model.ServiceResponse, err error)

	// GetRoleByID retrieves a single role by ID
	GetRoleByID(ctx context.Context, id uuid.UUID) (res model.ServiceResponse, err error)
}
