package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/helper"
	"github.com/gunawanpras/simple-rbac/internal/domain/permission/model"
	"github.com/gunawanpras/simple-rbac/util/constant"
)

// Create a new permission and return the generated ID
func (repo *PermissionRepository) CreatePermission(ctx context.Context, permission model.Permission) (id uuid.UUID, err error) {
	permission.ID = helper.UUIDHelper.New()

	repo.prepareCreatePermission()
	_, err = repo.statement.CreatePermission.ExecContext(ctx, permission.ID, permission.Name, permission.CreatedBy, permission.CreatedAt)
	if err != nil {
		return uuid.Nil, err
	}
	return permission.ID, nil
}

// Retrieve a list of all permissions
func (repo *PermissionRepository) ListPermissions(ctx context.Context) (permissions model.ListPermissions, err error) {
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
func (repo *PermissionRepository) GetPermissionByID(ctx context.Context, id uuid.UUID) (permission model.Permission, err error) {
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

func (repo *PermissionRepository) GetPermissionByName(ctx context.Context, name string) (permission model.Permission, err error) {
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
