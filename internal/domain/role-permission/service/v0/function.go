package v0

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/helper"
	"github.com/gunawanpras/simple-rbac/internal/domain/role-permission/model"
	"github.com/gunawanpras/simple-rbac/util/constant"
)

func (service *RolePermissionService) CreateRolePermission(ctx context.Context, roleID, permissionID uuid.UUID) (res model.ServiceResponse, err error) {
	rolePermission, err := service.repo.RolePermissionRepo.GetRolePermissionByPermissionID(ctx, roleID, permissionID)
	if err != nil {
		if err.Error() != constant.DbRecordNotFound {
			return model.ServiceResponse{
				Message: constant.RolePermissionCreateFailed,
			}, err
		}
	}

	if rolePermission.RoleID != uuid.Nil && rolePermission.PermissionID != uuid.Nil {
		return model.ServiceResponse{
			Message: constant.RolePermissionCreateFailed,
		}, errors.New(constant.RolePermissionAlreadyExists)
	}

	roleID, permissionID, err = service.repo.RolePermissionRepo.CreateRolePermission(ctx, model.RolePermission{
		RoleID:       roleID,
		PermissionID: permissionID,
		CreatedBy:    constant.SYSTEM,
		CreatedAt:    helper.TimeHelper.Now(),
	})
	if err != nil {
		return model.ServiceResponse{
			Message: constant.RolePermissionCreateFailed,
		}, err
	}

	res = model.ServiceResponse{
		Message: constant.RolePermissionCreateSuccess,
		Data: model.CreateRolePermissionResponse{
			RoleID:       roleID,
			PermissionID: permissionID,
		},
	}

	return res, nil
}

func (service *RolePermissionService) ListRolePermissions(ctx context.Context) (res model.ServiceResponse, err error) {
	rolePermissions, err := service.repo.RolePermissionRepo.ListRolePermissions(ctx)
	if err != nil {
		return model.ServiceResponse{
			Message: constant.RolePermissionGetFailed,
		}, err
	}

	res = model.ServiceResponse{
		Message: constant.RolePermissionGetSuccess,
		Data:    rolePermissions.ToResponse(),
	}

	return res, nil
}

func (service *RolePermissionService) ListRolePermissionsByRoleID(ctx context.Context, roleID uuid.UUID) (res model.ServiceResponse, err error) {
	rolePermissions, err := service.repo.RolePermissionRepo.ListRolePermissionsByRoleID(ctx, roleID)
	if err != nil {
		return model.ServiceResponse{
			Message: constant.RolePermissionGetFailed,
		}, err
	}

	res = model.ServiceResponse{
		Message: constant.RolePermissionGetSuccess,
		Data:    rolePermissions.ToResponse(),
	}

	return res, nil
}

func (service *RolePermissionService) GetRolePermissionByPermissionID(ctx context.Context, roleID, permissionID uuid.UUID) (res model.ServiceResponse, err error) {
	rolePermission, err := service.repo.RolePermissionRepo.GetRolePermissionByPermissionID(ctx, roleID, permissionID)
	if err != nil {
		return model.ServiceResponse{
			Message: constant.RolePermissionGetFailed,
		}, err
	}

	res = model.ServiceResponse{
		Message: constant.RolePermissionGetSuccess,
		Data:    rolePermission.ToResponse(),
	}

	return res, nil
}
