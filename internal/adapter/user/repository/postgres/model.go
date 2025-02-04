package postgres

import (
	"time"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/core/user/domain"
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

func (u User) Validate() bool {
	if u.ID == uuid.Nil {
		return false
	}

	if u.Name == "" {
		return false
	}

	if u.Email == "" {
		return false
	}

	if u.CreatedAt.IsZero() {
		return false
	}

	if u.CreatedBy == "" {
		return false
	}

	if u.UpdatedAt != nil && u.UpdatedAt.IsZero() {
		return false
	}

	if u.UpdatedBy != nil && *u.UpdatedBy == "" {
		return false
	}

	return true
}

func (u User) ToModel() domain.User {
	return domain.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		RoleID:    u.RoleID,
		CreatedAt: u.CreatedAt,
		CreatedBy: u.CreatedBy,
		UpdatedAt: u.UpdatedAt,
		UpdatedBy: u.UpdatedBy,
	}
}

type ListUsers []User

func (users ListUsers) Validate() bool {
	for _, u := range users {
		if u.ID == uuid.Nil {
			return false
		}

		if u.Name == "" {
			return false
		}

		if u.Email == "" {
			return false
		}

		if u.CreatedAt.IsZero() {
			return false
		}

		if u.CreatedBy == "" {
			return false
		}

		if u.UpdatedAt != nil && u.UpdatedAt.IsZero() {
			return false
		}

		if u.UpdatedBy != nil && *u.UpdatedBy == "" {
			return false
		}
	}

	return true
}

func (users ListUsers) ToModel() domain.ListUsers {
	var res domain.ListUsers
	for _, user := range users {
		u := domain.User{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			RoleID:    user.RoleID,
			CreatedAt: user.CreatedAt,
			CreatedBy: user.CreatedBy,
			UpdatedAt: user.UpdatedAt,
			UpdatedBy: user.UpdatedBy,
		}
		res = append(res, u)
	}
	return res
}
