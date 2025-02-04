package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gunawanpras/simple-rbac/internal/adapter/role/http/dto"
	"github.com/gunawanpras/simple-rbac/pkg/response"
	"github.com/gunawanpras/simple-rbac/pkg/util/constant"
	"github.com/gunawanpras/simple-rbac/pkg/validator"
)

func (handler *RoleHandler) CreateRole(c *fiber.Ctx) error {
	var (
		req dto.CreateRoleRequest
		res dto.CreateRoleResponse
	)

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
		return response.Error(c, constant.RoleCreateFailed, err, constant.RbacHttpStatusMappings)
	}

	res.ToResponse(resp)

	return response.OK(c, constant.RoleCreateSuccess, res, constant.RbacHttpStatusMappings)
}

func (handler *RoleHandler) ListRoles(c *fiber.Ctx) error {
	var res dto.ListRolesResponse

	ctx := c.UserContext()
	resp, err := handler.service.RoleService.ListRoles(ctx)
	if err != nil {
		return response.Error(c, constant.RoleGetFailed, err, constant.RbacHttpStatusMappings)
	}

	res.ToResponse(resp)

	return response.OK(c, constant.RoleGetSuccess, res, constant.RbacHttpStatusMappings)
}

func (handler *RoleHandler) GetRoleByID(c *fiber.Ctx) error {
	var (
		req dto.GetRoleByIDRequest
		res dto.GetRoleResponse
	)

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
		return response.Error(c, constant.RoleGetFailed, err, constant.RbacHttpStatusMappings)
	}

	res.ToResponse(resp)

	return response.OK(c, constant.RoleGetSuccess, res, constant.RbacHttpStatusMappings)
}
