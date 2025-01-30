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

type Role struct {
	ID        uuid.UUID  `db:"id"`
	Name      string     `db:"name"`
	CreatedAt time.Time  `db:"created_at"`
	CreatedBy string     `db:"created_by"`
	UpdatedAt *time.Time `db:"updated_at"`
	UpdatedBy *string    `db:"updated_by"`
}

func (role Role) ToResponse() GetRoleResponse {
	return GetRoleResponse{
		ID:        role.ID,
		Name:      role.Name,
		CreatedBy: role.CreatedBy,
		CreatedAt: role.CreatedAt.String(),
	}
}

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

type RolePermission struct {
	RoleID       uuid.UUID `db:"role_id"`
	PermissionID uuid.UUID `db:"permission_id"`
	CreatedAt    time.Time `db:"created_at"`
	CreatedBy    string    `db:"created_by"`
}

func (rolePermission RolePermission) ToResponse() GetRolePermissionResponse {
	return GetRolePermissionResponse{
		RoleID:       rolePermission.RoleID,
		PermissionID: rolePermission.PermissionID,
		CreatedBy:    rolePermission.CreatedBy,
		CreatedAt:    rolePermission.CreatedAt.String(),
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

type ListRoles []Role

func (roles ListRoles) ToResponse() ListRolesResponse {
	var res ListRolesResponse
	for _, role := range roles {
		res = append(res, GetRoleResponse{
			ID:        role.ID,
			Name:      role.Name,
			CreatedBy: role.CreatedBy,
			CreatedAt: role.CreatedAt.String(),
		})
	}
	return res
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

type ListRolePermissions []RolePermission

func (rolePermissions ListRolePermissions) ToResponse() ListRolePermissionsResponse {
	var res ListRolePermissionsResponse
	for _, rolePermission := range rolePermissions {
		res = append(res, GetRolePermissionResponse{
			RoleID:       rolePermission.RoleID,
			PermissionID: rolePermission.PermissionID,
			CreatedBy:    rolePermission.CreatedBy,
			CreatedAt:    rolePermission.CreatedAt.String(),
		})
	}
	return res
}
