package server

import (
	rbacRepo "github.com/gunawanpras/simple-rbac/internal/domain/rbac/repository"
	rbacRepoPg "github.com/gunawanpras/simple-rbac/internal/domain/rbac/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	RbacRepository rbacRepo.Repository
}

func NewRepository(db *sqlx.DB) Repository {

	rbacRepository := rbacRepoPg.New(rbacRepoPg.InitAttribute{
		DB: rbacRepoPg.DB{
			Db: db,
		},
	})

	return Repository{
		RbacRepository: rbacRepository,
	}
}
