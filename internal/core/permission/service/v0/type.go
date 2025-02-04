package v0

import (
	"github.com/gunawanpras/simple-rbac/internal/core/permission/port"
)

type (
	RepoAttribute struct {
		PermissionRepo port.Repository
	}

	Config struct{}

	PermissionService struct {
		repo   RepoAttribute
		config Config
	}

	InitAttribute struct {
		Repo   RepoAttribute
		Config Config
	}
)
