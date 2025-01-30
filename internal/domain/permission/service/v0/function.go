package v0

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/helper"
	"github.com/gunawanpras/simple-rbac/internal/domain/permission/model"
	"github.com/gunawanpras/simple-rbac/util/constant"
)

func (service *PermissionService) CreatePermission(ctx context.Context, name, description string) (res model.ServiceResponse, err error) {
	permission, err := service.repo.PermissionRepo.GetPermissionByName(ctx, name)
	if err != nil {
		if err.Error() != constant.DbRecordNotFound {
			return model.ServiceResponse{
				Message: constant.PermissionCreateFailed,
			}, err
		}
	}

	if permission.ID != uuid.Nil {
		return model.ServiceResponse{
			Message: constant.PermissionCreateFailed,
		}, errors.New(constant.PermissionAlreadyExists)
	}

	id, err := service.repo.PermissionRepo.CreatePermission(ctx, model.Permission{
		Name:        name,
		Description: &description,
		CreatedBy:   constant.SYSTEM,
		CreatedAt:   helper.TimeHelper.Now(),
	})

	if err != nil {
		return model.ServiceResponse{
			Message: constant.PermissionCreateFailed,
		}, err
	}

	res = model.ServiceResponse{
		Message: constant.PermissionCreateSuccess,
		Data: model.CreatePermissionResponse{
			ID: id,
		},
	}

	return res, nil
}

func (service *PermissionService) ListPermissions(ctx context.Context) (res model.ServiceResponse, err error) {
	permissions, err := service.repo.PermissionRepo.ListPermissions(ctx)
	if err != nil {
		return model.ServiceResponse{
			Message: constant.PermissionGetFailed,
		}, err
	}

	res = model.ServiceResponse{
		Message: constant.PermissionGetSuccess,
		Data:    permissions.ToResponse(),
	}

	return res, nil
}

func (service *PermissionService) GetPermissionByID(ctx context.Context, id uuid.UUID) (res model.ServiceResponse, err error) {
	permission, err := service.repo.PermissionRepo.GetPermissionByID(ctx, id)
	if err != nil {
		return model.ServiceResponse{
			Message: constant.PermissionGetFailed,
		}, err
	}

	res = model.ServiceResponse{
		Message: constant.PermissionGetSuccess,
		Data:    permission.ToResponse(),
	}

	return res, nil
}
