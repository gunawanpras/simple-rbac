package v0

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/core/permission/domain"
	"github.com/gunawanpras/simple-rbac/pkg/util/constant"
	"github.com/gunawanpras/simple-rbac/pkg/util/timeutil"
)

func (service *PermissionService) CreatePermission(ctx context.Context, name, description string) (res domain.Permission, err error) {
	permission, err := service.repo.PermissionRepo.GetPermissionByName(ctx, name)
	if err != nil {
		if err.Error() != constant.DataNotFound {
			return res, err
		}
	}

	if permission.ID != uuid.Nil {
		return res, errors.New(constant.PermissionAlreadyExists)
	}

	now := timeutil.TimeHelper.Now()

	id, err := service.repo.PermissionRepo.CreatePermission(ctx, domain.Permission{
		Name:        name,
		Description: &description,
		CreatedBy:   constant.SYSTEM,
		CreatedAt:   now,
	})

	if err != nil {
		return res, err
	}

	res = domain.Permission{
		ID:          id,
		Name:        name,
		Description: &description,
		CreatedBy:   constant.SYSTEM,
		CreatedAt:   now,
	}

	return res, nil
}

func (service *PermissionService) ListPermissions(ctx context.Context) (res domain.ListPermissions, err error) {
	permissions, err := service.repo.PermissionRepo.ListPermissions(ctx)
	if err != nil {
		return res, err
	}

	return permissions, nil
}

func (service *PermissionService) GetPermissionByID(ctx context.Context, id uuid.UUID) (res domain.Permission, err error) {
	permission, err := service.repo.PermissionRepo.GetPermissionByID(ctx, id)
	if err != nil {
		return res, err
	}

	return permission, nil
}
