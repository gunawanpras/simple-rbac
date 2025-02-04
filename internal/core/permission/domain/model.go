package domain

import (
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	ID          uuid.UUID
	Name        string
	Description *string
	CreatedAt   time.Time
	CreatedBy   string
}

type ListPermissions []Permission
