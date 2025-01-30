package server

import (
	permission "github.com/gunawanpras/simple-rbac/internal/domain/permission/handler"
	permissionHttp "github.com/gunawanpras/simple-rbac/internal/domain/permission/handler/http"
	rolePermission "github.com/gunawanpras/simple-rbac/internal/domain/role-permission/handler"
	rolePermissionHttp "github.com/gunawanpras/simple-rbac/internal/domain/role-permission/handler/http"
	role "github.com/gunawanpras/simple-rbac/internal/domain/role/handler"
	roleHttp "github.com/gunawanpras/simple-rbac/internal/domain/role/handler/http"
	user "github.com/gunawanpras/simple-rbac/internal/domain/user/handler"
	userHttp "github.com/gunawanpras/simple-rbac/internal/domain/user/handler/http"
)

type Handler struct {
	userHandler           user.Handler
	roleHandler           role.Handler
	permissionHandler     permission.Handler
	rolePermissionHandler rolePermission.Handler
}

func NewHandler(service Service) Handler {
	return Handler{
		userHandler:           userHttp.New(userHttp.InitAttribute{Service: userHttp.ServiceAttribute{UserService: service.UserService}}),
		roleHandler:           roleHttp.New(roleHttp.InitAttribute{Service: roleHttp.ServiceAttribute{RoleService: service.RoleService}}),
		permissionHandler:     permissionHttp.New(permissionHttp.InitAttribute{Service: permissionHttp.ServiceAttribute{PermissionService: service.PermissionService}}),
		rolePermissionHandler: rolePermissionHttp.New(rolePermissionHttp.InitAttribute{Service: rolePermissionHttp.ServiceAttribute{RolePermissionService: service.RolePermissionService}}),
	}
}
