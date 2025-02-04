package server

import (
	permission "github.com/gunawanpras/simple-rbac/internal/adapter/permission/http/handler"
	rolePermission "github.com/gunawanpras/simple-rbac/internal/adapter/role-permission/http/handler"
	role "github.com/gunawanpras/simple-rbac/internal/adapter/role/http/handler"
	user "github.com/gunawanpras/simple-rbac/internal/adapter/user/http/handler"
)

type Handler struct {
	userHandler           user.Handler
	roleHandler           role.Handler
	permissionHandler     permission.Handler
	rolePermissionHandler rolePermission.Handler
}

func NewHandler(service Service) Handler {
	return Handler{
		userHandler:           user.New(user.InitAttribute{Service: user.ServiceAttribute{UserService: service.UserService}}),
		roleHandler:           role.New(role.InitAttribute{Service: role.ServiceAttribute{RoleService: service.RoleService}}),
		permissionHandler:     permission.New(permission.InitAttribute{Service: permission.ServiceAttribute{PermissionService: service.PermissionService}}),
		rolePermissionHandler: rolePermission.New(rolePermission.InitAttribute{Service: rolePermission.ServiceAttribute{RolePermissionService: service.RolePermissionService}}),
	}
}
