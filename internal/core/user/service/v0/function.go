package v0

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/core/user/domain"
	"github.com/gunawanpras/simple-rbac/pkg/util/constant"
	"github.com/gunawanpras/simple-rbac/pkg/util/timeutil"
)

func (service *UserService) CreateUser(ctx context.Context, name, email, password string, roleId uuid.UUID) (res domain.User, err error) {
	user, err := service.repo.UserRepo.GetUserByEmail(ctx, email)
	if err != nil {
		if err.Error() != constant.DataNotFound {
			return res, err
		}
	}

	if user.ID != uuid.Nil {
		return res, errors.New(constant.UserAlreadyExists)
	}

	now := timeutil.TimeHelper.Now()

	id, err := service.repo.UserRepo.CreateUser(ctx, domain.User{
		Name:      name,
		Email:     email,
		Password:  password,
		RoleID:    &roleId,
		CreatedBy: constant.SYSTEM,
		CreatedAt: now,
	})

	if err != nil {
		return res, err
	}

	res = domain.User{
		ID:        id,
		Name:      name,
		Email:     email,
		RoleID:    &roleId,
		CreatedBy: constant.SYSTEM,
		CreatedAt: now,
	}

	return res, nil
}

func (service *UserService) ListUsers(ctx context.Context) (res domain.ListUsers, err error) {
	res, err = service.repo.UserRepo.ListUsers(ctx)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (service *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (res domain.User, err error) {
	res, err = service.repo.UserRepo.GetUserByID(ctx, id)
	if err != nil {
		return res, err
	}

	return res, nil
}
