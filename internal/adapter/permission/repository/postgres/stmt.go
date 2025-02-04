package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func (repo *PermissionRepository) prepareStatements() {
	repo.statement = StatementList{}
}

func (repo *PermissionRepository) prepareCreatePermission() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryCreatePermission); err != nil {
		log.Panic("[prepareCreatePermission] error:", err)
	}
	repo.statement.CreatePermission = stmt
}

func (repo *PermissionRepository) prepareListPermissions() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryListPermissions); err != nil {
		log.Panic("[prepareListPermissions] error:", err)
	}
	repo.statement.ListPermissions = stmt
}

func (repo *PermissionRepository) prepareGetPermissionByID() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryGetPermissionByID); err != nil {
		log.Panic("[prepareGetPermissionByID] error:", err)
	}
	repo.statement.GetPermissionByID = stmt
}

func (repo *PermissionRepository) prepareGetPermissionByName() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryGetPermissionByName); err != nil {
		log.Panic("[prepareGetPermissionByName] error:", err)
	}
	repo.statement.GetPermissionByName = stmt
}
