package handler

import (
	"github.com/gunawanpras/simple-rbac/internal/core/user/port"
)

type (
	UserHandler struct {
		service ServiceAttribute
	}

	ServiceAttribute struct {
		UserService port.Service
	}

	InitAttribute struct {
		Service ServiceAttribute
	}
)
