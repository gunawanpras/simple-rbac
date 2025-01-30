package httphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gunawanpras/simple-rbac/helper/response"
	"github.com/gunawanpras/simple-rbac/internal/domain/permission/model"
	"github.com/gunawanpras/simple-rbac/util/constant"
	"github.com/gunawanpras/simple-rbac/util/validator"
)

func (handler *PermissionHandler) CreatePermission(c *fiber.Ctx) error {
	var req model.CreatePermissionRequest

	ctx := c.UserContext()
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, constant.BindingParameterFailed, err, constant.GenericHttpStatusMappings)
	}

	errv := validator.Validate(req)
	if errv != nil {
		return response.ErrorValidator(c, errv)
	}

	resp, err := handler.service.PermissionService.CreatePermission(ctx, req.Name, req.Description)
	if err != nil {
		return response.Error(c, resp.Message, err, constant.RbacHttpStatusMappings)
	}

	return response.OK(c, resp.Message, resp.Data, constant.RbacHttpStatusMappings)
}

func (handler *PermissionHandler) ListPermissions(c *fiber.Ctx) error {
	ctx := c.UserContext()
	resp, err := handler.service.PermissionService.ListPermissions(ctx)
	if err != nil {
		return response.Error(c, resp.Message, err, constant.RbacHttpStatusMappings)
	}

	return response.OK(c, resp.Message, resp.Data, constant.RbacHttpStatusMappings)
}

func (handler *PermissionHandler) GetPermissionByID(c *fiber.Ctx) error {
	var req model.GetPermissionByIDRequest

	ctx := c.UserContext()
	if err := c.ParamsParser(&req); err != nil {
		return response.Error(c, constant.BindingParameterFailed, err, constant.GenericHttpStatusMappings)
	}

	errv := validator.Validate(req)
	if errv != nil {
		return response.ErrorValidator(c, errv)
	}

	resp, err := handler.service.PermissionService.GetPermissionByID(ctx, req.ID)
	if err != nil {
		return response.Error(c, resp.Message, err, constant.RbacHttpStatusMappings)
	}

	return response.OK(c, resp.Message, resp.Data, constant.RbacHttpStatusMappings)
}
