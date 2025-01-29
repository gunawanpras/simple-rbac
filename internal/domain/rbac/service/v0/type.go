package v0

import (
	rbacRepo "github.com/gunawanpras/simple-rbac/internal/domain/rbac/repository"
)

type (
	RepoAttribute struct {
		RbacPg rbacRepo.Repository
	}

	Config struct{}

	RbacService struct {
		repo   RepoAttribute
		config Config
	}

	InitAttribute struct {
		Repo   RepoAttribute
		Config Config
	}
)
