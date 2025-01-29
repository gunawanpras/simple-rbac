package server

import (
	rbacService "github.com/gunawanpras/simple-rbac/internal/domain/rbac/service"
	rbacServiceV0 "github.com/gunawanpras/simple-rbac/internal/domain/rbac/service/v0"
)

type Service struct {
	RbacService rbacService.Service
}

func NewService(repository Repository) Service {
	return Service{
		RbacService: rbacServiceV0.New(rbacServiceV0.InitAttribute{
			Repo: rbacServiceV0.RepoAttribute{
				RbacPg: repository.RbacRepository,
			},
		}),
	}
}
