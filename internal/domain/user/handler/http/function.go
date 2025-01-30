package httphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gunawanpras/simple-rbac/helper/response"
	"github.com/gunawanpras/simple-rbac/internal/domain/user/model"
	"github.com/gunawanpras/simple-rbac/util/constant"
	"github.com/gunawanpras/simple-rbac/util/validator"
)

func (handler *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req model.CreateUserRequest

	ctx := c.UserContext()
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, constant.BindingParameterFailed, err, constant.GenericHttpStatusMappings)
	}

	errv := validator.Validate(req)
	if errv != nil {
		return response.ErrorValidator(c, errv)
	}

	resp, err := handler.service.UserService.CreateUser(ctx, req.Name, req.Email, req.Password, req.RoleID)
	if err != nil {
		return response.Error(c, resp.Message, err, constant.RbacHttpStatusMappings)
	}

	return response.OK(c, resp.Message, resp.Data, constant.RbacHttpStatusMappings)
}

func (handler *UserHandler) ListUsers(c *fiber.Ctx) error {
	ctx := c.UserContext()
	resp, err := handler.service.UserService.ListUsers(ctx)
	if err != nil {
		return response.Error(c, resp.Message, err, constant.RbacHttpStatusMappings)
	}

	return response.OK(c, resp.Message, resp.Data, constant.RbacHttpStatusMappings)
}

func (handler *UserHandler) GetUserByID(c *fiber.Ctx) error {
	var req model.GetUserByIDRequest

	ctx := c.UserContext()
	if err := c.ParamsParser(&req); err != nil {
		return response.Error(c, constant.BindingParameterFailed, err, constant.GenericHttpStatusMappings)
	}

	errv := validator.Validate(req)
	if errv != nil {
		return response.ErrorValidator(c, errv)
	}

	resp, err := handler.service.UserService.GetUserByID(ctx, req.ID)
	if err != nil {
		return response.Error(c, resp.Message, err, constant.RbacHttpStatusMappings)
	}

	return response.OK(c, resp.Message, resp.Data, constant.RbacHttpStatusMappings)
}
