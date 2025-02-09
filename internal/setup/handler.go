package setup

import (
	permission "github.com/gunawanpras/simple-rbac/internal/adapter/permission/http/handler"
	rolePermission "github.com/gunawanpras/simple-rbac/internal/adapter/role-permission/http/handler"
	role "github.com/gunawanpras/simple-rbac/internal/adapter/role/http/handler"
	user "github.com/gunawanpras/simple-rbac/internal/adapter/user/http/handler"
)

type Handler struct {
	UserHandler           user.Handler
	RoleHandler           role.Handler
	PermissionHandler     permission.Handler
	RolePermissionHandler rolePermission.Handler
}

func NewHandler(service Service) Handler {
	return Handler{
		UserHandler:           user.New(user.InitAttribute{Service: user.ServiceAttribute{UserService: service.UserService}}),
		RoleHandler:           role.New(role.InitAttribute{Service: role.ServiceAttribute{RoleService: service.RoleService}}),
		PermissionHandler:     permission.New(permission.InitAttribute{Service: permission.ServiceAttribute{PermissionService: service.PermissionService}}),
		RolePermissionHandler: rolePermission.New(rolePermission.InitAttribute{Service: rolePermission.ServiceAttribute{RolePermissionService: service.RolePermissionService}}),
	}
}
