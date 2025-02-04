package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/core/role-permission/domain"
	"github.com/gunawanpras/simple-rbac/pkg/util/constant"
)

func (repo *RolePermissionRepository) CreateRolePermission(ctx context.Context, rolePermission domain.RolePermission) (roleID, permissionID uuid.UUID, err error) {
	repo.prepareCreateRolePermission()
	_, err = repo.statement.CreateRolePermission.ExecContext(ctx, rolePermission.RoleID, rolePermission.PermissionID, rolePermission.CreatedBy, rolePermission.CreatedAt)
	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}

	return rolePermission.RoleID, rolePermission.PermissionID, nil
}

func (repo *RolePermissionRepository) ListRolePermissions(ctx context.Context) (res domain.ListRolePermissions, err error) {
	var (
		rolePermission  RolePermission
		rolePermissions ListRolePermissions
	)

	repo.prepareListRolePermissions()
	rows, err := repo.statement.ListRolePermissions.QueryxContext(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(constant.DataNotFound)
		}

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		rolePermission = RolePermission{}
		if err = rows.StructScan(&rolePermission); err != nil {
			return nil, err
		}
		rolePermissions = append(rolePermissions, rolePermission)
	}

	if !rolePermissions.Validate() {
		return nil, errors.New(constant.DbReturnedMalformedData)
	}

	return rolePermissions.ToModel(), nil
}

func (repo *RolePermissionRepository) ListRolePermissionsByRoleID(ctx context.Context, roleID uuid.UUID) (res domain.ListRolePermissions, err error) {
	var (
		rolePermission  RolePermission
		rolePermissions ListRolePermissions
	)

	repo.prepareGetRolePermissionByRoleID()
	rows, err := repo.statement.ListRolePermissionByRoleID.QueryxContext(ctx, roleID)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.New(constant.DataNotFound)
		}

		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		rolePermission = RolePermission{}
		if err = rows.StructScan(&rolePermission); err != nil {
			return nil, err
		}
		rolePermissions = append(rolePermissions, rolePermission)
	}

	if !rolePermissions.Validate() {
		return nil, errors.New(constant.DbReturnedMalformedData)
	}

	return rolePermissions.ToModel(), nil
}

func (repo *RolePermissionRepository) GetRolePermissionByPermissionID(ctx context.Context, roleID, permissionID uuid.UUID) (res domain.RolePermission, err error) {
	var (
		rolePermission RolePermission
	)

	repo.prepareGetRolePermissionByPermissionID()
	err = repo.statement.GetRolePermissionByPermissionID.QueryRowxContext(ctx, roleID, permissionID).StructScan(&rolePermission)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.New(constant.DataNotFound)
		}

		return res, err
	}

	if !rolePermission.Validate() {
		return res, errors.New(constant.DbReturnedMalformedData)
	}

	return rolePermission.ToModel(), nil
}
