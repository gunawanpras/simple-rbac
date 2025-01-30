package server

import (
	permissionRepo "github.com/gunawanpras/simple-rbac/internal/domain/permission/repository"
	permissionRepoPg "github.com/gunawanpras/simple-rbac/internal/domain/permission/repository/postgres"
	rolePermissionRepo "github.com/gunawanpras/simple-rbac/internal/domain/role-permission/repository"
	rolePermissionPg "github.com/gunawanpras/simple-rbac/internal/domain/role-permission/repository/postgres"
	roleRepo "github.com/gunawanpras/simple-rbac/internal/domain/role/repository"
	roleRepoPg "github.com/gunawanpras/simple-rbac/internal/domain/role/repository/postgres"
	userRepo "github.com/gunawanpras/simple-rbac/internal/domain/user/repository"
	userRepoPg "github.com/gunawanpras/simple-rbac/internal/domain/user/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	UserRepository           userRepo.Repository
	RoleRepository           roleRepo.Repository
	PermissionRepository     permissionRepo.Repository
	RolePermissionRepository rolePermissionRepo.Repository
}

func NewRepository(db *sqlx.DB) Repository {

	userRepository := userRepoPg.New(userRepoPg.InitAttribute{
		DB: userRepoPg.DB{
			Db: db,
		},
	})

	roleRepository := roleRepoPg.New(roleRepoPg.InitAttribute{
		DB: roleRepoPg.DB{
			Db: db,
		},
	})

	permissionRepository := permissionRepoPg.New(permissionRepoPg.InitAttribute{
		DB: permissionRepoPg.DB{
			Db: db,
		},
	})

	rolePermissionRepository := rolePermissionPg.New(rolePermissionPg.InitAttribute{
		DB: rolePermissionPg.DB{
			Db: db,
		},
	})

	return Repository{
		UserRepository:           userRepository,
		RoleRepository:           roleRepository,
		PermissionRepository:     permissionRepository,
		RolePermissionRepository: rolePermissionRepository,
	}
}
