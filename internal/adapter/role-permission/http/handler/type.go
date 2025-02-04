package handler

import (
	"github.com/gunawanpras/simple-rbac/internal/core/role-permission/port"
)

type (
	RolePermissionHandler struct {
		service ServiceAttribute
	}

	ServiceAttribute struct {
		RolePermissionService port.Service
	}

	InitAttribute struct {
		Service ServiceAttribute
	}
)
