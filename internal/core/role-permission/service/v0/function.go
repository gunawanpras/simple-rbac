package v0

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/core/role-permission/domain"
	"github.com/gunawanpras/simple-rbac/pkg/util/constant"
	"github.com/gunawanpras/simple-rbac/pkg/util/timeutil"
)

func (service *RolePermissionService) CreateRolePermission(ctx context.Context, roleID, permissionID uuid.UUID) (res domain.RolePermission, err error) {
	rolePermission, err := service.repo.RolePermissionRepo.GetRolePermissionByPermissionID(ctx, roleID, permissionID)
	if err != nil {
		if err.Error() != constant.DataNotFound {
			return res, err
		}
	}

	if rolePermission.RoleID != uuid.Nil && rolePermission.PermissionID != uuid.Nil {
		return res, errors.New(constant.RolePermissionAlreadyExists)
	}

	now := timeutil.TimeHelper.Now()

	roleID, permissionID, err = service.repo.RolePermissionRepo.CreateRolePermission(ctx, domain.RolePermission{
		RoleID:       roleID,
		PermissionID: permissionID,
		CreatedBy:    constant.SYSTEM,
		CreatedAt:    now,
	})
	if err != nil {
		return res, err
	}

	res = domain.RolePermission{
		RoleID:       roleID,
		PermissionID: permissionID,
		CreatedBy:    constant.SYSTEM,
		CreatedAt:    now,
	}

	return res, nil
}

func (service *RolePermissionService) ListRolePermissions(ctx context.Context) (res domain.ListRolePermissions, err error) {
	res, err = service.repo.RolePermissionRepo.ListRolePermissions(ctx)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (service *RolePermissionService) ListRolePermissionsByRoleID(ctx context.Context, roleID uuid.UUID) (res domain.ListRolePermissions, err error) {
	res, err = service.repo.RolePermissionRepo.ListRolePermissionsByRoleID(ctx, roleID)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (service *RolePermissionService) GetRolePermissionByPermissionID(ctx context.Context, roleID, permissionID uuid.UUID) (res domain.RolePermission, err error) {
	res, err = service.repo.RolePermissionRepo.GetRolePermissionByPermissionID(ctx, roleID, permissionID)
	if err != nil {
		return res, err
	}

	return res, nil
}
