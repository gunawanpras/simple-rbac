package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/gunawanpras/simple-rbac/internal/core/user/domain"
	"github.com/gunawanpras/simple-rbac/pkg/util/constant"
	"github.com/gunawanpras/simple-rbac/pkg/util/uuidutil"
)

// Create a new user and return the generated ID
func (repo *UserRepository) CreateUser(ctx context.Context, user domain.User) (id uuid.UUID, err error) {
	user.ID = uuidutil.UUIDHelper.New()

	repo.prepareCreateUser()
	_, err = repo.statement.CreateUser.ExecContext(ctx, user.ID, user.Name, user.Email, user.Password, user.RoleID, user.CreatedBy, user.CreatedAt)
	if err != nil {
		return uuid.Nil, err
	}
	return user.ID, nil
}

// Retrieve a list of all users
func (repo *UserRepository) ListUsers(ctx context.Context) (res domain.ListUsers, err error) {
	var (
		user  User
		users ListUsers
	)

	repo.prepareListUsers()
	rows, err := repo.statement.ListUsers.QueryxContext(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(constant.DataNotFound)
		}

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user = User{}
		if err = rows.StructScan(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if !users.Validate() {
		return nil, errors.New(constant.DbReturnedMalformedData)
	}

	return users.ToModel(), nil
}

// Retrieve a single user by ID
func (repo *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (res domain.User, err error) {
	var user User

	repo.prepareGetUserByID()
	err = repo.statement.GetUserByID.QueryRowxContext(ctx, id).StructScan(&user)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.New(constant.DataNotFound)
		}

		return res, err
	}

	if !user.Validate() {
		return res, errors.New(constant.DbReturnedMalformedData)
	}

	return user.ToModel(), nil
}

func (repo *UserRepository) GetUserByEmail(ctx context.Context, email string) (res domain.User, err error) {
	var user User

	repo.prepareGetUserByEmail()
	err = repo.statement.GetUserByEmail.QueryRowxContext(ctx, email).StructScan(&user)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.New(constant.DataNotFound)
		}

		return res, err
	}

	if !user.Validate() {
		return res, errors.New(constant.DbReturnedMalformedData)
	}

	return user.ToModel(), nil
}
