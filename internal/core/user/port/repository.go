package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/core/user/domain"
)

type Repository interface {
	// Create a new user and return the generated ID
	CreateUser(ctx context.Context, user domain.User) (id uuid.UUID, err error)

	// Retrieve a list of all users
	ListUsers(ctx context.Context) (users domain.ListUsers, err error)

	// Retrieve a single user by ID
	GetUserByID(ctx context.Context, id uuid.UUID) (user domain.User, err error)

	// GetUserByEmail retrieves a single user by email
	GetUserByEmail(ctx context.Context, email string) (user domain.User, err error)
}
