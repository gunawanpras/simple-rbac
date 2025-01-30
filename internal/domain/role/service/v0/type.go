package v0

import (
	roleRepo "github.com/gunawanpras/simple-rbac/internal/domain/role/repository"
)

type (
	RepoAttribute struct {
		RoleRepo roleRepo.Repository
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
