package postgres

import "github.com/jmoiron/sqlx"

type RolePermissionRepository struct {
	db        DB
	statement StatementList
}

type DB struct {
	Db *sqlx.DB
}

type StatementList struct {
	CreateRolePermission            *sqlx.Stmt
	ListRolePermissions             *sqlx.Stmt
	ListRolePermissionByRoleID      *sqlx.Stmt
	GetRolePermissionByPermissionID *sqlx.Stmt
}

type InitAttribute struct {
	DB DB
}
