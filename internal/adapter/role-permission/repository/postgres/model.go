package postgres

import (
	"time"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/core/role-permission/domain"
)

type RolePermission struct {
	RoleID       uuid.UUID `db:"role_id"`
	PermissionID uuid.UUID `db:"permission_id"`
	CreatedAt    time.Time `db:"created_at"`
	CreatedBy    string    `db:"created_by"`
}

func (r *RolePermission) Validate() bool {
	if r.RoleID == uuid.Nil {
		return false
	}

	if r.PermissionID == uuid.Nil {
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

func (r *RolePermission) ToModel() domain.RolePermission {
	return domain.RolePermission{
		RoleID:       r.RoleID,
		PermissionID: r.PermissionID,
		CreatedAt:    r.CreatedAt,
		CreatedBy:    r.CreatedBy,
	}
}

type ListRolePermissions []RolePermission

func (l ListRolePermissions) Validate() bool {
	for _, rp := range l {
		if !rp.Validate() {
			return false
		}
	}
	return true
}

func (l ListRolePermissions) ToModel() domain.ListRolePermissions {
	var rolePermissions domain.ListRolePermissions
	for _, rp := range l {
		rolePermissions = append(rolePermissions, rp.ToModel())
	}
	return rolePermissions
}
