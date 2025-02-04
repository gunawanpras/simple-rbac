package postgres

import "github.com/jmoiron/sqlx"

type PermissionRepository struct {
	db        DB
	statement StatementList
}

type DB struct {
	Db *sqlx.DB
}

type StatementList struct {
	CreatePermission    *sqlx.Stmt
	ListPermissions     *sqlx.Stmt
	GetPermissionByID   *sqlx.Stmt
	GetPermissionByName *sqlx.Stmt
}

type InitAttribute struct {
	DB DB
}
