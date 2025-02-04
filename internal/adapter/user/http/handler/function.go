package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gunawanpras/simple-rbac/internal/adapter/user/http/dto"
	"github.com/gunawanpras/simple-rbac/pkg/response"
	"github.com/gunawanpras/simple-rbac/pkg/util/constant"
	"github.com/gunawanpras/simple-rbac/pkg/validator"
)

func (handler *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req dto.CreateUserRequest

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
		return response.Error(c, constant.UserCreateFailed, err, constant.RbacHttpStatusMappings)
	}

	respData := dto.CreateUserResponse{
		ID: resp.ID,
	}

	return response.OK(c, constant.UserCreateSuccess, respData, constant.RbacHttpStatusMappings)
}

func (handler *UserHandler) ListUsers(c *fiber.Ctx) error {
	var res dto.ListUsersResponse

	ctx := c.UserContext()
	resp, err := handler.service.UserService.ListUsers(ctx)
	if err != nil {
		return response.Error(c, constant.UserGetFailed, err, constant.RbacHttpStatusMappings)
	}

	res.ToResponse(resp)

	return response.OK(c, constant.UserGetSuccess, res, constant.RbacHttpStatusMappings)
}

func (handler *UserHandler) GetUserByID(c *fiber.Ctx) error {
	var (
		req dto.GetUserByIDRequest
		res dto.GetUserResponse
	)

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
		return response.Error(c, constant.UserGetFailed, err, constant.RbacHttpStatusMappings)
	}

	res.ToResponse(resp)

	return response.OK(c, constant.UserGetSuccess, res, constant.RbacHttpStatusMappings)
}
