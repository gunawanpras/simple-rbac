package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/core/role/domain"
)

type Repository interface {
	// Create a new role and return the generated ID
	CreateRole(ctx context.Context, role domain.Role) (id uuid.UUID, err error)

	// Retrieve a list of all roles
	ListRoles(ctx context.Context) (res domain.ListRoles, err error)

	// Retrieve a single role by ID
	GetRoleByID(ctx context.Context, id uuid.UUID) (res domain.Role, err error)

	// Retrieve a single role by name
	GetRoleByName(ctx context.Context, name string) (res domain.Role, err error)
}
