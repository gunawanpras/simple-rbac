package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func (repo *RbacRepository) prepareStatements() {
	repo.statement = StatementList{}
}

func (repo *RbacRepository) prepareCreateUser() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryCreateUser); err != nil {
		log.Panic("[prepareCreateUser] error:", err)
	}
	repo.statement.CreateUser = stmt
}

func (repo *RbacRepository) prepareListUsers() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryListUsers); err != nil {
		log.Panic("[prepareListUsers] error:", err)
	}
	repo.statement.ListUsers = stmt
}

func (repo *RbacRepository) prepareGetUserByID() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryGetUserByID); err != nil {
		log.Panic("[prepareGetUserByID] error:", err)
	}
	repo.statement.GetUserByID = stmt
}

func (repo *RbacRepository) prepareGetUserByEmail() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryGetUserByEmail); err != nil {
		log.Panic("[prepareGetUserByEmail] error:", err)
	}
	repo.statement.GetUserByEmail = stmt
}

func (repo *RbacRepository) prepareCreateRole() {
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

func (repo *RbacRepository) prepareListRoles() {
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

func (repo *RbacRepository) prepareGetRoleByID() {
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

func (repo *RbacRepository) prepareGetRoleByName() {
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

func (repo *RbacRepository) prepareCreatePermission() {
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

func (repo *RbacRepository) prepareListPermissions() {
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

func (repo *RbacRepository) prepareGetPermissionByID() {
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

func (repo *RbacRepository) prepareGetPermissionByName() {
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

func (repo *RbacRepository) prepareCreateRolePermission() {
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

func (repo *RbacRepository) prepareListRolePermissions() {
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

func (repo *RbacRepository) prepareGetRolePermissionByRoleID() {
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

func (repo *RbacRepository) prepareGetRolePermissionByPermissionID() {
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
