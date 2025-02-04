package v0

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/core/role/domain"
	"github.com/gunawanpras/simple-rbac/pkg/util/timeutil"

	"github.com/gunawanpras/simple-rbac/pkg/util/constant"
)

func (service *RoleService) CreateRole(ctx context.Context, name string) (res domain.Role, err error) {
	role, err := service.repo.RoleRepo.GetRoleByName(ctx, name)
	if err != nil {
		if err.Error() != constant.DataNotFound {
			return res, err
		}
	}

	if role.ID != uuid.Nil {
		return res, errors.New(constant.RoleAlreadyExists)
	}

	now := timeutil.TimeHelper.Now()

	id, err := service.repo.RoleRepo.CreateRole(ctx, domain.Role{
		Name:      name,
		CreatedBy: constant.SYSTEM,
		CreatedAt: now,
	})

	if err != nil {
		return res, err
	}

	res = domain.Role{
		ID:        id,
		Name:      name,
		CreatedBy: constant.SYSTEM,
		CreatedAt: now,
	}

	return res, nil
}

func (service *RoleService) ListRoles(ctx context.Context) (res domain.ListRoles, err error) {
	roles, err := service.repo.RoleRepo.ListRoles(ctx)
	if err != nil {
		return res, err
	}

	return roles, nil
}

func (service *RoleService) GetRoleByID(ctx context.Context, id uuid.UUID) (res domain.Role, err error) {
	role, err := service.repo.RoleRepo.GetRoleByID(ctx, id)
	if err != nil {
		return res, err
	}

	return role, nil
}
