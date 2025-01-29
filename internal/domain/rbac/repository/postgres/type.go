package postgres

import "github.com/jmoiron/sqlx"

type RbacRepository struct {
	db        DB
	statement StatementList
}

type DB struct {
	Db *sqlx.DB
}

type StatementList struct {
	CreateUser                      *sqlx.Stmt
	ListUsers                       *sqlx.Stmt
	GetUserByID                     *sqlx.Stmt
	GetUserByEmail                  *sqlx.Stmt
	CreateRole                      *sqlx.Stmt
	ListRoles                       *sqlx.Stmt
	GetRoleByID                     *sqlx.Stmt
	GetRoleByName                   *sqlx.Stmt
	CreatePermission                *sqlx.Stmt
	ListPermissions                 *sqlx.Stmt
	GetPermissionByID               *sqlx.Stmt
	GetPermissionByName             *sqlx.Stmt
	CreateRolePermission            *sqlx.Stmt
	ListRolePermissions             *sqlx.Stmt
	ListRolePermissionByRoleID      *sqlx.Stmt
	GetRolePermissionByPermissionID *sqlx.Stmt
}

type InitAttribute struct {
	DB DB
}
