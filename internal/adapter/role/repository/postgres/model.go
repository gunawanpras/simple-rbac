package postgres

import (
	"time"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/core/role/domain"
)

type Role struct {
	ID        uuid.UUID  `db:"id"`
	Name      string     `db:"name"`
	CreatedAt time.Time  `db:"created_at"`
	CreatedBy string     `db:"created_by"`
	UpdatedAt *time.Time `db:"updated_at"`
	UpdatedBy *string    `db:"updated_by"`
}

func (r Role) Validate() bool {
	if r.ID == uuid.Nil {
		return false
	}

	if r.Name == "" {
		return false
	}

	if r.CreatedAt.IsZero() {
		return false
	}

	if r.CreatedBy == "" {
		return false
	}

	return true
}

func (r Role) ToModel() domain.Role {
	return domain.Role{
		ID:        r.ID,
		Name:      r.Name,
		CreatedAt: r.CreatedAt,
		CreatedBy: r.CreatedBy,
		UpdatedAt: r.UpdatedAt,
		UpdatedBy: r.UpdatedBy,
	}
}

type ListRoles []Role

func (r ListRoles) Validate() bool {
	for _, role := range r {
		if !role.Validate() {
			return false
		}
	}
	return true
}

func (r ListRoles) ToModel() domain.ListRoles {
	var roles domain.ListRoles
	for _, role := range r {
		roles = append(roles, role.ToModel())
	}
	return roles
}
