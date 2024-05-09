package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	db "github.com/kongsakchai/catopia-backend/database"
	"github.com/kongsakchai/catopia-backend/domain"
	errs "github.com/kongsakchai/catopia-backend/domain/error"
)

type userRepository struct {
	db *db.Database
}

func NewUserRepository(db *db.Database) domain.UserRepository {
	return &userRepository{db}
}

func (u *userRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	getSql, args, err := sq.Select("*").From("users").Where(sq.Eq{"email": email}).ToSql()
	if err != nil {
		return nil, errs.NewError(errs.ErrUserGetByEmail, err)
	}

	user := &domain.User{}
	err = u.db.GetContext(ctx, user, getSql, args...)
	if err != nil {
		return nil, errs.NewError(errs.ErrUserGetByEmail, db.HandlerError(err))
	}

	return user, nil
}

func (u *userRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	getSql, args, err := sq.Select("*").From("users").Where(sq.Eq{"username": username}).ToSql()
	if err != nil {
		return nil, errs.NewError(errs.ErrUserGetByUsername, err)
	}

	user := &domain.User{}
	err = u.db.GetContext(ctx, user, getSql, args...)
	if err != nil {
		return nil, errs.NewError(errs.ErrUserGetByUsername, db.HandlerError(err))
	}

	return user, nil
}

func (u *userRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	getSql, args, err := sq.Select("*").From("users").Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		return nil, errs.NewError(errs.ErrUserGetByID, err)
	}

	user := &domain.User{}
	err = u.db.GetContext(ctx, user, getSql, args...)
	if err != nil {
		return nil, errs.NewError(errs.ErrUserGetByID, db.HandlerError(err))
	}

	return user, nil
}

func (u *userRepository) Create(ctx context.Context, user *domain.User) error {
	insertSql, args, err := sq.Insert("users").
		Columns("username", "password", "email", "salt", "gender", "profile", "date").
		Values(user.Username, user.Password, user.Email, user.Salt, user.Gender, user.Profile, user.Date).
		ToSql()

	if err != nil {
		return errs.NewError(errs.ErrUserCreate, err)
	}

	_, err = u.db.ExecContext(ctx, insertSql, args...)
	if err != nil {
		return errs.NewError(errs.ErrUserCreate, db.HandlerError(err))
	}

	return nil
}

func (u *userRepository) Update(ctx context.Context, user *domain.User) error {
	updateSql, args, err := sq.Update("users").
		Set("username", user.Username).
		Set("email", user.Email).
		Set("password", user.Password).
		Set("salt", user.Salt).
		Set("gender", user.Gender).
		Set("profile", user.Profile).
		Set("date", user.Date).
		Where(sq.Eq{"id": user.ID}).
		ToSql()

	if err != nil {
		return errs.NewError(errs.ErrUserUpdate, err)
	}

	_, err = u.db.ExecContext(ctx, updateSql, args...)
	if err != nil {
		return errs.NewError(errs.ErrUserUpdate, db.HandlerError(err))
	}

	return nil
}

func (u *userRepository) UpdatePassword(ctx context.Context, id int64, password string, salt string) error {
	updateSql, args, err := sq.Update("users").
		Set("password", password).
		Set("salt", salt).
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return errs.NewError(errs.ErrUserUpdatePassword, err)
	}

	_, err = u.db.ExecContext(ctx, updateSql, args...)
	if err != nil {
		return errs.NewError(errs.ErrUserUpdatePassword, db.HandlerError(err))
	}

	return nil
}
