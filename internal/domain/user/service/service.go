package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/domain/user/model"
)

type Service interface {
	// CreateUser creates a new user and return the generated ID
	CreateUser(ctx context.Context, name, email, password string, roleId uuid.UUID) (res model.ServiceResponse, err error)

	// ListUsers retrieves a list of all users
	ListUsers(ctx context.Context) (res model.ServiceResponse, err error)

	// GetUserByID retrieves a single user by ID
	GetUserByID(ctx context.Context, id uuid.UUID) (res model.ServiceResponse, err error)
}
