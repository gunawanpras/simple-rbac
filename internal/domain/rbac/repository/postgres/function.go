package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/helper"
	"github.com/gunawanpras/simple-rbac/internal/domain/rbac/model"
	"github.com/gunawanpras/simple-rbac/util/constant"
)

// Create a new user and return the generated ID
func (repo *RbacRepository) CreateUser(ctx context.Context, user model.User) (id uuid.UUID, err error) {
	user.ID = helper.UUIDHelper.New()

	repo.prepareCreateUser()
	_, err = repo.statement.CreateUser.ExecContext(ctx, user.ID, user.Name, user.Email, user.Password, user.RoleID, user.CreatedBy, user.CreatedAt)
	if err != nil {
		return uuid.Nil, err
	}
	return user.ID, nil
}

// Retrieve a list of all users
func (repo *RbacRepository) ListUsers(ctx context.Context) (users model.ListUsers, err error) {
	repo.prepareListUsers()
	rows, err := repo.statement.ListUsers.QueryxContext(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(constant.DbRecordNotFound)
		}

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err = rows.StructScan(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Retrieve a single user by ID
func (repo *RbacRepository) GetUserByID(ctx context.Context, id uuid.UUID) (user model.User, err error) {
	repo.prepareGetUserByID()
	err = repo.statement.GetUserByID.QueryRowxContext(ctx, id).StructScan(&user)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New(constant.DbRecordNotFound)
		}

		return user, err
	}

	return user, nil
}

func (repo *RbacRepository) GetUserByEmail(ctx context.Context, email string) (user model.User, err error) {
	repo.prepareGetUserByEmail()
	err = repo.statement.GetUserByEmail.QueryRowxContext(ctx, email).StructScan(&user)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New(constant.DbRecordNotFound)
		}

		return user, err
	}

	return user, nil
}

// Create a new role and return the generated ID
func (repo *RbacRepository) CreateRole(ctx context.Context, role model.Role) (id uuid.UUID, err error) {
	role.ID = helper.UUIDHelper.New()

	repo.prepareCreateRole()
	_, err = repo.statement.CreateRole.ExecContext(ctx, role.ID, role.Name, role.CreatedBy, role.CreatedAt)
	if err != nil {
		return uuid.Nil, err
	}
	return role.ID, nil
}

// Retrieve a list of all roles
func (repo *RbacRepository) ListRoles(ctx context.Context) (roles model.ListRoles, err error) {
	repo.prepareListRoles()
	rows, err := repo.statement.ListRoles.QueryxContext(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(constant.DbRecordNotFound)
		}

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var role model.Role
		if err = rows.StructScan(&role); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

// Retrieve a single role by ID
func (repo *RbacRepository) GetRoleByID(ctx context.Context, id uuid.UUID) (role model.Role, err error) {
	repo.prepareGetRoleByID()
	err = repo.statement.GetRoleByID.QueryRowxContext(ctx, id).StructScan(&role)
	if err != nil {
		if err == sql.ErrNoRows {
			return role, errors.New(constant.DbRecordNotFound)
		}

		return role, err
	}
	return role, nil
}

func (repo *RbacRepository) GetRoleByName(ctx context.Context, name string) (role model.Role, err error) {
	repo.prepareGetRoleByName()
	err = repo.statement.GetRoleByName.QueryRowxContext(ctx, name).StructScan(&role)
	if err != nil {
		if err == sql.ErrNoRows {
			return role, errors.New(constant.DbRecordNotFound)
		}

		return role, err
	}
	return role, nil
}

// Create a new permission and return the generated ID
func (repo *RbacRepository) CreatePermission(ctx context.Context, permission model.Permission) (id uuid.UUID, err error) {
	permission.ID = helper.UUIDHelper.New()

	repo.prepareCreatePermission()
	_, err = repo.statement.CreatePermission.ExecContext(ctx, permission.ID, permission.Name, permission.CreatedBy, permission.CreatedAt)
	if err != nil {
		return uuid.Nil, err
	}
	return permission.ID, nil
}

// Retrieve a list of all permissions
func (repo *RbacRepository) ListPermissions(ctx context.Context) (permissions model.ListPermissions, err error) {
	repo.prepareListPermissions()
	rows, err := repo.statement.ListPermissions.QueryxContext(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(constant.DbRecordNotFound)
		}

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var permission model.Permission
		if err = rows.StructScan(&permission); err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}
	return permissions, nil
}

