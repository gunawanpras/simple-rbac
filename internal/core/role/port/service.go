package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/core/role/domain"
)

type Service interface {
	// CreateRole creates a new role and return the generated ID

	CreateRole(ctx context.Context, name string) (res domain.Role, err error)
	// ListRoles retrieves a list of all roles
	ListRoles(ctx context.Context) (res domain.ListRoles, err error)

	// GetRoleByID retrieves a single role by ID
	GetRoleByID(ctx context.Context, id uuid.UUID) (res domain.Role, err error)
}
