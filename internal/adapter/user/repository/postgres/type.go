package postgres

import "github.com/jmoiron/sqlx"

type UserRepository struct {
	db        DB
	statement StatementList
}

type DB struct {
	Db *sqlx.DB
}

type StatementList struct {
	CreateUser     *sqlx.Stmt
	ListUsers      *sqlx.Stmt
	GetUserByID    *sqlx.Stmt
	GetUserByEmail *sqlx.Stmt
}

type InitAttribute struct {
	DB DB
}
