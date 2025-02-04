package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func (repo *RoleRepository) prepareStatements() {
	repo.statement = StatementList{}
}

func (repo *RoleRepository) prepareCreateRole() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryCreateRole); err != nil {
		log.Panic("[prepareCreateRole] error:", err)
	}
	repo.statement.CreateRole = stmt
}

func (repo *RoleRepository) prepareListRoles() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryListRoles); err != nil {
		log.Panic("[prepareListRoles] error:", err)
	}
	repo.statement.ListRoles = stmt
}

func (repo *RoleRepository) prepareGetRoleByID() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryGetRoleByID); err != nil {
		log.Panic("[prepareGetRoleByID] error:", err)
	}
	repo.statement.GetRoleByID = stmt
}

func (repo *RoleRepository) prepareGetRoleByName() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryGetRoleByName); err != nil {
		log.Panic("[prepareGetRoleByName] error:", err)
	}
	repo.statement.GetRoleByName = stmt
}
