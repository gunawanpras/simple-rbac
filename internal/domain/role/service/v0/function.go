package v0

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/helper"
	"github.com/gunawanpras/simple-rbac/internal/domain/role/model"
	"github.com/gunawanpras/simple-rbac/util/constant"
)

func (service *RoleService) CreateRole(ctx context.Context, name string) (res model.ServiceResponse, err error) {
	role, err := service.repo.RoleRepo.GetRoleByName(ctx, name)
	if err != nil {
		if err.Error() != constant.DbRecordNotFound {
			return model.ServiceResponse{
				Message: constant.RoleCreateFailed,
			}, err
		}
	}

	if role.ID != uuid.Nil {
		return model.ServiceResponse{
			Message: constant.RoleCreateFailed,
		}, errors.New(constant.RoleAlreadyExists)
	}

	id, err := service.repo.RoleRepo.CreateRole(ctx, model.Role{
		Name:      name,
		CreatedBy: constant.SYSTEM,
		CreatedAt: helper.TimeHelper.Now(),
	})

	if err != nil {
		return model.ServiceResponse{
			Message: constant.RoleCreateFailed,
		}, err
	}

	res = model.ServiceResponse{
		Message: constant.RoleCreateSuccess,
		Data: model.CreateRoleResponse{
			ID: id,
		},
	}

	return res, nil
}

func (service *RoleService) ListRoles(ctx context.Context) (res model.ServiceResponse, err error) {
	roles, err := service.repo.RoleRepo.ListRoles(ctx)
	if err != nil {
		return model.ServiceResponse{
			Message: constant.RoleGetFailed,
		}, err
	}

	res = model.ServiceResponse{
		Message: constant.RoleGetSuccess,
		Data:    roles.ToResponse(),
	}

	return res, nil
}

func (service *RoleService) GetRoleByID(ctx context.Context, id uuid.UUID) (res model.ServiceResponse, err error) {
	role, err := service.repo.RoleRepo.GetRoleByID(ctx, id)
	if err != nil {
		return model.ServiceResponse{
			Message: constant.RoleGetFailed,
		}, err
	}

	res = model.ServiceResponse{
		Message: constant.RoleGetSuccess,
		Data:    role.ToResponse(),
	}

	return res, nil
}
