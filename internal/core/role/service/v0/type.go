package v0

import (
	"github.com/gunawanpras/simple-rbac/internal/core/role/port"
)

type (
	RepoAttribute struct {
		RoleRepo port.Repository
	}

	Config struct{}

	RoleService struct {
		repo   RepoAttribute
		config Config
	}

	InitAttribute struct {
		Repo   RepoAttribute
		Config Config
	}
)
