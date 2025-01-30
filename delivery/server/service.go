package server

import (
	permissionService "github.com/gunawanpras/simple-rbac/internal/domain/permission/service"
	permissionServiceV0 "github.com/gunawanpras/simple-rbac/internal/domain/permission/service/v0"
	rolePermissionService "github.com/gunawanpras/simple-rbac/internal/domain/role-permission/service"
	rolePermissionServiceV0 "github.com/gunawanpras/simple-rbac/internal/domain/role-permission/service/v0"
	roleService "github.com/gunawanpras/simple-rbac/internal/domain/role/service"
	roleServiceV0 "github.com/gunawanpras/simple-rbac/internal/domain/role/service/v0"
	userService "github.com/gunawanpras/simple-rbac/internal/domain/user/service"
	userServiceV0 "github.com/gunawanpras/simple-rbac/internal/domain/user/service/v0"
)

type Service struct {
	UserService           userService.Service
	RoleService           roleService.Service
	PermissionService     permissionService.Service
	RolePermissionService rolePermissionService.Service
}

func NewService(repository Repository) Service {
	return Service{
		UserService: userServiceV0.New(userServiceV0.InitAttribute{
			Repo: userServiceV0.RepoAttribute{
				UserRepo: repository.UserRepository,
			},
		}),

		RoleService: roleServiceV0.New(roleServiceV0.InitAttribute{
			Repo: roleServiceV0.RepoAttribute{
				RoleRepo: repository.RoleRepository,
			},
		}),

		PermissionService: permissionServiceV0.New(permissionServiceV0.InitAttribute{
			Repo: permissionServiceV0.RepoAttribute{
				PermissionRepo: repository.PermissionRepository,
			},
		}),

		RolePermissionService: rolePermissionServiceV0.New(rolePermissionServiceV0.InitAttribute{
			Repo: rolePermissionServiceV0.RepoAttribute{
				RolePermissionRepo: repository.RolePermissionRepository,
			},
		}),
	}
}
