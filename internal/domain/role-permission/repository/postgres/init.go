package postgres

import (
	"fmt"
	"log"

	"github.com/gunawanpras/simple-rbac/internal/domain/role-permission/repository"
)

func New(attr InitAttribute) repository.Repository {
	if err := attr.validate(); err != nil {
		log.Panic(err)
	}

	repo := &RolePermissionRepository{
		db: attr.DB,
	}

	repo.prepareStatements()

	return repo
}

func (init InitAttribute) validate() error {
	if !init.DB.validate() {
		return fmt.Errorf("missing DB driver : %+v", init.DB)
	}

	return nil
}

func (db DB) validate() bool {
	return db.Db != nil
}
