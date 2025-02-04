package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func (repo *RolePermissionRepository) prepareStatements() {
	repo.statement = StatementList{}
}

func (repo *RolePermissionRepository) prepareCreateRolePermission() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryCreateRolePermission); err != nil {
		log.Panic("[prepareCreateRolePermission] error:", err)
	}
	repo.statement.CreateRolePermission = stmt
}

func (repo *RolePermissionRepository) prepareListRolePermissions() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryListRolePermissions); err != nil {
		log.Panic("[prepareListRolePermissions] error:", err)
	}
	repo.statement.ListRolePermissions = stmt
}

func (repo *RolePermissionRepository) prepareGetRolePermissionByRoleID() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryGetRolePermissionByRoleID); err != nil {
		log.Panic("[prepareGetRolePermissionByRoleID] error:", err)
	}
	repo.statement.ListRolePermissionByRoleID = stmt
}

func (repo *RolePermissionRepository) prepareGetRolePermissionByPermissionID() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryGetRolePermissionByPermissionID); err != nil {
		log.Panic("[prepareGetRolePermissionByPermissionID] error:", err)
	}
	repo.statement.GetRolePermissionByPermissionID = stmt
}
