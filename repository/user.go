package repository

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	db "github.com/kongsakchai/catopia-backend/database"
	"github.com/kongsakchai/catopia-backend/domain"
)

type userRepository struct {
	db *db.Database
}

func NewUserRepository() domain.UserRepository {
	db := db.GetDB()
	return &userRepository{db}
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*domain.UserModel, error) {
	sqlBuild := sq.Select("*").From("users").Where(sq.Eq{"email": email})

	query, arg, err := sqlBuild.ToSql()
	if err != nil {
		return nil, fmt.Errorf("get user by email: cannot build query: %w", err)
	}

	user := &domain.UserModel{}
	err = r.db.GetContext(ctx, user, query, arg...)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("get user by email: cannot execute query: %w", err)
	}

	return user, nil
}

func (r *userRepository) GetByID(ctx context.Context, id int) (*domain.UserModel, error) {
	sqlBuild := sq.Select("*").From("users").Where(sq.Eq{"id": id})

	query, arg, err := sqlBuild.ToSql()
	if err != nil {
		return nil, fmt.Errorf("get user by id: cannot build query: %w", err)
	}

	user := &domain.UserModel{}
	err = r.db.GetContext(ctx, user, query, arg...)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("get user by id: cannot execute query: %w", err)
	}

	return user, nil
}

func (r *userRepository) Create(ctx context.Context, user *domain.UserModel) error {
	sqlBuild := sq.Insert("users").
		Columns("username", "password", "email", "salt", "gender", "date").
		Values(sq.Expr(":username,:password,:email,:salt,:gender,:date"))

	query, _, err := sqlBuild.ToSql()
	if err != nil {
		return fmt.Errorf("create user: cannot build query: %w", err)
	}

	return r.db.UseTx(ctx, func(tx *db.Tx) error {
		_, err := tx.NamedExecContext(ctx, query, user)
		if err != nil {
			return fmt.Errorf("create user: cannot execute query: %w", err)
		}

		return nil
	})
}

func (r *userRepository) Update(ctx context.Context, user *domain.UserModel) error {
	sqlBuild := sq.Update("users").
		Set("username", user.Username).
		Set("password", user.Password).
		Set("email", user.Email).
		Set("salt", user.Salt).
		Set("gender", user.Gender).
		Set("date", user.Date).
		Where(sq.Eq{"id": user.ID})

	query, args, err := sqlBuild.ToSql()
	if err != nil {
		return fmt.Errorf("update user: cannot build query: %w", err)
	}

	return r.db.UseTx(ctx, func(tx *db.Tx) error {
		_, err := tx.ExecContext(ctx, query, args...)
		if err != nil {
			return fmt.Errorf("update user: cannot execute query: %w", err)
		}

		return nil
	})
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*domain.UserModel, error) {
	sqlBuild := sq.Select("*").From("users").Where(sq.Eq{"username": username})

	query, arg, err := sqlBuild.ToSql()
	if err != nil {
		return nil, fmt.Errorf("get user by username: cannot build query: %w", err)
	}

	user := &domain.UserModel{}
	err = r.db.GetContext(ctx, user, query, arg...)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("get user by username: cannot execute query: %w", err)
	}

	return user, nil
}
