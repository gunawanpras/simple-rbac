package dto

import (
	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/core/role/domain"
)

type CreateRoleResponse struct {
	ID uuid.UUID `json:"id"`
}

func (u *CreateRoleResponse) ToResponse(role domain.Role) {
	*u = CreateRoleResponse{
		ID: role.ID,
	}
}

type GetRoleResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedBy string    `json:"created_by"`
	CreatedAt string    `json:"created_at"`
}

func (u *GetRoleResponse) ToResponse(role domain.Role) {
	*u = GetRoleResponse{
		ID:        role.ID,
		Name:      role.Name,
		CreatedBy: role.CreatedBy,
		CreatedAt: role.CreatedAt.String(),
	}
}

type ListRolesResponse []GetRoleResponse

func (u *ListRolesResponse) ToResponse(roles domain.ListRoles) {
	for _, role := range roles {
		*u = append(*u, GetRoleResponse{
			ID:        role.ID,
			Name:      role.Name,
			CreatedBy: role.CreatedBy,
			CreatedAt: role.CreatedAt.String(),
		})
	}
}
