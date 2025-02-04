package handler

import (
	"github.com/gunawanpras/simple-rbac/internal/core/permission/port"
)

type (
	PermissionHandler struct {
		service ServiceAttribute
	}

	ServiceAttribute struct {
		PermissionService port.Service
	}

	InitAttribute struct {
		Service ServiceAttribute
	}
)
