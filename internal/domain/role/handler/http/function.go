package httphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gunawanpras/simple-rbac/helper/response"
	"github.com/gunawanpras/simple-rbac/internal/domain/role/model"
	"github.com/gunawanpras/simple-rbac/util/constant"
	"github.com/gunawanpras/simple-rbac/util/validator"
)

func (handler *RoleHandler) CreateRole(c *fiber.Ctx) error {
	var req model.CreateRoleRequest

	ctx := c.UserContext()
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, constant.BindingParameterFailed, err, constant.GenericHttpStatusMappings)
	}

	errv := validator.Validate(req)
	if errv != nil {
		return response.ErrorValidator(c, errv)
	}

	resp, err := handler.service.RoleService.CreateRole(ctx, req.Name)
	if err != nil {
		return response.Error(c, resp.Message, err, constant.RbacHttpStatusMappings)
	}

	return response.OK(c, resp.Message, resp.Data, constant.RbacHttpStatusMappings)
}

func (handler *RoleHandler) ListRoles(c *fiber.Ctx) error {
	ctx := c.UserContext()
	resp, err := handler.service.RoleService.ListRoles(ctx)
	if err != nil {
		return response.Error(c, resp.Message, err, constant.RbacHttpStatusMappings)
	}

	return response.OK(c, resp.Message, resp.Data, constant.RbacHttpStatusMappings)
}

func (handler *RoleHandler) GetRoleByID(c *fiber.Ctx) error {
	var req model.GetRoleByIDRequest

	ctx := c.UserContext()
	if err := c.ParamsParser(&req); err != nil {
		return response.Error(c, constant.BindingParameterFailed, err, constant.GenericHttpStatusMappings)
	}

	errv := validator.Validate(req)
	if errv != nil {
		return response.ErrorValidator(c, errv)
	}

	resp, err := handler.service.RoleService.GetRoleByID(ctx, req.ID)
	if err != nil {
		return response.Error(c, resp.Message, err, constant.RbacHttpStatusMappings)
	}

	return response.OK(c, resp.Message, resp.Data, constant.RbacHttpStatusMappings)
}
