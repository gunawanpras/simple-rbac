package model

import (
	"time"

	"github.com/google/uuid"
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

func (permission Permission) ToResponse() GetPermissionResponse {
	return GetPermissionResponse{
		ID:        permission.ID,
		Name:      permission.Name,
		CreatedBy: permission.CreatedBy,
		CreatedAt: permission.CreatedAt.String(),
	}
}

type ListPermissions []Permission

func (permissions ListPermissions) ToResponse() ListPermissionsResponse {
	var res ListPermissionsResponse
	for _, permission := range permissions {
		res = append(res, GetPermissionResponse{
			ID:        permission.ID,
			Name:      permission.Name,
			CreatedBy: permission.CreatedBy,
			CreatedAt: permission.CreatedAt.String(),
		})
	}
	return res
}
