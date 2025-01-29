package v0

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/helper"
	"github.com/gunawanpras/simple-rbac/internal/domain/rbac/model"
	"github.com/gunawanpras/simple-rbac/util/constant"
)

func (service *RbacService) CreateUser(ctx context.Context, name, email, password string, roleId uuid.UUID) (res model.ServiceResponse, err error) {
	user, err := service.repo.RbacPg.GetUserByEmail(ctx, email)
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

	id, err := service.repo.RbacPg.CreateUser(ctx, model.User{
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
	users, err := service.repo.RbacPg.ListUsers(ctx)
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
	user, err := service.repo.RbacPg.GetUserByID(ctx, id)
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

func (service *RbacService) CreateRole(ctx context.Context, name string) (res model.ServiceResponse, err error) {
	role, err := service.repo.RbacPg.GetRoleByName(ctx, name)
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

	id, err := service.repo.RbacPg.CreateRole(ctx, model.Role{
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

func (service *RbacService) ListRoles(ctx context.Context) (res model.ServiceResponse, err error) {
	roles, err := service.repo.RbacPg.ListRoles(ctx)
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

func (service *RbacService) GetRoleByID(ctx context.Context, id uuid.UUID) (res model.ServiceResponse, err error) {
	role, err := service.repo.RbacPg.GetRoleByID(ctx, id)
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

func (service *RbacService) CreatePermission(ctx context.Context, name, description string) (res model.ServiceResponse, err error) {
	permission, err := service.repo.RbacPg.GetPermissionByName(ctx, name)
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

	id, err := service.repo.RbacPg.CreatePermission(ctx, model.Permission{
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

func (service *RbacService) ListPermissions(ctx context.Context) (res model.ServiceResponse, err error) {
	permissions, err := service.repo.RbacPg.ListPermissions(ctx)
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

func (service *RbacService) GetPermissionByID(ctx context.Context, id uuid.UUID) (res model.ServiceResponse, err error) {
	permission, err := service.repo.RbacPg.GetPermissionByID(ctx, id)
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

func (service *RbacService) CreateRolePermission(ctx context.Context, roleID, permissionID uuid.UUID) (res model.ServiceResponse, err error) {
	rolePermission, err := service.repo.RbacPg.GetRolePermissionByPermissionID(ctx, roleID, permissionID)
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

	roleID, permissionID, err = service.repo.RbacPg.CreateRolePermission(ctx, model.RolePermission{
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

func (service *RbacService) ListRolePermissions(ctx context.Context) (res model.ServiceResponse, err error) {
	rolePermissions, err := service.repo.RbacPg.ListRolePermissions(ctx)
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

func (service *RbacService) ListRolePermissionsByRoleID(ctx context.Context, roleID uuid.UUID) (res model.ServiceResponse, err error) {
	rolePermissions, err := service.repo.RbacPg.ListRolePermissionsByRoleID(ctx, roleID)
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

func (service *RbacService) GetRolePermissionByPermissionID(ctx context.Context, roleID, permissionID uuid.UUID) (res model.ServiceResponse, err error) {
	rolePermission, err := service.repo.RbacPg.GetRolePermissionByPermissionID(ctx, roleID, permissionID)
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
