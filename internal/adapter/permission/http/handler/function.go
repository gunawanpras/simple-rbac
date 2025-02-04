package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gunawanpras/simple-rbac/internal/adapter/permission/http/dto"
	"github.com/gunawanpras/simple-rbac/pkg/response"
	"github.com/gunawanpras/simple-rbac/pkg/util/constant"
	"github.com/gunawanpras/simple-rbac/pkg/validator"
)

func (handler *PermissionHandler) CreatePermission(c *fiber.Ctx) error {
	var req dto.CreatePermissionRequest

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
		return response.Error(c, constant.PermissionCreateFailed, err, constant.RbacHttpStatusMappings)
	}

	respData := dto.CreatePermissionResponse{
		ID: resp.ID,
	}

	return response.OK(c, constant.PermissionCreateSuccess, respData, constant.RbacHttpStatusMappings)
}

func (handler *PermissionHandler) ListPermissions(c *fiber.Ctx) error {
	var res dto.ListPermissionsResponse

	ctx := c.UserContext()
	resp, err := handler.service.PermissionService.ListPermissions(ctx)
	if err != nil {
		return response.Error(c, constant.PermissionGetFailed, err, constant.RbacHttpStatusMappings)
	}

	res.ToResponse(resp)

	return response.OK(c, constant.PermissionGetSuccess, res, constant.RbacHttpStatusMappings)
}

func (handler *PermissionHandler) GetPermissionByID(c *fiber.Ctx) error {
	var (
		req dto.GetPermissionByIDRequest
		res dto.GetPermissionResponse
	)

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
		return response.Error(c, constant.PermissionGetFailed, err, constant.RbacHttpStatusMappings)
	}

	res.ToResponse(resp)

	return response.OK(c, constant.PermissionGetSuccess, res, constant.RbacHttpStatusMappings)
}
