package v0

import (
	permissionRepo "github.com/gunawanpras/simple-rbac/internal/domain/permission/repository"
)

type (
	RepoAttribute struct {
		PermissionRepo permissionRepo.Repository
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
