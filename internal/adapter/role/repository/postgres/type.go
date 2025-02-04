package postgres

import "github.com/jmoiron/sqlx"

type RoleRepository struct {
	db        DB
	statement StatementList
}

type DB struct {
	Db *sqlx.DB
}

type StatementList struct {
	CreateRole    *sqlx.Stmt
	ListRoles     *sqlx.Stmt
	GetRoleByID   *sqlx.Stmt
	GetRoleByName *sqlx.Stmt
}

type InitAttribute struct {
	DB DB
}
