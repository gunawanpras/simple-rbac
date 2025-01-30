package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `db:"id"`
	Name      string     `db:"name"`
	Email     string     `db:"email"`
	Password  string     `db:"password"`
	RoleID    *uuid.UUID `db:"role_id"`
	CreatedAt time.Time  `db:"created_at"`
	CreatedBy string     `db:"created_by"`
	UpdatedAt *time.Time `db:"updated_at"`
	UpdatedBy *string    `db:"updated_by"`
}

func (user User) ToResponse() GetUserResponse {
	return GetUserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		RoleID:    user.RoleID,
		CreatedBy: user.CreatedBy,
		CreatedAt: user.CreatedAt.String(),
	}
}

type ListUsers []User

func (users ListUsers) ToResponse() ListUsersResponse {
	var res ListUsersResponse
	for _, user := range users {
		res = append(res, GetUserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			RoleID:    user.RoleID,
			CreatedBy: user.CreatedBy,
			CreatedAt: user.CreatedAt.String(),
		})
	}
	return res
}
