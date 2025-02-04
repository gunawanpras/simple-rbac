package postgres

import (
	"time"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/core/permission/domain"
)

type Permission struct {
	ID          uuid.UUID  `db:"id"`
	Name        string     `db:"name"`
	Description *string    `db:"description"`
	CreatedAt   time.Time  `db:"created_at"`
	CreatedBy   string     `db:"created_by"`
	UpdatedAt   *time.Time `db:"updated_at"`
	UpdatedBy   *string    `db:"updated_by"`
}

func (p Permission) Validate() bool {
	if p.ID == uuid.Nil {
		return false
	}

	if p.Name == "" {
		return false
	}

	if p.Description != nil && *p.Description == "" {
		return false
	}

	if p.CreatedAt.IsZero() {
		return false
	}

	if p.CreatedBy == "" {
		return false
	}

	return true
}

func (p Permission) ToModel() domain.Permission {
	return domain.Permission{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		CreatedAt:   p.CreatedAt,
		CreatedBy:   p.CreatedBy,
	}
}

type ListPermissions []Permission

func (l ListPermissions) Validate() bool {
	for _, p := range l {
		if !p.Validate() {
			return false
		}
	}
	return true
}

func (l ListPermissions) ToModel() domain.ListPermissions {
	var permissions domain.ListPermissions
	for _, p := range l {
		permissions = append(permissions, p.ToModel())
	}
	return permissions
}
