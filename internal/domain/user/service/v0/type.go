package v0

import (
	userRepo "github.com/gunawanpras/simple-rbac/internal/domain/user/repository"
)

type (
	RepoAttribute struct {
		UserRepo userRepo.Repository
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