// Retrieve a single permission by ID
func (repo *RbacRepository) GetPermissionByID(ctx context.Context, id uuid.UUID) (permission model.Permission, err error) {
	repo.prepareGetPermissionByID()
	err = repo.statement.GetPermissionByID.QueryRowxContext(ctx, id).StructScan(&permission)
	if err != nil {
		if err == sql.ErrNoRows {
			return permission, errors.New(constant.DbRecordNotFound)
		}

		return permission, err
	}
	return permission, nil
}

func (repo *RbacRepository) GetPermissionByName(ctx context.Context, name string) (permission model.Permission, err error) {
	repo.prepareGetPermissionByName()
	err = repo.statement.GetPermissionByName.QueryRowxContext(ctx, name).StructScan(&permission)
	if err != nil {
		if err == sql.ErrNoRows {
			return permission, errors.New(constant.DbRecordNotFound)
		}

		return permission, err
	}
	return permission, nil
}

func (repo *RbacRepository) CreateRolePermission(ctx context.Context, rolePermission model.RolePermission) (roleID, permissionID uuid.UUID, err error) {
	repo.prepareCreateRolePermission()
	_, err = repo.statement.CreateRolePermission.ExecContext(ctx, rolePermission.RoleID, rolePermission.PermissionID, rolePermission.CreatedBy, rolePermission.CreatedAt)
	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}
	return rolePermission.RoleID, rolePermission.PermissionID, nil
}

func (repo *RbacRepository) ListRolePermissions(ctx context.Context) (rolePermissions model.ListRolePermissions, err error) {
	repo.prepareListRolePermissions()
	rows, err := repo.statement.ListRolePermissions.QueryxContext(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(constant.DbRecordNotFound)
		}

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var rolePermission model.RolePermission
		if err = rows.StructScan(&rolePermission); err != nil {
			return nil, err
		}
		rolePermissions = append(rolePermissions, rolePermission)
	}
	return rolePermissions, nil
}

func (repo *RbacRepository) ListRolePermissionsByRoleID(ctx context.Context, roleID uuid.UUID) (rolePermissions model.ListRolePermissions, err error) {
	repo.prepareGetRolePermissionByRoleID()
	rows, err := repo.statement.ListRolePermissionByRoleID.QueryxContext(ctx, roleID)
	if err != nil {
		if err == sql.ErrNoRows {
			return rolePermissions, errors.New(constant.DbRecordNotFound)
		}

		return rolePermissions, err
	}
	defer rows.Close()

	for rows.Next() {
		var rolePermission model.RolePermission
		if err = rows.StructScan(&rolePermission); err != nil {
			return nil, err
		}
		rolePermissions = append(rolePermissions, rolePermission)
	}

	return rolePermissions, nil
}

func (repo *RbacRepository) GetRolePermissionByPermissionID(ctx context.Context, roleID, permissionID uuid.UUID) (rolePermission model.RolePermission, err error) {
	repo.prepareGetRolePermissionByPermissionID()
	err = repo.statement.GetRolePermissionByPermissionID.QueryRowxContext(ctx, roleID, permissionID).StructScan(&rolePermission)
	if err != nil {
		if err == sql.ErrNoRows {
			return rolePermission, errors.New(constant.DbRecordNotFound)
		}

		return rolePermission, err
	}
	return rolePermission, nil
}
