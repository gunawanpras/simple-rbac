package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/domain/role-permission/model"
	"github.com/gunawanpras/simple-rbac/util/constant"
)

func (repo *RolePermissionRepository) CreateRolePermission(ctx context.Context, rolePermission model.RolePermission) (roleID, permissionID uuid.UUID, err error) {
	repo.prepareCreateRolePermission()
	_, err = repo.statement.CreateRolePermission.ExecContext(ctx, rolePermission.RoleID, rolePermission.PermissionID, rolePermission.CreatedBy, rolePermission.CreatedAt)
	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}
	return rolePermission.RoleID, rolePermission.PermissionID, nil
}

func (repo *RolePermissionRepository) ListRolePermissions(ctx context.Context) (rolePermissions model.ListRolePermissions, err error) {
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

func (repo *RolePermissionRepository) ListRolePermissionsByRoleID(ctx context.Context, roleID uuid.UUID) (rolePermissions model.ListRolePermissions, err error) {
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

func (repo *RolePermissionRepository) GetRolePermissionByPermissionID(ctx context.Context, roleID, permissionID uuid.UUID) (rolePermission model.RolePermission, err error) {
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
