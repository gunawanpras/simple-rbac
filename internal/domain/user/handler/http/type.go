package httphandler

import (
	"github.com/gunawanpras/simple-rbac/internal/domain/user/service"
)

type (
	UserHandler struct {
		service ServiceAttribute
	}

	ServiceAttribute struct {
		UserService service.Service
	}

	InitAttribute struct {
		Service ServiceAttribute
	}
)
