package handler

import "github.com/gunawanpras/simple-rbac/internal/core/role/port"

type (
	RoleHandler struct {
		service ServiceAttribute
	}

	ServiceAttribute struct {
		RoleService port.Service
	}

	InitAttribute struct {
		Service ServiceAttribute
	}
)
