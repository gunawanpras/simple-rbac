package httphandler

import (
	"github.com/gunawanpras/simple-rbac/internal/domain/permission/service"
)

type (
	PermissionHandler struct {
		service ServiceAttribute
	}

	ServiceAttribute struct {
		PermissionService service.Service
	}

	InitAttribute struct {
		Service ServiceAttribute
	}
)
