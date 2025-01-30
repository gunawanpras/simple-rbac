package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/helper"
	"github.com/gunawanpras/simple-rbac/internal/domain/role/model"
	"github.com/gunawanpras/simple-rbac/util/constant"
)

// Create a new role and return the generated ID
func (repo *RoleRepository) CreateRole(ctx context.Context, role model.Role) (id uuid.UUID, err error) {
	role.ID = helper.UUIDHelper.New()

	repo.prepareCreateRole()
	_, err = repo.statement.CreateRole.ExecContext(ctx, role.ID, role.Name, role.CreatedBy, role.CreatedAt)
	if err != nil {
		return uuid.Nil, err
	}
	return role.ID, nil
}

// Retrieve a list of all roles
func (repo *RoleRepository) ListRoles(ctx context.Context) (roles model.ListRoles, err error) {
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
func (repo *RoleRepository) GetRoleByID(ctx context.Context, id uuid.UUID) (role model.Role, err error) {
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

func (repo *RoleRepository) GetRoleByName(ctx context.Context, name string) (role model.Role, err error) {
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
