package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/domain/user/model"
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
}
