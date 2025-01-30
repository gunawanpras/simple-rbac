package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/helper"
	"github.com/gunawanpras/simple-rbac/internal/domain/user/model"
	"github.com/gunawanpras/simple-rbac/util/constant"
)

// Create a new user and return the generated ID
func (repo *UserRepository) CreateUser(ctx context.Context, user model.User) (id uuid.UUID, err error) {
	user.ID = helper.UUIDHelper.New()

	repo.prepareCreateUser()
	_, err = repo.statement.CreateUser.ExecContext(ctx, user.ID, user.Name, user.Email, user.Password, user.RoleID, user.CreatedBy, user.CreatedAt)
	if err != nil {
		return uuid.Nil, err
	}
	return user.ID, nil
}

// Retrieve a list of all users
func (repo *UserRepository) ListUsers(ctx context.Context) (users model.ListUsers, err error) {
	repo.prepareListUsers()
	rows, err := repo.statement.ListUsers.QueryxContext(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(constant.DbRecordNotFound)
		}

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err = rows.StructScan(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Retrieve a single user by ID
func (repo *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (user model.User, err error) {
	repo.prepareGetUserByID()
	err = repo.statement.GetUserByID.QueryRowxContext(ctx, id).StructScan(&user)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New(constant.DbRecordNotFound)
		}

		return user, err
	}

	return user, nil
}

func (repo *UserRepository) GetUserByEmail(ctx context.Context, email string) (user model.User, err error) {
	repo.prepareGetUserByEmail()
	err = repo.statement.GetUserByEmail.QueryRowxContext(ctx, email).StructScan(&user)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New(constant.DbRecordNotFound)
		}

		return user, err
	}

	return user, nil
}
