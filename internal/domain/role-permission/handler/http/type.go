package httphandler

import (
	"github.com/gunawanpras/simple-rbac/internal/domain/role-permission/service"
)

type (
	RolePermissionHandler struct {
		service ServiceAttribute
	}

	ServiceAttribute struct {
		RolePermissionService service.Service
	}

	InitAttribute struct {
		Service ServiceAttribute
	}
)
