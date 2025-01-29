package httphandler

import (
	"github.com/gunawanpras/simple-rbac/internal/domain/rbac/service"
)

type (
	RbacHandler struct {
		service ServiceAttribute
	}

	ServiceAttribute struct {
		RbacService service.Service
	}

	InitAttribute struct {
		Service ServiceAttribute
	}
)
