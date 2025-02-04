package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func (repo *UserRepository) prepareStatements() {
	repo.statement = StatementList{}
}

func (repo *UserRepository) prepareCreateUser() {
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

func (repo *UserRepository) prepareListUsers() {
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

func (repo *UserRepository) prepareGetUserByID() {
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

func (repo *UserRepository) prepareGetUserByEmail() {
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
