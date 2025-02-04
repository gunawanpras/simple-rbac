package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/core/permission/domain"
	"github.com/gunawanpras/simple-rbac/pkg/util/constant"
	"github.com/gunawanpras/simple-rbac/pkg/util/uuidutil"
)

// Create a new permission and return the generated ID
func (repo *PermissionRepository) CreatePermission(ctx context.Context, permission domain.Permission) (id uuid.UUID, err error) {
	permission.ID = uuidutil.UUIDHelper.New()

	repo.prepareCreatePermission()
	_, err = repo.statement.CreatePermission.ExecContext(ctx, permission.ID, permission.Name, permission.CreatedBy, permission.CreatedAt)
	if err != nil {
		return uuid.Nil, err
	}
	return permission.ID, nil
}

// Retrieve a list of all permissions
func (repo *PermissionRepository) ListPermissions(ctx context.Context) (res domain.ListPermissions, err error) {
	var (
		permission  Permission
		permissions ListPermissions
	)

	repo.prepareListPermissions()
	rows, err := repo.statement.ListPermissions.QueryxContext(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(constant.DataNotFound)
		}

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		permission = Permission{}
		if err = rows.StructScan(&permission); err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}

	if !permissions.Validate() {
		return nil, errors.New(constant.DbReturnedMalformedData)
	}
	return permissions.ToModel(), nil
}

// Retrieve a single permission by ID
func (repo *PermissionRepository) GetPermissionByID(ctx context.Context, id uuid.UUID) (res domain.Permission, err error) {
	var permission Permission

	repo.prepareGetPermissionByID()
	err = repo.statement.GetPermissionByID.QueryRowxContext(ctx, id).StructScan(&permission)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.New(constant.DataNotFound)
		}

		return res, err
	}

	if !permission.Validate() {
		return res, errors.New(constant.DbReturnedMalformedData)
	}

	return permission.ToModel(), nil
}

func (repo *PermissionRepository) GetPermissionByName(ctx context.Context, name string) (res domain.Permission, err error) {
	var permission Permission

	repo.prepareGetPermissionByName()
	err = repo.statement.GetPermissionByName.QueryRowxContext(ctx, name).StructScan(&permission)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.New(constant.DataNotFound)
		}

		return res, err
	}

	if !permission.Validate() {
		return res, errors.New(constant.DbReturnedMalformedData)
	}

	return permission.ToModel(), nil
}
