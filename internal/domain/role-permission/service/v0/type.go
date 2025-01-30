package v0

import (
	rolePermissionRepo "github.com/gunawanpras/simple-rbac/internal/domain/role-permission/repository"
)

type (
	RepoAttribute struct {
		RolePermissionRepo rolePermissionRepo.Repository
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
