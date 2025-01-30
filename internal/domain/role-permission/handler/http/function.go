package httphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/helper/response"
	"github.com/gunawanpras/simple-rbac/internal/domain/role-permission/model"
	"github.com/gunawanpras/simple-rbac/util/constant"
	"github.com/gunawanpras/simple-rbac/util/validator"
)

func (handler *RolePermissionHandler) CreateRolePermission(c *fiber.Ctx) error {
	var req model.CreateRolePermissionRequest

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
		return response.Error(c, resp.Message, err, constant.RbacHttpStatusMappings)
	}

	return response.OK(c, resp.Message, resp.Data, constant.RbacHttpStatusMappings)
}

func (handler *RolePermissionHandler) ListRolePermissions(c *fiber.Ctx) error {
	ctx := c.UserContext()
	resp, err := handler.service.RolePermissionService.ListRolePermissions(ctx)
	if err != nil {
		return response.Error(c, resp.Message, err, constant.RbacHttpStatusMappings)
	}

	return response.OK(c, resp.Message, resp.Data, constant.RbacHttpStatusMappings)
}

func (handler *RolePermissionHandler) ListRolePermissionsByRoleID(c *fiber.Ctx) error {
	var req model.ListRolePermissionsByRoleIDRequest

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
		return response.Error(c, resp.Message, err, constant.RbacHttpStatusMappings)
	}

	return response.OK(c, resp.Message, resp.Data, constant.RbacHttpStatusMappings)
}

func (handler *RolePermissionHandler) GetRolePermissionByPermissionID(c *fiber.Ctx) error {
	var req model.GetRolePermissionByPermissionIDRequest

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
		return response.Error(c, resp.Message, err, constant.RbacHttpStatusMappings)
	}

	return response.OK(c, resp.Message, resp.Data, constant.RbacHttpStatusMappings)
}
