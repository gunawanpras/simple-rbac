package dto

import (
	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/core/permission/domain"
)

type CreatePermissionResponse struct {
	ID uuid.UUID `json:"id"`
}

func (u *CreatePermissionResponse) ToResponse(permission domain.Permission) {
	*u = CreatePermissionResponse{
		ID: permission.ID,
	}
}

type GetPermissionResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedBy string    `json:"created_by"`
	CreatedAt string    `json:"created_at"`
}

func (u *GetPermissionResponse) ToResponse(permission domain.Permission) {
	*u = GetPermissionResponse{
		ID:        permission.ID,
		Name:      permission.Name,
		CreatedBy: permission.CreatedBy,
		CreatedAt: permission.CreatedAt.String(),
	}
}

type ListPermissionsResponse []GetPermissionResponse

func (us *ListPermissionsResponse) ToResponse(permissions domain.ListPermissions) {
	for _, permission := range permissions {
		*us = append(*us, GetPermissionResponse{
			ID:        permission.ID,
			Name:      permission.Name,
			CreatedBy: permission.CreatedBy,
			CreatedAt: permission.CreatedAt.String(),
		})
	}
}
