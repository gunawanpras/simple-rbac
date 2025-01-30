package httphandler

import (
	"github.com/gunawanpras/simple-rbac/internal/domain/role/service"
)

type (
	RoleHandler struct {
		service ServiceAttribute
	}

	ServiceAttribute struct {
		RoleService service.Service
	}

	InitAttribute struct {
		Service ServiceAttribute
	}
)
