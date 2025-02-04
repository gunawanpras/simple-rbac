package v0

import "github.com/gunawanpras/simple-rbac/internal/core/role-permission/port"

type (
	RepoAttribute struct {
		RolePermissionRepo port.Repository
	}

	Config struct{}

	RolePermissionService struct {
		repo   RepoAttribute
		config Config
	}

	InitAttribute struct {
		Repo   RepoAttribute
		Config Config
	}
)
