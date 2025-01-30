package v0

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/helper"
	"github.com/gunawanpras/simple-rbac/internal/domain/user/model"
	"github.com/gunawanpras/simple-rbac/util/constant"
)

func (service *RbacService) CreateUser(ctx context.Context, name, email, password string, roleId uuid.UUID) (res model.ServiceResponse, err error) {
	user, err := service.repo.UserRepo.GetUserByEmail(ctx, email)
	if err != nil {
		if err.Error() != constant.DbRecordNotFound {
			return model.ServiceResponse{
				Message: constant.UserCreateFailed,
			}, err
		}
	}

	if user.ID != uuid.Nil {
		return model.ServiceResponse{
			Message: constant.UserCreateFailed,
		}, errors.New(constant.UserAlreadyExists)
	}

	id, err := service.repo.UserRepo.CreateUser(ctx, model.User{
		Name:      name,
		Email:     email,
		Password:  password,
		RoleID:    &roleId,
		CreatedBy: constant.SYSTEM,
		CreatedAt: helper.TimeHelper.Now(),
	})

	if err != nil {
		return model.ServiceResponse{
			Message: constant.UserCreateFailed,
		}, err
	}

	res = model.ServiceResponse{
		Message: constant.UserCreateSuccess,
		Data: model.CreateUserResponse{
			ID: id,
		},
	}

	return res, nil
}

func (service *RbacService) ListUsers(ctx context.Context) (res model.ServiceResponse, err error) {
	users, err := service.repo.UserRepo.ListUsers(ctx)
	if err != nil {
		return model.ServiceResponse{
			Message: constant.UserGetFailed,
		}, err
	}

	res = model.ServiceResponse{
		Message: constant.UserGetSuccess,
		Data:    users.ToResponse(),
	}

	return res, nil
}

func (service *RbacService) GetUserByID(ctx context.Context, id uuid.UUID) (res model.ServiceResponse, err error) {
	user, err := service.repo.UserRepo.GetUserByID(ctx, id)
	if err != nil {
		return model.ServiceResponse{
			Message: constant.UserGetFailed,
		}, err
	}

	res = model.ServiceResponse{
		Message: constant.UserGetSuccess,
		Data:    user.ToResponse(),
	}

	return res, nil
}
