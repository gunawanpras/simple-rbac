package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/core/role/domain"
	"github.com/gunawanpras/simple-rbac/pkg/util/constant"
	"github.com/gunawanpras/simple-rbac/pkg/util/uuidutil"
)

// Create a new role and return the generated ID
func (repo *RoleRepository) CreateRole(ctx context.Context, role domain.Role) (id uuid.UUID, err error) {
	role.ID = uuidutil.UUIDHelper.New()

	repo.prepareCreateRole()
	_, err = repo.statement.CreateRole.ExecContext(ctx, role.ID, role.Name, role.CreatedBy, role.CreatedAt)
	if err != nil {
		return uuid.Nil, err
	}
	return role.ID, nil
}

// Retrieve a list of all roles
func (repo *RoleRepository) ListRoles(ctx context.Context) (res domain.ListRoles, err error) {
	var (
		role  Role
		roles ListRoles
	)

	repo.prepareListRoles()
	rows, err := repo.statement.ListRoles.QueryxContext(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(constant.DataNotFound)
		}

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		role = Role{}
		if err = rows.StructScan(&role); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	if !roles.Validate() {
		return nil, errors.New(constant.DataNotFound)
	}

	return roles.ToModel(), nil
}

// Retrieve a single role by ID
func (repo *RoleRepository) GetRoleByID(ctx context.Context, id uuid.UUID) (res domain.Role, err error) {
	var role Role

	repo.prepareGetRoleByID()
	err = repo.statement.GetRoleByID.QueryRowxContext(ctx, id).StructScan(&role)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.New(constant.DataNotFound)
		}

		return res, err
	}

	if !role.Validate() {
		return res, errors.New(constant.DbReturnedMalformedData)
	}

	return role.ToModel(), nil
}

func (repo *RoleRepository) GetRoleByName(ctx context.Context, name string) (res domain.Role, err error) {
	var role Role

	repo.prepareGetRoleByName()
	err = repo.statement.GetRoleByName.QueryRowxContext(ctx, name).StructScan(&role)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.New(constant.DataNotFound)
		}

		return res, err
	}

	if !role.Validate() {
		return res, errors.New(constant.DbReturnedMalformedData)
	}

	return role.ToModel(), nil
}
