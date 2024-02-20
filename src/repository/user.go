package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/kongsakchai/catopia-backend/src/model"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) Create(user *model.UserModel) error {
	sql, args, err := squirrel.Insert("users").
		Columns("username", "password", "email", "name", "salt", "gender", "date").
		Values(user.Username, user.Password, user.Email, user.Name, user.Salt, user.Gender, user.Date).
		ToSql()

	if err != nil {
		return err
	}

	_, err = u.db.Exec(sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) GetByUsername(username string) (*model.UserModel, error) {
	user := &model.UserModel{}
	sql, args, err := squirrel.Select("*").
		From("users").
		Where(squirrel.Eq{"username": username}).
		ToSql()

	if err != nil {
		return nil, err
	}

	err = u.db.Get(user, sql, args...)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetByEmail(email string) (*model.UserModel, error) {
	user := &model.UserModel{}
	sql, args, err := squirrel.Select("*").
		From("users").
		Where(squirrel.Eq{"email": email}).
		ToSql()

	if err != nil {
		return nil, err
	}

	err = u.db.Get(user, sql, args...)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetByID(id int) (*model.UserModel, error) {
	user := &model.UserModel{}
	sql, args, err := squirrel.Select("*").
		From("users").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		return nil, err
	}

	err = u.db.Get(user, sql, args...)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) Update(user *model.UserModel) error {
	sql, args, err := squirrel.Update("users").
		Set("username", user.Username).
		Set("password", user.Password).
		Set("email", user.Email).
		Set("name", user.Name).
		Set("salt", user.Salt).
		Set("gender", user.Gender).
		Set("date", user.Date).
		Where(squirrel.Eq{"id": user.ID}).
		ToSql()

	if err != nil {
		return err
	}

	_, err = u.db.Exec(sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) Delete(id int) error {
	sql, args, err := squirrel.Delete("users").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		return err
	}

	_, err = u.db.Exec(sql, args...)
	if err != nil {
		return err
	}

	return nil
}
