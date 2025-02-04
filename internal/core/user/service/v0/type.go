package v0

import (
	"github.com/gunawanpras/simple-rbac/internal/core/user/port"
)

type (
	RepoAttribute struct {
		UserRepo port.Repository
	}

	Config struct{}

	UserService struct {
		repo   RepoAttribute
		config Config
	}

	InitAttribute struct {
		Repo   RepoAttribute
		Config Config
	}
)
