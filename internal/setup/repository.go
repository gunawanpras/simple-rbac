package setup

import (
	permissionRepoPg "github.com/gunawanpras/simple-rbac/internal/adapter/permission/repository/postgres"
	rolePermissionPg "github.com/gunawanpras/simple-rbac/internal/adapter/role-permission/repository/postgres"
	roleRepoPg "github.com/gunawanpras/simple-rbac/internal/adapter/role/repository/postgres"
	userRepoPg "github.com/gunawanpras/simple-rbac/internal/adapter/user/repository/postgres"
	permissionRepo "github.com/gunawanpras/simple-rbac/internal/core/permission/port"
	rolePermissionRepo "github.com/gunawanpras/simple-rbac/internal/core/role-permission/port"
	roleRepo "github.com/gunawanpras/simple-rbac/internal/core/role/port"
	userRepo "github.com/gunawanpras/simple-rbac/internal/core/user/port"
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
