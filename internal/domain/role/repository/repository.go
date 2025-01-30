package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/domain/role/model"
)

type Repository interface {
	// Create a new role and return the generated ID
	CreateRole(ctx context.Context, role model.Role) (id uuid.UUID, err error)

	// Retrieve a list of all roles
	ListRoles(ctx context.Context) (roles model.ListRoles, err error)

	// Retrieve a single role by ID
	GetRoleByID(ctx context.Context, id uuid.UUID) (role model.Role, err error)

	// Retrieve a single role by name
	GetRoleByName(ctx context.Context, name string) (role model.Role, err error)
}
