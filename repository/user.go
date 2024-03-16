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
	db *db.DB
}

func NewUserRepository(db *db.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}
func (r *userRepository) FindByID(ctx context.Context, id int) (*domain.UserModel, error) {
	sqlBuilder := sq.
		Select("*").
		From("users").
		Where(sq.Eq{"id": id})

	query, args, err := sqlBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("at user.findByID build query : %w", err)
	}

	var user domain.UserModel
	err = r.db.GetContext(ctx, &user, query, args...)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("at user.findByID excute query : %w", err)
	}

	return &user, nil
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*domain.UserModel, error) {
	sqlBuilder := sq.
		Select("*").
		From("users").
		Where(sq.Eq{"username": username})

	query, args, err := sqlBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("at user.findByUsername build query : %w", err)
	}

	var user domain.UserModel
	err = r.db.GetContext(ctx, &user, query, args...)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("at user.findByUsername excute query : %w", err)
	}

	return &user, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*domain.UserModel, error) {
	sqlBuilder := sq.
		Select("*").
		From("users").
		Where(sq.Eq{"email": email})

	query, args, err := sqlBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("at user.findByEmail build query : %w", err)
	}

	var user domain.UserModel
	err = r.db.GetContext(ctx, &user, query, args...)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("at user.findByEmail excute query : %w", err)
	}

	return &user, nil
}

func (r *userRepository) Insert(ctx context.Context, user *domain.UserModel) (int, error) {
	sqlBuilder := sq.
		Insert("users").
		Columns("username", "email", "password", "salt", "gender", "date").
		Values(user.Username, user.Email, user.Password, user.Salt, user.Gender, user.Date)

	query, args, err := sqlBuilder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("at user.insert build query : %w", err)
	}

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return 0, fmt.Errorf("at user.insert excute query : %w", err)
	}

	return 0, nil
}

func (r *userRepository) Update(ctx context.Context, user *domain.UserModel) error {
	sqlBuilder := sq.Update("users").
		Set("gender", user.Gender).
		Set("date", user.Date).
		Set("email", user.Email).
		Where(sq.Eq{"id": user.ID})

	query, args, err := sqlBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("at user.update build query : %w", err)
	}

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("at user.update excute query : %w", err)
	}

	return nil
}

func (r *userRepository) Delete(ctx context.Context, id int) error {
	sqlBuilder := sq.Delete("users").Where(sq.Eq{"id": id})

	query, args, err := sqlBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("at user.delete build query : %w", err)
	}

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("at user.delete excute query : %w", err)
	}

	return nil
}
