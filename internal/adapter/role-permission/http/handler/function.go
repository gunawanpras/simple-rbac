package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/adapter/role-permission/http/dto"
	"github.com/gunawanpras/simple-rbac/pkg/response"
	"github.com/gunawanpras/simple-rbac/pkg/util/constant"
	"github.com/gunawanpras/simple-rbac/pkg/validator"
)

func (handler *RolePermissionHandler) CreateRolePermission(c *fiber.Ctx) error {
	var (
		req dto.CreateRolePermissionRequest
		res dto.CreateRolePermissionResponse
	)

	ctx := c.UserContext()
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, constant.BindingParameterFailed, err, constant.GenericHttpStatusMappings)
	}

	errv := validator.Validate(req)
	if errv != nil {
		return response.ErrorValidator(c, errv)
	}

	resp, err := handler.service.RolePermissionService.CreateRolePermission(ctx, req.RoleID, req.PermissionID)
	if err != nil {
		return response.Error(c, constant.RolePermissionCreateFailed, err, constant.RbacHttpStatusMappings)
	}

	res.ToResponse(resp)

	return response.OK(c, constant.RolePermissionCreateSuccess, res, constant.RbacHttpStatusMappings)
}

func (handler *RolePermissionHandler) ListRolePermissions(c *fiber.Ctx) error {
	var res dto.ListRolePermissionsResponse

	ctx := c.UserContext()
	resp, err := handler.service.RolePermissionService.ListRolePermissions(ctx)
	if err != nil {
		return response.Error(c, constant.RolePermissionCreateFailed, err, constant.RbacHttpStatusMappings)
	}

	res.ToResponse(resp)

	return response.OK(c, constant.RolePermissionCreateSuccess, res, constant.RbacHttpStatusMappings)
}

func (handler *RolePermissionHandler) ListRolePermissionsByRoleID(c *fiber.Ctx) error {
	var (
		req dto.ListRolePermissionsByRoleIDRequest
		res dto.ListRolePermissionsResponse
	)

	ctx := c.UserContext()
	if err := c.ParamsParser(&req); err != nil {
		return response.Error(c, constant.BindingParameterFailed, err, constant.GenericHttpStatusMappings)
	}

	roleID := c.Params("role_id")
	req.RoleID, _ = uuid.Parse(roleID)

	errv := validator.Validate(req)
	if errv != nil {
		return response.ErrorValidator(c, errv)
	}

	resp, err := handler.service.RolePermissionService.ListRolePermissionsByRoleID(ctx, req.RoleID)
	if err != nil {
		return response.Error(c, constant.RolePermissionGetFailed, err, constant.RbacHttpStatusMappings)
	}

	res.ToResponse(resp)

	return response.OK(c, constant.RolePermissionGetSuccess, res, constant.RbacHttpStatusMappings)
}

func (handler *RolePermissionHandler) GetRolePermissionByPermissionID(c *fiber.Ctx) error {
	var (
		req dto.GetRolePermissionByPermissionIDRequest
		res dto.GetRolePermissionResponse
	)

	ctx := c.UserContext()
	if err := c.ParamsParser(&req); err != nil {
		return response.Error(c, constant.BindingParameterFailed, err, constant.GenericHttpStatusMappings)
	}

	roleID := c.Params("role_id")
	req.RoleID, _ = uuid.Parse(roleID)

	permissionID := c.Params("permission_id")
	req.PermissionID, _ = uuid.Parse(permissionID)

	errv := validator.Validate(req)
	if errv != nil {
		return response.ErrorValidator(c, errv)
	}

	resp, err := handler.service.RolePermissionService.GetRolePermissionByPermissionID(ctx, req.RoleID, req.PermissionID)
	if err != nil {
		return response.Error(c, constant.RolePermissionGetFailed, err, constant.RbacHttpStatusMappings)
	}

	res.ToResponse(resp)

	return response.OK(c, constant.RolePermissionGetSuccess, res, constant.RbacHttpStatusMappings)
}
