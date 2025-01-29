package httphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/helper/response"
	"github.com/gunawanpras/simple-rbac/internal/domain/rbac/model"
	"github.com/gunawanpras/simple-rbac/util/constant"
	"github.com/gunawanpras/simple-rbac/util/validator"
)

func (handler *RbacHandler) CreateUser(c *fiber.Ctx) error {
	var req model.CreateUserRequest

	ctx := c.UserContext()
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, constant.BindingParameterFailed, err, constant.GenericHttpStatusMappings)
	}

	errv := validator.Validate(req)
	if errv != nil {
		return response.ErrorValidator(c, errv)
	}

	resp, err := handler.service.RbacService.CreateUser(ctx, req.Name, req.Email, req.Password, req.RoleID)
	if err != nil {
		return response.Error(c, resp.Message, err, constant.RbacHttpStatusMappings)
	}

	return response.OK(c, resp.Message, resp.Data, constant.RbacHttpStatusMappings)
}

func (handler *RbacHandler) ListUsers(c *fiber.Ctx) error {
	ctx := c.UserContext()
	resp, err := handler.service.RbacService.ListUsers(ctx)
	if err != nil {
		return response.Error(c, resp.Message, err, constant.RbacHttpStatusMappings)
	}

	return response.OK(c, resp.Message, resp.Data, constant.RbacHttpStatusMappings)
}

func (handler *RbacHandler) GetUserByID(c *fiber.Ctx) error {
	var req model.GetUserByIDRequest

	ctx := c.UserContext()
	if err := c.ParamsParser(&req); err != nil {
		return response.Error(c, constant.BindingParameterFailed, err, constant.GenericHttpStatusMappings)
	}

	errv := validator.Validate(req)
	if errv != nil {
		return response.ErrorValidator(c, errv)
	}

	resp, err := handler.service.RbacService.GetUserByID(ctx, req.ID)
	if err != nil {
		return response.Error(c, resp.Message, err, constant.RbacHttpStatusMappings)
	}

	return response.OK(c, resp.Message, resp.Data, constant.RbacHttpStatusMappings)
}

func (handler *RbacHandler) CreateRole(c *fiber.Ctx) error {
	var req model.CreateRoleRequest

	ctx := c.UserContext()
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, constant.BindingParameterFailed, err, constant.GenericHttpStatusMappings)
	}

	errv := validator.Validate(req)
	if errv != nil {
		return response.ErrorValidator(c, errv)
	}

	resp, err := handler.service.RbacService.CreateRole(ctx, req.Name)
	if err != nil {
		return response.Error(c, resp.Message, err, constant.RbacHttpStatusMappings)
	}

	return response.OK(c, resp.Message, resp.Data, constant.RbacHttpStatusMappings)
}

func (handler *RbacHandler) ListRoles(c *fiber.Ctx) error {
	ctx := c.UserContext()
	resp, err := handler.service.RbacService.ListRoles(ctx)
	if err != nil {
		return response.Error(c, resp.Message, err, constant.RbacHttpStatusMappings)
	}

	return response.OK(c, resp.Message, resp.Data, constant.RbacHttpStatusMappings)
}

func (handler *RbacHandler) GetRoleByID(c *fiber.Ctx) error {
	var req model.GetRoleByIDRequest

	ctx := c.UserContext()
	if err := c.ParamsParser(&req); err != nil {
		return response.Error(c, constant.BindingParameterFailed, err, constant.GenericHttpStatusMappings)
	}

	errv := validator.Validate(req)
	if errv != nil {
		return response.ErrorValidator(c, errv)
	}

	resp, err := handler.service.RbacService.GetRoleByID(ctx, req.ID)
	if err != nil {
		return response.Error(c, resp.Message, err, constant.RbacHttpStatusMappings)
	}

	return response.OK(c, resp.Message, resp.Data, constant.RbacHttpStatusMappings)
}

func (handler *RbacHandler) CreatePermission(c *fiber.Ctx) error {
	var req model.CreatePermissionRequest

	ctx := c.UserContext()
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, constant.BindingParameterFailed, err, constant.GenericHttpStatusMappings)
	}

	errv := validator.Validate(req)
	if errv != nil {
		return response.ErrorValidator(c, errv)
	}

	resp, err := handler.service.RbacService.CreatePermission(ctx, req.Name, req.Description)
	if err != nil {
		return response.Error(c, resp.Message, err, constant.RbacHttpStatusMappings)
	}

	return response.OK(c, resp.Message, resp.Data, constant.RbacHttpStatusMappings)
}

func (handler *RbacHandler) ListPermissions(c *fiber.Ctx) error {
	ctx := c.UserContext()
	resp, err := handler.service.RbacService.ListPermissions(ctx)
	if err != nil {
		return response.Error(c, resp.Message, err, constant.RbacHttpStatusMappings)
	}

	return response.OK(c, resp.Message, resp.Data, constant.RbacHttpStatusMappings)
}

func (handler *RbacHandler) GetPermissionByID(c *fiber.Ctx) error {
	var req model.GetPermissionByIDRequest

	ctx := c.UserContext()
	if err := c.ParamsParser(&req); err != nil {
		return response.Error(c, constant.BindingParameterFailed, err, constant.GenericHttpStatusMappings)
	}

	errv := validator.Validate(req)
	if errv != nil {
		return response.ErrorValidator(c, errv)
	}

	resp, err := handler.service.RbacService.GetPermissionByID(ctx, req.ID)
	if err != nil {
		return response.Error(c, resp.Message, err, constant.RbacHttpStatusMappings)
	}

	return response.OK(c, resp.Message, resp.Data, constant.RbacHttpStatusMappings)
}

func (handler *RbacHandler) CreateRolePermission(c *fiber.Ctx) error {
	var req model.CreateRolePermissionRequest

	ctx := c.UserContext()
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, constant.BindingParameterFailed, err, constant.GenericHttpStatusMappings)
	}

	errv := validator.Validate(req)
	if errv != nil {
		return response.ErrorValidator(c, errv)
	}

	resp, err := handler.service.RbacService.CreateRolePermission(ctx, req.RoleID, req.PermissionID)
	if err != nil {
		return response.Error(c, resp.Message, err, constant.RbacHttpStatusMappings)
	}

	return response.OK(c, resp.Message, resp.Data, constant.RbacHttpStatusMappings)
}

func (handler *RbacHandler) ListRolePermissions(c *fiber.Ctx) error {
	ctx := c.UserContext()
	resp, err := handler.service.RbacService.ListRolePermissions(ctx)
	if err != nil {
		return response.Error(c, resp.Message, err, constant.RbacHttpStatusMappings)
	}

	return response.OK(c, resp.Message, resp.Data, constant.RbacHttpStatusMappings)
}

func (handler *RbacHandler) ListRolePermissionsByRoleID(c *fiber.Ctx) error {
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

	resp, err := handler.service.RbacService.ListRolePermissionsByRoleID(ctx, req.RoleID)
	if err != nil {
		return response.Error(c, resp.Message, err, constant.RbacHttpStatusMappings)
	}

	return response.OK(c, resp.Message, resp.Data, constant.RbacHttpStatusMappings)
}

func (handler *RbacHandler) GetRolePermissionByPermissionID(c *fiber.Ctx) error {
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

	resp, err := handler.service.RbacService.GetRolePermissionByPermissionID(ctx, req.RoleID, req.PermissionID)
	if err != nil {
		return response.Error(c, resp.Message, err, constant.RbacHttpStatusMappings)
	}

	return response.OK(c, resp.Message, resp.Data, constant.RbacHttpStatusMappings)
}
