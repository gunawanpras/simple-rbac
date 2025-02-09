package setup

import (
	permissionService "github.com/gunawanpras/simple-rbac/internal/core/permission/port"
	permissionServiceV0 "github.com/gunawanpras/simple-rbac/internal/core/permission/service/v0"
	rolePermissionService "github.com/gunawanpras/simple-rbac/internal/core/role-permission/port"
	rolePermissionServiceV0 "github.com/gunawanpras/simple-rbac/internal/core/role-permission/service/v0"
	roleService "github.com/gunawanpras/simple-rbac/internal/core/role/port"
	roleServiceV0 "github.com/gunawanpras/simple-rbac/internal/core/role/service/v0"
	userService "github.com/gunawanpras/simple-rbac/internal/core/user/port"
	userServiceV0 "github.com/gunawanpras/simple-rbac/internal/core/user/service/v0"
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
