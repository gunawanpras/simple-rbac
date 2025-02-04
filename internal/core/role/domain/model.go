package domain

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt *time.Time
	UpdatedBy *string
}

type ListRoles []Role
