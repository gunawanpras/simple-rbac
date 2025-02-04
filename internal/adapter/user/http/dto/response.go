package dto

import (
	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/core/user/domain"
)

type CreateUserResponse struct {
	ID uuid.UUID `json:"id"`
}

type GetUserResponse struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	RoleID    *uuid.UUID `json:"role_id"`
	CreatedBy string     `json:"created_by"`
	CreatedAt string     `json:"created_at"`
}

func (u *GetUserResponse) ToResponse(user domain.User) {
	*u = GetUserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		RoleID:    user.RoleID,
		CreatedBy: user.CreatedBy,
		CreatedAt: user.CreatedAt.String(),
	}
}

type ListUsersResponse []GetUserResponse

func (us *ListUsersResponse) ToResponse(users domain.ListUsers) {
	for _, user := range users {
		*us = append(*us, GetUserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			RoleID:    user.RoleID,
			CreatedBy: user.CreatedBy,
			CreatedAt: user.CreatedAt.String(),
		})
	}
}
